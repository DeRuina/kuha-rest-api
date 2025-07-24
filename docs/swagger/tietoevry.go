package swagger

// TietoevryUserUpsertInput represents the POST body schema for upserting a user
type TietoevryUserUpsertInput struct {
	ID                        string   `json:"id" example:"7cffe6e0-3f28-43b6-b511-d836d3a9f7b5"`
	SporttiID                 int32    `json:"sportti_id" example:"12345"`
	ProfileGender             *string  `json:"profile_gender" example:"male"`
	ProfileBirthdate          *string  `json:"profile_birthdate" example:"1990-08-15"` // ISO 8601
	ProfileWeight             *float64 `json:"profile_weight" example:"72.5"`
	ProfileHeight             *float64 `json:"profile_height" example:"178.2"`
	ProfileRestingHeartRate   *int32   `json:"profile_resting_heart_rate" example:"60"`
	ProfileMaximumHeartRate   *int32   `json:"profile_maximum_heart_rate" example:"190"`
	ProfileAerobicThreshold   *int32   `json:"profile_aerobic_threshold" example:"140"`
	ProfileAnaerobicThreshold *int32   `json:"profile_anaerobic_threshold" example:"165"`
	ProfileVo2max             *int32   `json:"profile_vo2max" example:"50"`
}

type HRZone struct {
	ExerciseID    string `json:"exercise_id" example:"2d4f6aee-b62c-408e-85e1-07bd78f383a7"`
	ZoneIndex     int32  `json:"zone_index" example:"2"`
	SecondsInZone int32  `json:"seconds_in_zone" example:"90"`
	LowerLimit    int32  `json:"lower_limit" example:"120"`
	UpperLimit    int32  `json:"upper_limit" example:"140"`
	CreatedAt     string `json:"created_at" example:"2024-07-21T08:00:00Z"`
	UpdatedAt     string `json:"updated_at" example:"2024-07-21T08:00:00Z"`
}

type Sample struct {
	ID            string    `json:"id" example:"3f28c7a1-2ea3-438c-9b35-099c4372da49"`
	UserID        string    `json:"user_id" example:"7cffe6e0-3f28-43b6-b511-d836d3a9f7b5"`
	ExerciseID    string    `json:"exercise_id" example:"2d4f6aee-b62c-408e-85e1-07bd78f383a7"`
	SampleType    string    `json:"sample_type" example:"heart_rate"`
	RecordingRate int32     `json:"recording_rate" example:"1"`
	Samples       []float64 `json:"samples"`
	Source        string    `json:"source" example:"garmin"`
}

type Section struct {
	ID          string `json:"id" example:"1a2b3c4d-0000-0000-0000-111122223333"`
	UserID      string `json:"user_id" example:"7cffe6e0-3f28-43b6-b511-d836d3a9f7b5"`
	ExerciseID  string `json:"exercise_id" example:"2d4f6aee-b62c-408e-85e1-07bd78f383a7"`
	CreatedAt   string `json:"created_at" example:"2024-07-21T08:00:00Z"`
	UpdatedAt   string `json:"updated_at" example:"2024-07-21T08:05:00Z"`
	StartTime   string `json:"start_time" example:"2024-07-21T08:10:00Z"`
	EndTime     string `json:"end_time" example:"2024-07-21T08:30:00Z"`
	SectionType string `json:"section_type" example:"warmup"`
	Name        string `json:"name" example:"Warm-up"`
	Comment     string `json:"comment" example:"Felt great"`
	Source      string `json:"source" example:"polar"`
	RawID       string `json:"raw_id" example:"abc-123"`
	RawData     string `json:"raw_data" example:"{\"extra\":\"data\"}"`
}

type TietoevryExerciseUpsertInput struct {
	ID                string   `json:"id" example:"2d4f6aee-b62c-408e-85e1-07bd78f383a7"`
	CreatedAt         string   `json:"created_at" example:"2024-07-21T08:00:00Z"`
	UpdatedAt         string   `json:"updated_at" example:"2024-07-21T08:00:00Z"`
	UserID            string   `json:"user_id" example:"7cffe6e0-3f28-43b6-b511-d836d3a9f7b5"`
	StartTime         string   `json:"start_time" example:"2024-07-21T08:10:00Z"`
	Duration          int64    `json:"duration" example:"3600"`
	Comment           *string  `json:"comment" example:"Morning run"`
	SportType         *string  `json:"sport_type" example:"running"`
	DetailedSportType *string  `json:"detailed_sport_type" example:"trail"`
	Distance          *float64 `json:"distance" example:"8.5"`
	AvgHeartRate      *float64 `json:"avg_heart_rate" example:"145.2"`
	MaxHeartRate      *float64 `json:"max_heart_rate" example:"165.3"`
	Trimp             *float64 `json:"trimp" example:"100.5"`
	SprintCount       *int32   `json:"sprint_count" example:"2"`
	AvgSpeed          *float64 `json:"avg_speed" example:"3.2"`
	MaxSpeed          *float64 `json:"max_speed" example:"4.8"`
	Source            string   `json:"source" example:"garmin"`
	Status            *string  `json:"status" example:"completed"`
	Calories          *int32   `json:"calories" example:"450"`
	TrainingLoad      *int32   `json:"training_load" example:"55"`
	RawID             *string  `json:"raw_id" example:"xyz789"`
	Feeling           *int32   `json:"feeling" example:"4"`
	Recovery          *int32   `json:"recovery" example:"3"`
	RPE               *int32   `json:"rpe" example:"7"`
	RawData           *string  `json:"raw_data" example:"{\"power\":300}"`

	HRZones  []HRZone  `json:"hr_zones"`
	Samples  []Sample  `json:"samples"`
	Sections []Section `json:"sections"`
}

type TietoevrySymptomInput struct {
	ID             string  `json:"id" validate:"required,uuid4" example:"2d4f6aee-b62c-408e-85e1-07bd78f383a7"`
	UserID         string  `json:"user_id" validate:"required,uuid4" example:"7cffe6e0-3f28-43b6-b511-d836d3a9f7b5"`
	Date           string  `json:"date" validate:"required" example:"2024-01-15"`
	Symptom        string  `json:"symptom" validate:"required" example:"knee pain"`
	Severity       int32   `json:"severity" validate:"required" example:"7"`
	Comment        *string `json:"comment" example:"Pain started after morning run"`
	Source         string  `json:"source" validate:"required" example:"polar"`
	CreatedAt      string  `json:"created_at" validate:"required" example:"2024-01-15T10:30:00Z"`
	UpdatedAt      string  `json:"updated_at" validate:"required" example:"2024-01-15T10:30:00Z"`
	RawID          *string `json:"raw_id" example:"raw_symptom_123"`
	OriginalID     *string `json:"original_id" example:""`
	Recovered      *bool   `json:"recovered" example:"true"`
	PainIndex      *int32  `json:"pain_index" example:"8"`
	Side           *string `json:"side" example:"left"`
	Category       *string `json:"category" example:"joint"`
	AdditionalData *string `json:"additional_data" example:"{\"intensity\": \"moderate\", \"duration\": \"30min\"}"`
}

// Bulk input types
type TietoevryExercisesBulkInput struct {
	Exercises []TietoevryExerciseUpsertInput `json:"exercises" validate:"required,dive"`
}

type TietoevrySymptomsBulkInput struct {
	Symptoms []TietoevrySymptomInput `json:"symptoms" validate:"required,dive"`
}

type TietoevryMeasurementInput struct {
	ID             string   `json:"id" example:"b872a58e-1234-4baf-8c19-5f2ad34d9ef1"`
	CreatedAt      string   `json:"created_at" example:"2024-07-23T08:00:00Z"`
	UpdatedAt      string   `json:"updated_at" example:"2024-07-23T08:10:00Z"`
	UserID         string   `json:"user_id" example:"7cffe6e0-3f28-43b6-b511-d836d3a9f7b5"`
	Date           string   `json:"date" example:"2024-07-22"`
	Name           string   `json:"name" example:"weight"`
	NameType       string   `json:"name_type" example:"numeric"`
	Source         string   `json:"source" example:"oura"`
	Value          string   `json:"value" example:"72.5"`
	ValueNumeric   *float64 `json:"value_numeric" example:"72.5"`
	Comment        *string  `json:"comment" example:"Post training"`
	RawID          *string  `json:"raw_id" example:"ouraid-xyz-123"`
	RawData        *string  `json:"raw_data" example:"{\"unit\": \"kg\"}"`
	AdditionalInfo *string  `json:"additional_info" example:"{\"sensor_accuracy\": \"high\"}"`
}

type TietoevryMeasurementsBulkInput struct {
	Measurements []TietoevryMeasurementInput `json:"measurements" validate:"required,dive"`
}

type TietoevryTestResultInput struct {
	ID                          string  `json:"id" example:"d19f2c63-fc5a-4aeb-8b90-5fc7df5d1c0c"`
	UserID                      string  `json:"user_id" example:"7cffe6e0-3f28-43b6-b511-d836d3a9f7b5"`
	TypeID                      string  `json:"type_id" example:"23d9453b-8d43-46f5-bb27-1db29a88456a"`
	TypeType                    *string `json:"type_type" example:"threshold"`
	TypeResultType              string  `json:"type_result_type" example:"float"`
	TypeName                    *string `json:"type_name" example:"Aerobic Threshold Test"`
	Timestamp                   string  `json:"timestamp" example:"2024-07-23T08:00:00Z"`
	Name                        *string `json:"name" example:"Threshold Test 1"`
	Comment                     *string `json:"comment" example:"User reached fatigue quickly"`
	Data                        string  `json:"data" example:"{\"vo2\": 45.5, \"lactate\": 2.1}"`
	CreatedAt                   string  `json:"created_at" example:"2024-07-23T08:10:00Z"`
	UpdatedAt                   string  `json:"updated_at" example:"2024-07-23T08:10:01Z"`
	TestEventID                 *string `json:"test_event_id" example:"8abf0123-4567-89ab-cdef-1234567890ab"`
	TestEventName               *string `json:"test_event_name" example:"Summer Test Event"`
	TestEventDate               *string `json:"test_event_date" example:"2024-07-22"`
	TestEventTemplateTestID     *string `json:"test_event_template_test_id" example:"4b9f2cde-8aaa-4e01-8103-8d1cabc712aa"`
	TestEventTemplateTestName   *string `json:"test_event_template_test_name" example:"VO2 Max Lab Test"`
	TestEventTemplateTestLimits *string `json:"test_event_template_test_limits" example:"{\"max\": 70, \"min\": 45}"`
}

type TietoevryTestResultsBulkInput struct {
	TestResults []TietoevryTestResultInput `json:"test_results" validate:"required,dive"`
}

type TietoevryQuestionnaireAnswerInput struct {
	UserID                  string  `json:"user_id" example:"7cffe6e0-3f28-43b6-b511-d836d3a9f7b5"`
	QuestionnaireInstanceID string  `json:"questionnaire_instance_id" example:"8c94dfeb-8f3e-4e61-9a66-64559eb354d2"`
	QuestionnaireNameFi     *string `json:"questionnaire_name_fi" example:"Unikysely"`
	QuestionnaireNameEn     *string `json:"questionnaire_name_en" example:"Sleep Survey"`
	QuestionnaireKey        string  `json:"questionnaire_key" example:"sleep_quality"`
	QuestionID              string  `json:"question_id" example:"123e4567-e89b-4d3a-a456-426614174000"`
	QuestionLabelFi         *string `json:"question_label_fi" example:"Kuinka hyvin nukuit?"`
	QuestionLabelEn         *string `json:"question_label_en" example:"How well did you sleep?"`
	QuestionType            string  `json:"question_type" example:"scale"`
	OptionID                *string `json:"option_id" example:"c3e8d954-bb0a-4d77-9885-f9e7cabc1234"`
	OptionValue             *int32  `json:"option_value" example:"4"`
	OptionLabelFi           *string `json:"option_label_fi" example:"Hyvin"`
	OptionLabelEn           *string `json:"option_label_en" example:"Well"`
	FreeText                *string `json:"free_text" example:"Had vivid dreams"`
	CreatedAt               string  `json:"created_at" example:"2024-07-23T09:00:00Z"`
	UpdatedAt               string  `json:"updated_at" example:"2024-07-23T09:00:10Z"`
	Value                   *string `json:"value" example:"{\"scale\": 4, \"note\": \"slightly tired\"}"`
}

type TietoevryQuestionnaireAnswersBulkInput struct {
	Questionnaires []TietoevryQuestionnaireAnswerInput `json:"questionnaires" validate:"required,dive"`
}
