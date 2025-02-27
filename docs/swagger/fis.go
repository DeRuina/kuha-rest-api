package swagger

// Athletes
type Athlete struct {
	FirstName string `json:"first_name" example:"John"`
	LastName  string `json:"last_name" example:"Doe"`
	FisCode   int    `json:"fis_code" example:"123456"`
}

type AthleteListResponse struct {
	Data []Athlete `json:"athletes"`
}

// Nations
type NationsResponse struct {
	Data []string `json:"nations" example:"USA,NOR,GER"`
}
