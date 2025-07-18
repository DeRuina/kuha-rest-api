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
