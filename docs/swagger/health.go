package swagger

// Health
type HealthStatusResponse struct {
	API           string `json:"api" example:"ok"`
	DBArchinisis  string `json:"db_archinisis" example:"ok"`
	DBAuth        string `json:"db_auth" example:"ok"`
	DBFIS         string `json:"db_fis" example:"ok"`
	DBKAMK        string `json:"db_kamk" example:"ok"`
	DBKlab        string `json:"db_klab" example:"ok"`
	DBTietoevry   string `json:"db_tietoevry" example:"ok"`
	DBUTV         string `json:"db_utv" example:"ok"`
	Env           string `json:"env" example:"development"`
	Goroutines    int    `json:"goroutines" example:"21"`
	Redis         string `json:"redis" example:"ok"`
	UptimeSeconds int64  `json:"uptime_seconds" example:"149039"`
	Version       string `json:"version" example:"1.2.1"`
}
