package fisapi

import (
	"time"

	fissqlc "github.com/DeRuina/KUHA-REST-API/internal/db/fis"
	"github.com/DeRuina/KUHA-REST-API/internal/store/fis"
	"github.com/DeRuina/KUHA-REST-API/internal/utils"
)

type RacesCCQuery struct {
	SeasonCode     []string `form:"seasoncode"`
	DisciplineCode []string `form:"disciplinecode"`
	CatCode        []string `form:"catcode"`
}

type InsertRaceCCInput struct {
	Raceid            int32   `json:"raceid" validate:"required"`
	Eventid           *int32  `json:"eventid"`
	Seasoncode        *int32  `json:"seasoncode"`
	Racecodex         *int32  `json:"racecodex"`
	Disciplineid      *string `json:"disciplineid"`
	Disciplinecode    *string `json:"disciplinecode"`
	Catcode           *string `json:"catcode"`
	Catcode2          *string `json:"catcode2"`
	Catcode3          *string `json:"catcode3"`
	Catcode4          *string `json:"catcode4"`
	Gender            *string `json:"gender"`
	Racedate          *string `json:"racedate"`       // YYYY-MM-DD
	Starteventdate    *string `json:"starteventdate"` // YYYY-MM-DD
	Description       *string `json:"description"`
	Place             *string `json:"place"`
	Nationcode        *string `json:"nationcode"`
	Receiveddate      *string `json:"receiveddate"` // RFC3339 or date
	Validdate         *string `json:"validdate"`    // RFC3339 or date
	Td1id             *int32  `json:"td1id"`
	Td1name           *string `json:"td1name"`
	Td1nation         *string `json:"td1nation"`
	Td1code           *int32  `json:"td1code"`
	Td2id             *int32  `json:"td2id"`
	Td2name           *string `json:"td2name"`
	Td2nation         *string `json:"td2nation"`
	Td2code           *int32  `json:"td2code"`
	Calstatuscode     *string `json:"calstatuscode"`
	Procstatuscode    *string `json:"procstatuscode"`
	Displaystatus     *string `json:"displaystatus"`
	Fisinterncomment  *string `json:"fisinterncomment"`
	Webcomment        *string `json:"webcomment"`
	Pursuit           *string `json:"pursuit"`
	Masse             *string `json:"masse"`
	Relay             *string `json:"relay"`
	Distance          *string `json:"distance"`
	Hill              *string `json:"hill"`
	Style             *string `json:"style"`
	Qualif            *string `json:"qualif"`
	Finale            *string `json:"finale"`
	Homol             *string `json:"homol"`
	Published         *int32  `json:"published"`
	Validforfispoints *int32  `json:"validforfispoints"`
	Usedfislist       *string `json:"usedfislist"`
	Tolist            *string `json:"tolist"`
	Discforlistcode   *string `json:"discforlistcode"`
	Calculatedpenalty *string `json:"calculatedpenalty"`
	Appliedpenalty    *string `json:"appliedpenalty"`
	Appliedscala      *string `json:"appliedscala"`
	Penscafixed       *string `json:"penscafixed"`
	Version           *int32  `json:"version"`
	Nationraceid      *int32  `json:"nationraceid"`
	Provraceid        *int32  `json:"provraceid"`
	Msql7evid         *int32  `json:"msql7evid"`
	Mssql7id          *int32  `json:"mssql7id"`
	Topbanner         *string `json:"topbanner"`
	Bottombanner      *string `json:"bottombanner"`
	Toplogo           *string `json:"toplogo"`
	Bottomlogo        *string `json:"bottomlogo"`
	Gallery           *string `json:"gallery"`
	Indi              *int32  `json:"indi"`
	Team              *int32  `json:"team"`
	Tabcount          *int32  `json:"tabcount"`
	Columncount       *int32  `json:"columncount"`
	Level             *string `json:"level"`
	Hloc1             *string `json:"hloc1"`
	Hloc2             *string `json:"hloc2"`
	Hloc3             *string `json:"hloc3"`
	Hcet1             *string `json:"hcet1"`
	Hcet2             *string `json:"hcet2"`
	Hcet3             *string `json:"hcet3"`
	Live              *int32  `json:"live"`
	Livestatus1       *string `json:"livestatus1"`
	Livestatus2       *string `json:"livestatus2"`
	Livestatus3       *string `json:"livestatus3"`
	Liveinfo1         *string `json:"liveinfo1"`
	Liveinfo2         *string `json:"liveinfo2"`
	Liveinfo3         *string `json:"liveinfo3"`
	Passwd            *string `json:"passwd"`
	Timinglogo        *string `json:"timinglogo"`
	Results           *int32  `json:"results"`
	Pdf               *int32  `json:"pdf"`
	Noepr             *int32  `json:"noepr"`
	Tddoc             *int32  `json:"tddoc"`
	Timingreport      *int32  `json:"timingreport"`
	SpecialCupPoints  *int32  `json:"special_cup_points"`
	SkipWcsl          *int32  `json:"skip_wcsl"`
	Validforowg       *int32  `json:"validforowg"`
	Lastupdate        *string `json:"lastupdate"` // RFC3339
}

type UpdateRaceCCInput = InsertRaceCCInput

func mapInsertRaceCCInput(in InsertRaceCCInput) (fis.InsertRaceCCClean, error) {
	var rd, sed, recv, vald, lup *time.Time
	var err error
	if in.Racedate != nil {
		rd, err = utils.ParseDatePtr(in.Racedate)
		if err != nil {
			return fis.InsertRaceCCClean{}, err
		}
	}
	if in.Starteventdate != nil {
		sed, err = utils.ParseDatePtr(in.Starteventdate)
		if err != nil {
			return fis.InsertRaceCCClean{}, err
		}
	}
	if in.Receiveddate != nil {
		recv, err = utils.ParseTimestampPtr(in.Receiveddate)
		if err != nil {
			return fis.InsertRaceCCClean{}, err
		}
	}
	if in.Validdate != nil {
		vald, err = utils.ParseTimestampPtr(in.Validdate)
		if err != nil {
			return fis.InsertRaceCCClean{}, err
		}
	}
	if in.Lastupdate != nil {
		lup, err = utils.ParseTimestampPtr(in.Lastupdate)
		if err != nil {
			return fis.InsertRaceCCClean{}, err
		}
	}
	return fis.InsertRaceCCClean{
		Raceid:            in.Raceid,
		Eventid:           in.Eventid,
		Seasoncode:        in.Seasoncode,
		Racecodex:         in.Racecodex,
		Disciplineid:      in.Disciplineid,
		Disciplinecode:    in.Disciplinecode,
		Catcode:           in.Catcode,
		Catcode2:          in.Catcode2,
		Catcode3:          in.Catcode3,
		Catcode4:          in.Catcode4,
		Gender:            in.Gender,
		Racedate:          rd,
		Starteventdate:    sed,
		Description:       in.Description,
		Place:             in.Place,
		Nationcode:        in.Nationcode,
		Receiveddate:      recv,
		Validdate:         vald,
		Td1id:             in.Td1id,
		Td1name:           in.Td1name,
		Td1nation:         in.Td1nation,
		Td1code:           in.Td1code,
		Td2id:             in.Td2id,
		Td2name:           in.Td2name,
		Td2nation:         in.Td2nation,
		Td2code:           in.Td2code,
		Calstatuscode:     in.Calstatuscode,
		Procstatuscode:    in.Procstatuscode,
		Displaystatus:     in.Displaystatus,
		Fisinterncomment:  in.Fisinterncomment,
		Webcomment:        in.Webcomment,
		Pursuit:           in.Pursuit,
		Masse:             in.Masse,
		Relay:             in.Relay,
		Distance:          in.Distance,
		Hill:              in.Hill,
		Style:             in.Style,
		Qualif:            in.Qualif,
		Finale:            in.Finale,
		Homol:             in.Homol,
		Published:         in.Published,
		Validforfispoints: in.Validforfispoints,
		Usedfislist:       in.Usedfislist,
		Tolist:            in.Tolist,
		Discforlistcode:   in.Discforlistcode,
		Calculatedpenalty: in.Calculatedpenalty,
		Appliedpenalty:    in.Appliedpenalty,
		Appliedscala:      in.Appliedscala,
		Penscafixed:       in.Penscafixed,
		Version:           in.Version,
		Nationraceid:      in.Nationraceid,
		Provraceid:        in.Provraceid,
		Msql7evid:         in.Msql7evid,
		Mssql7id:          in.Mssql7id,
		Topbanner:         in.Topbanner,
		Bottombanner:      in.Bottombanner,
		Toplogo:           in.Toplogo,
		Bottomlogo:        in.Bottomlogo,
		Gallery:           in.Gallery,
		Indi:              in.Indi,
		Team:              in.Team,
		Tabcount:          in.Tabcount,
		Columncount:       in.Columncount,
		Level:             in.Level,
		Hloc1:             in.Hloc1,
		Hloc2:             in.Hloc2,
		Hloc3:             in.Hloc3,
		Hcet1:             in.Hcet1,
		Hcet2:             in.Hcet2,
		Hcet3:             in.Hcet3,
		Live:              in.Live,
		Livestatus1:       in.Livestatus1,
		Livestatus2:       in.Livestatus2,
		Livestatus3:       in.Livestatus3,
		Liveinfo1:         in.Liveinfo1,
		Liveinfo2:         in.Liveinfo2,
		Liveinfo3:         in.Liveinfo3,
		Passwd:            in.Passwd,
		Timinglogo:        in.Timinglogo,
		Results:           in.Results,
		Pdf:               in.Pdf,
		Noepr:             in.Noepr,
		Tddoc:             in.Tddoc,
		Timingreport:      in.Timingreport,
		SpecialCupPoints:  in.SpecialCupPoints,
		SkipWcsl:          in.SkipWcsl,
		Validforowg:       in.Validforowg,
		Lastupdate:        lup,
	}, nil
}

func mapUpdateRaceCCInput(in UpdateRaceCCInput) (fis.UpdateRaceCCClean, error) {
	// same parser as insert
	clean, err := mapInsertRaceCCInput(InsertRaceCCInput(in))
	return fis.UpdateRaceCCClean(clean), err
}

type FISRaceCCFullResponse struct {
	Raceid            int32   `json:"raceid"`
	Eventid           *int32  `json:"eventid"`
	Seasoncode        *int32  `json:"seasoncode"`
	Racecodex         *int32  `json:"racecodex"`
	Disciplineid      *string `json:"disciplineid"`
	Disciplinecode    *string `json:"disciplinecode"`
	Catcode           *string `json:"catcode"`
	Catcode2          *string `json:"catcode2"`
	Catcode3          *string `json:"catcode3"`
	Catcode4          *string `json:"catcode4"`
	Gender            *string `json:"gender"`
	Racedate          *string `json:"racedate"`
	Starteventdate    *string `json:"starteventdate"`
	Description       *string `json:"description"`
	Place             *string `json:"place"`
	Nationcode        *string `json:"nationcode"`
	Td1id             *int32  `json:"td1id"`
	Td1name           *string `json:"td1name"`
	Td1nation         *string `json:"td1nation"`
	Td1code           *int32  `json:"td1code"`
	Td2id             *int32  `json:"td2id"`
	Td2name           *string `json:"td2name"`
	Td2nation         *string `json:"td2nation"`
	Td2code           *int32  `json:"td2code"`
	Calstatuscode     *string `json:"calstatuscode"`
	Procstatuscode    *string `json:"procstatuscode"`
	Receiveddate      *string `json:"receiveddate"`
	Pursuit           *string `json:"pursuit"`
	Masse             *string `json:"masse"`
	Relay             *string `json:"relay"`
	Distance          *string `json:"distance"`
	Hill              *string `json:"hill"`
	Style             *string `json:"style"`
	Qualif            *string `json:"qualif"`
	Finale            *string `json:"finale"`
	Homol             *string `json:"homol"`
	Webcomment        *string `json:"webcomment"`
	Displaystatus     *string `json:"displaystatus"`
	Fisinterncomment  *string `json:"fisinterncomment"`
	Published         *int32  `json:"published"`
	Validforfispoints *int32  `json:"validforfispoints"`
	Usedfislist       *string `json:"usedfislist"`
	Tolist            *string `json:"tolist"`
	Discforlistcode   *string `json:"discforlistcode"`
	Calculatedpenalty *string `json:"calculatedpenalty"`
	Appliedpenalty    *string `json:"appliedpenalty"`
	Appliedscala      *string `json:"appliedscala"`
	Penscafixed       *string `json:"penscafixed"`
	Version           *int32  `json:"version"`
	Nationraceid      *int32  `json:"nationraceid"`
	Provraceid        *int32  `json:"provraceid"`
	Msql7evid         *int32  `json:"msql7evid"`
	Mssql7id          *int32  `json:"mssql7id"`
	Results           *int32  `json:"results"`
	Pdf               *int32  `json:"pdf"`
	Topbanner         *string `json:"topbanner"`
	Bottombanner      *string `json:"bottombanner"`
	Toplogo           *string `json:"toplogo"`
	Bottomlogo        *string `json:"bottomlogo"`
	Gallery           *string `json:"gallery"`
	Indi              *int32  `json:"indi"`
	Team              *int32  `json:"team"`
	Tabcount          *int32  `json:"tabcount"`
	Columncount       *int32  `json:"columncount"`
	Level             *string `json:"level"`
	Hloc1             *string `json:"hloc1"`
	Hloc2             *string `json:"hloc2"`
	Hloc3             *string `json:"hloc3"`
	Hcet1             *string `json:"hcet1"`
	Hcet2             *string `json:"hcet2"`
	Hcet3             *string `json:"hcet3"`
	Live              *int32  `json:"live"`
	Livestatus1       *string `json:"livestatus1"`
	Livestatus2       *string `json:"livestatus2"`
	Livestatus3       *string `json:"livestatus3"`
	Liveinfo1         *string `json:"liveinfo1"`
	Liveinfo2         *string `json:"liveinfo2"`
	Liveinfo3         *string `json:"liveinfo3"`
	Passwd            *string `json:"passwd"`
	Timinglogo        *string `json:"timinglogo"`
	Validdate         *string `json:"validdate"`
	Noepr             *int32  `json:"noepr"`
	Tddoc             *int32  `json:"tddoc"`
	Timingreport      *int32  `json:"timingreport"`
	SpecialCupPoints  *int32  `json:"special_cup_points"`
	SkipWcsl          *int32  `json:"skip_wcsl"`
	Validforowg       *int32  `json:"validforowg"`
	Lastupdate        *string `json:"lastupdate"`
}

func FISRaceCCFullFromSqlc(row fissqlc.ARacecc) FISRaceCCFullResponse {
	var (
		raceDateStr   *string
		startEventStr *string
		receivedStr   *string
		validStr      *string
		lastUpdateStr *string
	)

	if row.Racedate.Valid {
		raceDateStr = utils.FormatDatePtr(row.Racedate)
	}
	if row.Starteventdate.Valid {
		startEventStr = utils.FormatDatePtr(row.Starteventdate)
	}
	if row.Receiveddate.Valid {
		receivedStr = utils.FormatDatePtr(row.Receiveddate)
	}
	if row.Validdate.Valid {
		validStr = utils.FormatDatePtr(row.Validdate)
	}
	if row.Lastupdate.Valid {
		lastUpdateStr = utils.FormatTimestampPtr(row.Lastupdate)
	}

	return FISRaceCCFullResponse{
		Raceid:            row.Raceid,
		Eventid:           utils.Int32PtrOrNil(row.Eventid),
		Seasoncode:        utils.Int32PtrOrNil(row.Seasoncode),
		Racecodex:         utils.Int32PtrOrNil(row.Racecodex),
		Disciplineid:      utils.StringPtrOrNil(row.Disciplineid),
		Disciplinecode:    utils.StringPtrOrNil(row.Disciplinecode),
		Catcode:           utils.StringPtrOrNil(row.Catcode),
		Catcode2:          utils.StringPtrOrNil(row.Catcode2),
		Catcode3:          utils.StringPtrOrNil(row.Catcode3),
		Catcode4:          utils.StringPtrOrNil(row.Catcode4),
		Gender:            utils.StringPtrOrNil(row.Gender),
		Racedate:          raceDateStr,
		Starteventdate:    startEventStr,
		Description:       utils.StringPtrOrNil(row.Description),
		Place:             utils.StringPtrOrNil(row.Place),
		Nationcode:        utils.StringPtrOrNil(row.Nationcode),
		Td1id:             utils.Int32PtrOrNil(row.Td1id),
		Td1name:           utils.StringPtrOrNil(row.Td1name),
		Td1nation:         utils.StringPtrOrNil(row.Td1nation),
		Td1code:           utils.Int32PtrOrNil(row.Td1code),
		Td2id:             utils.Int32PtrOrNil(row.Td2id),
		Td2name:           utils.StringPtrOrNil(row.Td2name),
		Td2nation:         utils.StringPtrOrNil(row.Td2nation),
		Td2code:           utils.Int32PtrOrNil(row.Td2code),
		Calstatuscode:     utils.StringPtrOrNil(row.Calstatuscode),
		Procstatuscode:    utils.StringPtrOrNil(row.Procstatuscode),
		Receiveddate:      receivedStr,
		Pursuit:           utils.StringPtrOrNil(row.Pursuit),
		Masse:             utils.StringPtrOrNil(row.Masse),
		Relay:             utils.StringPtrOrNil(row.Relay),
		Distance:          utils.StringPtrOrNil(row.Distance),
		Hill:              utils.StringPtrOrNil(row.Hill),
		Style:             utils.StringPtrOrNil(row.Style),
		Qualif:            utils.StringPtrOrNil(row.Qualif),
		Finale:            utils.StringPtrOrNil(row.Finale),
		Homol:             utils.StringPtrOrNil(row.Homol),
		Webcomment:        utils.StringPtrOrNil(row.Webcomment),
		Displaystatus:     utils.StringPtrOrNil(row.Displaystatus),
		Fisinterncomment:  utils.StringPtrOrNil(row.Fisinterncomment),
		Published:         utils.Int32PtrOrNil(row.Published),
		Validforfispoints: utils.Int32PtrOrNil(row.Validforfispoints),
		Usedfislist:       utils.StringPtrOrNil(row.Usedfislist),
		Tolist:            utils.StringPtrOrNil(row.Tolist),
		Discforlistcode:   utils.StringPtrOrNil(row.Discforlistcode),
		Calculatedpenalty: utils.StringPtrOrNil(row.Calculatedpenalty),
		Appliedpenalty:    utils.StringPtrOrNil(row.Appliedpenalty),
		Appliedscala:      utils.StringPtrOrNil(row.Appliedscala),
		Penscafixed:       utils.StringPtrOrNil(row.Penscafixed),
		Version:           utils.Int32PtrOrNil(row.Version),
		Nationraceid:      utils.Int32PtrOrNil(row.Nationraceid),
		Provraceid:        utils.Int32PtrOrNil(row.Provraceid),
		Msql7evid:         utils.Int32PtrOrNil(row.Msql7evid),
		Mssql7id:          utils.Int32PtrOrNil(row.Mssql7id),
		Results:           utils.Int32PtrOrNil(row.Results),
		Pdf:               utils.Int32PtrOrNil(row.Pdf),
		Topbanner:         utils.StringPtrOrNil(row.Topbanner),
		Bottombanner:      utils.StringPtrOrNil(row.Bottombanner),
		Toplogo:           utils.StringPtrOrNil(row.Toplogo),
		Bottomlogo:        utils.StringPtrOrNil(row.Bottomlogo),
		Gallery:           utils.StringPtrOrNil(row.Gallery),
		Indi:              utils.Int32PtrOrNil(row.Indi),
		Team:              utils.Int32PtrOrNil(row.Team),
		Tabcount:          utils.Int32PtrOrNil(row.Tabcount),
		Columncount:       utils.Int32PtrOrNil(row.Columncount),
		Level:             utils.StringPtrOrNil(row.Level),
		Hloc1:             utils.StringPtrOrNil(row.Hloc1),
		Hloc2:             utils.StringPtrOrNil(row.Hloc2),
		Hloc3:             utils.StringPtrOrNil(row.Hloc3),
		Hcet1:             utils.StringPtrOrNil(row.Hcet1),
		Hcet2:             utils.StringPtrOrNil(row.Hcet2),
		Hcet3:             utils.StringPtrOrNil(row.Hcet3),
		Live:              utils.Int32PtrOrNil(row.Live),
		Livestatus1:       utils.StringPtrOrNil(row.Livestatus1),
		Livestatus2:       utils.StringPtrOrNil(row.Livestatus2),
		Livestatus3:       utils.StringPtrOrNil(row.Livestatus3),
		Liveinfo1:         utils.StringPtrOrNil(row.Liveinfo1),
		Liveinfo2:         utils.StringPtrOrNil(row.Liveinfo2),
		Liveinfo3:         utils.StringPtrOrNil(row.Liveinfo3),
		Passwd:            utils.StringPtrOrNil(row.Passwd),
		Timinglogo:        utils.StringPtrOrNil(row.Timinglogo),
		Validdate:         validStr,
		Noepr:             utils.Int32PtrOrNil(row.Noepr),
		Tddoc:             utils.Int32PtrOrNil(row.Tddoc),
		Timingreport:      utils.Int32PtrOrNil(row.Timingreport),
		SpecialCupPoints:  utils.Int32PtrOrNil(row.SpecialCupPoints),
		SkipWcsl:          utils.Int32PtrOrNil(row.SkipWcsl),
		Validforowg:       utils.Int32PtrOrNil(row.Validforowg),
		Lastupdate:        lastUpdateStr,
	}
}
