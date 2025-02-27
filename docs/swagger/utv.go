package swagger

// Oura
type OuraDatesResponse struct {
	Dates []string `json:"dates" example:"2024-02-25,2024-02-26,2024-03-25"`
}

type OuraTypesResponse struct {
	Types []string `json:"types" example:"heartrate,sleep,daily_activity"`
}

type SleepData struct {
	ID                string `json:"id" example:"82496435-af96-462e-90c9-806d38418e87"`
	Day               string `json:"day" example:"2023-12-01"`
	Type              string `json:"type" example:"long_sleep"`
	AwakeTime         int    `json:"awake_time" example:"1530"`
	DeepSleepDuration int    `json:"deep_sleep_duration" example:"6630"`
	AverageHeartRate  int    `json:"average_heart_rate" example:"47"`
	BedtimeStart      string `json:"bedtime_start" example:"2023-11-30T23:23:59+02:00"`
	BedtimeEnd        string `json:"bedtime_end" example:"2023-12-01T06:21:59+02:00"`
}

type HeartRateData struct {
	BPM       int    `json:"bpm" example:"43"`
	Source    string `json:"source" example:"rest"`
	Timestamp string `json:"timestamp" example:"2023-12-01T00:00:43+00:00"`
}

type DailyActivity struct {
	Day           string `json:"day" example:"2023-12-01"`
	Steps         int    `json:"steps" example:"6653"`
	TotalCalories int    `json:"total_calories" example:"4272"`
	SedentaryTime int    `json:"sedentary_time" example:"27480"`
}

type OuraData struct {
	Sleep         []SleepData     `json:"sleep"`
	HeartRate     []HeartRateData `json:"heartrate"`
	DailyActivity []DailyActivity `json:"daily_activity"`
}

type OuraDataResponse struct {
	Data OuraData `json:"data"`
}
