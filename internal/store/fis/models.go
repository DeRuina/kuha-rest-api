package fis

import (
	"database/sql"
	"time"

	fissqlc "github.com/DeRuina/KUHA-REST-API/internal/db/fis"
	"github.com/DeRuina/KUHA-REST-API/internal/utils"
)

type AthleteRow struct {
	Firstname *string `json:"Firstname,omitempty"`
	Lastname  *string `json:"Lastname,omitempty"`
	Fiscode   *int32  `json:"Fiscode,omitempty"`
}

type InsertCompetitorClean struct {
	Competitorid       *int32
	Personid           *int32
	Ipcid              *int32
	Fiscode            *int32
	Birthdate          *time.Time
	StatusDate         *time.Time
	Fee                *string
	Dateofcreation     *time.Time
	Injury             *int32
	Version            *int32
	Compidmssql        *int32
	Carving            *int32
	Photo              *int32
	Notallowed         *int32
	Published          *int32
	Team               *int32
	PhotoBig           *int32
	Lastupdate         *time.Time
	Statusnextlist     *string
	Alternatenamecheck *string
	Deletedat          *string
	Doped              *string
	Createdby          *string
	Categorycode       *string
	Classname          *string
	Data               *string
	Lastupdateby       *string
	Disciplines        *string
	Type               *string
	Sectorcode         *string
	Classcode          *string
	Lastname           *string
	Firstname          *string
	Gender             *string
	Natteam            *string
	Nationcode         *string
	Nationalcode       *string
	Skiclub            *string
	Association        *string
	Status             *string
	StatusOld          *string
	StatusBy           *string
	Tragroup           *string
}

type UpdateCompetitorClean struct {
	Competitorid       int32
	Personid           *int32
	Ipcid              *int32
	Type               *string
	Sectorcode         *string
	Fiscode            *int32
	Lastname           *string
	Firstname          *string
	Gender             *string
	Birthdate          *time.Time
	Nationcode         *string
	Nationalcode       *string
	Skiclub            *string
	Association        *string
	Status             *string
	StatusOld          *string
	StatusBy           *string
	StatusDate         *time.Time
	Statusnextlist     *string
	Alternatenamecheck *string
	Fee                *string
	Dateofcreation     *time.Time
	Createdby          *string
	Injury             *int32
	Version            *int32
	Compidmssql        *int32
	Carving            *int32
	Photo              *int32
	Notallowed         *int32
	Natteam            *string
	Tragroup           *string
	Published          *int32
	Doped              *string
	Team               *int32
	PhotoBig           *int32
	Data               *string
	Lastupdateby       *string
	Disciplines        *string
	Lastupdate         *time.Time
	Deletedat          *string
	Categorycode       *string
	Classname          *string
	Classcode          *string
}

func mapInsertToParams(in InsertCompetitorClean) fissqlc.InsertCompetitorParams {
	return fissqlc.InsertCompetitorParams{
		Competitorid:       utils.NullInt32Ptr(in.Competitorid),
		Personid:           utils.NullInt32Ptr(in.Personid),
		Ipcid:              utils.NullInt32Ptr(in.Ipcid),
		Fiscode:            utils.NullInt32Ptr(in.Fiscode),
		Birthdate:          utils.NullTimePtr(in.Birthdate),
		StatusDate:         utils.NullTimePtr(in.StatusDate),
		Fee:                utils.NullStringPtr(in.Fee),
		Dateofcreation:     utils.NullTimePtr(in.Dateofcreation),
		Injury:             utils.NullInt32Ptr(in.Injury),
		Version:            utils.NullInt32Ptr(in.Version),
		Compidmssql:        utils.NullInt32Ptr(in.Compidmssql),
		Carving:            utils.NullInt32Ptr(in.Carving),
		Photo:              utils.NullInt32Ptr(in.Photo),
		Notallowed:         utils.NullInt32Ptr(in.Notallowed),
		Published:          utils.NullInt32Ptr(in.Published),
		Team:               utils.NullInt32Ptr(in.Team),
		PhotoBig:           utils.NullInt32Ptr(in.PhotoBig),
		Lastupdate:         utils.NullTimePtr(in.Lastupdate),
		Statusnextlist:     utils.NullStringPtr(in.Statusnextlist),
		Alternatenamecheck: utils.NullStringPtr(in.Alternatenamecheck),
		Deletedat:          utils.NullStringPtr(in.Deletedat),
		Doped:              utils.NullStringPtr(in.Doped),
		Createdby:          utils.NullStringPtr(in.Createdby),
		Categorycode:       utils.NullStringPtr(in.Categorycode),
		Classname:          utils.NullStringPtr(in.Classname),
		Data:               utils.NullStringPtr(in.Data),
		Lastupdateby:       utils.NullStringPtr(in.Lastupdateby),
		Disciplines:        utils.NullStringPtr(in.Disciplines),
		Type:               utils.NullStringPtr(in.Type),
		Sectorcode:         utils.NullStringPtr(in.Sectorcode),
		Classcode:          utils.NullStringPtr(in.Classcode),
		Lastname:           utils.NullStringPtr(in.Lastname),
		Firstname:          utils.NullStringPtr(in.Firstname),
		Gender:             utils.NullStringPtr(in.Gender),
		Natteam:            utils.NullStringPtr(in.Natteam),
		Nationcode:         utils.NullStringPtr(in.Nationcode),
		Nationalcode:       utils.NullStringPtr(in.Nationalcode),
		Skiclub:            utils.NullStringPtr(in.Skiclub),
		Association:        utils.NullStringPtr(in.Association),
		Status:             utils.NullStringPtr(in.Status),
		StatusOld:          utils.NullStringPtr(in.StatusOld),
		StatusBy:           utils.NullStringPtr(in.StatusBy),
		Tragroup:           utils.NullStringPtr(in.Tragroup),
	}
}

func mapUpdateToParams(in UpdateCompetitorClean) fissqlc.UpdateCompetitorByIDParams {
	return fissqlc.UpdateCompetitorByIDParams{
		Competitorid:       sql.NullInt32{Int32: in.Competitorid, Valid: true},
		Personid:           utils.NullInt32Ptr(in.Personid),
		Ipcid:              utils.NullInt32Ptr(in.Ipcid),
		Type:               utils.NullStringPtr(in.Type),
		Sectorcode:         utils.NullStringPtr(in.Sectorcode),
		Fiscode:            utils.NullInt32Ptr(in.Fiscode),
		Lastname:           utils.NullStringPtr(in.Lastname),
		Firstname:          utils.NullStringPtr(in.Firstname),
		Gender:             utils.NullStringPtr(in.Gender),
		Birthdate:          utils.NullTimePtr(in.Birthdate),
		Nationcode:         utils.NullStringPtr(in.Nationcode),
		Nationalcode:       utils.NullStringPtr(in.Nationalcode),
		Skiclub:            utils.NullStringPtr(in.Skiclub),
		Association:        utils.NullStringPtr(in.Association),
		Status:             utils.NullStringPtr(in.Status),
		StatusOld:          utils.NullStringPtr(in.StatusOld),
		StatusBy:           utils.NullStringPtr(in.StatusBy),
		StatusDate:         utils.NullTimePtr(in.StatusDate),
		Statusnextlist:     utils.NullStringPtr(in.Statusnextlist),
		Alternatenamecheck: utils.NullStringPtr(in.Alternatenamecheck),
		Fee:                utils.NullStringPtr(in.Fee),
		Dateofcreation:     utils.NullTimePtr(in.Dateofcreation),
		Createdby:          utils.NullStringPtr(in.Createdby),
		Injury:             utils.NullInt32Ptr(in.Injury),
		Version:            utils.NullInt32Ptr(in.Version),
		Compidmssql:        utils.NullInt32Ptr(in.Compidmssql),
		Carving:            utils.NullInt32Ptr(in.Carving),
		Photo:              utils.NullInt32Ptr(in.Photo),
		Notallowed:         utils.NullInt32Ptr(in.Notallowed),
		Natteam:            utils.NullStringPtr(in.Natteam),
		Tragroup:           utils.NullStringPtr(in.Tragroup),
		Published:          utils.NullInt32Ptr(in.Published),
		Doped:              utils.NullStringPtr(in.Doped),
		Team:               utils.NullInt32Ptr(in.Team),
		PhotoBig:           utils.NullInt32Ptr(in.PhotoBig),
		Data:               utils.NullStringPtr(in.Data),
		Lastupdateby:       utils.NullStringPtr(in.Lastupdateby),
		Disciplines:        utils.NullStringPtr(in.Disciplines),
		Lastupdate:         utils.NullTimePtr(in.Lastupdate),
		Deletedat:          utils.NullStringPtr(in.Deletedat),
		Categorycode:       utils.NullStringPtr(in.Categorycode),
		Classname:          utils.NullStringPtr(in.Classname),
		Classcode:          utils.NullStringPtr(in.Classcode),
	}
}
