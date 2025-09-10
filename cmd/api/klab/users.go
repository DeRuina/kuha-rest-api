package klabapi

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/DeRuina/KUHA-REST-API/docs/swagger"
	"github.com/DeRuina/KUHA-REST-API/internal/auth/authz"
	"github.com/DeRuina/KUHA-REST-API/internal/store/cache"
	"github.com/DeRuina/KUHA-REST-API/internal/store/klab"
	"github.com/DeRuina/KUHA-REST-API/internal/utils"
)

type UserDataHandler struct {
	store klab.Users
	cache *cache.Storage
}

func NewUserDataHandler(store klab.Users, cache *cache.Storage) *UserDataHandler {
	return &UserDataHandler{store: store, cache: cache}
}

type KlabUserParams struct {
	ID string `validate:"required,numeric"`
}

// GetSporttiIDs godoc
//
//	@Summary		Get all Sportti IDs
//	@Description	Returns a list of all Sportti IDs
//	@Tags			KLAB - User
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	swagger.UserDataKlabResponse
//	@Failure		401	{object}	swagger.UnauthorizedResponse
//	@Failure		403	{object}	swagger.ForbiddenResponse
//	@Failure		500	{object}	swagger.InternalServerErrorResponse
//	@Failure		503	{object}	swagger.ServiceUnavailableResponse
//	@Security		BearerAuth
//	@Router			/klab/sport-ids [get]
func (h *UserDataHandler) GetSporttiIDs(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}

	sporttiIDs, err := h.store.GetAllSporttiIDs(r.Context())
	if err != nil {
		utils.InternalServerError(w, r, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]any{
		"sportti_ids": sporttiIDs,
	})
}

// GetUser godoc
//
//	@Summary		Get customer by Sportti ID
//	@Description	Retrieve a single KLAB customer by sportti_id
//	@Tags			KLAB - User
//	@Accept			json
//	@Produce		json
//	@Param			id	query		integer	true	"Sportti ID"
//	@Success		200	{object}	swagger.UserKlabResponse
//	@Failure		400	{object}	swagger.ValidationErrorResponse
//	@Failure		401	{object}	swagger.UnauthorizedResponse
//	@Failure		403	{object}	swagger.ForbiddenResponse
//	@Failure		404	{object}	swagger.NotFoundResponse
//	@Failure		500	{object}	swagger.InternalServerErrorResponse
//	@Failure		503	{object}	swagger.ServiceUnavailableResponse
//	@Security		BearerAuth
//	@Router			/klab/user [get]
func (h *UserDataHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}

	if err := utils.ValidateParams(r, []string{"id"}); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	params := KlabUserParams{
		ID: r.URL.Query().Get("id"),
	}

	if err := utils.GetValidator().Struct(params); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	sporttiID, err := utils.ParseSporttiID(params.ID)
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	idcustomer, err := h.store.GetCustomerIDBySporttiID(r.Context(), sporttiID)
	if errors.Is(err, sql.ErrNoRows) {
		utils.NotFoundResponse(w, r, err)
		return
	}
	if err != nil {
		utils.InternalServerError(w, r, err)
		return
	}

	row, err := h.store.GetCustomerByID(r.Context(), idcustomer)
	if errors.Is(err, sql.ErrNoRows) {
		utils.NotFoundResponse(w, r, err)
		return
	}
	if err != nil {
		utils.InternalServerError(w, r, err)
		return
	}

	resp := swagger.KlabCustomerResponse{
		Idcustomer:         row.Idcustomer,
		Firstname:          row.Firstname,
		Lastname:           row.Lastname,
		Idgroups:           utils.Int32PtrOrNil(row.Idgroups),
		Dob:                utils.FormatDatePtr(row.Dob),
		Sex:                utils.Int32PtrOrNil(row.Sex),
		DobYear:            utils.Int32PtrOrNil(row.DobYear),
		DobMonth:           utils.Int32PtrOrNil(row.DobMonth),
		DobDay:             utils.Int32PtrOrNil(row.DobDay),
		PidNumber:          utils.StringPtrOrNil(row.PidNumber),
		Company:            utils.StringPtrOrNil(row.Company),
		Occupation:         utils.StringPtrOrNil(row.Occupation),
		Education:          utils.StringPtrOrNil(row.Education),
		Address:            utils.StringPtrOrNil(row.Address),
		PhoneHome:          utils.StringPtrOrNil(row.PhoneHome),
		PhoneWork:          utils.StringPtrOrNil(row.PhoneWork),
		PhoneMobile:        utils.StringPtrOrNil(row.PhoneMobile),
		Faxno:              utils.StringPtrOrNil(row.Faxno),
		Email:              utils.StringPtrOrNil(row.Email),
		Username:           utils.StringPtrOrNil(row.Username),
		Password:           utils.StringPtrOrNil(row.Password),
		Readonly:           utils.Int32PtrOrNil(row.Readonly),
		Warnings:           utils.Int32PtrOrNil(row.Warnings),
		AllowToSave:        utils.Int32PtrOrNil(row.AllowToSave),
		AllowToCloud:       utils.Int32PtrOrNil(row.AllowToCloud),
		Flag2:              utils.Int32PtrOrNil(row.Flag2),
		Idsport:            utils.Int32PtrOrNil(row.Idsport),
		Medication:         utils.StringPtrOrNil(row.Medication),
		Addinfo:            utils.StringPtrOrNil(row.Addinfo),
		TeamName:           utils.StringPtrOrNil(row.TeamName),
		Add1:               utils.Int32PtrOrNil(row.Add1),
		Athlete:            utils.Int32PtrOrNil(row.Athlete),
		Add10:              utils.StringPtrOrNil(row.Add10),
		Add20:              utils.StringPtrOrNil(row.Add20),
		Updatemode:         utils.Int32PtrOrNil(row.Updatemode),
		WeightKg:           utils.Float64PtrOrNil(row.WeightKg),
		HeightCm:           utils.Float64PtrOrNil(row.HeightCm),
		DateModified:       utils.Float64PtrOrNil(row.DateModified),
		RecomTestlevel:     utils.Int32PtrOrNil(row.RecomTestlevel),
		CreatedBy:          utils.Int64PtrOrNil(row.CreatedBy),
		ModBy:              utils.Int64PtrOrNil(row.ModBy),
		ModDate:            utils.FormatTimestampPtr(row.ModDate),
		Deleted:            utils.Int16AsInt32PtrOrNil(row.Deleted),
		CreatedDate:        utils.FormatTimestampPtr(row.CreatedDate),
		Modded:             utils.Int16AsInt32PtrOrNil(row.Modded),
		AllowAnonymousData: utils.BoolPtrOrNil(row.AllowAnonymousData),
		Locked:             utils.Int16AsInt32PtrOrNil(row.Locked),
		AllowToSprintai:    utils.Int32PtrOrNil(row.AllowToSprintai),
		TosprintaiFrom:     utils.FormatDatePtr(row.TosprintaiFrom),
		StatSent:           utils.FormatDatePtr(row.StatSent),
		SporttiID:          utils.StringPtrOrNil(row.SporttiID),
	}

	utils.WriteJSON(w, http.StatusOK, map[string]any{
		"customer": resp,
	})
}
