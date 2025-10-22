package tietoevryapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/DeRuina/KUHA-REST-API/docs/swagger"
	"github.com/DeRuina/KUHA-REST-API/internal/auth/authz"
	tietoevrysqlc "github.com/DeRuina/KUHA-REST-API/internal/db/tietoevry"
	"github.com/DeRuina/KUHA-REST-API/internal/store/cache"
	"github.com/DeRuina/KUHA-REST-API/internal/store/tietoevry"
	"github.com/DeRuina/KUHA-REST-API/internal/utils"
	"github.com/google/uuid"
)

type TietoevryExerciseHandler struct {
	store tietoevry.Exercises
	cache *cache.Storage
}

func NewTietoevryExerciseHandler(store tietoevry.Exercises, cache *cache.Storage) *TietoevryExerciseHandler {
	return &TietoevryExerciseHandler{store: store, cache: cache}
}

// request model

type HRZoneInput struct {
	ExerciseID    string `json:"exercise_id"`
	ZoneIndex     int32  `json:"zone_index"`
	SecondsInZone int32  `json:"seconds_in_zone"`
	LowerLimit    int32  `json:"lower_limit"`
	UpperLimit    int32  `json:"upper_limit"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}

type SampleInput struct {
	ID            string    `json:"id"`
	UserID        string    `json:"user_id"`
	ExerciseID    string    `json:"exercise_id"`
	SampleType    string    `json:"sample_type"`
	RecordingRate int32     `json:"recording_rate"`
	Samples       []float64 `json:"samples"`
	Source        string    `json:"source"`
}

type SectionInput struct {
	ID          string  `json:"id"`
	UserID      string  `json:"user_id"`
	ExerciseID  string  `json:"exercise_id"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
	StartTime   string  `json:"start_time"`
	EndTime     string  `json:"end_time"`
	SectionType *string `json:"section_type"`
	Name        *string `json:"name"`
	Comment     *string `json:"comment"`
	Source      string  `json:"source"`
	RawID       *string `json:"raw_id"`
	RawData     *string `json:"raw_data"`
}

type TietoevryExerciseUpsertInput struct {
	ID                string   `json:"id" validate:"required,uuid4"`
	CreatedAt         string   `json:"created_at" validate:"required"`
	UpdatedAt         string   `json:"updated_at" validate:"required"`
	UserID            string   `json:"user_id" validate:"required,uuid4"`
	StartTime         string   `json:"start_time" validate:"required"`
	Duration          string   `json:"duration" validate:"required"`
	Comment           *string  `json:"comment"`
	SportType         *string  `json:"sport_type"`
	DetailedSportType *string  `json:"detailed_sport_type"`
	Distance          *float64 `json:"distance"`
	AvgHeartRate      *float64 `json:"avg_heart_rate"`
	MaxHeartRate      *float64 `json:"max_heart_rate"`
	Trimp             *float64 `json:"trimp"`
	SprintCount       *int32   `json:"sprint_count"`
	AvgSpeed          *float64 `json:"avg_speed"`
	MaxSpeed          *float64 `json:"max_speed"`
	Source            string   `json:"source" validate:"required"`
	Status            *string  `json:"status"`
	Calories          *int32   `json:"calories"`
	TrainingLoad      *int32   `json:"training_load"`
	RawID             *string  `json:"raw_id"`
	Feeling           *int32   `json:"feeling"`
	Recovery          *int32   `json:"recovery"`
	RPE               *int32   `json:"rpe"`
	RawData           *string  `json:"raw_data"`

	HRZones  []HRZoneInput  `json:"hr_zones"`
	Samples  []SampleInput  `json:"samples"`
	Sections []SectionInput `json:"sections"`
}

type TietoevryExercisesBulkInput struct {
	Exercises []TietoevryExerciseUpsertInput `json:"exercises" validate:"required,dive"`
}

type TietoevryExerciseParams struct {
	UserID string `form:"user_id" validate:"required,uuid4"`
}

// InsertExercise godoc
//
//	@Summary		Insert exercise (bulk)
//	@Description	Insert multiple exercise bundles for user (idempotent)
//	@Tags			Tietoevry - Exercise
//	@Accept			json
//	@Produce		json
//	@Param			exercise	body	swagger.TietoevryExercisesBulkInput	true	"Exercise data"
//	@Success		201			"Exercises processed successfully (idempotent operation)"
//	@Failure		400			{object}	swagger.ValidationErrorResponse
//	@Failure		401			{object}	swagger.UnauthorizedResponse
//	@Failure		403			{object}	swagger.ForbiddenResponse
//	@Failure		409			{object}	swagger.ConflictResponse
//	@Failure		500			{object}	swagger.InternalServerErrorResponse
//	@Failure		503			{object}	swagger.ServiceUnavailableResponse
//	@Security		BearerAuth
//	@Router			/tietoevry/exercises [post]
func (h *TietoevryExerciseHandler) InsertExercisesBulk(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}

	var input TietoevryExercisesBulkInput
	if err := utils.ReadJSON(w, r, &input); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}
	if err := utils.GetValidator().Struct(input); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	// Prevalidation
	userIDs := make([]uuid.UUID, len(input.Exercises))
	for i, exercise := range input.Exercises {
		userID, err := utils.ParseUUID(exercise.UserID)
		if err != nil {
			utils.BadRequestResponse(w, r, err)
			return
		}
		userIDs[i] = userID
	}

	if err := h.store.ValidateUsersExist(r.Context(), userIDs); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	// Convert all exercises to ExercisePayload
	exercises := make([]tietoevry.ExercisePayload, len(input.Exercises))

	for i, exercise := range input.Exercises {
		// Parse UUIDs and timestamps
		exerciseID, err := utils.ParseUUID(exercise.ID)
		if err != nil {
			utils.BadRequestResponse(w, r, err)
			return
		}
		userID, err := utils.ParseUUID(exercise.UserID)
		if err != nil {
			utils.BadRequestResponse(w, r, err)
			return
		}

		createdAt, err := utils.ParseTimestamp(exercise.CreatedAt)
		if err != nil {
			utils.BadRequestResponse(w, r, err)
			return
		}
		updatedAt, err := utils.ParseTimestamp(exercise.UpdatedAt)
		if err != nil {
			utils.BadRequestResponse(w, r, err)
			return
		}
		startTime, err := utils.ParseTimestamp(exercise.StartTime)
		if err != nil {
			utils.BadRequestResponse(w, r, err)
			return
		}

		rawData := utils.ParseRawJSON(exercise.RawData)

		arg := tietoevrysqlc.InsertExerciseParams{
			ID:                exerciseID,
			CreatedAt:         createdAt,
			UpdatedAt:         updatedAt,
			UserID:            userID,
			StartTime:         startTime,
			Duration:          exercise.Duration,
			Comment:           utils.NullStringPtr(exercise.Comment),
			SportType:         utils.NullStringPtr(exercise.SportType),
			DetailedSportType: utils.NullStringPtr(exercise.DetailedSportType),
			Distance:          utils.NullFloat64Ptr(exercise.Distance),
			AvgHeartRate:      utils.NullFloat64Ptr(exercise.AvgHeartRate),
			MaxHeartRate:      utils.NullFloat64Ptr(exercise.MaxHeartRate),
			Trimp:             utils.NullFloat64Ptr(exercise.Trimp),
			SprintCount:       utils.NullInt32Ptr(exercise.SprintCount),
			AvgSpeed:          utils.NullFloat64Ptr(exercise.AvgSpeed),
			MaxSpeed:          utils.NullFloat64Ptr(exercise.MaxSpeed),
			Source:            exercise.Source,
			Status:            utils.NullStringPtr(exercise.Status),
			Calories:          utils.NullInt32Ptr(exercise.Calories),
			TrainingLoad:      utils.NullInt32Ptr(exercise.TrainingLoad),
			RawID:             utils.NullStringPtr(exercise.RawID),
			Feeling:           utils.NullInt32Ptr(exercise.Feeling),
			Recovery:          utils.NullInt32Ptr(exercise.Recovery),
			Rpe:               utils.NullInt32Ptr(exercise.RPE),
			RawData:           rawData,
		}

		var hrZones []tietoevrysqlc.InsertExerciseHRZoneParams
		for _, z := range exercise.HRZones {
			exerciseID, _ := utils.ParseUUID(z.ExerciseID)
			createdAt, _ := utils.ParseTimestamp(z.CreatedAt)
			updatedAt, _ := utils.ParseTimestamp(z.UpdatedAt)

			hrZones = append(hrZones, tietoevrysqlc.InsertExerciseHRZoneParams{
				ExerciseID:    exerciseID,
				ZoneIndex:     z.ZoneIndex,
				SecondsInZone: z.SecondsInZone,
				LowerLimit:    z.LowerLimit,
				UpperLimit:    z.UpperLimit,
				CreatedAt:     createdAt,
				UpdatedAt:     updatedAt,
			})
		}

		var samples []tietoevrysqlc.InsertExerciseSampleParams
		for _, s := range exercise.Samples {
			id, _ := utils.ParseUUID(s.ID)
			userID, _ := utils.ParseUUID(s.UserID)
			exerciseID, _ := utils.ParseUUID(s.ExerciseID)

			samples = append(samples, tietoevrysqlc.InsertExerciseSampleParams{
				ID:            id,
				UserID:        userID,
				ExerciseID:    exerciseID,
				SampleType:    s.SampleType,
				RecordingRate: s.RecordingRate,
				Samples:       s.Samples,
				Source:        s.Source,
			})
		}

		var sections []tietoevrysqlc.InsertExerciseSectionParams
		for _, sec := range exercise.Sections {
			id, _ := utils.ParseUUID(sec.ID)
			userID, _ := utils.ParseUUID(sec.UserID)
			exerciseID, _ := utils.ParseUUID(sec.ExerciseID)
			createdAt, _ := utils.ParseTimestamp(sec.CreatedAt)
			updatedAt, _ := utils.ParseTimestamp(sec.UpdatedAt)
			startTime, _ := utils.ParseTimestamp(sec.StartTime)
			endTime, _ := utils.ParseTimestamp(sec.EndTime)
			rawData := utils.ParseRawJSON(sec.RawData)

			sections = append(sections, tietoevrysqlc.InsertExerciseSectionParams{
				ID:          id,
				UserID:      userID,
				ExerciseID:  exerciseID,
				CreatedAt:   createdAt,
				UpdatedAt:   updatedAt,
				StartTime:   startTime,
				EndTime:     endTime,
				SectionType: utils.NullStringPtr(sec.SectionType),
				Name:        utils.NullStringPtr(sec.Name),
				Comment:     utils.NullStringPtr(sec.Comment),
				Source:      sec.Source,
				RawID:       utils.NullStringPtr(sec.RawID),
				RawData:     rawData,
			})
		}

		exercises[i] = tietoevry.ExercisePayload{
			Exercise: arg,
			HRZones:  hrZones,
			Samples:  samples,
			Sections: sections,
		}
	}

	if err := h.store.InsertExercisesBulk(r.Context(), exercises); err != nil {
		utils.HandleDatabaseError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// GetExercises godoc
//
//	@Summary		Get exercises by user ID
//	@Description	Get all exercises (HR_Zones, Samples, Sections) for a specific user
//	@Tags			Tietoevry - Exercise
//	@Accept			json
//	@Produce		json
//	@Param			user_id	query		string	true	"User ID (UUID)"
//	@Success		200		{object}	swagger.TietoevryExerciseResponse
//	@Failure		400		{object}	swagger.ValidationErrorResponse
//	@Failure		401		{object}	swagger.UnauthorizedResponse
//	@Failure		403		{object}	swagger.ForbiddenResponse
//	@Failure		500		{object}	swagger.InternalServerErrorResponse
//	@Failure		503		{object}	swagger.ServiceUnavailableResponse
//	@Security		BearerAuth
//	@Router			/tietoevry/exercises [get]
func (h *TietoevryExerciseHandler) GetExercises(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}

	if err := utils.ValidateParams(r, []string{"user_id"}); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	params := TietoevryExerciseParams{
		UserID: r.URL.Query().Get("user_id"),
	}

	if err := utils.GetValidator().Struct(params); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	cacheKey := fmt.Sprintf("tietoevry:exercises:%s", params.UserID)
	if h.cache != nil {
		if cached, err := h.cache.Get(r.Context(), cacheKey); err == nil && cached != "" {
			utils.WriteJSON(w, http.StatusOK, json.RawMessage(cached))
			return
		}
	}

	userID, err := utils.ParseUUID(params.UserID)
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	exercises, err := h.store.GetExercisesByUser(r.Context(), userID)
	if err != nil {
		utils.InternalServerError(w, r, err)
		return
	}

	if len(exercises) == 0 {
		utils.WriteJSON(w, http.StatusOK, map[string]any{
			"exercises": []swagger.TietoevryExerciseUpsertInput{},
		})
		return
	}

	var output []swagger.TietoevryExerciseUpsertInput
	for _, ex := range exercises {
		hrZones, _ := h.store.GetExerciseHRZones(r.Context(), ex.ID)
		samples, _ := h.store.GetExerciseSamples(r.Context(), ex.ID)
		sections, _ := h.store.GetExerciseSections(r.Context(), ex.ID)

		out := swagger.TietoevryExerciseUpsertInput{
			ID:                ex.ID.String(),
			CreatedAt:         ex.CreatedAt.Format(time.RFC3339),
			UpdatedAt:         ex.UpdatedAt.Format(time.RFC3339),
			UserID:            ex.UserID.String(),
			StartTime:         ex.StartTime.Format(time.RFC3339),
			Duration:          ex.Duration,
			Comment:           utils.StringPtrOrNil(ex.Comment),
			SportType:         utils.StringPtrOrNil(ex.SportType),
			DetailedSportType: utils.StringPtrOrNil(ex.DetailedSportType),
			Distance:          utils.Float64PtrOrNil(ex.Distance),
			AvgHeartRate:      utils.Float64PtrOrNil(ex.AvgHeartRate),
			MaxHeartRate:      utils.Float64PtrOrNil(ex.MaxHeartRate),
			Trimp:             utils.Float64PtrOrNil(ex.Trimp),
			SprintCount:       utils.Int32PtrOrNil(ex.SprintCount),
			AvgSpeed:          utils.Float64PtrOrNil(ex.AvgSpeed),
			MaxSpeed:          utils.Float64PtrOrNil(ex.MaxSpeed),
			Source:            ex.Source,
			Status:            utils.StringPtrOrNil(ex.Status),
			Calories:          utils.Int32PtrOrNil(ex.Calories),
			TrainingLoad:      utils.Int32PtrOrNil(ex.TrainingLoad),
			RawID:             utils.StringPtrOrNil(ex.RawID),
			Feeling:           utils.Int32PtrOrNil(ex.Feeling),
			Recovery:          utils.Int32PtrOrNil(ex.Recovery),
			RPE:               utils.Int32PtrOrNil(ex.Rpe),
			RawData:           utils.RawMessagePtrOrNil(ex.RawData),
		}

		// HR Zones
		for _, z := range hrZones {
			out.HRZones = append(out.HRZones, swagger.HRZone{
				ExerciseID:    z.ExerciseID.String(),
				ZoneIndex:     z.ZoneIndex,
				SecondsInZone: z.SecondsInZone,
				LowerLimit:    z.LowerLimit,
				UpperLimit:    z.UpperLimit,
				CreatedAt:     z.CreatedAt.Format(time.RFC3339),
				UpdatedAt:     z.UpdatedAt.Format(time.RFC3339),
			})
		}

		// Samples
		for _, s := range samples {
			out.Samples = append(out.Samples, swagger.Sample{
				ID:            s.ID.String(),
				UserID:        s.UserID.String(),
				ExerciseID:    s.ExerciseID.String(),
				SampleType:    s.SampleType,
				RecordingRate: s.RecordingRate,
				Samples:       s.Samples,
				Source:        s.Source,
			})
		}

		// Sections
		for _, sec := range sections {
			out.Sections = append(out.Sections, swagger.Section{
				ID:          sec.ID.String(),
				UserID:      sec.UserID.String(),
				ExerciseID:  sec.ExerciseID.String(),
				CreatedAt:   sec.CreatedAt.Format(time.RFC3339),
				UpdatedAt:   sec.UpdatedAt.Format(time.RFC3339),
				StartTime:   sec.StartTime.Format(time.RFC3339),
				EndTime:     sec.EndTime.Format(time.RFC3339),
				SectionType: utils.StringPtrOrNil(sec.SectionType),
				Name:        utils.StringPtrOrNil(sec.Name),
				Comment:     utils.StringPtrOrNil(sec.Comment),
				Source:      sec.Source,
				RawID:       utils.StringPtrOrNil(sec.RawID),
				RawData:     utils.RawMessagePtrOrNil(sec.RawData),
			})
		}

		output = append(output, out)
	}

	resp := map[string]any{"exercises": output}
	cache.SetCacheJSON(r.Context(), h.cache, cacheKey, resp, 3*time.Minute)
	utils.WriteJSON(w, http.StatusOK, resp)
}
