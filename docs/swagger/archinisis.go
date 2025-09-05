package swagger

type UserDataArchinisisResponse struct {
	Users []string `json:"sportti_ids" example:"12345,67890,54321"`
}

type RaceReportSessionsResponse struct {
	RaceReport []int32 `json:"race_report" example:"101,102,103"`
}
