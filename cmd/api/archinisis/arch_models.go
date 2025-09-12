package archapi

import (
	"time"

	archsqlc "github.com/DeRuina/KUHA-REST-API/internal/db/archinisis"
	"github.com/DeRuina/KUHA-REST-API/internal/utils"
)

type ArchDataUpsertInput struct {
	NationalID   string                 `json:"national_id"    validate:"required,numeric"`
	FirstName    *string                `json:"first_name"     validate:"omitempty"`
	LastName     *string                `json:"last_name"      validate:"omitempty"`
	Initials     *string                `json:"initials"       validate:"omitempty"`
	DateOfBirth  *string                `json:"date_of_birth"  validate:"omitempty"`
	Height       *float64               `json:"height"         validate:"omitempty"`
	Weight       *float64               `json:"weight"         validate:"omitempty"`
	Measurements []ArchMeasurementInput `json:"measurements"   validate:"omitempty,dive"`
}

type ArchMeasurementInput struct {
	Comment            *string `json:"comment"`
	Discipline         *string `json:"discipline"`
	MeasurementGroupID int32   `json:"measurement_group_id" validate:"required,gt=0"`
	MeasurementID      int32   `json:"measurement_id"       validate:"required,gt=0"`
	NbSegments         *int32  `json:"nb_segments"`
	Place              *string `json:"place"`
	RaceID             *int32  `json:"race_id"`
	SessionName        *string `json:"session_name"`
	StartTime          *string `json:"start_time"`
	StopTime           *string `json:"stop_time"`
}

// Mapping
func mapAthleteToParams(in ArchDataUpsertInput, sid string) (archsqlc.UpsertAthleteParams, error) {
	var dob *time.Time
	var err error
	if in.DateOfBirth != nil {
		dob, err = utils.ParseDatePtr(in.DateOfBirth)
		if err != nil {
			return archsqlc.UpsertAthleteParams{}, err
		}
	}

	return archsqlc.UpsertAthleteParams{
		NationalID:  sid,
		FirstName:   utils.NullStringPtr(in.FirstName),
		LastName:    utils.NullStringPtr(in.LastName),
		Initials:    utils.NullStringPtr(in.Initials),
		DateOfBirth: utils.NullTimePtr(dob),
		Height:      utils.NullNumericFromFloat64Ptr(in.Height),
		Weight:      utils.NullNumericFromFloat64Ptr(in.Weight),
	}, nil
}

func mapMeasurementToParams(in ArchMeasurementInput, sid string) (archsqlc.UpsertMeasurementParams, error) {
	var (
		start *time.Time
		stop  *time.Time
		err   error
	)
	if in.StartTime != nil && *in.StartTime != "" {
		start, err = utils.ParseTimestampPtrFlexible(in.StartTime)
		if err != nil {
			return archsqlc.UpsertMeasurementParams{}, err
		}
	}
	if in.StopTime != nil && *in.StopTime != "" {
		stop, err = utils.ParseTimestampPtrFlexible(in.StopTime)
		if err != nil {
			return archsqlc.UpsertMeasurementParams{}, err
		}
	}

	return archsqlc.UpsertMeasurementParams{
		MeasurementGroupID: in.MeasurementGroupID,
		MeasurementID:      utils.NullInt32Ptr(&in.MeasurementID),
		NationalID:         utils.NullString(sid),
		Discipline:         utils.NullStringPtr(in.Discipline),
		SessionName:        utils.NullStringPtr(in.SessionName),
		Place:              utils.NullStringPtr(in.Place),
		RaceID:             utils.NullInt32Ptr(in.RaceID),
		StartTime:          utils.NullTimePtr(start),
		StopTime:           utils.NullTimePtr(stop),
		NbSegments:         utils.NullInt32Ptr(in.NbSegments),
		Comment:            utils.NullStringPtr(in.Comment),
	}, nil
}
