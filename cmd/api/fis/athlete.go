package fisapi

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/DeRuina/KUHA-REST-API/internal/auth/authz"
	"github.com/DeRuina/KUHA-REST-API/internal/store/cache"
	"github.com/DeRuina/KUHA-REST-API/internal/store/fis"
	"github.com/DeRuina/KUHA-REST-API/internal/utils"
)

type AthleteHandler struct {
	store fis.Athlete
	cache *cache.Storage
}

func NewAthleteHandler(store fis.Athlete, cache *cache.Storage) *AthleteHandler {
	return &AthleteHandler{store: store, cache: cache}
}

// GetAthletesBySporttiID godoc
//
//	@Summary	Get all athletes for a given SporttiID
//	@Tags		FIS - Athlete
//	@Accept		json
//	@Produce	json
//	@Param		sporttiid	query		integer	true	"SporttiID"
//	@Success	200			{object}	swagger.FISAthletesResponse
//	@Failure	400			{object}	swagger.ValidationErrorResponse
//	@Failure	401			{object}	swagger.UnauthorizedResponse
//	@Failure	403			{object}	swagger.ForbiddenResponse
//	@Failure	404			{object}	swagger.NotFoundResponse
//	@Failure	500			{object}	swagger.InternalServerErrorResponse
//	@Failure	503			{object}	swagger.ServiceUnavailableResponse
//	@Security	BearerAuth
//	@Router		/fis/fiscode [get]
func (h *AthleteHandler) GetAthletesBySporttiID(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}
	if err := utils.ValidateParams(r, []string{"sporttiid"}); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	params := SporttiIDParam{
		SporttiID: r.URL.Query().Get("sporttiid"),
	}
	if err := utils.GetValidator().Struct(params); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	sporttiID, err := utils.ParsePositiveInt32(params.SporttiID)
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	rows, err := h.store.GetAthletesBySporttiID(r.Context(), sporttiID)
	if err != nil {
		utils.InternalServerError(w, r, err)
		return
	}

	if len(rows) == 0 {
		utils.NotFoundResponse(w, r,
			fmt.Errorf("no athletes found for SporttiID %d", sporttiID),
		)
		return
	}

	body := map[string]any{"athletes": rows}
	utils.WriteJSON(w, http.StatusOK, body)
}

// InsertAthlete godoc
//
//	@Summary		Add new athlete
//	@Description	Inserts a new athlete into athlete table
//	@Tags			FIS - Athlete Management
//	@Accept			json
//	@Produce		json
//	@Param			athlete	body	swagger.FISInsertAthleteExample	true	"Athlete payload"
//	@Success		201		"Created"
//	@Failure		400		{object}	swagger.ValidationErrorResponse
//	@Failure		401		{object}	swagger.UnauthorizedResponse
//	@Failure		403		{object}	swagger.ForbiddenResponse
//	@Failure		409		{object}	swagger.ConflictResponse
//	@Failure		500		{object}	swagger.InternalServerErrorResponse
//	@Failure		503		{object}	swagger.ServiceUnavailableResponse
//	@Security		BearerAuth
//	@Router			/fis/athlete [post]
func (h *AthleteHandler) InsertAthlete(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}

	var in InsertAthleteInput
	if err := utils.ReadJSON(w, r, &in); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}
	if err := utils.GetValidator().Struct(in); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	clean := mapInsertAthleteInput(in)
	if err := h.store.InsertAthlete(r.Context(), clean); err != nil {
		utils.HandleDatabaseError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// UpdateAthlete godoc
//
//	@Summary		Update athlete by fiscode
//	@Description	Updates an existing athlete in athlete table
//	@Tags			FIS - Athlete Management
//	@Accept			json
//	@Produce		json
//	@Param			athlete	body	swagger.FISUpdateAthleteExample	true	"Athlete payload"
//	@Success		200		"Updated"
//	@Failure		400		{object}	swagger.ValidationErrorResponse
//	@Failure		401		{object}	swagger.UnauthorizedResponse
//	@Failure		403		{object}	swagger.ForbiddenResponse
//	@Failure		404		{object}	swagger.NotFoundResponse
//	@Failure		500		{object}	swagger.InternalServerErrorResponse
//	@Failure		503		{object}	swagger.ServiceUnavailableResponse
//	@Security		BearerAuth
//	@Router			/fis/athlete [put]
func (h *AthleteHandler) UpdateAthlete(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}

	var in UpdateAthleteInput
	if err := utils.ReadJSON(w, r, &in); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}
	if err := utils.GetValidator().Struct(in); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	clean := mapUpdateAthleteInput(in)
	if err := h.store.UpdateAthleteByFiscode(r.Context(), clean); err != nil {
		if err == sql.ErrNoRows {
			utils.NotFoundResponse(w, r, err)
			return
		}
		utils.HandleDatabaseError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// DeleteAthlete godoc
//
//	@Summary		Delete athlete
//	@Description	Deletes an athlete by FIS code
//	@Tags			FIS - Athlete Management
//	@Accept			json
//	@Produce		json
//	@Param			fiscode	query	integer	true	"FIS code"
//	@Success		200		"Deleted"
//	@Failure		400		{object}	swagger.ValidationErrorResponse
//	@Failure		401		{object}	swagger.UnauthorizedResponse
//	@Failure		403		{object}	swagger.ForbiddenResponse
//	@Failure		404		{object}	swagger.NotFoundResponse
//	@Failure		500		{object}	swagger.InternalServerErrorResponse
//	@Failure		503		{object}	swagger.ServiceUnavailableResponse
//	@Security		BearerAuth
//	@Router			/fis/athlete [delete]
func (h *AthleteHandler) DeleteAthlete(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}
	if err := utils.ValidateParams(r, []string{"fiscode"}); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	params := AthleteFiscodeParam{
		Fiscode: r.URL.Query().Get("fiscode"),
	}
	if err := utils.GetValidator().Struct(params); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	fiscode, err := utils.ParsePositiveInt32(params.Fiscode)
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	if err := h.store.DeleteAthleteByFiscode(r.Context(), fiscode); err != nil {
		if err == sql.ErrNoRows {
			utils.NotFoundResponse(w, r, err)
			return
		}
		utils.HandleDatabaseError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
