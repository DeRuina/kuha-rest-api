package fis

import (
	"time"

	fissqlc "github.com/DeRuina/KUHA-REST-API/internal/db/fis"
	"github.com/DeRuina/KUHA-REST-API/internal/utils"
)

// Clean structs for INSERT/UPDATE on a_resultcc

type InsertResultCCClean struct {
	Recid          int32
	Raceid         *int32
	Competitorid   *int32
	Status         *string
	Reason         *string
	Position       *string // numeric(10,5) in DB, but handled as string
	Pf             *int32
	Status2        *string
	Bib            *string // numeric(10,5) in DB, but handled as string
	Bibcolor       *string
	Fiscode        *int32
	Competitorname *string
	Nationcode     *string
	Stage          *string
	Level          *string
	Heat           *string
	Timer1         *string
	Timer2         *string
	Timer3         *string
	Timetot        *string
	Valid          *string // numeric(10,5) in DB, but handled as string
	Racepoints     *string // numeric(10,5) in DB, but handled as string
	Cuppoints      *string // numeric(10,5) in DB, but handled as string
	Bonustime      *string
	Bonuscuppoints *string
	Version        *string
	Rg1            *string
	Rg2            *string
	Lastupdate     *time.Time
}

type UpdateResultCCClean = InsertResultCCClean

func mapInsertResultCCToParams(in InsertResultCCClean) fissqlc.InsertResultCCParams {
	return fissqlc.InsertResultCCParams{
		Recid:          in.Recid,
		Raceid:         utils.NullInt32Ptr(in.Raceid),
		Competitorid:   utils.NullInt32Ptr(in.Competitorid),
		Status:         utils.NullStringPtr(in.Status),
		Reason:         utils.NullStringPtr(in.Reason),
		Position:       utils.NullStringPtr(in.Position),
		Pf:             utils.NullInt32Ptr(in.Pf),
		Status2:        utils.NullStringPtr(in.Status2),
		Bib:            utils.NullStringPtr(in.Bib),
		Bibcolor:       utils.NullStringPtr(in.Bibcolor),
		Fiscode:        utils.NullInt32Ptr(in.Fiscode),
		Competitorname: utils.NullStringPtr(in.Competitorname),
		Nationcode:     utils.NullStringPtr(in.Nationcode),
		Stage:          utils.NullStringPtr(in.Stage),
		Level:          utils.NullStringPtr(in.Level),
		Heat:           utils.NullStringPtr(in.Heat),
		Timer1:         utils.NullStringPtr(in.Timer1),
		Timer2:         utils.NullStringPtr(in.Timer2),
		Timer3:         utils.NullStringPtr(in.Timer3),
		Timetot:        utils.NullStringPtr(in.Timetot),
		Valid:          utils.NullStringPtr(in.Valid),
		Racepoints:     utils.NullStringPtr(in.Racepoints),
		Cuppoints:      utils.NullStringPtr(in.Cuppoints),
		Bonustime:      utils.NullStringPtr(in.Bonustime),
		Bonuscuppoints: utils.NullStringPtr(in.Bonuscuppoints),
		Version:        utils.NullStringPtr(in.Version),
		Rg1:            utils.NullStringPtr(in.Rg1),
		Rg2:            utils.NullStringPtr(in.Rg2),
		Lastupdate:     utils.NullTimePtr(in.Lastupdate),
	}
}

func mapUpdateResultCCToParams(in UpdateResultCCClean) fissqlc.UpdateResultCCByRecIDParams {
	p := mapInsertResultCCToParams(InsertResultCCClean(in))
	return fissqlc.UpdateResultCCByRecIDParams{
		Recid:          p.Recid,
		Raceid:         p.Raceid,
		Competitorid:   p.Competitorid,
		Status:         p.Status,
		Reason:         p.Reason,
		Position:       p.Position,
		Pf:             p.Pf,
		Status2:        p.Status2,
		Bib:            p.Bib,
		Bibcolor:       p.Bibcolor,
		Fiscode:        p.Fiscode,
		Competitorname: p.Competitorname,
		Nationcode:     p.Nationcode,
		Stage:          p.Stage,
		Level:          p.Level,
		Heat:           p.Heat,
		Timer1:         p.Timer1,
		Timer2:         p.Timer2,
		Timer3:         p.Timer3,
		Timetot:        p.Timetot,
		Valid:          p.Valid,
		Racepoints:     p.Racepoints,
		Cuppoints:      p.Cuppoints,
		Bonustime:      p.Bonustime,
		Bonuscuppoints: p.Bonuscuppoints,
		Version:        p.Version,
		Rg1:            p.Rg1,
		Rg2:            p.Rg2,
		Lastupdate:     p.Lastupdate,
	}
}

type InsertResultJPClean struct {
	Recid          int32
	Raceid         *int32
	Competitorid   *int32
	Status         *string
	Status2        *string
	Position       *int32
	Bib            *int32
	Fiscode        *int32
	Competitorname *string
	Nationcode     *string
	Level          *string
	Heat           *string
	Stage          *string

	J1r1     *string
	J2r1     *string
	J3r1     *string
	J4r1     *string
	J5r1     *string
	Speedr1  *string
	Distr1   *string
	Disptsr1 *string
	Judptsr1 *string
	Totrun1  *string
	Posr1    *string
	Statusr1 *string

	J1r2     *string
	J2r2     *string
	J3r2     *string
	J4r2     *string
	J5r2     *string
	Speedr2  *string
	Distr2   *string
	Disptsr2 *string
	Judptsr2 *string
	Totrun2  *string
	Posr2    *string
	Statusr2 *string

	J1r3     *string
	J2r3     *string
	J3r3     *string
	J4r3     *string
	J5r3     *string
	Speedr3  *string
	Distr3   *string
	Disptsr3 *string
	Judptsr3 *string
	Totrun3  *string
	Posr3    *string
	Statusr3 *string

	J1r4     *string
	J2r4     *string
	J3r4     *string
	J4r4     *string
	J5r4     *string
	Speedr4  *string
	Distr4   *string
	Disptsr4 *string
	Judptsr4 *string

	Gater1    *string
	Gater2    *string
	Gater3    *string
	Gater4    *string
	Gateptsr1 *string
	Gateptsr2 *string
	Gateptsr3 *string
	Gateptsr4 *string

	Windr1    *string
	Windr2    *string
	Windr3    *string
	Windr4    *string
	Windptsr1 *string
	Windptsr2 *string
	Windptsr3 *string
	Windptsr4 *string

	Reason     *string
	Totrun4    *string
	Tot        *string
	Valid      *int32
	Racepoints *string
	Cuppoints  *string
	Version    *string
	Lastupdate *time.Time
	Posr4      *string
	Statusr4   *string
}

type UpdateResultJPClean = InsertResultJPClean

func mapInsertResultJPToParams(in InsertResultJPClean) fissqlc.InsertResultJPParams {
	return fissqlc.InsertResultJPParams{
		Recid:          in.Recid,
		Raceid:         utils.NullInt32Ptr(in.Raceid),
		Competitorid:   utils.NullInt32Ptr(in.Competitorid),
		Status:         utils.NullStringPtr(in.Status),
		Status2:        utils.NullStringPtr(in.Status2),
		Position:       utils.NullInt32Ptr(in.Position),
		Bib:            utils.NullInt32Ptr(in.Bib),
		Fiscode:        utils.NullInt32Ptr(in.Fiscode),
		Competitorname: utils.NullStringPtr(in.Competitorname),
		Nationcode:     utils.NullStringPtr(in.Nationcode),
		Level:          utils.NullStringPtr(in.Level),
		Heat:           utils.NullStringPtr(in.Heat),
		Stage:          utils.NullStringPtr(in.Stage),

		J1r1:     utils.NullStringPtr(in.J1r1),
		J2r1:     utils.NullStringPtr(in.J2r1),
		J3r1:     utils.NullStringPtr(in.J3r1),
		J4r1:     utils.NullStringPtr(in.J4r1),
		J5r1:     utils.NullStringPtr(in.J5r1),
		Speedr1:  utils.NullStringPtr(in.Speedr1),
		Distr1:   utils.NullStringPtr(in.Distr1),
		Disptsr1: utils.NullStringPtr(in.Disptsr1),
		Judptsr1: utils.NullStringPtr(in.Judptsr1),
		Totrun1:  utils.NullStringPtr(in.Totrun1),
		Posr1:    utils.NullStringPtr(in.Posr1),
		Statusr1: utils.NullStringPtr(in.Statusr1),

		J1r2:     utils.NullStringPtr(in.J1r2),
		J2r2:     utils.NullStringPtr(in.J2r2),
		J3r2:     utils.NullStringPtr(in.J3r2),
		J4r2:     utils.NullStringPtr(in.J4r2),
		J5r2:     utils.NullStringPtr(in.J5r2),
		Speedr2:  utils.NullStringPtr(in.Speedr2),
		Distr2:   utils.NullStringPtr(in.Distr2),
		Disptsr2: utils.NullStringPtr(in.Disptsr2),
		Judptsr2: utils.NullStringPtr(in.Judptsr2),
		Totrun2:  utils.NullStringPtr(in.Totrun2),
		Posr2:    utils.NullStringPtr(in.Posr2),
		Statusr2: utils.NullStringPtr(in.Statusr2),

		J1r3:     utils.NullStringPtr(in.J1r3),
		J2r3:     utils.NullStringPtr(in.J2r3),
		J3r3:     utils.NullStringPtr(in.J3r3),
		J4r3:     utils.NullStringPtr(in.J4r3),
		J5r3:     utils.NullStringPtr(in.J5r3),
		Speedr3:  utils.NullStringPtr(in.Speedr3),
		Distr3:   utils.NullStringPtr(in.Distr3),
		Disptsr3: utils.NullStringPtr(in.Disptsr3),
		Judptsr3: utils.NullStringPtr(in.Judptsr3),
		Totrun3:  utils.NullStringPtr(in.Totrun3),
		Posr3:    utils.NullStringPtr(in.Posr3),
		Statusr3: utils.NullStringPtr(in.Statusr3),

		J1r4:     utils.NullStringPtr(in.J1r4),
		J2r4:     utils.NullStringPtr(in.J2r4),
		J3r4:     utils.NullStringPtr(in.J3r4),
		J4r4:     utils.NullStringPtr(in.J4r4),
		J5r4:     utils.NullStringPtr(in.J5r4),
		Speedr4:  utils.NullStringPtr(in.Speedr4),
		Distr4:   utils.NullStringPtr(in.Distr4),
		Disptsr4: utils.NullStringPtr(in.Disptsr4),
		Judptsr4: utils.NullStringPtr(in.Judptsr4),

		Gater1:    utils.NullStringPtr(in.Gater1),
		Gater2:    utils.NullStringPtr(in.Gater2),
		Gater3:    utils.NullStringPtr(in.Gater3),
		Gater4:    utils.NullStringPtr(in.Gater4),
		Gateptsr1: utils.NullStringPtr(in.Gateptsr1),
		Gateptsr2: utils.NullStringPtr(in.Gateptsr2),
		Gateptsr3: utils.NullStringPtr(in.Gateptsr3),
		Gateptsr4: utils.NullStringPtr(in.Gateptsr4),

		Windr1:    utils.NullStringPtr(in.Windr1),
		Windr2:    utils.NullStringPtr(in.Windr2),
		Windr3:    utils.NullStringPtr(in.Windr3),
		Windr4:    utils.NullStringPtr(in.Windr4),
		Windptsr1: utils.NullStringPtr(in.Windptsr1),
		Windptsr2: utils.NullStringPtr(in.Windptsr2),
		Windptsr3: utils.NullStringPtr(in.Windptsr3),
		Windptsr4: utils.NullStringPtr(in.Windptsr4),

		Reason:     utils.NullStringPtr(in.Reason),
		Totrun4:    utils.NullStringPtr(in.Totrun4),
		Tot:        utils.NullStringPtr(in.Tot),
		Valid:      utils.NullInt32Ptr(in.Valid),
		Racepoints: utils.NullStringPtr(in.Racepoints),
		Cuppoints:  utils.NullStringPtr(in.Cuppoints),
		Version:    utils.NullStringPtr(in.Version),
		Lastupdate: utils.NullTimePtr(in.Lastupdate),
		Posr4:      utils.NullStringPtr(in.Posr4),
		Statusr4:   utils.NullStringPtr(in.Statusr4),
	}
}

func mapUpdateResultJPToParams(in UpdateResultJPClean) fissqlc.UpdateResultJPByRecIDParams {
	p := mapInsertResultJPToParams(InsertResultJPClean(in))
	return fissqlc.UpdateResultJPByRecIDParams{
		Recid:          p.Recid,
		Raceid:         p.Raceid,
		Competitorid:   p.Competitorid,
		Status:         p.Status,
		Status2:        p.Status2,
		Position:       p.Position,
		Bib:            p.Bib,
		Fiscode:        p.Fiscode,
		Competitorname: p.Competitorname,
		Nationcode:     p.Nationcode,
		Level:          p.Level,
		Heat:           p.Heat,
		Stage:          p.Stage,
		J1r1:           p.J1r1,
		J2r1:           p.J2r1,
		J3r1:           p.J3r1,
		J4r1:           p.J4r1,
		J5r1:           p.J5r1,
		Speedr1:        p.Speedr1,
		Distr1:         p.Distr1,
		Disptsr1:       p.Disptsr1,
		Judptsr1:       p.Judptsr1,
		Totrun1:        p.Totrun1,
		Posr1:          p.Posr1,
		Statusr1:       p.Statusr1,
		J1r2:           p.J1r2,
		J2r2:           p.J2r2,
		J3r2:           p.J3r2,
		J4r2:           p.J4r2,
		J5r2:           p.J5r2,
		Speedr2:        p.Speedr2,
		Distr2:         p.Distr2,
		Disptsr2:       p.Disptsr2,
		Judptsr2:       p.Judptsr2,
		Totrun2:        p.Totrun2,
		Posr2:          p.Posr2,
		Statusr2:       p.Statusr2,
		J1r3:           p.J1r3,
		J2r3:           p.J2r3,
		J3r3:           p.J3r3,
		J4r3:           p.J4r3,
		J5r3:           p.J5r3,
		Speedr3:        p.Speedr3,
		Distr3:         p.Distr3,
		Disptsr3:       p.Disptsr3,
		Judptsr3:       p.Judptsr3,
		Totrun3:        p.Totrun3,
		Posr3:          p.Posr3,
		Statusr3:       p.Statusr3,
		J1r4:           p.J1r4,
		J2r4:           p.J2r4,
		J3r4:           p.J3r4,
		J4r4:           p.J4r4,
		J5r4:           p.J5r4,
		Speedr4:        p.Speedr4,
		Distr4:         p.Distr4,
		Disptsr4:       p.Disptsr4,
		Judptsr4:       p.Judptsr4,
		Gater1:         p.Gater1,
		Gater2:         p.Gater2,
		Gater3:         p.Gater3,
		Gater4:         p.Gater4,
		Gateptsr1:      p.Gateptsr1,
		Gateptsr2:      p.Gateptsr2,
		Gateptsr3:      p.Gateptsr3,
		Gateptsr4:      p.Gateptsr4,
		Windr1:         p.Windr1,
		Windr2:         p.Windr2,
		Windr3:         p.Windr3,
		Windr4:         p.Windr4,
		Windptsr1:      p.Windptsr1,
		Windptsr2:      p.Windptsr2,
		Windptsr3:      p.Windptsr3,
		Windptsr4:      p.Windptsr4,
		Reason:         p.Reason,
		Totrun4:        p.Totrun4,
		Tot:            p.Tot,
		Valid:          p.Valid,
		Racepoints:     p.Racepoints,
		Cuppoints:      p.Cuppoints,
		Version:        p.Version,
		Lastupdate:     p.Lastupdate,
		Posr4:          p.Posr4,
		Statusr4:       p.Statusr4,
	}
}
