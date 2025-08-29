package klabapi

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/DeRuina/KUHA-REST-API/internal/auth/authz"
	"github.com/DeRuina/KUHA-REST-API/internal/store/cache"
	"github.com/DeRuina/KUHA-REST-API/internal/store/klab"
	"github.com/DeRuina/KUHA-REST-API/internal/utils"
)

type KlabDataHandler struct {
	store klab.Data
	cache *cache.Storage
}

func NewKlabDataHandler(store klab.Data, cache *cache.Storage) *KlabDataHandler {
	return &KlabDataHandler{store: store, cache: cache}
}

type KlabDataBulkInput map[string]KlabDataBundleInput

type KlabDataBundleInput struct {
	Customer        []KlabCustomerInput    `json:"customer"         validate:"omitempty,dive"`
	MeasurementList []KlabMeasurementInput `json:"measurement_list" validate:"omitempty,dive"`
	DirTest         []KlabDirTestInput     `json:"dirtest"          validate:"omitempty,dive"`
	DirTestSteps    []KlabDirTestStepInput `json:"dirteststeps"     validate:"omitempty,dive"`
	DirReport       []KlabDirReportInput   `json:"dirreport"        validate:"omitempty,dive"`
	DirRawData      []KlabDirRawDataInput  `json:"dirrawdata"       validate:"omitempty,dive"`
	DirResults      []KlabDirResultsInput  `json:"dirresults"       validate:"omitempty,dive"`
}

// InsertKlabDataBulk godoc
//
//	@Summary		Insert k-Lab data (one customer per request)
//	@Description	Insert a full k-Lab bundle for a single customer. The JSON root object key must be the customer id.
//	@Tags			KLAB - Data
//	@Accept			json
//	@Produce		json
//	@Param			data	body	swagger.KlabDataBulkDoc	true	"klab data"
//	@Success		201		"Data processed successfully"
//	@Failure		400		{object}	swagger.ValidationErrorResponse
//	@Failure		401		{object}	swagger.UnauthorizedResponse
//	@Failure		403		{object}	swagger.ForbiddenResponse
//	@Failure		500		{object}	swagger.InternalServerErrorResponse
//	@Failure		503		{object}	swagger.ServiceUnavailableResponse
//	@Security		BearerAuth
//	@Router			/klab/data [post]
func (h *KlabDataHandler) InsertKlabDataBulk(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}

	var input KlabDataBulkInput
	if err := utils.ReadJSON(w, r, &input); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	if len(input) != 1 {
		utils.BadRequestResponse(w, r, fmt.Errorf("payload must contain exactly one customer key at the top level"))
		return
	}

	var (
		keyStr string
		bundle KlabDataBundleInput
	)
	for k, v := range input {
		keyStr, bundle = k, v
		break
	}

	id, err := utils.ParsePositiveInt32(keyStr)
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	if err := utils.GetValidator().Struct(bundle); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	var p klab.KlabDataPayload

	// 1) customer rows
	for _, c := range bundle.Customer {
		arg, err := mapCustomerToParams(c, id) // forces/aligns idcustomer
		if err != nil {
			utils.BadRequestResponse(w, r, err)
			return
		}
		p.Customers = append(p.Customers, arg)
	}

	// 2) measurement_list
	for _, m := range bundle.MeasurementList {
		m.IdCustomer = &id
		arg, err := mapMeasurementToParams(m)
		if err != nil {
			utils.BadRequestResponse(w, r, err)
			return
		}
		p.Measurements = append(p.Measurements, arg)
	}

	// 3) child tables
	for _, t := range bundle.DirTest {
		arg, err := mapDirTestToParams(t)
		if err != nil {
			utils.BadRequestResponse(w, r, err)
			return
		}
		p.DirTests = append(p.DirTests, arg)
	}
	for _, st := range bundle.DirTestSteps {
		arg, err := mapDirTestStepToParams(st)
		if err != nil {
			utils.BadRequestResponse(w, r, err)
			return
		}
		p.DirTestSteps = append(p.DirTestSteps, arg)
	}
	for _, rp := range bundle.DirReport {
		arg, err := mapDirReportToParams(rp)
		if err != nil {
			utils.BadRequestResponse(w, r, err)
			return
		}
		p.DirReports = append(p.DirReports, arg)
	}
	for _, rd := range bundle.DirRawData {
		arg, err := mapDirRawDataToParams(rd)
		if err != nil {
			utils.BadRequestResponse(w, r, err)
			return
		}
		p.DirRawData = append(p.DirRawData, arg)
	}
	for _, rs := range bundle.DirResults {
		arg, err := mapDirResultsToParams(rs)
		if err != nil {
			utils.BadRequestResponse(w, r, err)
			return
		}
		p.DirResults = append(p.DirResults, arg)
	}

	if err := h.store.InsertKlabDataBulk(r.Context(), []klab.KlabDataPayload{p}); err != nil {
		utils.HandleDatabaseError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// GetKlabData godoc
//
//	@Summary		Get kLab data by customer ID
//	@Description	Returns measurement_list + all child tables for the given customer (no customer row)
//	@Tags			KLAB - Data
//	@Accept			json
//	@Produce		json
//	@Param			id	query		string	true	"Customer ID (integer)"
//	@Success		200	{object}	swagger.KlabDataResponse
//	@Failure		400	{object}	swagger.ValidationErrorResponse
//	@Failure		401	{object}	swagger.UnauthorizedResponse
//	@Failure		403	{object}	swagger.ForbiddenResponse
//	@Failure		404	{object}	swagger.NotFoundResponse
//	@Failure		500	{object}	swagger.InternalServerErrorResponse
//	@Failure		503	{object}	swagger.ServiceUnavailableResponse
//	@Security		BearerAuth
//	@Router			/klab/data [get]
func (h *KlabDataHandler) GetKlabData(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}

	if err := utils.ValidateParams(r, []string{"id"}); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	idStr := r.URL.Query().Get("id")
	id, err := utils.ParsePositiveInt32(idStr)
	if err != nil {
		utils.BadRequestResponse(w, r, utils.ErrInvalidParameter)
		return
	}

	res, err := h.store.GetDataByCustomerIDNoCustomer(r.Context(), id)
	if err == sql.ErrNoRows {
		utils.NotFoundResponse(w, r, err)
		return
	}
	if err != nil {
		utils.InternalServerError(w, r, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]any{
		"customer_id":  res.CustomerID,
		"measurements": res.Measurements,
		"dirtest":      res.DirTests,
		"dirteststeps": res.DirTestSteps,
		"dirreport":    res.DirReports,
		"dirrawdata":   res.DirRawData,
		"dirresults":   res.DirResults,
	})
}
