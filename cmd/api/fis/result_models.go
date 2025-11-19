package fisapi

import (
	"time"

	fissqlc "github.com/DeRuina/KUHA-REST-API/internal/db/fis"
	"github.com/DeRuina/KUHA-REST-API/internal/store/fis"
	"github.com/DeRuina/KUHA-REST-API/internal/utils"
)

type AthleteResultsCCQuery struct {
	SeasonCode     []string `form:"seasoncode"`
	DisciplineCode []string `form:"disciplinecode"`
	CatCode        []string `form:"catcode"`
}

type InsertResultCCInput struct {
	Recid          int32   `json:"recid" validate:"required"`
	Raceid         *int32  `json:"raceid"`
	Competitorid   *int32  `json:"competitorid"`
	Status         *string `json:"status"`
	Reason         *string `json:"reason"`
	Position       *string `json:"position"`
	Pf             *int32  `json:"pf"`
	Status2        *string `json:"status2"`
	Bib            *string `json:"bib"`
	Bibcolor       *string `json:"bibcolor"`
	Fiscode        *int32  `json:"fiscode"`
	Competitorname *string `json:"competitorname"`
	Nationcode     *string `json:"nationcode"`
	Stage          *string `json:"stage"`
	Level          *string `json:"level"`
	Heat           *string `json:"heat"`
	Timer1         *string `json:"timer1"`
	Timer2         *string `json:"timer2"`
	Timer3         *string `json:"timer3"`
	Timetot        *string `json:"timetot"`
	Valid          *string `json:"valid"`
	Racepoints     *string `json:"racepoints"`
	Cuppoints      *string `json:"cuppoints"`
	Bonustime      *string `json:"bonustime"`
	Bonuscuppoints *string `json:"bonuscuppoints"`
	Version        *string `json:"version"`
	Rg1            *string `json:"rg1"`
	Rg2            *string `json:"rg2"`
	Lastupdate     *string `json:"lastupdate"` // RFC3339
}

type UpdateResultCCInput = InsertResultCCInput

func mapInsertResultCCInput(in InsertResultCCInput) (fis.InsertResultCCClean, error) {
	var lup *time.Time
	var err error

	if in.Lastupdate != nil {
		lup, err = utils.ParseTimestampPtr(in.Lastupdate)
		if err != nil {
			return fis.InsertResultCCClean{}, err
		}
	}

	return fis.InsertResultCCClean{
		Recid:          in.Recid,
		Raceid:         in.Raceid,
		Competitorid:   in.Competitorid,
		Status:         in.Status,
		Reason:         in.Reason,
		Position:       in.Position,
		Pf:             in.Pf,
		Status2:        in.Status2,
		Bib:            in.Bib,
		Bibcolor:       in.Bibcolor,
		Fiscode:        in.Fiscode,
		Competitorname: in.Competitorname,
		Nationcode:     in.Nationcode,
		Stage:          in.Stage,
		Level:          in.Level,
		Heat:           in.Heat,
		Timer1:         in.Timer1,
		Timer2:         in.Timer2,
		Timer3:         in.Timer3,
		Timetot:        in.Timetot,
		Valid:          in.Valid,
		Racepoints:     in.Racepoints,
		Cuppoints:      in.Cuppoints,
		Bonustime:      in.Bonustime,
		Bonuscuppoints: in.Bonuscuppoints,
		Version:        in.Version,
		Rg1:            in.Rg1,
		Rg2:            in.Rg2,
		Lastupdate:     lup,
	}, nil
}

func mapUpdateResultCCInput(in UpdateResultCCInput) (fis.UpdateResultCCClean, error) {
	clean, err := mapInsertResultCCInput(InsertResultCCInput(in))
	return fis.UpdateResultCCClean(clean), err
}

type FISResultCCFullResponse struct {
	Recid          int32   `json:"recid"`
	Raceid         *int32  `json:"raceid"`
	Competitorid   *int32  `json:"competitorid"`
	Status         *string `json:"status"`
	Reason         *string `json:"reason"`
	Position       *string `json:"position"`
	Pf             *int32  `json:"pf"`
	Status2        *string `json:"status2"`
	Bib            *string `json:"bib"`
	Bibcolor       *string `json:"bibcolor"`
	Fiscode        *int32  `json:"fiscode"`
	Competitorname *string `json:"competitorname"`
	Nationcode     *string `json:"nationcode"`
	Stage          *string `json:"stage"`
	Level          *string `json:"level"`
	Heat           *string `json:"heat"`
	Timer1         *string `json:"timer1"`
	Timer2         *string `json:"timer2"`
	Timer3         *string `json:"timer3"`
	Timetot        *string `json:"timetot"`
	Valid          *string `json:"valid"`
	Racepoints     *string `json:"racepoints"`
	Cuppoints      *string `json:"cuppoints"`
	Bonustime      *string `json:"bonustime"`
	Bonuscuppoints *string `json:"bonuscuppoints"`
	Version        *string `json:"version"`
	Rg1            *string `json:"rg1"`
	Rg2            *string `json:"rg2"`
	Lastupdate     *string `json:"lastupdate"`
}

func FISResultCCFullFromSqlc(row fissqlc.AResultcc) FISResultCCFullResponse {
	var lastUpdateStr *string
	if row.Lastupdate.Valid {
		lastUpdateStr = utils.FormatTimestampPtr(row.Lastupdate)
	}

	return FISResultCCFullResponse{
		Recid:          row.Recid,
		Raceid:         utils.Int32PtrOrNil(row.Raceid),
		Competitorid:   utils.Int32PtrOrNil(row.Competitorid),
		Status:         utils.StringPtrOrNil(row.Status),
		Reason:         utils.StringPtrOrNil(row.Reason),
		Position:       utils.StringPtrOrNil(row.Position),
		Pf:             utils.Int32PtrOrNil(row.Pf),
		Status2:        utils.StringPtrOrNil(row.Status2),
		Bib:            utils.StringPtrOrNil(row.Bib),
		Bibcolor:       utils.StringPtrOrNil(row.Bibcolor),
		Fiscode:        utils.Int32PtrOrNil(row.Fiscode),
		Competitorname: utils.StringPtrOrNil(row.Competitorname),
		Nationcode:     utils.StringPtrOrNil(row.Nationcode),
		Stage:          utils.StringPtrOrNil(row.Stage),
		Level:          utils.StringPtrOrNil(row.Level),
		Heat:           utils.StringPtrOrNil(row.Heat),
		Timer1:         utils.StringPtrOrNil(row.Timer1),
		Timer2:         utils.StringPtrOrNil(row.Timer2),
		Timer3:         utils.StringPtrOrNil(row.Timer3),
		Timetot:        utils.StringPtrOrNil(row.Timetot),
		Valid:          utils.StringPtrOrNil(row.Valid),
		Racepoints:     utils.StringPtrOrNil(row.Racepoints),
		Cuppoints:      utils.StringPtrOrNil(row.Cuppoints),
		Bonustime:      utils.StringPtrOrNil(row.Bonustime),
		Bonuscuppoints: utils.StringPtrOrNil(row.Bonuscuppoints),
		Version:        utils.StringPtrOrNil(row.Version),
		Rg1:            utils.StringPtrOrNil(row.Rg1),
		Rg2:            utils.StringPtrOrNil(row.Rg2),
		Lastupdate:     lastUpdateStr,
	}
}

type FISAthleteResultCCRow struct {
	Recid          int32   `json:"recid"`
	Raceid         *int32  `json:"raceid"`
	Position       *string `json:"position"`
	Timetot        *string `json:"timetot"`
	Competitorid   *int32  `json:"competitorid"`
	Racedate       *string `json:"racedate"`
	Seasoncode     *int32  `json:"seasoncode"`
	Disciplinecode *string `json:"disciplinecode"`
	Catcode        *string `json:"catcode"`
	Place          *string `json:"place"`
}

func FISAthleteResultCCFromSqlc(row fissqlc.GetAthleteResultsCCRow) FISAthleteResultCCRow {
	var raceDateStr *string
	if row.Racedate.Valid {
		raceDateStr = utils.FormatDatePtr(row.Racedate)
	}
	return FISAthleteResultCCRow{
		Recid:          row.Recid,
		Raceid:         utils.Int32PtrOrNil(row.Raceid),
		Position:       utils.StringPtrOrNil(row.Position),
		Timetot:        utils.StringPtrOrNil(row.Timetot),
		Competitorid:   utils.Int32PtrOrNil(row.Competitorid),
		Racedate:       raceDateStr,
		Seasoncode:     utils.Int32PtrOrNil(row.Seasoncode),
		Disciplinecode: utils.StringPtrOrNil(row.Disciplinecode),
		Catcode:        utils.StringPtrOrNil(row.Catcode),
		Place:          utils.StringPtrOrNil(row.Place),
	}
}

type AthleteResultsJPQuery struct {
	SeasonCode     []string `form:"seasoncode"`
	DisciplineCode []string `form:"disciplinecode"`
	CatCode        []string `form:"catcode"`
}

type InsertResultJPInput struct {
	Recid          int32   `json:"recid" validate:"required"`
	Raceid         *int32  `json:"raceid"`
	Competitorid   *int32  `json:"competitorid"`
	Status         *string `json:"status"`
	Status2        *string `json:"status2"`
	Position       *int32  `json:"position"`
	Bib            *int32  `json:"bib"`
	Fiscode        *int32  `json:"fiscode"`
	Competitorname *string `json:"competitorname"`
	Nationcode     *string `json:"nationcode"`
	Level          *string `json:"level"`
	Heat           *string `json:"heat"`
	Stage          *string `json:"stage"`

	J1r1     *string `json:"j1r1"`
	J2r1     *string `json:"j2r1"`
	J3r1     *string `json:"j3r1"`
	J4r1     *string `json:"j4r1"`
	J5r1     *string `json:"j5r1"`
	Speedr1  *string `json:"speedr1"`
	Distr1   *string `json:"distr1"`
	Disptsr1 *string `json:"disptsr1"`
	Judptsr1 *string `json:"judptsr1"`
	Totrun1  *string `json:"totrun1"`
	Posr1    *string `json:"posr1"`
	Statusr1 *string `json:"statusr1"`

	J1r2     *string `json:"j1r2"`
	J2r2     *string `json:"j2r2"`
	J3r2     *string `json:"j3r2"`
	J4r2     *string `json:"j4r2"`
	J5r2     *string `json:"j5r2"`
	Speedr2  *string `json:"speedr2"`
	Distr2   *string `json:"distr2"`
	Disptsr2 *string `json:"disptsr2"`
	Judptsr2 *string `json:"judptsr2"`
	Totrun2  *string `json:"totrun2"`
	Posr2    *string `json:"posr2"`
	Statusr2 *string `json:"statusr2"`

	J1r3     *string `json:"j1r3"`
	J2r3     *string `json:"j2r3"`
	J3r3     *string `json:"j3r3"`
	J4r3     *string `json:"j4r3"`
	J5r3     *string `json:"j5r3"`
	Speedr3  *string `json:"speedr3"`
	Distr3   *string `json:"distr3"`
	Disptsr3 *string `json:"disptsr3"`
	Judptsr3 *string `json:"judptsr3"`
	Totrun3  *string `json:"totrun3"`
	Posr3    *string `json:"posr3"`
	Statusr3 *string `json:"statusr3"`

	J1r4     *string `json:"j1r4"`
	J2r4     *string `json:"j2r4"`
	J3r4     *string `json:"j3r4"`
	J4r4     *string `json:"j4r4"`
	J5r4     *string `json:"j5r4"`
	Speedr4  *string `json:"speedr4"`
	Distr4   *string `json:"distr4"`
	Disptsr4 *string `json:"disptsr4"`
	Judptsr4 *string `json:"judptsr4"`

	Gater1    *string `json:"gater1"`
	Gater2    *string `json:"gater2"`
	Gater3    *string `json:"gater3"`
	Gater4    *string `json:"gater4"`
	Gateptsr1 *string `json:"gateptsr1"`
	Gateptsr2 *string `json:"gateptsr2"`
	Gateptsr3 *string `json:"gateptsr3"`
	Gateptsr4 *string `json:"gateptsr4"`

	Windr1    *string `json:"windr1"`
	Windr2    *string `json:"windr2"`
	Windr3    *string `json:"windr3"`
	Windr4    *string `json:"windr4"`
	Windptsr1 *string `json:"windptsr1"`
	Windptsr2 *string `json:"windptsr2"`
	Windptsr3 *string `json:"windptsr3"`
	Windptsr4 *string `json:"windptsr4"`

	Reason     *string `json:"reason"`
	Totrun4    *string `json:"totrun4"`
	Tot        *string `json:"tot"`
	Valid      *int32  `json:"valid"`
	Racepoints *string `json:"racepoints"`
	Cuppoints  *string `json:"cuppoints"`
	Version    *string `json:"version"`
	Lastupdate *string `json:"lastupdate"` // RFC3339
	Posr4      *string `json:"posr4"`
	Statusr4   *string `json:"statusr4"`
}

type UpdateResultJPInput = InsertResultJPInput

func mapInsertResultJPInput(in InsertResultJPInput) (fis.InsertResultJPClean, error) {
	var lup *time.Time
	var err error

	if in.Lastupdate != nil {
		lup, err = utils.ParseTimestampPtr(in.Lastupdate)
		if err != nil {
			return fis.InsertResultJPClean{}, err
		}
	}

	return fis.InsertResultJPClean{
		Recid:          in.Recid,
		Raceid:         in.Raceid,
		Competitorid:   in.Competitorid,
		Status:         in.Status,
		Status2:        in.Status2,
		Position:       in.Position,
		Bib:            in.Bib,
		Fiscode:        in.Fiscode,
		Competitorname: in.Competitorname,
		Nationcode:     in.Nationcode,
		Level:          in.Level,
		Heat:           in.Heat,
		Stage:          in.Stage,

		J1r1:     in.J1r1,
		J2r1:     in.J2r1,
		J3r1:     in.J3r1,
		J4r1:     in.J4r1,
		J5r1:     in.J5r1,
		Speedr1:  in.Speedr1,
		Distr1:   in.Distr1,
		Disptsr1: in.Disptsr1,
		Judptsr1: in.Judptsr1,
		Totrun1:  in.Totrun1,
		Posr1:    in.Posr1,
		Statusr1: in.Statusr1,

		J1r2:     in.J1r2,
		J2r2:     in.J2r2,
		J3r2:     in.J3r2,
		J4r2:     in.J4r2,
		J5r2:     in.J5r2,
		Speedr2:  in.Speedr2,
		Distr2:   in.Distr2,
		Disptsr2: in.Disptsr2,
		Judptsr2: in.Judptsr2,
		Totrun2:  in.Totrun2,
		Posr2:    in.Posr2,
		Statusr2: in.Statusr2,

		J1r3:     in.J1r3,
		J2r3:     in.J2r3,
		J3r3:     in.J3r3,
		J4r3:     in.J4r3,
		J5r3:     in.J5r3,
		Speedr3:  in.Speedr3,
		Distr3:   in.Distr3,
		Disptsr3: in.Disptsr3,
		Judptsr3: in.Judptsr3,
		Totrun3:  in.Totrun3,
		Posr3:    in.Posr3,
		Statusr3: in.Statusr3,

		J1r4:     in.J1r4,
		J2r4:     in.J2r4,
		J3r4:     in.J3r4,
		J4r4:     in.J4r4,
		J5r4:     in.J5r4,
		Speedr4:  in.Speedr4,
		Distr4:   in.Distr4,
		Disptsr4: in.Disptsr4,
		Judptsr4: in.Judptsr4,

		Gater1:    in.Gater1,
		Gater2:    in.Gater2,
		Gater3:    in.Gater3,
		Gater4:    in.Gater4,
		Gateptsr1: in.Gateptsr1,
		Gateptsr2: in.Gateptsr2,
		Gateptsr3: in.Gateptsr3,
		Gateptsr4: in.Gateptsr4,

		Windr1:    in.Windr1,
		Windr2:    in.Windr2,
		Windr3:    in.Windr3,
		Windr4:    in.Windr4,
		Windptsr1: in.Windptsr1,
		Windptsr2: in.Windptsr2,
		Windptsr3: in.Windptsr3,
		Windptsr4: in.Windptsr4,

		Reason:     in.Reason,
		Totrun4:    in.Totrun4,
		Tot:        in.Tot,
		Valid:      in.Valid,
		Racepoints: in.Racepoints,
		Cuppoints:  in.Cuppoints,
		Version:    in.Version,
		Lastupdate: lup,
		Posr4:      in.Posr4,
		Statusr4:   in.Statusr4,
	}, nil
}

func mapUpdateResultJPInput(in UpdateResultJPInput) (fis.UpdateResultJPClean, error) {
	clean, err := mapInsertResultJPInput(InsertResultJPInput(in))
	return fis.UpdateResultJPClean(clean), err
}

type FISResultJPFullResponse struct {
	Recid          int32   `json:"recid"`
	Raceid         *int32  `json:"raceid"`
	Competitorid   *int32  `json:"competitorid"`
	Status         *string `json:"status"`
	Status2        *string `json:"status2"`
	Position       *int32  `json:"position"`
	Bib            *int32  `json:"bib"`
	Fiscode        *int32  `json:"fiscode"`
	Competitorname *string `json:"competitorname"`
	Nationcode     *string `json:"nationcode"`
	Level          *string `json:"level"`
	Heat           *string `json:"heat"`
	Stage          *string `json:"stage"`

	J1r1     *string `json:"j1r1"`
	J2r1     *string `json:"j2r1"`
	J3r1     *string `json:"j3r1"`
	J4r1     *string `json:"j4r1"`
	J5r1     *string `json:"j5r1"`
	Speedr1  *string `json:"speedr1"`
	Distr1   *string `json:"distr1"`
	Disptsr1 *string `json:"disptsr1"`
	Judptsr1 *string `json:"judptsr1"`
	Totrun1  *string `json:"totrun1"`
	Posr1    *string `json:"posr1"`
	Statusr1 *string `json:"statusr1"`

	J1r2     *string `json:"j1r2"`
	J2r2     *string `json:"j2r2"`
	J3r2     *string `json:"j3r2"`
	J4r2     *string `json:"j4r2"`
	J5r2     *string `json:"j5r2"`
	Speedr2  *string `json:"speedr2"`
	Distr2   *string `json:"distr2"`
	Disptsr2 *string `json:"disptsr2"`
	Judptsr2 *string `json:"judptsr2"`
	Totrun2  *string `json:"totrun2"`
	Posr2    *string `json:"posr2"`
	Statusr2 *string `json:"statusr2"`

	J1r3     *string `json:"j1r3"`
	J2r3     *string `json:"j2r3"`
	J3r3     *string `json:"j3r3"`
	J4r3     *string `json:"j4r3"`
	J5r3     *string `json:"j5r3"`
	Speedr3  *string `json:"speedr3"`
	Distr3   *string `json:"distr3"`
	Disptsr3 *string `json:"disptsr3"`
	Judptsr3 *string `json:"judptsr3"`
	Totrun3  *string `json:"totrun3"`
	Posr3    *string `json:"posr3"`
	Statusr3 *string `json:"statusr3"`

	J1r4     *string `json:"j1r4"`
	J2r4     *string `json:"j2r4"`
	J3r4     *string `json:"j3r4"`
	J4r4     *string `json:"j4r4"`
	J5r4     *string `json:"j5r4"`
	Speedr4  *string `json:"speedr4"`
	Distr4   *string `json:"distr4"`
	Disptsr4 *string `json:"disptsr4"`
	Judptsr4 *string `json:"judptsr4"`

	Gater1    *string `json:"gater1"`
	Gater2    *string `json:"gater2"`
	Gater3    *string `json:"gater3"`
	Gater4    *string `json:"gater4"`
	Gateptsr1 *string `json:"gateptsr1"`
	Gateptsr2 *string `json:"gateptsr2"`
	Gateptsr3 *string `json:"gateptsr3"`
	Gateptsr4 *string `json:"gateptsr4"`

	Windr1    *string `json:"windr1"`
	Windr2    *string `json:"windr2"`
	Windr3    *string `json:"windr3"`
	Windr4    *string `json:"windr4"`
	Windptsr1 *string `json:"windptsr1"`
	Windptsr2 *string `json:"windptsr2"`
	Windptsr3 *string `json:"windptsr3"`
	Windptsr4 *string `json:"windptsr4"`

	Reason     *string `json:"reason"`
	Totrun4    *string `json:"totrun4"`
	Tot        *string `json:"tot"`
	Valid      *int32  `json:"valid"`
	Racepoints *string `json:"racepoints"`
	Cuppoints  *string `json:"cuppoints"`
	Version    *string `json:"version"`
	Lastupdate *string `json:"lastupdate"`
	Posr4      *string `json:"posr4"`
	Statusr4   *string `json:"statusr4"`
}

func FISResultJPFullFromSqlc(row fissqlc.AResultjp) FISResultJPFullResponse {
	var lastUpdateStr *string
	if row.Lastupdate.Valid {
		lastUpdateStr = utils.FormatTimestampPtr(row.Lastupdate)
	}

	return FISResultJPFullResponse{
		Recid:          row.Recid,
		Raceid:         utils.Int32PtrOrNil(row.Raceid),
		Competitorid:   utils.Int32PtrOrNil(row.Competitorid),
		Status:         utils.StringPtrOrNil(row.Status),
		Status2:        utils.StringPtrOrNil(row.Status2),
		Position:       utils.Int32PtrOrNil(row.Position),
		Bib:            utils.Int32PtrOrNil(row.Bib),
		Fiscode:        utils.Int32PtrOrNil(row.Fiscode),
		Competitorname: utils.StringPtrOrNil(row.Competitorname),
		Nationcode:     utils.StringPtrOrNil(row.Nationcode),
		Level:          utils.StringPtrOrNil(row.Level),
		Heat:           utils.StringPtrOrNil(row.Heat),
		Stage:          utils.StringPtrOrNil(row.Stage),

		J1r1:     utils.StringPtrOrNil(row.J1r1),
		J2r1:     utils.StringPtrOrNil(row.J2r1),
		J3r1:     utils.StringPtrOrNil(row.J3r1),
		J4r1:     utils.StringPtrOrNil(row.J4r1),
		J5r1:     utils.StringPtrOrNil(row.J5r1),
		Speedr1:  utils.StringPtrOrNil(row.Speedr1),
		Distr1:   utils.StringPtrOrNil(row.Distr1),
		Disptsr1: utils.StringPtrOrNil(row.Disptsr1),
		Judptsr1: utils.StringPtrOrNil(row.Judptsr1),
		Totrun1:  utils.StringPtrOrNil(row.Totrun1),
		Posr1:    utils.StringPtrOrNil(row.Posr1),
		Statusr1: utils.StringPtrOrNil(row.Statusr1),

		J1r2:     utils.StringPtrOrNil(row.J1r2),
		J2r2:     utils.StringPtrOrNil(row.J2r2),
		J3r2:     utils.StringPtrOrNil(row.J3r2),
		J4r2:     utils.StringPtrOrNil(row.J4r2),
		J5r2:     utils.StringPtrOrNil(row.J5r2),
		Speedr2:  utils.StringPtrOrNil(row.Speedr2),
		Distr2:   utils.StringPtrOrNil(row.Distr2),
		Disptsr2: utils.StringPtrOrNil(row.Disptsr2),
		Judptsr2: utils.StringPtrOrNil(row.Judptsr2),
		Totrun2:  utils.StringPtrOrNil(row.Totrun2),
		Posr2:    utils.StringPtrOrNil(row.Posr2),
		Statusr2: utils.StringPtrOrNil(row.Statusr2),

		J1r3:     utils.StringPtrOrNil(row.J1r3),
		J2r3:     utils.StringPtrOrNil(row.J2r3),
		J3r3:     utils.StringPtrOrNil(row.J3r3),
		J4r3:     utils.StringPtrOrNil(row.J4r3),
		J5r3:     utils.StringPtrOrNil(row.J5r3),
		Speedr3:  utils.StringPtrOrNil(row.Speedr3),
		Distr3:   utils.StringPtrOrNil(row.Distr3),
		Disptsr3: utils.StringPtrOrNil(row.Disptsr3),
		Judptsr3: utils.StringPtrOrNil(row.Judptsr3),
		Totrun3:  utils.StringPtrOrNil(row.Totrun3),
		Posr3:    utils.StringPtrOrNil(row.Posr3),
		Statusr3: utils.StringPtrOrNil(row.Statusr3),

		J1r4:     utils.StringPtrOrNil(row.J1r4),
		J2r4:     utils.StringPtrOrNil(row.J2r4),
		J3r4:     utils.StringPtrOrNil(row.J3r4),
		J4r4:     utils.StringPtrOrNil(row.J4r4),
		J5r4:     utils.StringPtrOrNil(row.J5r4),
		Speedr4:  utils.StringPtrOrNil(row.Speedr4),
		Distr4:   utils.StringPtrOrNil(row.Distr4),
		Disptsr4: utils.StringPtrOrNil(row.Disptsr4),
		Judptsr4: utils.StringPtrOrNil(row.Judptsr4),

		Gater1:    utils.StringPtrOrNil(row.Gater1),
		Gater2:    utils.StringPtrOrNil(row.Gater2),
		Gater3:    utils.StringPtrOrNil(row.Gater3),
		Gater4:    utils.StringPtrOrNil(row.Gater4),
		Gateptsr1: utils.StringPtrOrNil(row.Gateptsr1),
		Gateptsr2: utils.StringPtrOrNil(row.Gateptsr2),
		Gateptsr3: utils.StringPtrOrNil(row.Gateptsr3),
		Gateptsr4: utils.StringPtrOrNil(row.Gateptsr4),

		Windr1:    utils.StringPtrOrNil(row.Windr1),
		Windr2:    utils.StringPtrOrNil(row.Windr2),
		Windr3:    utils.StringPtrOrNil(row.Windr3),
		Windr4:    utils.StringPtrOrNil(row.Windr4),
		Windptsr1: utils.StringPtrOrNil(row.Windptsr1),
		Windptsr2: utils.StringPtrOrNil(row.Windptsr2),
		Windptsr3: utils.StringPtrOrNil(row.Windptsr3),
		Windptsr4: utils.StringPtrOrNil(row.Windptsr4),

		Reason:     utils.StringPtrOrNil(row.Reason),
		Totrun4:    utils.StringPtrOrNil(row.Totrun4),
		Tot:        utils.StringPtrOrNil(row.Tot),
		Valid:      utils.Int32PtrOrNil(row.Valid),
		Racepoints: utils.StringPtrOrNil(row.Racepoints),
		Cuppoints:  utils.StringPtrOrNil(row.Cuppoints),
		Version:    utils.StringPtrOrNil(row.Version),
		Lastupdate: lastUpdateStr,
		Posr4:      utils.StringPtrOrNil(row.Posr4),
		Statusr4:   utils.StringPtrOrNil(row.Statusr4),
	}
}

type FISAthleteResultJPRow struct {
	Raceid         *int32  `json:"raceid"`
	Position       *int32  `json:"position"`
	Racedate       *string `json:"racedate"`
	Seasoncode     *int32  `json:"seasoncode"`
	Disciplinecode *string `json:"disciplinecode"`
	Catcode        *string `json:"catcode"`
	Place          *string `json:"place"`

	Posr1     *string `json:"posr1"`
	Speedr1   *string `json:"speedr1"`
	Distr1    *string `json:"distr1"`
	Judptsr1  *string `json:"judptsr1"`
	Windr1    *string `json:"windr1"`
	Windptsr1 *string `json:"windptsr1"`
	Gater1    *string `json:"gater1"`

	Posr2     *string `json:"posr2"`
	Speedr2   *string `json:"speedr2"`
	Distr2    *string `json:"distr2"`
	Judptsr2  *string `json:"judptsr2"`
	Windr2    *string `json:"windr2"`
	Windptsr2 *string `json:"windptsr2"`
	Gater2    *string `json:"gater2"`

	Totrun1 *string `json:"totrun1"`
	Totrun2 *string `json:"totrun2"`
}

func FISAthleteResultJPFromSqlc(row fissqlc.GetAthleteResultsJPRow) FISAthleteResultJPRow {
	var raceDateStr *string
	if row.Racedate.Valid {
		raceDateStr = utils.FormatDatePtr(row.Racedate)
	}

	return FISAthleteResultJPRow{
		Raceid:         utils.Int32PtrOrNil(row.Raceid),
		Position:       utils.Int32PtrOrNil(row.Position),
		Racedate:       raceDateStr,
		Seasoncode:     utils.Int32PtrOrNil(row.Seasoncode),
		Disciplinecode: utils.StringPtrOrNil(row.Disciplinecode),
		Catcode:        utils.StringPtrOrNil(row.Catcode),
		Place:          utils.StringPtrOrNil(row.Place),

		Posr1:     utils.StringPtrOrNil(row.Posr1),
		Speedr1:   utils.StringPtrOrNil(row.Speedr1),
		Distr1:    utils.StringPtrOrNil(row.Distr1),
		Judptsr1:  utils.StringPtrOrNil(row.Judptsr1),
		Windr1:    utils.StringPtrOrNil(row.Windr1),
		Windptsr1: utils.StringPtrOrNil(row.Windptsr1),
		Gater1:    utils.StringPtrOrNil(row.Gater1),

		Posr2:     utils.StringPtrOrNil(row.Posr2),
		Speedr2:   utils.StringPtrOrNil(row.Speedr2),
		Distr2:    utils.StringPtrOrNil(row.Distr2),
		Judptsr2:  utils.StringPtrOrNil(row.Judptsr2),
		Windr2:    utils.StringPtrOrNil(row.Windr2),
		Windptsr2: utils.StringPtrOrNil(row.Windptsr2),
		Gater2:    utils.StringPtrOrNil(row.Gater2),

		Totrun1: utils.StringPtrOrNil(row.Totrun1),
		Totrun2: utils.StringPtrOrNil(row.Totrun2),
	}
}
