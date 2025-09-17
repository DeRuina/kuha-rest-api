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

type GarminPostDataInput struct {
	UserID string               `json:"user_id" example:"f47ac10b-58cc-4372-a567-0e02b2c3d479"`
	Date   string               `json:"date" example:"2025-06-12"`
	Data   GarminPayloadExample `json:"data"`
}

type GarminPayloadExample struct {
	Dailies []GarminDailyExample `json:"dailies"`
}

type GarminDailyExample struct {
	Steps        int    `json:"steps" example:"5000"`
	CalendarDate string `json:"calendarDate" example:"2025-06-12"`
}

type OuraPostDataInput struct {
	UserID string             `json:"user_id" example:"f47ac10b-58cc-4372-a567-0e02b2c3d479"`
	Date   string             `json:"date" example:"2025-06-12"`
	Data   OuraPayloadExample `json:"data"`
}

type OuraPayloadExample struct {
	DailyActivity []OuraDailyExample `json:"daily_activity"`
}

type OuraDailyExample struct {
	ID    string `json:"id" example:"232b8559-9175-43f5-88c2-7f727b1f3ced"`
	Day   string `json:"day" example:"2021-12-10"`
	Steps int    `json:"steps" example:"1074"`
	Score int    `json:"score" example:"93"`
}

type SuuntoPostDataInput struct {
	UserID string               `json:"user_id" example:"f47ac10b-58cc-4372-a567-0e02b2c3d479"`
	Date   string               `json:"date" example:"2025-06-12"`
	Data   SuuntoPayloadExample `json:"data"`
}

type SuuntoPayloadExample struct {
	Workout map[string]SuuntoWorkout `json:"workout"`
}

type SuuntoWorkout struct {
	WorkoutContent string                `json:"workout_content" example:"65ed97d4dddc530a5389f8f9"`
	WorkoutSummary SuuntoWorkoutSummary1 `json:"workout_summary"`
}

type SuuntoWorkoutSummary1 struct {
	Steps         int     `json:"stepCount" example:"4469"`
	TotalTime     float64 `json:"totalTime" example:"3707.57"`
	TotalDistance float64 `json:"totalDistance" example:"14927.06"`
}

type PolarPostDataInput struct {
	UserID string              `json:"user_id" example:"f47ac10b-58cc-4372-a567-0e02b2c3d479"`
	Date   string              `json:"date" example:"2025-06-12"`
	Data   PolarPayloadExample `json:"data"`
}

type PolarPayloadExample struct {
	Alertness PolarAlertness `json:"alertness"`
}

type PolarAlertness struct {
	Grade      float64 `json:"grade" example:"5.1"`
	GradeType  string  `json:"grade_type" example:"GRADE_TYPE_PRIMARY"`
	SleepType  string  `json:"sleep_type" example:"SLEEP_TYPE_PRIMARY"`
	ResultType string  `json:"result_type" example:"ALERTNESS_TYPE_HISTORY"`
}

type SleepSample struct {
	ID   string `json:"id" example:"123"`
	Day  string `json:"day" example:"2025-06-12"`
	Type string `json:"type" example:"long_sleep"`
}

type DeviceData struct {
	Sleep []SleepSample `json:"sleep"`
}

type LatestDataResponse struct {
	Device string     `json:"device" example:"garmin"`
	Date   string     `json:"date" example:"2025-06-12"`
	Data   DeviceData `json:"data"`
}

type PolarTokenDetails struct {
	XUserID         int    `json:"x_user_id" example:"45071318"`
	AccessToken     string `json:"access_token" example:"abc123token"`
	ExpiresIn       int    `json:"expires_in" example:"315359999"`
	MemberID        string `json:"member_id" example:"7112b057-2b8d-47ab-a8f1-140dc09664cf"`
	TokenRefreshed  string `json:"token_last_refreshed" example:"2025-03-03 10:53"`
	DataLastFetched string `json:"data_last_fetched" example:"2025-06-18 10:54"`
}

type PolarTokenInput struct {
	UserID string            `json:"user_id" example:"208e2ffb-ac68-4980-a8a6-b7e0136e0798"`
	Data   PolarTokenDetails `json:"data"`
}

type PolarStatusResponse struct {
	Connected bool `json:"connected" example:"true"`
	Data      bool `json:"data" example:"true"`
}

type PolarTokenByIDResponse struct {
	UserID string            `json:"user_id" example:"208e2ffb-ac68-4980-a8a6-b7e0136e0798"`
	Data   PolarTokenDetails `json:"data"`
}

type OuraStatusResponse struct {
	Connected bool `json:"connected" example:"true"`
	Data      bool `json:"data" example:"true"`
}

type OuraPersonalInfo struct {
	ID            string  `json:"id" example:"a3f4d23bc10e4911b2878d23f09aa320"`
	Age           int     `json:"age" example:"38"`
	Email         *string `json:"email" example:"user42@example.com"`
	Height        float64 `json:"height" example:"1.75"`
	Weight        float64 `json:"weight" example:"76.5"`
	BiologicalSex string  `json:"biological_sex" example:"female"`
}

type OuraTokenDetails struct {
	AccessToken     string           `json:"access_token" example:"XYZ123ABC789TOKEN456"`
	RefreshToken    string           `json:"refresh_token" example:"REFRESH987XYZ654TOKEN"`
	PersonalInfo    OuraPersonalInfo `json:"personal_info"`
	TokenRefreshed  string           `json:"token_last_refreshed" example:"2025-07-10 12:30"`
	DataLastFetched string           `json:"data_last_fetched" example:"2025-07-11 08:45"`
}

type OuraTokenInput struct {
	UserID  string           `json:"user_id" example:"d5e1c624-2f39-4410-b9d6-842fc3226b7e"`
	Details OuraTokenDetails `json:"details"`
}

type OuraTokenByIDResponse struct {
	UserID string           `json:"user_id" example:"d5e1c624-2f39-4410-b9d6-842fc3226b7e"`
	Data   OuraTokenDetails `json:"data"`
}

type SuuntoStatusResponse struct {
	Connected bool `json:"connected" example:"true"`
	Data      bool `json:"data" example:"true"`
}

type SuuntoTokenInput struct {
	UserID  string             `json:"user_id" example:"7cffe6e0-3f28-43b6-b511-d836d3a9f7b5"`
	Details SuuntoTokenDetails `json:"details"`
}

type SuuntoTokenDetails struct {
	AccessToken     string `json:"access_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.random123tokenxyz"`
	RefreshToken    string `json:"refresh_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.refresh456tokenabc"`
	UK              string `json:"uk" example:"2x9+gFLEyAxpb6WyYqNMcqzwZKjVcqrnt/RT3gg9LPhh18ZxKpy7f5MDELzUR6FBMz4XZ1TAp7vDYx8vdwZaBw=="`
	JTI             string `json:"jti" example:"vghrTs9QdFrMB8EVTNNVoRTLkxu"`
	UKV             string `json:"ukv" example:"2"`
	User            string `json:"user" example:"randomuser88"`
	Scope           string `json:"scope" example:"activity"`
	ExpiresIn       int    `json:"expires_in" example:"86400"`
	TokenType       string `json:"token_type" example:"bearer"`
	TokenRefreshed  string `json:"token_last_refreshed" example:"2025-07-10 05:23"`
	DataLastFetched string `json:"data_last_fetched" example:"2025-07-11 09:45"`
}

type SuuntoTokenByUsernameResponse struct {
	UserID string             `json:"user_id" example:"7cffe6e0-3f28-43b6-b511-d836d3a9f7b5"`
	Data   SuuntoTokenDetails `json:"data"`
}

type GarminStatusResponse struct {
	Connected bool `json:"connected" example:"true"`
	Data      bool `json:"data" example:"true"`
}

type GarminTokenDetails struct {
	AccessToken       string `json:"access_token" example:"a1f2b3c4-d5e6-7890-1234-56789abcdef0"`
	AccessTokenSecret string `json:"access_token_secret" example:"s3cr3tT0kenValu3XyZ123"`
	GarminUserID      string `json:"garmin_user_id" example:"d3a1c9e8-7b2f-4c99-b77e-8e67cd1a2b10"`
}

type GarminTokenInput struct {
	UserID  string             `json:"user_id" example:"e19c1832-d7f3-4d65-90ea-33a3d7f6d6df"`
	Details GarminTokenDetails `json:"details"`
}

type GarminUserIDResponse struct {
	UserID string `json:"user_id" example:"e19c1832-d7f3-4d65-90ea-33a3d7f6d6df"`
}

type GarminTokenExistsResponse struct {
	Exists bool `json:"exists" example:"true"`
}

type KlabStatusResponse struct {
	Connected bool `json:"connected" example:"true"`
}

type ArchinisisStatusResponse struct {
	Connected bool `json:"connected" example:"true"`
}

type KlabTokenDetails struct {
	SportID      int    `json:"sport_id" example:"51678043"`
	SuomisportID string `json:"suomisport_id" example:"138538"`
}

type ArchinisisTokenDetails struct {
	SportID      int    `json:"sport_id" example:"51678043"`
	SuomisportID string `json:"suomisport_id" example:"138538"`
}

type KlabTokenInput struct {
	UserID  string           `json:"user_id" example:"208e2ffb-ac68-4980-a8b6-b7e0136e4172"`
	Details KlabTokenDetails `json:"details"`
}

type ArchinisisTokenInput struct {
	UserID  string                 `json:"user_id" example:"208e2ffb-ac68-4980-a8b6-b7e0136e4172"`
	Details ArchinisisTokenDetails `json:"details"`
}

type UserDataDetails struct {
	ContactInfo struct {
		SportID            int    `json:"sport_id" example:"87432910"`
		FirstName          string `json:"first_name" example:"Arttu"`
		MiddleName         string `json:"middle_name" example:"Ilmari"`
		LastName           string `json:"last_name" example:"Virtanen"`
		NickName           string `json:"nick_name" example:"Artzi"`
		Phone              string `json:"phone" example:"+358401234567"`
		MobilePhone        string `json:"mobile_phone" example:"+358451112233"`
		Address            string `json:"address" example:"Keskuskatu 12 A, 00100 Helsinki"`
		PhoneVisible       bool   `json:"phone_visible" example:"false"`
		MobilePhoneVisible bool   `json:"mobile_phone_visible" example:"false"`
		AddressVisible     bool   `json:"address_visible" example:"true"`
		SportIDVisible     bool   `json:"sport_id_visible" example:"false"`
		FirstNameVisible   bool   `json:"first_name_visible" example:"true"`
		LastNameVisible    bool   `json:"last_name_visible" example:"true"`
		MiddleNameVisible  bool   `json:"middle_name_visible" example:"false"`
		FirstNameRequired  bool   `json:"first_name_required" example:"true"`
		LastNameRequired   bool   `json:"last_name_required" example:"true"`
		NickNameRequired   bool   `json:"nick_name_required" example:"false"`
	} `json:"contact_info"`
}

type UserDataInput struct {
	Data UserDataDetails `json:"data"`
}

type UserDataResponse struct {
	Data UserDataDetails `json:"data"`
}

type UserIDResponse struct {
	UserID string `json:"user_id" example:"dcabe48a-3578-4743-93ba-001409c82a82"`
}

type NotFoundResponse struct {
	Error string `json:"error" example:"resource not found"`
}

type CoachtechStatusResponse struct {
	Data bool `json:"data" example:"true"`
}

type CoachtechData struct {
	Example  string `json:"example" example:"example_data"`
	Example1 string `json:"example1" example:"example_data_1"`
	Example2 string `json:"example2" example:"example_data_2"`
	Example3 string `json:"example3" example:"example_data_3"`
	Example4 string `json:"example4" example:"example_data_4"`
}

type CoachtechInsertInput struct {
	UserID      string        `json:"user_id" example:"1c2f6ad2-dc8c-4c44-85e2-381d70b093ef"`
	CoachtechID int32         `json:"coachtech_id" example:"123"`
	SummaryDate string        `json:"summary_date" example:"2025-07-14"`
	TestID      string        `json:"test_id" example:"endurance_test_1"`
	Data        CoachtechData `json:"data"`
}
type DeviceInfoConnectedWithData struct {
	Connected  bool `json:"connected" example:"true"`
	DataExists bool `json:"data_exists" example:"true"`
}

type DeviceInfoConnectedNoData struct {
	Connected  bool `json:"connected" example:"true"`
	DataExists bool `json:"data_exists" example:"false"`
}

type DeviceInfoNotConnectedNoData struct {
	Connected  bool `json:"connected" example:"false"`
	DataExists bool `json:"data_exists" example:"false"`
}

type DeviceStatusResponse struct {
	Garmin DeviceInfoConnectedWithData  `json:"garmin"`
	Oura   DeviceInfoConnectedNoData    `json:"oura"`
	Polar  DeviceInfoNotConnectedNoData `json:"polar"`
	Suunto DeviceInfoConnectedWithData  `json:"suunto"`
}

type SourceCacheUpsertInput struct {
	Source string   `json:"source" example:"garmin"`
	Data   []string `json:"data" example:"hr,steps"`
}

type SourceCacheSingleResponse struct {
	Source string   `json:"source" example:"garmin"`
	Data   []string `json:"data" example:"hr,steps"`
}

type SourceCacheItem struct {
	Source string   `json:"source" example:"garmin"`
	Data   []string `json:"data" example:"hr,steps"`
}
