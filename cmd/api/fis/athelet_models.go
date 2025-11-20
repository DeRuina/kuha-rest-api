package fisapi

import (
	"github.com/DeRuina/KUHA-REST-API/internal/store/fis"
)

type SporttiIDParam struct {
	SporttiID string `form:"sporttiid" validate:"required,numeric"`
}

type AthleteFiscodeParam struct {
	Fiscode string `form:"fiscode" validate:"required,numeric"`
}

type InsertAthleteInput struct {
	Fiscode   int32   `json:"fiscode" validate:"required"`
	Sporttiid *int32  `json:"sporttiid"`
	Firstname *string `json:"firstname"`
	Lastname  *string `json:"lastname"`
}

type UpdateAthleteInput struct {
	Fiscode   int32   `json:"fiscode" validate:"required"`
	Sporttiid *int32  `json:"sporttiid"`
	Firstname *string `json:"firstname"`
	Lastname  *string `json:"lastname"`
}

func mapInsertAthleteInput(in InsertAthleteInput) fis.InsertAthleteClean {
	return fis.InsertAthleteClean{
		Fiscode:   in.Fiscode,
		Sporttiid: in.Sporttiid,
		Firstname: in.Firstname,
		Lastname:  in.Lastname,
	}
}

func mapUpdateAthleteInput(in UpdateAthleteInput) fis.UpdateAthleteClean {
	return fis.UpdateAthleteClean{
		Fiscode:   in.Fiscode,
		Sporttiid: in.Sporttiid,
		Firstname: in.Firstname,
		Lastname:  in.Lastname,
	}
}
