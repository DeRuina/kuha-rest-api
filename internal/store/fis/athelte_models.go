package fis

import (
	fissqlc "github.com/DeRuina/KUHA-REST-API/internal/db/fis"
	"github.com/DeRuina/KUHA-REST-API/internal/utils"
)

type InsertAthleteClean struct {
	Fiscode   int32
	Sporttiid *int32
	Firstname *string
	Lastname  *string
}

type UpdateAthleteClean struct {
	Fiscode   int32
	Sporttiid *int32
	Firstname *string
	Lastname  *string
}

func mapInsertAthleteToParams(in InsertAthleteClean) fissqlc.InsertAthleteParams {
	return fissqlc.InsertAthleteParams{
		Fiscode:   in.Fiscode,
		Sporttiid: utils.NullInt32Ptr(in.Sporttiid),
		Firstname: utils.NullStringPtr(in.Firstname),
		Lastname:  utils.NullStringPtr(in.Lastname),
	}
}

func mapUpdateAthleteToParams(in UpdateAthleteClean) fissqlc.UpdateAthleteByFiscodeParams {
	return fissqlc.UpdateAthleteByFiscodeParams{
		Fiscode:   in.Fiscode,
		Sporttiid: utils.NullInt32Ptr(in.Sporttiid),
		Firstname: utils.NullStringPtr(in.Firstname),
		Lastname:  utils.NullStringPtr(in.Lastname),
	}
}
