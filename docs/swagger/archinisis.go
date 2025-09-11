package swagger

type UserDataArchinisisResponse struct {
	Users []string `json:"sportti_ids" example:"12345,67890,54321"`
}

type RaceReportSessionsResponse struct {
	RaceReport []int32 `json:"race_report" example:"101,102,103"`
}

type ArchRaceReportUpsertRequest struct {
	SporttiID  string `json:"sportti_id" example:"27578816"`
	SessionID  int32  `json:"session_id" example:"1842"`
	RaceReport string `json:"race_report" example:"<!DOCTYPE html><html><head><title>Race</title></head><body><h1>Report</h1></body></html>"`
}
