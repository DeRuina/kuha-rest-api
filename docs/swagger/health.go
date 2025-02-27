package swagger

// Health
type HealthStatusResponse struct {
	Env    string `json:"env" example:"production"`
	Status string `json:"status" example:"ok"`
}
