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

// InsertExercise godoc
//
//	@Summary		Insert exercise
//	@Description	Insert a new exercise bundle (main + zones + samples + sections)
//	@Tags			Tietoevry - Exercise
//	@Accept			json
//	@Produce		json
//	@Param			exercise	body	swagger.TietoevryExerciseUpsertInput	true	"Exercise data"
//	@Success		201			"created"
//	@Failure		400			{object}	swagger.ValidationErrorResponse
//	@Failure		403			{object}	swagger.ForbiddenResponse
//	@Failure		500			{object}	swagger.InternalServerErrorResponse
//	@Security		BearerAuth
//	@Router			/tietoevry/exercises [post]
func (h *TietoevryExerciseHandler) UpsertExercise(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}

	var input TietoevryExerciseUpsertInput
	if err := utils.ReadJSON(w, r, &input); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}
	if err := utils.GetValidator().Struct(input); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	// Parse UUIDs and timestamps
	exerciseID, err := utils.ParseUUID(input.ID)
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}
	userID, err := utils.ParseUUID(input.UserID)
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	createdAt, err := utils.ParseTimestamp(input.CreatedAt)
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}
	updatedAt, err := utils.ParseTimestamp(input.UpdatedAt)
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}
	startTime, err := utils.ParseTimestamp(input.StartTime)
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	rawData := utils.ParseRawJSON(input.RawData)

	arg := tietoevrysqlc.InsertExerciseParams{
		ID:                exerciseID,
		CreatedAt:         createdAt,
		UpdatedAt:         updatedAt,
		UserID:            userID,
		StartTime:         startTime,
		Duration:          input.Duration,
		Comment:           utils.NullStringPtr(input.Comment),
		SportType:         utils.NullStringPtr(input.SportType),
		DetailedSportType: utils.NullStringPtr(input.DetailedSportType),
		Distance:          utils.NullFloat64Ptr(input.Distance),
		AvgHeartRate:      utils.NullFloat64Ptr(input.AvgHeartRate),
		MaxHeartRate:      utils.NullFloat64Ptr(input.MaxHeartRate),
		Trimp:             utils.NullFloat64Ptr(input.Trimp),
		SprintCount:       utils.NullInt32Ptr(input.SprintCount),
		AvgSpeed:          utils.NullFloat64Ptr(input.AvgSpeed),
		MaxSpeed:          utils.NullFloat64Ptr(input.MaxSpeed),
		Source:            input.Source,
		Status:            utils.NullStringPtr(input.Status),
		Calories:          utils.NullInt32Ptr(input.Calories),
		TrainingLoad:      utils.NullInt32Ptr(input.TrainingLoad),
		RawID:             utils.NullStringPtr(input.RawID),
		Feeling:           utils.NullInt32Ptr(input.Feeling),
		Recovery:          utils.NullInt32Ptr(input.Recovery),
		Rpe:               utils.NullInt32Ptr(input.RPE),
		RawData:           rawData,
	}

	var hrZones []tietoevrysqlc.InsertExerciseHRZoneParams
	for _, z := range input.HRZones {
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
	for _, s := range input.Samples {
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
	for _, sec := range input.Sections {
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

	payload := tietoevry.ExercisePayload{
		Exercise: arg,
		HRZones:  hrZones,
		Samples:  samples,
		Sections: sections,
	}

	if err := h.store.InsertExerciseBundle(r.Context(), payload); err != nil {
		utils.HandleDatabaseError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
