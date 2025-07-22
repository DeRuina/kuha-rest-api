package tietoevryapi

import (
	"fmt"
	"net/http"

	"github.com/DeRuina/KUHA-REST-API/internal/auth/authz"
	tietoevrysqlc "github.com/DeRuina/KUHA-REST-API/internal/db/tietoevry"
	"github.com/DeRuina/KUHA-REST-API/internal/store/cache"
	"github.com/DeRuina/KUHA-REST-API/internal/store/tietoevry"
	"github.com/DeRuina/KUHA-REST-API/internal/utils"
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
	Duration          int64    `json:"duration" validate:"required"`
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

// InsertExercise godoc
//
//	@Summary		Insert exercise (bulk)
//	@Description	Insert multiple exercise bundles with idempotent behavior
//	@Tags			Tietoevry - Exercise
//	@Accept			json
//	@Produce		json
//	@Param			exercise	body	swagger.TietoevryExercisesBulkInput	true	"Exercise data"
//	@Success		201			"Exercises processed successfully (idempotent operation)"
//	@Failure		400			{object}	swagger.ValidationErrorResponse
//	@Failure		403			{object}	swagger.ForbiddenResponse
//	@Failure		500			{object}	swagger.InternalServerErrorResponse
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
