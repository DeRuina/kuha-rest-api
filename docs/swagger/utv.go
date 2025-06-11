package swagger

// Oura
type DatesResponse struct {
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

// Polar

type PolarTypesResponse struct {
	Types []string `json:"types" example:"alertness,circadian_bedtime,heart_rate"`
}

type PolarAlertnessData struct {
	Grade     float64 `json:"grade" example:"8.5"`
	Validity  string  `json:"validity" example:"VALIDITY_VALID"`
	GradeType string  `json:"grade_type" example:"GRADE_TYPE_PRIMARY"`
}

type PolarHeartRateSample struct {
	HeartRate  int    `json:"heart_rate" example:"72"`
	SampleTime string `json:"sample_time" example:"00:10:00"`
}

type PolarHeartRateData struct {
	Date          string                 `json:"date" example:"2023-12-01"`
	HeartRateData []PolarHeartRateSample `json:"heart_rate_samples"`
}

type PolarData struct {
	Alertness PolarAlertnessData `json:"alertness"`
	HeartRate PolarHeartRateData `json:"heart_rate"`
}

type PolarDataResponse struct {
	Data PolarData `json:"data"`
}

// Suunto

type SuuntoTypesResponse struct {
	Types []string `json:"types" example:"workout"`
}

type SuuntoWorkoutSummary struct {
	AvgPace   float64 `json:"avgPace" example:"4.50"`
	AvgSpeed  float64 `json:"avgSpeed" example:"3.80"`
	StepCount int     `json:"stepCount" example:"5000"`
	TotalTime float64 `json:"totalTime" example:"3600"`
	StartTime int64   `json:"startTime" example:"1710065996000"`
}

type SuuntoWorkoutData struct {
	WorkoutSummary SuuntoWorkoutSummary `json:"workout_summary"`
}

type SuuntoData struct {
	Workout map[string]SuuntoWorkoutData `json:"workout"`
}

type SuuntoDataResponse struct {
	Data SuuntoData `json:"data"`
}

// Garmin
type GarminTypesResponse struct {
	Types []string `json:"types" example:"epochs, dailies, hrv"`
}

type GarminEpoch struct {
	MET                   float64 `json:"met" example:"1.0"`
	Steps                 int     `json:"steps" example:"0"`
	Intensity             string  `json:"intensity" example:"SEDENTARY"`
	SummaryID             string  `json:"summaryId" example:"x32c7d85-67b67080-9"`
	ActivityType          string  `json:"activityType" example:"UNMONITORED"`
	DistanceInMeters      float64 `json:"distanceInMeters" example:"0.0"`
	DurationInSeconds     int     `json:"durationInSeconds" example:"900"`
	ActiveKilocalories    int     `json:"activeKilocalories" example:"0"`
	MaxMotionIntensity    float64 `json:"maxMotionIntensity" example:"0.0"`
	StartTimeInSeconds    int64   `json:"startTimeInSeconds" example:"1740009600"`
	ActiveTimeInSeconds   int     `json:"activeTimeInSeconds" example:"0"`
	MeanMotionIntensity   float64 `json:"meanMotionIntensity" example:"0.0"`
	StartTimeOffsetInSecs int     `json:"startTimeOffsetInSeconds" example:"7200"`
}

type GarminDaily struct {
	Steps                              int     `json:"steps" example:"0"`
	StepsGoal                          int     `json:"stepsGoal" example:"10000"`
	SummaryID                          string  `json:"summaryId" example:"x32c7d85-67b65460-128f4-0"`
	ActivityType                       string  `json:"activityType" example:"GENERIC"`
	CalendarDate                       string  `json:"calendarDate" example:"2025-02-20"`
	BMRKilocalories                    int     `json:"bmrKilocalories" example:"1967"`
	DistanceInMeters                   float64 `json:"distanceInMeters" example:"0.0"`
	DurationInSeconds                  int     `json:"durationInSeconds" example:"76020"`
	FloorsClimbedGoal                  int     `json:"floorsClimbedGoal" example:"10"`
	ActiveKilocalories                 int     `json:"activeKilocalories" example:"0"`
	StartTimeInSeconds                 int64   `json:"startTimeInSeconds" example:"1740002400"`
	ActiveTimeInSeconds                int     `json:"activeTimeInSeconds" example:"0"`
	StartTimeOffsetInSeconds           int     `json:"startTimeOffsetInSeconds" example:"7200"`
	IntensityDurationGoalInSeconds     int     `json:"intensityDurationGoalInSeconds" example:"36000"`
	ModerateIntensityDurationInSeconds int     `json:"moderateIntensityDurationInSeconds" example:"0"`
	VigorousIntensityDurationInSeconds int     `json:"vigorousIntensityDurationInSeconds" example:"0"`
}

type GarminData struct {
	Epochs  GarminEpoch `json:"epochs"`
	Dailies GarminDaily `json:"dailies"`
}

type GarminDataResponse struct {
	Data GarminData `json:"data"`
}
