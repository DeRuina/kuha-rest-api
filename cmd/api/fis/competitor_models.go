package fisapi

import (
	"time"

	fissqlc "github.com/DeRuina/KUHA-REST-API/internal/db/fis"
	"github.com/DeRuina/KUHA-REST-API/internal/store/fis"
	"github.com/DeRuina/KUHA-REST-API/internal/utils"
)

type NationsBySectorResponse struct {
	Nations []string `json:"nations"`
}

// Query params
type SectorParam struct {
	Sector string `form:"sectorcode" validate:"required,oneof=JP NK CC"`
}

type CompetitorIDParam struct {
	CompetitorID string `form:"id" validate:"required,numeric"`
}

type InsertCompetitorInput struct {
	Competitorid       int32   `json:"competitorid" validate:"required"`
	Personid           *int32  `json:"personid"`
	Ipcid              *int32  `json:"ipcid"`
	Fiscode            *int32  `json:"fiscode"`
	Birthdate          *string `json:"birthdate"`
	StatusDate         *string `json:"status_date"`
	Fee                *string `json:"fee"`
	Dateofcreation     *string `json:"dateofcreation"`
	Injury             *int32  `json:"injury"`
	Version            *int32  `json:"version"`
	Compidmssql        *int32  `json:"compidmssql"`
	Carving            *int32  `json:"carving"`
	Photo              *int32  `json:"photo"`
	Notallowed         *int32  `json:"notallowed"`
	Published          *int32  `json:"published"`
	Team               *int32  `json:"team"`
	PhotoBig           *int32  `json:"photo_big"`
	Lastupdate         *string `json:"lastupdate"`
	Statusnextlist     *string `json:"statusnextlist"`
	Alternatenamecheck *string `json:"alternatenamecheck"`
	Deletedat          *string `json:"deletedat"`
	Doped              *string `json:"doped"`
	Createdby          *string `json:"createdby"`
	Categorycode       *string `json:"categorycode"`
	Classname          *string `json:"classname"`
	Data               *string `json:"data"`
	Lastupdateby       *string `json:"lastupdateby"`
	Disciplines        *string `json:"disciplines"`
	Type               *string `json:"type"`
	Sectorcode         *string `json:"sectorcode"`
	Classcode          *string `json:"classcode"`
	Lastname           *string `json:"lastname"`
	Firstname          *string `json:"firstname"`
	Gender             *string `json:"gender"`
	Natteam            *string `json:"natteam"`
	Nationcode         *string `json:"nationcode"`
	Nationalcode       *string `json:"nationalcode"`
	Skiclub            *string `json:"skiclub"`
	Association        *string `json:"association"`
	Status             *string `json:"status"`
	StatusOld          *string `json:"status_old"`
	StatusBy           *string `json:"status_by"`
	Tragroup           *string `json:"tragroup"`
}

type UpdateCompetitorInput struct {
	Competitorid       int32   `json:"competitorid" validate:"required"`
	Personid           *int32  `json:"personid"`
	Ipcid              *int32  `json:"ipcid"`
	Type               *string `json:"type"`
	Sectorcode         *string `json:"sectorcode"`
	Fiscode            *int32  `json:"fiscode"`
	Lastname           *string `json:"lastname"`
	Firstname          *string `json:"firstname"`
	Gender             *string `json:"gender"`
	Birthdate          *string `json:"birthdate"`
	Nationcode         *string `json:"nationcode"`
	Nationalcode       *string `json:"nationalcode"`
	Skiclub            *string `json:"skiclub"`
	Association        *string `json:"association"`
	Status             *string `json:"status"`
	StatusOld          *string `json:"status_old"`
	StatusBy           *string `json:"status_by"`
	StatusDate         *string `json:"status_date"`
	Statusnextlist     *string `json:"statusnextlist"`
	Alternatenamecheck *string `json:"alternatenamecheck"`
	Fee                *string `json:"fee"`
	Dateofcreation     *string `json:"dateofcreation"`
	Createdby          *string `json:"createdby"`
	Injury             *int32  `json:"injury"`
	Version            *int32  `json:"version"`
	Compidmssql        *int32  `json:"compidmssql"`
	Carving            *int32  `json:"carving"`
	Photo              *int32  `json:"photo"`
	Notallowed         *int32  `json:"notallowed"`
	Natteam            *string `json:"natteam"`
	Tragroup           *string `json:"tragroup"`
	Published          *int32  `json:"published"`
	Doped              *string `json:"doped"`
	Team               *int32  `json:"team"`
	PhotoBig           *int32  `json:"photo_big"`
	Data               *string `json:"data"`
	Lastupdateby       *string `json:"lastupdateby"`
	Disciplines        *string `json:"disciplines"`
	Lastupdate         *string `json:"lastupdate"`
	Deletedat          *string `json:"deletedat"`
	Categorycode       *string `json:"categorycode"`
	Classname          *string `json:"classname"`
	Classcode          *string `json:"classcode"`
}

func mapInsertInput(in InsertCompetitorInput) (fis.InsertCompetitorClean, error) {
	var birth, sdate, doc, lup *time.Time
	var err error

	if in.Birthdate != nil {
		d, e := utils.ParseDatePtr(in.Birthdate)
		if e != nil {
			return fis.InsertCompetitorClean{}, e
		}
		birth = d
	}
	if in.StatusDate != nil {
		sdate, err = utils.ParseTimestampPtr(in.StatusDate)
		if err != nil {
			return fis.InsertCompetitorClean{}, err
		}
	}
	if in.Dateofcreation != nil {
		doc, err = utils.ParseDatePtr(in.Dateofcreation)
		if err != nil {
			return fis.InsertCompetitorClean{}, err
		}
	}
	if in.Lastupdate != nil {
		lup, err = utils.ParseTimestampPtr(in.Lastupdate)
		if err != nil {
			return fis.InsertCompetitorClean{}, err
		}
	}

	return fis.InsertCompetitorClean{
		Competitorid:       in.Competitorid,
		Personid:           in.Personid,
		Ipcid:              in.Ipcid,
		Fiscode:            in.Fiscode,
		Birthdate:          birth,
		StatusDate:         sdate,
		Fee:                in.Fee,
		Dateofcreation:     doc,
		Injury:             in.Injury,
		Version:            in.Version,
		Compidmssql:        in.Compidmssql,
		Carving:            in.Carving,
		Photo:              in.Photo,
		Notallowed:         in.Notallowed,
		Published:          in.Published,
		Team:               in.Team,
		PhotoBig:           in.PhotoBig,
		Lastupdate:         lup,
		Statusnextlist:     in.Statusnextlist,
		Alternatenamecheck: in.Alternatenamecheck,
		Deletedat:          in.Deletedat,
		Doped:              in.Doped,
		Createdby:          in.Createdby,
		Categorycode:       in.Categorycode,
		Classname:          in.Classname,
		Data:               in.Data,
		Lastupdateby:       in.Lastupdateby,
		Disciplines:        in.Disciplines,
		Type:               in.Type,
		Sectorcode:         in.Sectorcode,
		Classcode:          in.Classcode,
		Lastname:           in.Lastname,
		Firstname:          in.Firstname,
		Gender:             in.Gender,
		Natteam:            in.Natteam,
		Nationcode:         in.Nationcode,
		Nationalcode:       in.Nationalcode,
		Skiclub:            in.Skiclub,
		Association:        in.Association,
		Status:             in.Status,
		StatusOld:          in.StatusOld,
		StatusBy:           in.StatusBy,
		Tragroup:           in.Tragroup,
	}, nil
}

func mapUpdateInput(in UpdateCompetitorInput) (fis.UpdateCompetitorClean, error) {
	var birth, sdate, doc, lup *time.Time
	var err error

	if in.Birthdate != nil {
		d, e := utils.ParseDatePtr(in.Birthdate)
		if e != nil {
			return fis.UpdateCompetitorClean{}, e
		}
		birth = d
	}
	if in.StatusDate != nil {
		sdate, err = utils.ParseTimestampPtr(in.StatusDate)
		if err != nil {
			return fis.UpdateCompetitorClean{}, err
		}
	}
	if in.Dateofcreation != nil {
		doc, err = utils.ParseDatePtr(in.Dateofcreation)
		if err != nil {
			return fis.UpdateCompetitorClean{}, err
		}
	}
	if in.Lastupdate != nil {
		lup, err = utils.ParseTimestampPtr(in.Lastupdate)
		if err != nil {
			return fis.UpdateCompetitorClean{}, err
		}
	}

	return fis.UpdateCompetitorClean{
		Competitorid:       in.Competitorid,
		Personid:           in.Personid,
		Ipcid:              in.Ipcid,
		Type:               in.Type,
		Sectorcode:         in.Sectorcode,
		Fiscode:            in.Fiscode,
		Lastname:           in.Lastname,
		Firstname:          in.Firstname,
		Gender:             in.Gender,
		Birthdate:          birth,
		Nationcode:         in.Nationcode,
		Nationalcode:       in.Nationalcode,
		Skiclub:            in.Skiclub,
		Association:        in.Association,
		Status:             in.Status,
		StatusOld:          in.StatusOld,
		StatusBy:           in.StatusBy,
		StatusDate:         sdate,
		Statusnextlist:     in.Statusnextlist,
		Alternatenamecheck: in.Alternatenamecheck,
		Fee:                in.Fee,
		Dateofcreation:     doc,
		Createdby:          in.Createdby,
		Injury:             in.Injury,
		Version:            in.Version,
		Compidmssql:        in.Compidmssql,
		Carving:            in.Carving,
		Photo:              in.Photo,
		Notallowed:         in.Notallowed,
		Natteam:            in.Natteam,
		Tragroup:           in.Tragroup,
		Published:          in.Published,
		Doped:              in.Doped,
		Team:               in.Team,
		PhotoBig:           in.PhotoBig,
		Data:               in.Data,
		Lastupdateby:       in.Lastupdateby,
		Disciplines:        in.Disciplines,
		Lastupdate:         lup,
		Deletedat:          in.Deletedat,
		Categorycode:       in.Categorycode,
		Classname:          in.Classname,
		Classcode:          in.Classcode,
	}, nil
}

type FISCompetitorResponse struct {
	Competitorid       int32   `json:"competitorid"`
	Personid           *int32  `json:"personid"`
	Ipcid              *int32  `json:"ipcid"`
	Type               *string `json:"type"`
	Sectorcode         *string `json:"sectorcode"`
	Fiscode            *int32  `json:"fiscode"`
	Lastname           *string `json:"lastname"`
	Firstname          *string `json:"firstname"`
	Gender             *string `json:"gender"`
	Birthdate          *string `json:"birthdate"`      // YYYY-MM-DD
	StatusDate         *string `json:"status_date"`    // RFC3339
	Dateofcreation     *string `json:"dateofcreation"` // YYYY-MM-DD
	Lastupdate         *string `json:"lastupdate"`     // RFC3339
	Nationcode         *string `json:"nationcode"`
	Nationalcode       *string `json:"nationalcode"`
	Skiclub            *string `json:"skiclub"`
	Association        *string `json:"association"`
	Status             *string `json:"status"`
	StatusOld          *string `json:"status_old"`
	StatusBy           *string `json:"status_by"`
	Statusnextlist     *string `json:"statusnextlist"`
	Alternatenamecheck *string `json:"alternatenamecheck"`
	Fee                *string `json:"fee"`
	Createdby          *string `json:"createdby"`
	Injury             *int32  `json:"injury"`
	Version            *int32  `json:"version"`
	Compidmssql        *int32  `json:"compidmssql"`
	Carving            *int32  `json:"carving"`
	Photo              *int32  `json:"photo"`
	Notallowed         *int32  `json:"notallowed"`
	Natteam            *string `json:"natteam"`
	Tragroup           *string `json:"tragroup"`
	Published          *int32  `json:"published"`
	Doped              *string `json:"doped"`
	Team               *int32  `json:"team"`
	PhotoBig           *int32  `json:"photo_big"`
	Data               *string `json:"data"`
	Lastupdateby       *string `json:"lastupdateby"`
	Disciplines        *string `json:"disciplines"`
	Deletedat          *string `json:"deletedat"`
	Categorycode       *string `json:"categorycode"`
	Classname          *string `json:"classname"`
	Classcode          *string `json:"classcode"`
}

func FISCompetitorFullFromSqlc(row fissqlc.ACompetitor) FISCompetitorResponse {
	var (
		birthStr  *string
		statusStr *string
		createStr *string
		updateStr *string
	)

	if row.Birthdate.Valid {
		birthStr = utils.FormatDatePtr(row.Birthdate)
	}
	if row.StatusDate.Valid {
		statusStr = utils.FormatTimestampPtr(row.StatusDate)
	}
	if row.Dateofcreation.Valid {
		createStr = utils.FormatDatePtr(row.Dateofcreation)
	}
	if row.Lastupdate.Valid {
		updateStr = utils.FormatTimestampPtr(row.Lastupdate)
	}

	return FISCompetitorResponse{
		Competitorid:       row.Competitorid,
		Personid:           utils.Int32PtrOrNil(row.Personid),
		Ipcid:              utils.Int32PtrOrNil(row.Ipcid),
		Type:               utils.StringPtrOrNil(row.Type),
		Sectorcode:         utils.StringPtrOrNil(row.Sectorcode),
		Fiscode:            utils.Int32PtrOrNil(row.Fiscode),
		Lastname:           utils.StringPtrOrNil(row.Lastname),
		Firstname:          utils.StringPtrOrNil(row.Firstname),
		Gender:             utils.StringPtrOrNil(row.Gender),
		Birthdate:          birthStr,
		Nationcode:         utils.StringPtrOrNil(row.Nationcode),
		Nationalcode:       utils.StringPtrOrNil(row.Nationalcode),
		Skiclub:            utils.StringPtrOrNil(row.Skiclub),
		Association:        utils.StringPtrOrNil(row.Association),
		Status:             utils.StringPtrOrNil(row.Status),
		StatusOld:          utils.StringPtrOrNil(row.StatusOld),
		StatusBy:           utils.StringPtrOrNil(row.StatusBy),
		StatusDate:         statusStr,
		Statusnextlist:     utils.StringPtrOrNil(row.Statusnextlist),
		Alternatenamecheck: utils.StringPtrOrNil(row.Alternatenamecheck),
		Fee:                utils.StringPtrOrNil(row.Fee),
		Dateofcreation:     createStr,
		Createdby:          utils.StringPtrOrNil(row.Createdby),
		Injury:             utils.Int32PtrOrNil(row.Injury),
		Version:            utils.Int32PtrOrNil(row.Version),
		Compidmssql:        utils.Int32PtrOrNil(row.Compidmssql),
		Carving:            utils.Int32PtrOrNil(row.Carving),
		Photo:              utils.Int32PtrOrNil(row.Photo),
		Notallowed:         utils.Int32PtrOrNil(row.Notallowed),
		Natteam:            utils.StringPtrOrNil(row.Natteam),
		Tragroup:           utils.StringPtrOrNil(row.Tragroup),
		Published:          utils.Int32PtrOrNil(row.Published),
		Doped:              utils.StringPtrOrNil(row.Doped),
		Team:               utils.Int32PtrOrNil(row.Team),
		PhotoBig:           utils.Int32PtrOrNil(row.PhotoBig),
		Data:               utils.StringPtrOrNil(row.Data),
		Lastupdateby:       utils.StringPtrOrNil(row.Lastupdateby),
		Disciplines:        utils.StringPtrOrNil(row.Disciplines),
		Lastupdate:         updateStr,
		Deletedat:          utils.StringPtrOrNil(row.Deletedat),
		Categorycode:       utils.StringPtrOrNil(row.Categorycode),
		Classname:          utils.StringPtrOrNil(row.Classname),
		Classcode:          utils.StringPtrOrNil(row.Classcode),
	}
}
