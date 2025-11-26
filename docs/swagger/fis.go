package swagger

type FISAthleteItem struct {
	Firstname *string `json:"firstname,omitempty" example:"Iivo"`
	Lastname  *string `json:"lastname,omitempty" example:"Niskanen"`
	Fiscode   *int32  `json:"fiscode,omitempty" example:"342001"`
}

type FISAthletesResponse struct {
	Athletes []FISAthleteItem `json:"athletes"`
}

type FISNationsBySectorResponse struct {
	Nations []string `json:"nations" example:"FIN,NOR,SWE"`
}

type FISInsertCompetitorExample struct {
	Competitorid       int32   `json:"competitorid" example:"123456"`
	Personid           *int32  `json:"personid,omitempty" example:"98765"`
	Ipcid              *int32  `json:"ipcid,omitempty" example:"0"`
	Fiscode            *int32  `json:"fiscode,omitempty" example:"342001"`
	Birthdate          *string `json:"birthdate,omitempty" example:"1992-01-12"`             // YYYY-MM-DD
	StatusDate         *string `json:"status_date,omitempty" example:"2025-10-27T08:30:00Z"` // RFC3339
	Fee                *string `json:"fee,omitempty" example:"10.00000"`
	Dateofcreation     *string `json:"dateofcreation,omitempty" example:"2025-10-20"` // YYYY-MM-DD
	Injury             *int32  `json:"injury,omitempty" example:"0"`
	Version            *int32  `json:"version,omitempty" example:"1"`
	Compidmssql        *int32  `json:"compidmssql,omitempty" example:"0"`
	Carving            *int32  `json:"carving,omitempty" example:"0"`
	Photo              *int32  `json:"photo,omitempty" example:"0"`
	Notallowed         *int32  `json:"notallowed,omitempty" example:"0"`
	Published          *int32  `json:"published,omitempty" example:"1"`
	Team               *int32  `json:"team,omitempty" example:"0"`
	PhotoBig           *int32  `json:"photo_big,omitempty" example:"0"`
	Lastupdate         *string `json:"lastupdate,omitempty" example:"2025-10-27T08:30:00Z"` // RFC3339
	Statusnextlist     *string `json:"statusnextlist,omitempty" example:"A"`
	Alternatenamecheck *string `json:"alternatenamecheck,omitempty" example:"OK"`
	Deletedat          *string `json:"deletedat,omitempty" example:""`
	Doped              *string `json:"doped,omitempty" example:"NO"`
	Createdby          *string `json:"createdby,omitempty" example:"admin@fis.int"`
	Categorycode       *string `json:"categorycode,omitempty" example:"SEN"`
	Classname          *string `json:"classname,omitempty" example:"Senior"`
	Data               *string `json:"data,omitempty" example:"{}"`
	Lastupdateby       *string `json:"lastupdateby,omitempty" example:"system"`
	Disciplines        *string `json:"disciplines,omitempty" example:"DISTANCE,SPRINT"`
	Type               *string `json:"type,omitempty" example:"ATHLETE"`
	Sectorcode         *string `json:"sectorcode,omitempty" example:"CC"` // JP/NK/CC
	Classcode          *string `json:"classcode,omitempty" example:"A"`
	Lastname           *string `json:"lastname,omitempty" example:"Niskanen"`
	Firstname          *string `json:"firstname,omitempty" example:"Iivo"`
	Gender             *string `json:"gender,omitempty" example:"M"`
	Natteam            *string `json:"natteam,omitempty" example:"FIN-A"`
	Nationcode         *string `json:"nationcode,omitempty" example:"FIN"`
	Nationalcode       *string `json:"nationalcode,omitempty" example:"246"`
	Skiclub            *string `json:"skiclub,omitempty" example:"Ounasvaara Hiihtoseura"`
	Association        *string `json:"association,omitempty" example:"FIN"`
	Status             *string `json:"status,omitempty" example:"ACTIVE"`
	StatusOld          *string `json:"status_old,omitempty" example:""`
	StatusBy           *string `json:"status_by,omitempty" example:"FIS"`
	Tragroup           *string `json:"tragroup,omitempty" example:"A"`
}

type FISUpdateCompetitorExample struct {
	Competitorid       int32   `json:"competitorid" example:"123456"`
	Personid           *int32  `json:"personid,omitempty" example:"98765"`
	Ipcid              *int32  `json:"ipcid,omitempty" example:"0"`
	Fiscode            *int32  `json:"fiscode,omitempty" example:"342001"`
	Birthdate          *string `json:"birthdate,omitempty" example:"1992-01-12"`             // YYYY-MM-DD
	StatusDate         *string `json:"status_date,omitempty" example:"2025-10-27T08:30:00Z"` // RFC3339
	Fee                *string `json:"fee,omitempty" example:"10.00000"`
	Dateofcreation     *string `json:"dateofcreation,omitempty" example:"2025-10-20"` // YYYY-MM-DD
	Injury             *int32  `json:"injury,omitempty" example:"0"`
	Version            *int32  `json:"version,omitempty" example:"1"`
	Compidmssql        *int32  `json:"compidmssql,omitempty" example:"0"`
	Carving            *int32  `json:"carving,omitempty" example:"0"`
	Photo              *int32  `json:"photo,omitempty" example:"0"`
	Notallowed         *int32  `json:"notallowed,omitempty" example:"0"`
	Published          *int32  `json:"published,omitempty" example:"1"`
	Team               *int32  `json:"team,omitempty" example:"0"`
	PhotoBig           *int32  `json:"photo_big,omitempty" example:"0"`
	Lastupdate         *string `json:"lastupdate,omitempty" example:"2025-10-27T08:30:00Z"` // RFC3339
	Statusnextlist     *string `json:"statusnextlist,omitempty" example:"A"`
	Alternatenamecheck *string `json:"alternatenamecheck,omitempty" example:"OK"`
	Deletedat          *string `json:"deletedat,omitempty" example:""`
	Doped              *string `json:"doped,omitempty" example:"NO"`
	Createdby          *string `json:"createdby,omitempty" example:"admin@fis.int"`
	Categorycode       *string `json:"categorycode,omitempty" example:"SEN"`
	Classname          *string `json:"classname,omitempty" example:"Senior"`
	Data               *string `json:"data,omitempty" example:"{}"`
	Lastupdateby       *string `json:"lastupdateby,omitempty" example:"system"`
	Disciplines        *string `json:"disciplines,omitempty" example:"DISTANCE,SPRINT"`
	Type               *string `json:"type,omitempty" example:"ATHLETE"`
	Sectorcode         *string `json:"sectorcode,omitempty" example:"DD"` // JP/NK/CC
	Classcode          *string `json:"classcode,omitempty" example:"A"`
	Lastname           *string `json:"lastname,omitempty" example:"Niskanen"`
	Firstname          *string `json:"firstname,omitempty" example:"Iivo"`
	Gender             *string `json:"gender,omitempty" example:"M"`
	Natteam            *string `json:"natteam,omitempty" example:"FIN-A"`
	Nationcode         *string `json:"nationcode,omitempty" example:"FIN"`
	Nationalcode       *string `json:"nationalcode,omitempty" example:"246"`
	Skiclub            *string `json:"skiclub,omitempty" example:"Ounasvaara Hiihtoseura"`
	Association        *string `json:"association,omitempty" example:"FIN"`
	Status             *string `json:"status,omitempty" example:"ACTIVE"`
	StatusOld          *string `json:"status_old,omitempty" example:""`
	StatusBy           *string `json:"status_by,omitempty" example:"FIS"`
	Tragroup           *string `json:"tragroup,omitempty" example:"A"`
}

type FISCompetitor struct {
	Competitorid   *int32  `json:"competitorid,omitempty" example:"123456"`
	Personid       *int32  `json:"personid,omitempty" example:"98765"`
	Ipcid          *int32  `json:"ipcid,omitempty" example:"0"`
	Type           *string `json:"type,omitempty" example:"ATHLETE"`
	Sectorcode     *string `json:"sectorcode,omitempty" example:"CC"`
	Fiscode        *int32  `json:"fiscode,omitempty" example:"342001"`
	Lastname       *string `json:"lastname,omitempty" example:"Niskanen"`
	Firstname      *string `json:"firstname,omitempty" example:"Iivo"`
	Gender         *string `json:"gender,omitempty" example:"M"`
	Birthdate      *string `json:"birthdate,omitempty" example:"1992-01-12"`
	StatusDate     *string `json:"status_date,omitempty" example:"2025-10-27T08:30:00Z"`
	Dateofcreation *string `json:"dateofcreation,omitempty" example:"2025-10-20"`
	Lastupdate     *string `json:"lastupdate,omitempty" example:"2025-10-27T08:30:00Z"`
	Nationcode     *string `json:"nationcode,omitempty" example:"FIN"`
	Nationalcode   *string `json:"nationalcode,omitempty" example:"246"`
	Skiclub        *string `json:"skiclub,omitempty" example:"Ounasvaara Hiihtoseura"`
	Association    *string `json:"association,omitempty" example:"FIN"`
	Status         *string `json:"status,omitempty" example:"ACTIVE"`
}

type FISLastCompetitorResponse struct {
	Competitor FISCompetitor `json:"competitor"`
}

type FISSeasonsCCResponse struct {
	Seasons []int32 `json:"seasons" swaggertype:"array,integer" example:"2026,2025,2024"`
}

type FISDisciplinesCCResponse struct {
	Disciplines []string `json:"disciplines" swaggertype:"array,string" example:"SP,DI,PU"`
}

type FISCategoriesCCResponse struct {
	Categories []string `json:"categories" swaggertype:"array,string" example:"WC,WSC,COC"`
}

type FISRaceCC struct {
	Raceid            int32   `json:"raceid" example:"123456"`
	Eventid           *int32  `json:"eventid,omitempty" example:"7890"`
	Seasoncode        *int32  `json:"seasoncode,omitempty" example:"2025"`
	Racecodex         *int32  `json:"racecodex,omitempty" example:"1001"`
	Disciplineid      *string `json:"disciplineid,omitempty" example:"DIST"`
	Disciplinecode    *string `json:"disciplinecode,omitempty" example:"SP"`
	Catcode           *string `json:"catcode,omitempty" example:"WC"`
	Catcode2          *string `json:"catcode2,omitempty"`
	Catcode3          *string `json:"catcode3,omitempty"`
	Catcode4          *string `json:"catcode4,omitempty"`
	Gender            *string `json:"gender,omitempty" example:"M"`
	Racedate          *string `json:"racedate,omitempty" format:"date" example:"2025-02-14"`
	Starteventdate    *string `json:"starteventdate,omitempty" format:"date" example:"2025-02-13"`
	Description       *string `json:"description,omitempty" example:"World Cup Sprint"`
	Place             *string `json:"place,omitempty" example:"Lahti"`
	Nationcode        *string `json:"nationcode,omitempty" example:"FIN"`
	Td1id             *int32  `json:"td1id,omitempty"`
	Td1name           *string `json:"td1name,omitempty"`
	Td1nation         *string `json:"td1nation,omitempty"`
	Td1code           *int32  `json:"td1code,omitempty"`
	Td2id             *int32  `json:"td2id,omitempty"`
	Td2name           *string `json:"td2name,omitempty"`
	Td2nation         *string `json:"td2nation,omitempty"`
	Td2code           *int32  `json:"td2code,omitempty"`
	Calstatuscode     *string `json:"calstatuscode,omitempty"`
	Procstatuscode    *string `json:"procstatuscode,omitempty"`
	Receiveddate      *string `json:"receiveddate,omitempty" format:"date-time"`
	Pursuit           *string `json:"pursuit,omitempty"`
	Masse             *string `json:"masse,omitempty"`
	Relay             *string `json:"relay,omitempty"`
	Distance          *string `json:"distance,omitempty"`
	Hill              *string `json:"hill,omitempty"`
	Style             *string `json:"style,omitempty"`
	Qualif            *string `json:"qualif,omitempty"`
	Finale            *string `json:"finale,omitempty"`
	Homol             *string `json:"homol,omitempty"`
	Webcomment        *string `json:"webcomment,omitempty"`
	Displaystatus     *string `json:"displaystatus,omitempty"`
	Fisinterncomment  *string `json:"fisinterncomment,omitempty"`
	Published         *int32  `json:"published,omitempty" example:"1"`
	Validforfispoints *int32  `json:"validforfispoints,omitempty" example:"1"`
	Usedfislist       *string `json:"usedfislist,omitempty"`
	Tolist            *string `json:"tolist,omitempty"`
	Discforlistcode   *string `json:"discforlistcode,omitempty"`
	Calculatedpenalty *string `json:"calculatedpenalty,omitempty"`
	Appliedpenalty    *string `json:"appliedpenalty,omitempty"`
	Appliedscala      *string `json:"appliedscala,omitempty"`
	Penscafixed       *string `json:"penscafixed,omitempty"`
	Version           *int32  `json:"version,omitempty"`
	Nationraceid      *int32  `json:"nationraceid,omitempty"`
	Provraceid        *int32  `json:"provraceid,omitempty"`
	Msql7evid         *int32  `json:"msql7evid,omitempty"`
	Mssql7id          *int32  `json:"mssql7id,omitempty"`
	Results           *int32  `json:"results,omitempty"`
	Pdf               *int32  `json:"pdf,omitempty"`
	Topbanner         *string `json:"topbanner,omitempty"`
	Bottombanner      *string `json:"bottombanner,omitempty"`
	Toplogo           *string `json:"toplogo,omitempty"`
	Bottomlogo        *string `json:"bottomlogo,omitempty"`
	Gallery           *string `json:"gallery,omitempty"`
	Indi              *int32  `json:"indi,omitempty"`
	Team              *int32  `json:"team,omitempty"`
	Tabcount          *int32  `json:"tabcount,omitempty"`
	Columncount       *int32  `json:"columncount,omitempty"`
	Level             *string `json:"level,omitempty"`
	Hloc1             *string `json:"hloc1,omitempty"`
	Hloc2             *string `json:"hloc2,omitempty"`
	Hloc3             *string `json:"hloc3,omitempty"`
	Hcet1             *string `json:"hcet1,omitempty"`
	Hcet2             *string `json:"hcet2,omitempty"`
	Hcet3             *string `json:"hcet3,omitempty"`
	Live              *int32  `json:"live,omitempty"`
	Livestatus1       *string `json:"livestatus1,omitempty"`
	Livestatus2       *string `json:"livestatus2,omitempty"`
	Livestatus3       *string `json:"livestatus3,omitempty"`
	Liveinfo1         *string `json:"liveinfo1,omitempty"`
	Liveinfo2         *string `json:"liveinfo2,omitempty"`
	Liveinfo3         *string `json:"liveinfo3,omitempty"`
	Passwd            *string `json:"passwd,omitempty"`
	Timinglogo        *string `json:"timinglogo,omitempty"`
	Validdate         *string `json:"validdate,omitempty" format:"date-time"`
	Noepr             *int32  `json:"noepr,omitempty"`
	Tddoc             *int32  `json:"tddoc,omitempty"`
	Timingreport      *int32  `json:"timingreport,omitempty"`
	SpecialCupPoints  *int32  `json:"special_cup_points,omitempty"`
	SkipWcsl          *int32  `json:"skip_wcsl,omitempty"`
	Validforowg       *int32  `json:"validforowg,omitempty"`
	Lastupdate        *string `json:"lastupdate,omitempty" format:"date-time" example:"2025-02-14T18:30:00Z"`
}

type FISRacesCCResponse struct {
	Races []FISRaceCC `json:"races"`
}

type FISLastRaceCCResponse struct {
	Race FISRaceCC `json:"race"`
}

type FISInsertRaceCCExample struct {
	Raceid            int32   `json:"raceid" validate:"required" example:"123456"`
	Eventid           *int32  `json:"eventid,omitempty" example:"7890"`
	Seasoncode        *int32  `json:"seasoncode,omitempty" example:"2025"`
	Racecodex         *int32  `json:"racecodex,omitempty" example:"1001"`
	Disciplineid      *string `json:"disciplineid,omitempty" example:"DIST"`
	Disciplinecode    *string `json:"disciplinecode,omitempty" example:"SP"`
	Catcode           *string `json:"catcode,omitempty" example:"WC"`
	Catcode2          *string `json:"catcode2,omitempty"`
	Catcode3          *string `json:"catcode3,omitempty"`
	Catcode4          *string `json:"catcode4,omitempty"`
	Gender            *string `json:"gender,omitempty" example:"M"`
	Racedate          *string `json:"racedate,omitempty" format:"date" example:"2025-02-14"`
	Starteventdate    *string `json:"starteventdate,omitempty" format:"date" example:"2025-02-13"`
	Description       *string `json:"description,omitempty" example:"World Cup Sprint"`
	Place             *string `json:"place,omitempty" example:"Lahti"`
	Nationcode        *string `json:"nationcode,omitempty" example:"FIN"`
	Receiveddate      *string `json:"receiveddate,omitempty" format:"date-time"`
	Validdate         *string `json:"validdate,omitempty" format:"date-time"`
	Td1id             *int32  `json:"td1id,omitempty"`
	Td1name           *string `json:"td1name,omitempty"`
	Td1nation         *string `json:"td1nation,omitempty"`
	Td1code           *int32  `json:"td1code,omitempty"`
	Td2id             *int32  `json:"td2id,omitempty"`
	Td2name           *string `json:"td2name,omitempty"`
	Td2nation         *string `json:"td2nation,omitempty"`
	Td2code           *int32  `json:"td2code,omitempty"`
	Calstatuscode     *string `json:"calstatuscode,omitempty"`
	Procstatuscode    *string `json:"procstatuscode,omitempty"`
	Displaystatus     *string `json:"displaystatus,omitempty"`
	Fisinterncomment  *string `json:"fisinterncomment,omitempty"`
	Webcomment        *string `json:"webcomment,omitempty"`
	Pursuit           *string `json:"pursuit,omitempty"`
	Masse             *string `json:"masse,omitempty"`
	Relay             *string `json:"relay,omitempty"`
	Distance          *string `json:"distance,omitempty"`
	Hill              *string `json:"hill,omitempty"`
	Style             *string `json:"style,omitempty"`
	Qualif            *string `json:"qualif,omitempty"`
	Finale            *string `json:"finale,omitempty"`
	Homol             *string `json:"homol,omitempty"`
	Published         *int32  `json:"published,omitempty" example:"1"`
	Validforfispoints *int32  `json:"validforfispoints,omitempty" example:"1"`
	Usedfislist       *string `json:"usedfislist,omitempty"`
	Tolist            *string `json:"tolist,omitempty"`
	Discforlistcode   *string `json:"discforlistcode,omitempty"`
	Calculatedpenalty *string `json:"calculatedpenalty,omitempty"`
	Appliedpenalty    *string `json:"appliedpenalty,omitempty"`
	Appliedscala      *string `json:"appliedscala,omitempty"`
	Penscafixed       *string `json:"penscafixed,omitempty"`
	Version           *int32  `json:"version,omitempty"`
	Nationraceid      *int32  `json:"nationraceid,omitempty"`
	Provraceid        *int32  `json:"provraceid,omitempty"`
	Msql7evid         *int32  `json:"msql7evid,omitempty"`
	Mssql7id          *int32  `json:"mssql7id,omitempty"`
	Topbanner         *string `json:"topbanner,omitempty"`
	Bottombanner      *string `json:"bottombanner,omitempty"`
	Toplogo           *string `json:"toplogo,omitempty"`
	Bottomlogo        *string `json:"bottomlogo,omitempty"`
	Gallery           *string `json:"gallery,omitempty"`
	Indi              *int32  `json:"indi,omitempty"`
	Team              *int32  `json:"team,omitempty"`
	Tabcount          *int32  `json:"tabcount,omitempty"`
	Columncount       *int32  `json:"columncount,omitempty"`
	Level             *string `json:"level,omitempty"`
	Hloc1             *string `json:"hloc1,omitempty"`
	Hloc2             *string `json:"hloc2,omitempty"`
	Hloc3             *string `json:"hloc3,omitempty"`
	Hcet1             *string `json:"hcet1,omitempty"`
	Hcet2             *string `json:"hcet2,omitempty"`
	Hcet3             *string `json:"hcet3,omitempty"`
	Live              *int32  `json:"live,omitempty"`
	Livestatus1       *string `json:"livestatus1,omitempty"`
	Livestatus2       *string `json:"livestatus2,omitempty"`
	Livestatus3       *string `json:"livestatus3,omitempty"`
	Liveinfo1         *string `json:"liveinfo1,omitempty"`
	Liveinfo2         *string `json:"liveinfo2,omitempty"`
	Liveinfo3         *string `json:"liveinfo3,omitempty"`
	Passwd            *string `json:"passwd,omitempty"`
	Timinglogo        *string `json:"timinglogo,omitempty"`
	Results           *int32  `json:"results,omitempty"`
	Pdf               *int32  `json:"pdf,omitempty"`
	Noepr             *int32  `json:"noepr,omitempty"`
	Tddoc             *int32  `json:"tddoc,omitempty"`
	Timingreport      *int32  `json:"timingreport,omitempty"`
	SpecialCupPoints  *int32  `json:"special_cup_points,omitempty"`
	SkipWcsl          *int32  `json:"skip_wcsl,omitempty"`
	Validforowg       *int32  `json:"validforowg,omitempty"`
	Lastupdate        *string `json:"lastupdate,omitempty" format:"date-time" example:"2025-02-14T18:30:00Z"`
}

type FISUpdateRaceCCExample struct {
	Raceid            int32   `json:"raceid" validate:"required" example:"123456"`
	Eventid           *int32  `json:"eventid,omitempty" example:"7890"`
	Seasoncode        *int32  `json:"seasoncode,omitempty" example:"2025"`
	Racecodex         *int32  `json:"racecodex,omitempty" example:"1001"`
	Disciplineid      *string `json:"disciplineid,omitempty" example:"DDD"`
	Disciplinecode    *string `json:"disciplinecode,omitempty" example:"SP"`
	Catcode           *string `json:"catcode,omitempty" example:"WC"`
	Catcode2          *string `json:"catcode2,omitempty"`
	Catcode3          *string `json:"catcode3,omitempty"`
	Catcode4          *string `json:"catcode4,omitempty"`
	Gender            *string `json:"gender,omitempty" example:"M"`
	Racedate          *string `json:"racedate,omitempty" format:"date" example:"2025-02-14"`
	Starteventdate    *string `json:"starteventdate,omitempty" format:"date" example:"2025-02-13"`
	Description       *string `json:"description,omitempty" example:"World Cup Sprint (updated)"`
	Place             *string `json:"place,omitempty" example:"Lahti"`
	Nationcode        *string `json:"nationcode,omitempty" example:"FIN"`
	Receiveddate      *string `json:"receiveddate,omitempty" format:"date-time"`
	Validdate         *string `json:"validdate,omitempty" format:"date-time"`
	Td1id             *int32  `json:"td1id,omitempty"`
	Td1name           *string `json:"td1name,omitempty"`
	Td1nation         *string `json:"td1nation,omitempty"`
	Td1code           *int32  `json:"td1code,omitempty"`
	Td2id             *int32  `json:"td2id,omitempty"`
	Td2name           *string `json:"td2name,omitempty"`
	Td2nation         *string `json:"td2nation,omitempty"`
	Td2code           *int32  `json:"td2code,omitempty"`
	Calstatuscode     *string `json:"calstatuscode,omitempty"`
	Procstatuscode    *string `json:"procstatuscode,omitempty"`
	Displaystatus     *string `json:"displaystatus,omitempty"`
	Fisinterncomment  *string `json:"fisinterncomment,omitempty"`
	Webcomment        *string `json:"webcomment,omitempty"`
	Pursuit           *string `json:"pursuit,omitempty"`
	Masse             *string `json:"masse,omitempty"`
	Relay             *string `json:"relay,omitempty"`
	Distance          *string `json:"distance,omitempty"`
	Hill              *string `json:"hill,omitempty"`
	Style             *string `json:"style,omitempty"`
	Qualif            *string `json:"qualif,omitempty"`
	Finale            *string `json:"finale,omitempty"`
	Homol             *string `json:"homol,omitempty"`
	Published         *int32  `json:"published,omitempty" example:"1"`
	Validforfispoints *int32  `json:"validforfispoints,omitempty" example:"1"`
	Usedfislist       *string `json:"usedfislist,omitempty"`
	Tolist            *string `json:"tolist,omitempty"`
	Discforlistcode   *string `json:"discforlistcode,omitempty"`
	Calculatedpenalty *string `json:"calculatedpenalty,omitempty"`
	Appliedpenalty    *string `json:"appliedpenalty,omitempty"`
	Appliedscala      *string `json:"appliedscala,omitempty"`
	Penscafixed       *string `json:"penscafixed,omitempty"`
	Version           *int32  `json:"version,omitempty"`
	Nationraceid      *int32  `json:"nationraceid,omitempty"`
	Provraceid        *int32  `json:"provraceid,omitempty"`
	Msql7evid         *int32  `json:"msql7evid,omitempty"`
	Mssql7id          *int32  `json:"mssql7id,omitempty"`
	Topbanner         *string `json:"topbanner,omitempty"`
	Bottombanner      *string `json:"bottombanner,omitempty"`
	Toplogo           *string `json:"toplogo,omitempty"`
	Bottomlogo        *string `json:"bottomlogo,omitempty"`
	Gallery           *string `json:"gallery,omitempty"`
	Indi              *int32  `json:"indi,omitempty"`
	Team              *int32  `json:"team,omitempty"`
	Tabcount          *int32  `json:"tabcount,omitempty"`
	Columncount       *int32  `json:"columncount,omitempty"`
	Level             *string `json:"level,omitempty"`
	Hloc1             *string `json:"hloc1,omitempty"`
	Hloc2             *string `json:"hloc2,omitempty"`
	Hloc3             *string `json:"hloc3,omitempty"`
	Hcet1             *string `json:"hcet1,omitempty"`
	Hcet2             *string `json:"hcet2,omitempty"`
	Hcet3             *string `json:"hcet3,omitempty"`
	Live              *int32  `json:"live,omitempty"`
	Livestatus1       *string `json:"livestatus1,omitempty"`
	Livestatus2       *string `json:"livestatus2,omitempty"`
	Livestatus3       *string `json:"livestatus3,omitempty"`
	Liveinfo1         *string `json:"liveinfo1,omitempty"`
	Liveinfo2         *string `json:"liveinfo2,omitempty"`
	Liveinfo3         *string `json:"liveinfo3,omitempty"`
	Passwd            *string `json:"passwd,omitempty"`
	Timinglogo        *string `json:"timinglogo,omitempty"`
	Results           *int32  `json:"results,omitempty"`
	Pdf               *int32  `json:"pdf,omitempty"`
	Noepr             *int32  `json:"noepr,omitempty"`
	Tddoc             *int32  `json:"tddoc,omitempty"`
	Timingreport      *int32  `json:"timingreport,omitempty"`
	SpecialCupPoints  *int32  `json:"special_cup_points,omitempty"`
	SkipWcsl          *int32  `json:"skip_wcsl,omitempty"`
	Validforowg       *int32  `json:"validforowg,omitempty"`
	Lastupdate        *string `json:"lastupdate,omitempty" format:"date-time" example:"2025-02-15T08:00:00Z"`
}

type FISSeasonsJPResponse struct {
	Seasons []int32 `json:"seasons" swaggertype:"array,integer" example:"2026,2025,2024"`
}

type FISDisciplinesJPResponse struct {
	Disciplines []string `json:"disciplines" swaggertype:"array,string" example:"NH,LH,FH"`
}

type FISCategoriesJPResponse struct {
	Categories []string `json:"categories" swaggertype:"array,string" example:"WC,COC,WSC"`
}

type FISRaceJP struct {
	Raceid            int32   `json:"raceid" example:"234567"`
	Eventid           *int32  `json:"eventid,omitempty" example:"8901"`
	Seasoncode        *int32  `json:"seasoncode,omitempty" example:"2025"`
	Racecodex         *int32  `json:"racecodex,omitempty" example:"2001"`
	Disciplineid      *string `json:"disciplineid,omitempty" example:"HS140"`
	Disciplinecode    *string `json:"disciplinecode,omitempty" example:"LH"`
	Catcode           *string `json:"catcode,omitempty" example:"WC"`
	Catcode2          *string `json:"catcode2,omitempty"`
	Catcode3          *string `json:"catcode3,omitempty"`
	Catcode4          *string `json:"catcode4,omitempty"`
	Gender            *string `json:"gender,omitempty" example:"M"`
	Racedate          *string `json:"racedate,omitempty" format:"date" example:"2025-01-25"`
	Starteventdate    *string `json:"starteventdate,omitempty" format:"date" example:"2025-01-24"`
	Description       *string `json:"description,omitempty" example:"World Cup HS140 Individual"`
	Place             *string `json:"place,omitempty" example:"Planica"`
	Nationcode        *string `json:"nationcode,omitempty" example:"SLO"`
	Td1id             *int32  `json:"td1id,omitempty"`
	Td1name           *string `json:"td1name,omitempty"`
	Td1nation         *string `json:"td1nation,omitempty"`
	Td1code           *int32  `json:"td1code,omitempty"`
	Td2id             *int32  `json:"td2id,omitempty"`
	Td2name           *string `json:"td2name,omitempty"`
	Td2nation         *string `json:"td2nation,omitempty"`
	Td2code           *int32  `json:"td2code,omitempty"`
	Calstatuscode     *string `json:"calstatuscode,omitempty"`
	Procstatuscode    *string `json:"procstatuscode,omitempty"`
	Receiveddate      *string `json:"receiveddate,omitempty" format:"date-time"`
	Pursuit           *string `json:"pursuit,omitempty"`
	Masse             *string `json:"masse,omitempty"`
	Relay             *string `json:"relay,omitempty"`
	Distance          *string `json:"distance,omitempty"`
	Hill              *int32  `json:"hill,omitempty" example:"140"`
	Style             *string `json:"style,omitempty" example:"CLASSIC"`
	Qualif            *string `json:"qualif,omitempty" example:"Q"`
	Finale            *string `json:"finale,omitempty"`
	Homol             *string `json:"homol,omitempty"`
	Webcomment        *string `json:"webcomment,omitempty"`
	Displaystatus     *string `json:"displaystatus,omitempty"`
	Fisinterncomment  *string `json:"fisinterncomment,omitempty"`
	Published         *int32  `json:"published,omitempty" example:"1"`
	Validforfispoints *int32  `json:"validforfispoints,omitempty" example:"1"`
	Usedfislist       *string `json:"usedfislist,omitempty"`
	Tolist            *string `json:"tolist,omitempty"`
	Discforlistcode   *string `json:"discforlistcode,omitempty"`
	Calculatedpenalty *string `json:"calculatedpenalty,omitempty"`
	Appliedpenalty    *string `json:"appliedpenalty,omitempty"`
	Appliedscala      *string `json:"appliedscala,omitempty"`
	Penscafixed       *string `json:"penscafixed,omitempty"`
	Version           *int32  `json:"version,omitempty"`
	Nationraceid      *int32  `json:"nationraceid,omitempty"`
	Provraceid        *int32  `json:"provraceid,omitempty"`
	Msql7evid         *int32  `json:"msql7evid,omitempty"`
	Mssql7id          *int32  `json:"mssql7id,omitempty"`
	Results           *int32  `json:"results,omitempty"`
	Pdf               *int32  `json:"pdf,omitempty"`
	Topbanner         *string `json:"topbanner,omitempty"`
	Bottombanner      *string `json:"bottombanner,omitempty"`
	Toplogo           *string `json:"toplogo,omitempty"`
	Bottomlogo        *string `json:"bottomlogo,omitempty"`
	Gallery           *string `json:"gallery,omitempty"`
	Indi              *int32  `json:"indi,omitempty"`
	Team              *int32  `json:"team,omitempty"`
	Tabcount          *int32  `json:"tabcount,omitempty"`
	Columncount       *int32  `json:"columncount,omitempty"`
	Level             *string `json:"level,omitempty"`
	Hloc1             *string `json:"hloc1,omitempty" format:"date-time" example:"2025-01-25T15:00:00Z"`
	Hloc2             *string `json:"hloc2,omitempty" format:"date-time" example:"2025-01-25T15:30:00Z"`
	Hloc3             *string `json:"hloc3,omitempty" format:"date-time" example:"2025-01-25T16:00:00Z"`
	Hcet1             *string `json:"hcet1,omitempty" format:"date-time" example:"2025-01-25T15:00:00Z"`
	Hcet2             *string `json:"hcet2,omitempty" format:"date-time" example:"2025-01-25T15:30:00Z"`
	Hcet3             *string `json:"hcet3,omitempty" format:"date-time" example:"2025-01-25T16:00:00Z"`
	Live              *int32  `json:"live,omitempty"`
	Livestatus1       *string `json:"livestatus1,omitempty"`
	Livestatus2       *string `json:"livestatus2,omitempty"`
	Livestatus3       *string `json:"livestatus3,omitempty"`
	Liveinfo1         *string `json:"liveinfo1,omitempty"`
	Liveinfo2         *string `json:"liveinfo2,omitempty"`
	Liveinfo3         *string `json:"liveinfo3,omitempty"`
	Passwd            *string `json:"passwd,omitempty"`
	Timinglogo        *string `json:"timinglogo,omitempty"`
	Validdate         *string `json:"validdate,omitempty" format:"date-time"`
	Noepr             *int32  `json:"noepr,omitempty"`
	Tddoc             *int32  `json:"tddoc,omitempty"`
	Timingreport      *int32  `json:"timingreport,omitempty"`
	SpecialCupPoints  *int32  `json:"special_cup_points,omitempty"`
	SkipWcsl          *int32  `json:"skip_wcsl,omitempty"`
	Lastupdate        *string `json:"lastupdate,omitempty" format:"date-time" example:"2025-01-25T18:30:00Z"`
	Validforowg       *string `json:"validforowg,omitempty" example:"Y"`
}

type FISRacesJPResponse struct {
	Races []FISRaceJP `json:"races"`
}

type FISLastRaceJPResponse struct {
	Race FISRaceJP `json:"race"`
}

type FISInsertRaceJPExample struct {
	Raceid            int32   `json:"raceid" validate:"required" example:"234567"`
	Eventid           *int32  `json:"eventid,omitempty" example:"8901"`
	Seasoncode        *int32  `json:"seasoncode,omitempty" example:"2025"`
	Racecodex         *int32  `json:"racecodex,omitempty" example:"2001"`
	Disciplineid      *string `json:"disciplineid,omitempty" example:"HS140"`
	Disciplinecode    *string `json:"disciplinecode,omitempty" example:"LH"`
	Catcode           *string `json:"catcode,omitempty" example:"WC"`
	Catcode2          *string `json:"catcode2,omitempty"`
	Catcode3          *string `json:"catcode3,omitempty"`
	Catcode4          *string `json:"catcode4,omitempty"`
	Gender            *string `json:"gender,omitempty" example:"M"`
	Racedate          *string `json:"racedate,omitempty" format:"date" example:"2025-01-25"`
	Starteventdate    *string `json:"starteventdate,omitempty" format:"date" example:"2025-01-24"`
	Description       *string `json:"description,omitempty" example:"World Cup HS140 Individual"`
	Place             *string `json:"place,omitempty" example:"Planica"`
	Nationcode        *string `json:"nationcode,omitempty" example:"SLO"`
	Receiveddate      *string `json:"receiveddate,omitempty" format:"date-time" example:"2025-01-20T10:00:00Z"`
	Td1id             *int32  `json:"td1id,omitempty"`
	Td1name           *string `json:"td1name,omitempty" example:"John Doe"`
	Td1nation         *string `json:"td1nation,omitempty" example:"AUT"`
	Td1code           *int32  `json:"td1code,omitempty" example:"101"`
	Td2id             *int32  `json:"td2id,omitempty"`
	Td2name           *string `json:"td2name,omitempty" example:"Jane Doe"`
	Td2nation         *string `json:"td2nation,omitempty" example:"GER"`
	Td2code           *int32  `json:"td2code,omitempty" example:"102"`
	Calstatuscode     *string `json:"calstatuscode,omitempty" example:"OFF"`
	Procstatuscode    *string `json:"procstatuscode,omitempty" example:"P"`
	Pursuit           *string `json:"pursuit,omitempty"`
	Masse             *string `json:"masse,omitempty"`
	Relay             *string `json:"relay,omitempty"`
	Distance          *string `json:"distance,omitempty"`
	Hill              *int32  `json:"hill,omitempty" example:"140"`
	Style             *string `json:"style,omitempty" example:"CLASSIC"`
	Qualif            *string `json:"qualif,omitempty" example:"Q"`
	Finale            *string `json:"finale,omitempty"`
	Homol             *string `json:"homol,omitempty"`
	Webcomment        *string `json:"webcomment,omitempty"`
	Displaystatus     *string `json:"displaystatus,omitempty" example:"OK"`
	Fisinterncomment  *string `json:"fisinterncomment,omitempty"`
	Published         *int32  `json:"published,omitempty" example:"1"`
	Validforfispoints *int32  `json:"validforfispoints,omitempty" example:"1"`
	Usedfislist       *string `json:"usedfislist,omitempty"`
	Tolist            *string `json:"tolist,omitempty"`
	Discforlistcode   *string `json:"discforlistcode,omitempty"`
	Calculatedpenalty *string `json:"calculatedpenalty,omitempty"`
	Appliedpenalty    *string `json:"appliedpenalty,omitempty"`
	Appliedscala      *string `json:"appliedscala,omitempty"`
	Penscafixed       *string `json:"penscafixed,omitempty"`
	Version           *int32  `json:"version,omitempty" example:"1"`
	Nationraceid      *int32  `json:"nationraceid,omitempty"`
	Provraceid        *int32  `json:"provraceid,omitempty"`
	Msql7evid         *int32  `json:"msql7evid,omitempty"`
	Mssql7id          *int32  `json:"mssql7id,omitempty"`
	Topbanner         *string `json:"topbanner,omitempty"`
	Bottombanner      *string `json:"bottombanner,omitempty"`
	Toplogo           *string `json:"toplogo,omitempty"`
	Bottomlogo        *string `json:"bottomlogo,omitempty"`
	Gallery           *string `json:"gallery,omitempty"`
	Indi              *int32  `json:"indi,omitempty" example:"1"`
	Team              *int32  `json:"team,omitempty" example:"0"`
	Tabcount          *int32  `json:"tabcount,omitempty"`
	Columncount       *int32  `json:"columncount,omitempty"`
	Level             *string `json:"level,omitempty" example:"WC"`
	Hloc1             *string `json:"hloc1,omitempty" format:"date-time" example:"2025-01-25T15:00:00Z"`
	Hloc2             *string `json:"hloc2,omitempty" format:"date-time" example:"2025-01-25T15:30:00Z"`
	Hloc3             *string `json:"hloc3,omitempty" format:"date-time" example:"2025-01-25T16:00:00Z"`
	Hcet1             *string `json:"hcet1,omitempty" format:"date-time" example:"2025-01-25T15:00:00Z"`
	Hcet2             *string `json:"hcet2,omitempty" format:"date-time" example:"2025-01-25T15:30:00Z"`
	Hcet3             *string `json:"hcet3,omitempty" format:"date-time" example:"2025-01-25T16:00:00Z"`
	Live              *int32  `json:"live,omitempty" example:"1"`
	Livestatus1       *string `json:"livestatus1,omitempty" example:"LIVE"`
	Livestatus2       *string `json:"livestatus2,omitempty"`
	Livestatus3       *string `json:"livestatus3,omitempty"`
	Liveinfo1         *string `json:"liveinfo1,omitempty"`
	Liveinfo2         *string `json:"liveinfo2,omitempty"`
	Liveinfo3         *string `json:"liveinfo3,omitempty"`
	Passwd            *string `json:"passwd,omitempty"`
	Timinglogo        *string `json:"timinglogo,omitempty"`
	Validdate         *string `json:"validdate,omitempty" format:"date-time"`
	Noepr             *int32  `json:"noepr,omitempty"`
	Tddoc             *int32  `json:"tddoc,omitempty"`
	Timingreport      *int32  `json:"timingreport,omitempty"`
	SpecialCupPoints  *int32  `json:"special_cup_points,omitempty"`
	SkipWcsl          *int32  `json:"skip_wcsl,omitempty"`
	Validforowg       *string `json:"validforowg,omitempty" example:"Y"`
	Lastupdate        *string `json:"lastupdate,omitempty" format:"date-time" example:"2025-01-25T18:30:00Z"`
}

type FISUpdateRaceJPExample struct {
	Raceid            int32   `json:"raceid" validate:"required" example:"234567"`
	Eventid           *int32  `json:"eventid,omitempty" example:"8901"`
	Seasoncode        *int32  `json:"seasoncode,omitempty" example:"2027"`
	Racecodex         *int32  `json:"racecodex,omitempty" example:"2001"`
	Disciplineid      *string `json:"disciplineid,omitempty" example:"HS140"`
	Disciplinecode    *string `json:"disciplinecode,omitempty" example:"LH"`
	Catcode           *string `json:"catcode,omitempty" example:"WC"`
	Catcode2          *string `json:"catcode2,omitempty"`
	Catcode3          *string `json:"catcode3,omitempty"`
	Catcode4          *string `json:"catcode4,omitempty"`
	Gender            *string `json:"gender,omitempty" example:"M"`
	Racedate          *string `json:"racedate,omitempty" format:"date" example:"2025-01-25"`
	Starteventdate    *string `json:"starteventdate,omitempty" format:"date" example:"2025-01-24"`
	Description       *string `json:"description,omitempty" example:"World Cup HS140 Individual (updated)"`
	Place             *string `json:"place,omitempty" example:"Planica"`
	Nationcode        *string `json:"nationcode,omitempty" example:"SLO"`
	Receiveddate      *string `json:"receiveddate,omitempty" format:"date-time"`
	Td1id             *int32  `json:"td1id,omitempty"`
	Td1name           *string `json:"td1name,omitempty"`
	Td1nation         *string `json:"td1nation,omitempty"`
	Td1code           *int32  `json:"td1code,omitempty"`
	Td2id             *int32  `json:"td2id,omitempty"`
	Td2name           *string `json:"td2name,omitempty"`
	Td2nation         *string `json:"td2nation,omitempty"`
	Td2code           *int32  `json:"td2code,omitempty"`
	Calstatuscode     *string `json:"calstatuscode,omitempty"`
	Procstatuscode    *string `json:"procstatuscode,omitempty"`
	Pursuit           *string `json:"pursuit,omitempty"`
	Masse             *string `json:"masse,omitempty"`
	Relay             *string `json:"relay,omitempty"`
	Distance          *string `json:"distance,omitempty"`
	Hill              *int32  `json:"hill,omitempty" example:"140"`
	Style             *string `json:"style,omitempty"`
	Qualif            *string `json:"qualif,omitempty"`
	Finale            *string `json:"finale,omitempty"`
	Homol             *string `json:"homol,omitempty"`
	Webcomment        *string `json:"webcomment,omitempty"`
	Displaystatus     *string `json:"displaystatus,omitempty"`
	Fisinterncomment  *string `json:"fisinterncomment,omitempty"`
	Published         *int32  `json:"published,omitempty" example:"1"`
	Validforfispoints *int32  `json:"validforfispoints,omitempty" example:"1"`
	Usedfislist       *string `json:"usedfislist,omitempty"`
	Tolist            *string `json:"tolist,omitempty"`
	Discforlistcode   *string `json:"discforlistcode,omitempty"`
	Calculatedpenalty *string `json:"calculatedpenalty,omitempty"`
	Appliedpenalty    *string `json:"appliedpenalty,omitempty"`
	Appliedscala      *string `json:"appliedscala,omitempty"`
	Penscafixed       *string `json:"penscafixed,omitempty"`
	Version           *int32  `json:"version,omitempty"`
	Nationraceid      *int32  `json:"nationraceid,omitempty"`
	Provraceid        *int32  `json:"provraceid,omitempty"`
	Msql7evid         *int32  `json:"msql7evid,omitempty"`
	Mssql7id          *int32  `json:"mssql7id,omitempty"`
	Topbanner         *string `json:"topbanner,omitempty"`
	Bottombanner      *string `json:"bottombanner,omitempty"`
	Toplogo           *string `json:"toplogo,omitempty"`
	Bottomlogo        *string `json:"bottomlogo,omitempty"`
	Gallery           *string `json:"gallery,omitempty"`
	Indi              *int32  `json:"indi,omitempty"`
	Team              *int32  `json:"team,omitempty"`
	Tabcount          *int32  `json:"tabcount,omitempty"`
	Columncount       *int32  `json:"columncount,omitempty"`
	Level             *string `json:"level,omitempty"`
	Hloc1             *string `json:"hloc1,omitempty" format:"date-time" example:"2025-01-25T15:00:00Z"`
	Hloc2             *string `json:"hloc2,omitempty" format:"date-time" example:"2025-01-25T15:30:00Z"`
	Hloc3             *string `json:"hloc3,omitempty" format:"date-time" example:"2025-01-25T16:00:00Z"`
	Hcet1             *string `json:"hcet1,omitempty" format:"date-time" example:"2025-01-25T15:00:00Z"`
	Hcet2             *string `json:"hcet2,omitempty" format:"date-time" example:"2025-01-25T15:30:00Z"`
	Hcet3             *string `json:"hcet3,omitempty" format:"date-time" example:"2025-01-25T16:00:00Z"`
	Livestatus1       *string `json:"livestatus1,omitempty"`
	Livestatus2       *string `json:"livestatus2,omitempty"`
	Livestatus3       *string `json:"livestatus3,omitempty"`
	Liveinfo1         *string `json:"liveinfo1,omitempty"`
	Liveinfo2         *string `json:"liveinfo2,omitempty"`
	Liveinfo3         *string `json:"liveinfo3,omitempty"`
	Passwd            *string `json:"passwd,omitempty"`
	Timinglogo        *string `json:"timinglogo,omitempty"`
	Validdate         *string `json:"validdate,omitempty" format:"date-time"`
	Noepr             *int32  `json:"noepr,omitempty"`
	Tddoc             *int32  `json:"tddoc,omitempty"`
	Timingreport      *int32  `json:"timingreport,omitempty"`
	SpecialCupPoints  *int32  `json:"special_cup_points,omitempty"`
	SkipWcsl          *int32  `json:"skip_wcsl,omitempty"`
	Validforowg       *string `json:"validforowg,omitempty"`
	Lastupdate        *string `json:"lastupdate,omitempty" format:"date-time" example:"2025-01-26T08:00:00Z"`
}

type FISSeasonsNKResponse struct {
	Seasons []int32 `json:"seasons" swaggertype:"array,integer" example:"2026,2025,2024"`
}

type FISDisciplinesNKResponse struct {
	Disciplines []string `json:"disciplines" swaggertype:"array,string" example:"NH,LH,GH"`
}

type FISCategoriesNKResponse struct {
	Categories []string `json:"categories" swaggertype:"array,string" example:"WC,WSC,COC"`
}

type FISRaceNK struct {
	Raceid         int32   `json:"raceid" example:"8801"`
	Eventid        *int32  `json:"eventid,omitempty" example:"9001"`
	Seasoncode     *int32  `json:"seasoncode,omitempty" example:"2025"`
	Racecodex      *int32  `json:"racecodex,omitempty" example:"3001"`
	Disciplineid   *string `json:"disciplineid,omitempty" example:"NH10K"`
	Disciplinecode *string `json:"disciplinecode,omitempty" example:"NC"`
	Catcode        *string `json:"catcode,omitempty" example:"WC"`
	Catcode2       *string `json:"catcode2,omitempty"`
	Catcode3       *string `json:"catcode3,omitempty"`
	Catcode4       *string `json:"catcode4,omitempty"`
	Gender         *string `json:"gender,omitempty" example:"M"`
	Racedate       *string `json:"racedate,omitempty" format:"date" example:"2025-02-15"`
	Starteventdate *string `json:"starteventdate,omitempty" format:"date" example:"2025-02-14"`
	Description    *string `json:"description,omitempty" example:"World Cup NH 10 km"`
	Place          *string `json:"place,omitempty" example:"Seefeld"`
	Nationcode     *string `json:"nationcode,omitempty" example:"AUT"`

	Td1id          *int32  `json:"td1id,omitempty" example:"101"`
	Td1name        *string `json:"td1name,omitempty" example:"John Doe"`
	Td1nation      *string `json:"td1nation,omitempty" example:"AUT"`
	Td1code        *int32  `json:"td1code,omitempty" example:"1001"`
	Td2id          *int32  `json:"td2id,omitempty" example:"102"`
	Td2name        *string `json:"td2name,omitempty" example:"Jane Doe"`
	Td2nation      *string `json:"td2nation,omitempty" example:"GER"`
	Td2code        *int32  `json:"td2code,omitempty" example:"1002"`
	Calstatuscode  *string `json:"calstatuscode,omitempty" example:"OFF"`
	Procstatuscode *string `json:"procstatuscode,omitempty" example:"P"`

	Receiveddate *string `json:"receiveddate,omitempty" format:"date-time" example:"2025-02-10T10:00:00Z"`
	Pursuit      *string `json:"pursuit,omitempty"`
	Masse        *string `json:"masse,omitempty"`
	Relay        *string `json:"relay,omitempty"`
	Distance     *string `json:"distance,omitempty" example:"10km"`
	Hill         *int32  `json:"hill,omitempty" example:"109"`
	Style        *string `json:"style,omitempty" example:"FREE"`
	Qualif       *string `json:"qualif,omitempty" example:"Q"`
	Finale       *string `json:"finale,omitempty"`
	Homol        *string `json:"homol,omitempty"`

	Webcomment       *string `json:"webcomment,omitempty"`
	Displaystatus    *string `json:"displaystatus,omitempty" example:"OK"`
	Fisinterncomment *string `json:"fisinterncomment,omitempty"`

	Published         *int32  `json:"published,omitempty" example:"1"`
	Validforfispoints *int32  `json:"validforfispoints,omitempty" example:"1"`
	Usedfislist       *string `json:"usedfislist,omitempty"`
	Tolist            *string `json:"tolist,omitempty"`
	Discforlistcode   *string `json:"discforlistcode,omitempty"`
	Calculatedpenalty *string `json:"calculatedpenalty,omitempty"`
	Appliedpenalty    *string `json:"appliedpenalty,omitempty"`
	Appliedscala      *string `json:"appliedscala,omitempty"`
	Penscafixed       *string `json:"penscafixed,omitempty"`
	Version           *int32  `json:"version,omitempty" example:"1"`
	Nationraceid      *int32  `json:"nationraceid,omitempty" example:"0"`
	Provraceid        *int32  `json:"provraceid,omitempty" example:"0"`
	Msql7evid         *int32  `json:"msql7evid,omitempty" example:"0"`
	Mssql7id          *int32  `json:"mssql7id,omitempty" example:"0"`

	Results      *int32  `json:"results,omitempty" example:"0"`
	Pdf          *int32  `json:"pdf,omitempty" example:"0"`
	Topbanner    *string `json:"topbanner,omitempty"`
	Bottombanner *string `json:"bottombanner,omitempty"`
	Toplogo      *string `json:"toplogo,omitempty"`
	Bottomlogo   *string `json:"bottomlogo,omitempty"`
	Gallery      *string `json:"gallery,omitempty"`

	Indi        *int32  `json:"indi,omitempty" example:"1"`
	Team        *int32  `json:"team,omitempty" example:"0"`
	Tabcount    *int32  `json:"tabcount,omitempty" example:"0"`
	Columncount *int32  `json:"columncount,omitempty" example:"0"`
	Level       *string `json:"level,omitempty" example:"WC"`

	Hloc1 *string `json:"hloc1,omitempty" format:"date-time" example:"2025-02-15T09:30:00Z"`
	Hloc2 *string `json:"hloc2,omitempty" format:"date-time" example:"2025-02-15T10:00:00Z"`
	Hloc3 *string `json:"hloc3,omitempty" format:"date-time" example:"2025-02-15T10:30:00Z"`
	Hcet1 *string `json:"hcet1,omitempty" format:"date-time" example:"2025-02-15T11:00:00Z"`
	Hcet2 *string `json:"hcet2,omitempty" format:"date-time" example:"2025-02-15T11:30:00Z"`
	Hcet3 *string `json:"hcet3,omitempty" format:"date-time" example:"2025-02-15T12:00:00Z"`

	Live        *int32  `json:"live,omitempty" example:"1"`
	Livestatus1 *string `json:"livestatus1,omitempty" example:"LIVE"`
	Livestatus2 *string `json:"livestatus2,omitempty"`
	Livestatus3 *string `json:"livestatus3,omitempty"`
	Liveinfo1   *string `json:"liveinfo1,omitempty"`
	Liveinfo2   *string `json:"liveinfo2,omitempty"`
	Liveinfo3   *string `json:"liveinfo3,omitempty"`

	Passwd     *string `json:"passwd,omitempty"`
	Timinglogo *string `json:"timinglogo,omitempty"`

	Validdate    *string `json:"validdate,omitempty" format:"date-time" example:"2025-02-16T00:00:00Z"`
	Noepr        *int32  `json:"noepr,omitempty" example:"0"`
	Tddoc        *int32  `json:"tddoc,omitempty" example:"0"`
	Timingreport *int32  `json:"timingreport,omitempty" example:"0"`

	SpecialCupPoints *int32 `json:"special_cup_points,omitempty" example:"0"`
	SkipWcsl         *int32 `json:"skip_wcsl,omitempty" example:"0"`
	Validforowg      *int32 `json:"validforowg,omitempty" example:"1"`

	Lastupdate *string `json:"lastupdate,omitempty" format:"date-time" example:"2025-02-15T18:30:00Z"`
}

type FISRacesNKResponse struct {
	Races []FISRaceNK `json:"races"`
}

type FISLastRaceNKResponse struct {
	Race FISRaceNK `json:"race"`
}

type FISInsertRaceNKExample struct {
	Raceid         int32   `json:"raceid" validate:"required" example:"8801"`
	Eventid        *int32  `json:"eventid,omitempty" example:"9001"`
	Seasoncode     *int32  `json:"seasoncode,omitempty" example:"2025"`
	Racecodex      *int32  `json:"racecodex,omitempty" example:"3001"`
	Disciplineid   *string `json:"disciplineid,omitempty" example:"NH10K"`
	Disciplinecode *string `json:"disciplinecode,omitempty" example:"NC"`
	Catcode        *string `json:"catcode,omitempty" example:"WC"`
	Catcode2       *string `json:"catcode2,omitempty"`
	Catcode3       *string `json:"catcode3,omitempty"`
	Catcode4       *string `json:"catcode4,omitempty"`
	Gender         *string `json:"gender,omitempty" example:"M"`
	Racedate       *string `json:"racedate,omitempty" format:"date" example:"2025-02-15"`
	Starteventdate *string `json:"starteventdate,omitempty" format:"date" example:"2025-02-14"`
	Description    *string `json:"description,omitempty" example:"World Cup NH 10 km"`
	Place          *string `json:"place,omitempty" example:"Seefeld"`
	Nationcode     *string `json:"nationcode,omitempty" example:"AUT"`

	Td1id          *int32  `json:"td1id,omitempty" example:"101"`
	Td1name        *string `json:"td1name,omitempty" example:"John Doe"`
	Td1nation      *string `json:"td1nation,omitempty" example:"AUT"`
	Td1code        *int32  `json:"td1code,omitempty" example:"1001"`
	Td2id          *int32  `json:"td2id,omitempty" example:"102"`
	Td2name        *string `json:"td2name,omitempty" example:"Jane Doe"`
	Td2nation      *string `json:"td2nation,omitempty" example:"GER"`
	Td2code        *int32  `json:"td2code,omitempty" example:"1002"`
	Calstatuscode  *string `json:"calstatuscode,omitempty" example:"OFF"`
	Procstatuscode *string `json:"procstatuscode,omitempty" example:"P"`

	Receiveddate *string `json:"receiveddate,omitempty" format:"date-time" example:"2025-02-10T10:00:00Z"`
	Pursuit      *string `json:"pursuit,omitempty"`
	Masse        *string `json:"masse,omitempty"`
	Relay        *string `json:"relay,omitempty"`
	Distance     *string `json:"distance,omitempty" example:"10km"`
	Hill         *int32  `json:"hill,omitempty" example:"109"`
	Style        *string `json:"style,omitempty" example:"FREE"`
	Qualif       *string `json:"qualif,omitempty" example:"Q"`
	Finale       *string `json:"finale,omitempty"`
	Homol        *string `json:"homol,omitempty"`

	Webcomment       *string `json:"webcomment,omitempty"`
	Displaystatus    *string `json:"displaystatus,omitempty" example:"OK"`
	Fisinterncomment *string `json:"fisinterncomment,omitempty"`

	Published         *int32  `json:"published,omitempty" example:"1"`
	Validforfispoints *int32  `json:"validforfispoints,omitempty" example:"1"`
	Usedfislist       *string `json:"usedfislist,omitempty"`
	Tolist            *string `json:"tolist,omitempty"`
	Discforlistcode   *string `json:"discforlistcode,omitempty"`
	Calculatedpenalty *string `json:"calculatedpenalty,omitempty"`
	Appliedpenalty    *string `json:"appliedpenalty,omitempty"`
	Appliedscala      *string `json:"appliedscala,omitempty"`
	Penscafixed       *string `json:"penscafixed,omitempty"`
	Version           *int32  `json:"version,omitempty" example:"1"`
	Nationraceid      *int32  `json:"nationraceid,omitempty" example:"0"`
	Provraceid        *int32  `json:"provraceid,omitempty" example:"0"`
	Msql7evid         *int32  `json:"msql7evid,omitempty" example:"0"`
	Mssql7id          *int32  `json:"mssql7id,omitempty" example:"0"`

	Topbanner    *string `json:"topbanner,omitempty"`
	Bottombanner *string `json:"bottombanner,omitempty"`
	Toplogo      *string `json:"toplogo,omitempty"`
	Bottomlogo   *string `json:"bottomlogo,omitempty"`
	Gallery      *string `json:"gallery,omitempty"`

	Indi        *int32  `json:"indi,omitempty" example:"1"`
	Team        *int32  `json:"team,omitempty" example:"0"`
	Tabcount    *int32  `json:"tabcount,omitempty" example:"0"`
	Columncount *int32  `json:"columncount,omitempty" example:"0"`
	Level       *string `json:"level,omitempty" example:"WC"`

	Hloc1 *string `json:"hloc1,omitempty" format:"date-time" example:"2025-02-15T09:30:00Z"`
	Hloc2 *string `json:"hloc2,omitempty" format:"date-time" example:"2025-02-15T10:00:00Z"`
	Hloc3 *string `json:"hloc3,omitempty" format:"date-time" example:"2025-02-15T10:30:00Z"`
	Hcet1 *string `json:"hcet1,omitempty" format:"date-time" example:"2025-02-15T11:00:00Z"`
	Hcet2 *string `json:"hcet2,omitempty" format:"date-time" example:"2025-02-15T11:30:00Z"`
	Hcet3 *string `json:"hcet3,omitempty" format:"date-time" example:"2025-02-15T12:00:00Z"`

	Live        *int32  `json:"live,omitempty" example:"1"`
	Livestatus1 *string `json:"livestatus1,omitempty" example:"LIVE"`
	Livestatus2 *string `json:"livestatus2,omitempty"`
	Livestatus3 *string `json:"livestatus3,omitempty"`
	Liveinfo1   *string `json:"liveinfo1,omitempty"`
	Liveinfo2   *string `json:"liveinfo2,omitempty"`
	Liveinfo3   *string `json:"liveinfo3,omitempty"`

	Passwd     *string `json:"passwd,omitempty"`
	Timinglogo *string `json:"timinglogo,omitempty"`

	Validdate    *string `json:"validdate,omitempty" format:"date-time" example:"2025-02-16T00:00:00Z"`
	Noepr        *int32  `json:"noepr,omitempty" example:"0"`
	Tddoc        *int32  `json:"tddoc,omitempty" example:"0"`
	Timingreport *int32  `json:"timingreport,omitempty" example:"0"`

	SpecialCupPoints *int32 `json:"special_cup_points,omitempty" example:"0"`
	SkipWcsl         *int32 `json:"skip_wcsl,omitempty" example:"0"`
	Validforowg      *int32 `json:"validforowg,omitempty" example:"1"`

	Lastupdate *string `json:"lastupdate,omitempty" format:"date-time" example:"2025-02-15T18:30:00Z"`
}

type FISUpdateRaceNKExample struct {
	Raceid         int32   `json:"raceid" validate:"required" example:"8801"`
	Eventid        *int32  `json:"eventid,omitempty" example:"9001"`
	Seasoncode     *int32  `json:"seasoncode,omitempty" example:"2025"`
	Racecodex      *int32  `json:"racecodex,omitempty" example:"3001"`
	Disciplineid   *string `json:"disciplineid,omitempty" example:"NH10K"`
	Disciplinecode *string `json:"disciplinecode,omitempty" example:"NC"`
	Catcode        *string `json:"catcode,omitempty" example:"WC"`
	Catcode2       *string `json:"catcode2,omitempty"`
	Catcode3       *string `json:"catcode3,omitempty"`
	Catcode4       *string `json:"catcode4,omitempty"`
	Gender         *string `json:"gender,omitempty" example:"M"`
	Racedate       *string `json:"racedate,omitempty" format:"date" example:"2025-02-15"`
	Starteventdate *string `json:"starteventdate,omitempty" format:"date" example:"2025-02-14"`
	Description    *string `json:"description,omitempty" example:"World Cup NH 10 km (updated)"`
	Place          *string `json:"place,omitempty" example:"Seefeld"`
	Nationcode     *string `json:"nationcode,omitempty" example:"AUT"`

	Td1id          *int32  `json:"td1id,omitempty" example:"101"`
	Td1name        *string `json:"td1name,omitempty" example:"John Doe"`
	Td1nation      *string `json:"td1nation,omitempty" example:"AUT"`
	Td1code        *int32  `json:"td1code,omitempty" example:"1001"`
	Td2id          *int32  `json:"td2id,omitempty" example:"102"`
	Td2name        *string `json:"td2name,omitempty" example:"Jane Doe"`
	Td2nation      *string `json:"td2nation,omitempty" example:"GER"`
	Td2code        *int32  `json:"td2code,omitempty" example:"1002"`
	Calstatuscode  *string `json:"calstatuscode,omitempty" example:"OFF"`
	Procstatuscode *string `json:"procstatuscode,omitempty" example:"P"`

	Receiveddate *string `json:"receiveddate,omitempty" format:"date-time" example:"2025-02-10T10:00:00Z"`
	Pursuit      *string `json:"pursuit,omitempty"`
	Masse        *string `json:"masse,omitempty"`
	Relay        *string `json:"relay,omitempty"`
	Distance     *string `json:"distance,omitempty" example:"10km"`
	Hill         *int32  `json:"hill,omitempty" example:"109"`
	Style        *string `json:"style,omitempty" example:"FREE"`
	Qualif       *string `json:"qualif,omitempty" example:"Q"`
	Finale       *string `json:"finale,omitempty"`
	Homol        *string `json:"homol,omitempty"`

	Webcomment       *string `json:"webcomment,omitempty"`
	Displaystatus    *string `json:"displaystatus,omitempty" example:"OK"`
	Fisinterncomment *string `json:"fisinterncomment,omitempty"`

	Published         *int32  `json:"published,omitempty" example:"1"`
	Validforfispoints *int32  `json:"validforfispoints,omitempty" example:"1"`
	Usedfislist       *string `json:"usedfislist,omitempty"`
	Tolist            *string `json:"tolist,omitempty"`
	Discforlistcode   *string `json:"discforlistcode,omitempty"`
	Calculatedpenalty *string `json:"calculatedpenalty,omitempty"`
	Appliedpenalty    *string `json:"appliedpenalty,omitempty"`
	Appliedscala      *string `json:"appliedscala,omitempty"`
	Penscafixed       *string `json:"penscafixed,omitempty"`
	Version           *int32  `json:"version,omitempty" example:"1"`
	Nationraceid      *int32  `json:"nationraceid,omitempty" example:"0"`
	Provraceid        *int32  `json:"provraceid,omitempty" example:"0"`
	Msql7evid         *int32  `json:"msql7evid,omitempty" example:"0"`
	Mssql7id          *int32  `json:"mssql7id,omitempty" example:"0"`

	Topbanner    *string `json:"topbanner,omitempty"`
	Bottombanner *string `json:"bottombanner,omitempty"`
	Toplogo      *string `json:"toplogo,omitempty"`
	Bottomlogo   *string `json:"bottomlogo,omitempty"`
	Gallery      *string `json:"gallery,omitempty"`

	Indi        *int32  `json:"indi,omitempty" example:"1"`
	Team        *int32  `json:"team,omitempty" example:"0"`
	Tabcount    *int32  `json:"tabcount,omitempty" example:"0"`
	Columncount *int32  `json:"columncount,omitempty" example:"0"`
	Level       *string `json:"level,omitempty" example:"WC"`

	Hloc1 *string `json:"hloc1,omitempty" format:"date-time" example:"2025-02-15T09:30:00Z"`
	Hloc2 *string `json:"hloc2,omitempty" format:"date-time" example:"2025-02-15T10:00:00Z"`
	Hloc3 *string `json:"hloc3,omitempty" format:"date-time" example:"2025-02-15T10:30:00Z"`
	Hcet1 *string `json:"hcet1,omitempty" format:"date-time" example:"2025-02-15T11:00:00Z"`
	Hcet2 *string `json:"hcet2,omitempty" format:"date-time" example:"2025-02-15T11:30:00Z"`
	Hcet3 *string `json:"hcet3,omitempty" format:"date-time" example:"2025-02-15T12:00:00Z"`

	Live        *int32  `json:"live,omitempty" example:"1"`
	Livestatus1 *string `json:"livestatus1,omitempty" example:"LIVE"`
	Livestatus2 *string `json:"livestatus2,omitempty"`
	Livestatus3 *string `json:"livestatus3,omitempty"`
	Liveinfo1   *string `json:"liveinfo1,omitempty"`
	Liveinfo2   *string `json:"liveinfo2,omitempty"`
	Liveinfo3   *string `json:"liveinfo3,omitempty"`

	Passwd     *string `json:"passwd,omitempty"`
	Timinglogo *string `json:"timinglogo,omitempty"`

	Validdate    *string `json:"validdate,omitempty" format:"date-time" example:"2025-02-16T00:00:00Z"`
	Noepr        *int32  `json:"noepr,omitempty" example:"0"`
	Tddoc        *int32  `json:"tddoc,omitempty" example:"0"`
	Timingreport *int32  `json:"timingreport,omitempty" example:"0"`

	SpecialCupPoints *int32 `json:"special_cup_points,omitempty" example:"0"`
	SkipWcsl         *int32 `json:"skip_wcsl,omitempty" example:"0"`
	Validforowg      *int32 `json:"validforowg,omitempty" example:"1"`

	Lastupdate *string `json:"lastupdate,omitempty" format:"date-time" example:"2025-02-16T08:00:00Z"`
}

type FISResultCC struct {
	Recid          int32   `json:"recid" example:"12345"`
	Raceid         *int32  `json:"raceid,omitempty" example:"98765"`
	Competitorid   *int32  `json:"competitorid,omitempty" example:"11111"`
	Status         *string `json:"status,omitempty" example:"OK"`
	Reason         *string `json:"reason,omitempty" example:""`
	Position       *string `json:"position,omitempty" example:"1.00000"`
	Pf             *int32  `json:"pf,omitempty" example:"0"`
	Status2        *string `json:"status2,omitempty" example:""`
	Bib            *string `json:"bib,omitempty" example:"10.00000"`
	Bibcolor       *string `json:"bibcolor,omitempty" example:"RED"`
	Fiscode        *int32  `json:"fiscode,omitempty" example:"1234567"`
	Competitorname *string `json:"competitorname,omitempty" example:"DOE John"`
	Nationcode     *string `json:"nationcode,omitempty" example:"NOR"`
	Stage          *string `json:"stage,omitempty" example:"F"`
	Level          *string `json:"level,omitempty" example:"WC"`
	Heat           *string `json:"heat,omitempty" example:"1"`
	Timer1         *string `json:"timer1,omitempty" example:"03:10.5"`
	Timer2         *string `json:"timer2,omitempty" example:"06:22.0"`
	Timer3         *string `json:"timer3,omitempty" example:""`
	Timetot        *string `json:"timetot,omitempty" example:"26:30.2"`
	Valid          *string `json:"valid,omitempty" example:"1.00000"`
	Racepoints     *string `json:"racepoints,omitempty" example:"2.34"`
	Cuppoints      *string `json:"cuppoints,omitempty" example:"100.00000"`
	Bonustime      *string `json:"bonustime,omitempty" example:"00:10.0"`
	Bonuscuppoints *string `json:"bonuscuppoints,omitempty" example:"15"`
	Version        *string `json:"version,omitempty" example:"1"`
	Rg1            *string `json:"rg1,omitempty" example:""`
	Rg2            *string `json:"rg2,omitempty" example:""`
	Lastupdate     *string `json:"lastupdate,omitempty" example:"2025-01-01T12:00:00Z"`
}

type FISAthleteResultCC struct {
	Recid          int32   `json:"recid" example:"12345"`
	Raceid         *int32  `json:"raceid,omitempty" example:"98765"`
	Position       *string `json:"position,omitempty" example:"1.00000"`
	Timetot        *string `json:"timetot,omitempty" example:"26:30.2"`
	Competitorid   *int32  `json:"competitorid,omitempty" example:"11111"`
	Racedate       *string `json:"racedate,omitempty" example:"2025-02-15"`
	Seasoncode     *int32  `json:"seasoncode,omitempty" example:"2025"`
	Disciplinecode *string `json:"disciplinecode,omitempty" example:"DSPR"`
	Catcode        *string `json:"catcode,omitempty" example:"WC"`
	Place          *string `json:"place,omitempty" example:"Oslo"`
}

type FISLastResultCCResponse struct {
	Result FISResultCC `json:"result"`
}

type FISRaceResultsCCResponse struct {
	Results []FISResultCC `json:"results"`
}

type FISAthleteResultsCCResponse struct {
	Results []FISAthleteResultCC `json:"results"`
}

type FISInsertResultCCExample struct {
	Recid          int32   `json:"recid" example:"12345"`
	Raceid         *int32  `json:"raceid" example:"98765"`
	Competitorid   *int32  `json:"competitorid" example:"11111"`
	Status         *string `json:"status" example:"OK"`
	Reason         *string `json:"reason" example:""`
	Position       *string `json:"position" example:"1.00000"`
	Pf             *int32  `json:"pf" example:"0"`
	Status2        *string `json:"status2" example:""`
	Bib            *string `json:"bib" example:"10.00000"`
	Bibcolor       *string `json:"bibcolor" example:"RED"`
	Fiscode        *int32  `json:"fiscode" example:"1234567"`
	Competitorname *string `json:"competitorname" example:"DOE John"`
	Nationcode     *string `json:"nationcode" example:"NOR"`
	Stage          *string `json:"stage" example:"F"`
	Level          *string `json:"level" example:"WC"`
	Heat           *string `json:"heat" example:"1"`
	Timer1         *string `json:"timer1" example:"03:10.5"`
	Timer2         *string `json:"timer2" example:"06:22.0"`
	Timer3         *string `json:"timer3" example:""`
	Timetot        *string `json:"timetot" example:"26:30.2"`
	Valid          *string `json:"valid" example:"1.00000"`
	Racepoints     *string `json:"racepoints" example:"2.34"`
	Cuppoints      *string `json:"cuppoints" example:"100.00000"`
	Bonustime      *string `json:"bonustime" example:"00:10.0"`
	Bonuscuppoints *string `json:"bonuscuppoints" example:"15"`
	Version        *string `json:"version" example:"1"`
	Rg1            *string `json:"rg1" example:""`
	Rg2            *string `json:"rg2" example:""`
	Lastupdate     *string `json:"lastupdate" example:"2025-01-01T12:00:00Z"`
}

type FISUpdateResultCCExample = FISInsertResultCCExample

type FISResultJP struct {
	Recid          int32   `json:"recid" example:"50001"`
	Raceid         *int32  `json:"raceid,omitempty" example:"234567"`
	Competitorid   *int32  `json:"competitorid,omitempty" example:"345678"`
	Status         *string `json:"status,omitempty" example:"OK"`
	Status2        *string `json:"status2,omitempty" example:""`
	Position       *int32  `json:"position,omitempty" example:"1"`
	Bib            *int32  `json:"bib,omitempty" example:"15"`
	Fiscode        *int32  `json:"fiscode,omitempty" example:"1234567"`
	Competitorname *string `json:"competitorname,omitempty" example:"KOBAYASHI Ryoyu"`
	Nationcode     *string `json:"nationcode,omitempty" example:"JPN"`
	Level          *string `json:"level,omitempty" example:"WC"`
	Heat           *string `json:"heat,omitempty" example:"1"`
	Stage          *string `json:"stage,omitempty" example:"F"`

	J1r1 *string `json:"j1r1,omitempty"`
	J2r1 *string `json:"j2r1,omitempty"`
	J3r1 *string `json:"j3r1,omitempty"`
	J4r1 *string `json:"j4r1,omitempty"`
	J5r1 *string `json:"j5r1,omitempty"`

	Speedr1  *string `json:"speedr1,omitempty" example:"88.5"`
	Distr1   *string `json:"distr1,omitempty" example:"138.0"`
	Disptsr1 *string `json:"disptsr1,omitempty" example:"78.5"`
	Judptsr1 *string `json:"judptsr1,omitempty" example:"57.0"`
	Totrun1  *string `json:"totrun1,omitempty" example:"135.5"`
	Posr1    *string `json:"posr1,omitempty" example:"1"`
	Statusr1 *string `json:"statusr1,omitempty" example:"OK"`

	J1r2 *string `json:"j1r2,omitempty"`
	J2r2 *string `json:"j2r2,omitempty"`
	J3r2 *string `json:"j3r2,omitempty"`
	J4r2 *string `json:"j4r2,omitempty"`
	J5r2 *string `json:"j5r2,omitempty"`

	Speedr2  *string `json:"speedr2,omitempty" example:"88.0"`
	Distr2   *string `json:"distr2,omitempty" example:"137.0"`
	Disptsr2 *string `json:"disptsr2,omitempty" example:"77.9"`
	Judptsr2 *string `json:"judptsr2,omitempty" example:"56.5"`
	Totrun2  *string `json:"totrun2,omitempty" example:"134.4"`
	Posr2    *string `json:"posr2,omitempty" example:"1"`
	Statusr2 *string `json:"statusr2,omitempty" example:"OK"`

	J1r3 *string `json:"j1r3,omitempty"`
	J2r3 *string `json:"j2r3,omitempty"`
	J3r3 *string `json:"j3r3,omitempty"`
	J4r3 *string `json:"j4r3,omitempty"`
	J5r3 *string `json:"j5r3,omitempty"`

	Speedr3  *string `json:"speedr3,omitempty"`
	Distr3   *string `json:"distr3,omitempty"`
	Disptsr3 *string `json:"disptsr3,omitempty"`
	Judptsr3 *string `json:"judptsr3,omitempty"`
	Totrun3  *string `json:"totrun3,omitempty"`
	Posr3    *string `json:"posr3,omitempty"`
	Statusr3 *string `json:"statusr3,omitempty"`

	J1r4 *string `json:"j1r4,omitempty"`
	J2r4 *string `json:"j2r4,omitempty"`
	J3r4 *string `json:"j3r4,omitempty"`
	J4r4 *string `json:"j4r4,omitempty"`
	J5r4 *string `json:"j5r4,omitempty"`

	Speedr4  *string `json:"speedr4,omitempty"`
	Distr4   *string `json:"distr4,omitempty"`
	Disptsr4 *string `json:"disptsr4,omitempty"`
	Judptsr4 *string `json:"judptsr4,omitempty"`

	Gater1    *string `json:"gater1,omitempty" example:"8"`
	Gater2    *string `json:"gater2,omitempty" example:"8"`
	Gater3    *string `json:"gater3,omitempty"`
	Gater4    *string `json:"gater4,omitempty"`
	Gateptsr1 *string `json:"gateptsr1,omitempty" example:"0.0"`
	Gateptsr2 *string `json:"gateptsr2,omitempty" example:"0.0"`
	Gateptsr3 *string `json:"gateptsr3,omitempty"`
	Gateptsr4 *string `json:"gateptsr4,omitempty"`

	Windr1    *string `json:"windr1,omitempty" example:"-0.5"`
	Windr2    *string `json:"windr2,omitempty" example:"-0.3"`
	Windr3    *string `json:"windr3,omitempty"`
	Windr4    *string `json:"windr4,omitempty"`
	Windptsr1 *string `json:"windptsr1,omitempty" example:"1.2"`
	Windptsr2 *string `json:"windptsr2,omitempty" example:"0.8"`
	Windptsr3 *string `json:"windptsr3,omitempty"`
	Windptsr4 *string `json:"windptsr4,omitempty"`

	Reason     *string `json:"reason,omitempty" example:""`
	Totrun4    *string `json:"totrun4,omitempty"`
	Tot        *string `json:"tot,omitempty" example:"269.9"`
	Valid      *int32  `json:"valid,omitempty" example:"1"`
	Racepoints *string `json:"racepoints,omitempty" example:"0.0"`
	Cuppoints  *string `json:"cuppoints,omitempty" example:"100"`
	Version    *string `json:"version,omitempty" example:"1"`
	Lastupdate *string `json:"lastupdate,omitempty" example:"2025-01-25T18:30:00Z"`
	Posr4      *string `json:"posr4,omitempty"`
	Statusr4   *string `json:"statusr4,omitempty"`
}

type FISAthleteResultJP struct {
	Raceid         *int32  `json:"raceid,omitempty" example:"234567"`
	Position       *int32  `json:"position,omitempty" example:"1"`
	Racedate       *string `json:"racedate,omitempty" example:"2025-01-25"`
	Seasoncode     *int32  `json:"seasoncode,omitempty" example:"2025"`
	Disciplinecode *string `json:"disciplinecode,omitempty" example:"LH"`
	Catcode        *string `json:"catcode,omitempty" example:"WC"`
	Place          *string `json:"place,omitempty" example:"Planica"`

	Posr1     *string `json:"posr1,omitempty" example:"1"`
	Speedr1   *string `json:"speedr1,omitempty" example:"88.5"`
	Distr1    *string `json:"distr1,omitempty" example:"138.0"`
	Judptsr1  *string `json:"judptsr1,omitempty" example:"57.0"`
	Windr1    *string `json:"windr1,omitempty" example:"-0.5"`
	Windptsr1 *string `json:"windptsr1,omitempty" example:"1.2"`
	Gater1    *string `json:"gater1,omitempty" example:"8"`

	Posr2     *string `json:"posr2,omitempty" example:"1"`
	Speedr2   *string `json:"speedr2,omitempty" example:"88.0"`
	Distr2    *string `json:"distr2,omitempty" example:"137.0"`
	Judptsr2  *string `json:"judptsr2,omitempty" example:"56.5"`
	Windr2    *string `json:"windr2,omitempty" example:"-0.3"`
	Windptsr2 *string `json:"windptsr2,omitempty" example:"0.8"`
	Gater2    *string `json:"gater2,omitempty" example:"8"`

	Totrun1 *string `json:"totrun1,omitempty" example:"135.5"`
	Totrun2 *string `json:"totrun2,omitempty" example:"134.4"`
}

type FISLastResultJPResponse struct {
	Result FISResultJP `json:"result"`
}

type FISRaceResultsJPResponse struct {
	Results []FISResultJP `json:"results"`
}

type FISAthleteResultsJPResponse struct {
	Results []FISAthleteResultJP `json:"results"`
}

type FISInsertResultJPExample struct {
	Recid          int32   `json:"recid" example:"50001"`
	Raceid         *int32  `json:"raceid" example:"234567"`
	Competitorid   *int32  `json:"competitorid" example:"345678"`
	Status         *string `json:"status" example:"OK"`
	Status2        *string `json:"status2" example:""`
	Position       *int32  `json:"position" example:"1"`
	Bib            *int32  `json:"bib" example:"15"`
	Fiscode        *int32  `json:"fiscode" example:"1234567"`
	Competitorname *string `json:"competitorname" example:"KOBAYASHI Ryoyu"`
	Nationcode     *string `json:"nationcode" example:"JPN"`
	Level          *string `json:"level" example:"WC"`
	Heat           *string `json:"heat" example:"1"`
	Stage          *string `json:"stage" example:"F"`

	J1r1 *string `json:"j1r1,omitempty" example:"19.0"`
	J2r1 *string `json:"j2r1,omitempty" example:"19.0"`
	J3r1 *string `json:"j3r1,omitempty" example:"19.0"`
	J4r1 *string `json:"j4r1,omitempty" example:"19.0"`
	J5r1 *string `json:"j5r1,omitempty" example:"19.0"`

	Speedr1  *string `json:"speedr1" example:"88.5"`
	Distr1   *string `json:"distr1" example:"138.0"`
	Disptsr1 *string `json:"disptsr1" example:"78.5"`
	Judptsr1 *string `json:"judptsr1" example:"57.0"`
	Totrun1  *string `json:"totrun1" example:"135.5"`
	Posr1    *string `json:"posr1" example:"1"`
	Statusr1 *string `json:"statusr1" example:"OK"`

	J1r2 *string `json:"j1r2,omitempty" example:"18.5"`
	J2r2 *string `json:"j2r2,omitempty" example:"18.5"`
	J3r2 *string `json:"j3r2,omitempty" example:"18.5"`
	J4r2 *string `json:"j4r2,omitempty" example:"18.5"`
	J5r2 *string `json:"j5r2,omitempty" example:"18.5"`

	Speedr2  *string `json:"speedr2" example:"88.0"`
	Distr2   *string `json:"distr2" example:"137.0"`
	Disptsr2 *string `json:"disptsr2" example:"77.9"`
	Judptsr2 *string `json:"judptsr2" example:"56.5"`
	Totrun2  *string `json:"totrun2" example:"134.4"`
	Posr2    *string `json:"posr2" example:"1"`
	Statusr2 *string `json:"statusr2" example:"OK"`

	J1r3 *string `json:"j1r3,omitempty"`
	J2r3 *string `json:"j2r3,omitempty"`
	J3r3 *string `json:"j3r3,omitempty"`
	J4r3 *string `json:"j4r3,omitempty"`
	J5r3 *string `json:"j5r3,omitempty"`

	Speedr3  *string `json:"speedr3,omitempty"`
	Distr3   *string `json:"distr3,omitempty"`
	Disptsr3 *string `json:"disptsr3,omitempty"`
	Judptsr3 *string `json:"judptsr3,omitempty"`
	Totrun3  *string `json:"totrun3,omitempty"`
	Posr3    *string `json:"posr3,omitempty"`
	Statusr3 *string `json:"statusr3,omitempty"`

	J1r4 *string `json:"j1r4,omitempty"`
	J2r4 *string `json:"j2r4,omitempty"`
	J3r4 *string `json:"j3r4,omitempty"`
	J4r4 *string `json:"j4r4,omitempty"`
	J5r4 *string `json:"j5r4,omitempty"`

	Speedr4  *string `json:"speedr4,omitempty"`
	Distr4   *string `json:"distr4,omitempty"`
	Disptsr4 *string `json:"disptsr4,omitempty"`
	Judptsr4 *string `json:"judptsr4,omitempty"`

	Gater1    *string `json:"gater1,omitempty" example:"8"`
	Gater2    *string `json:"gater2,omitempty" example:"8"`
	Gater3    *string `json:"gater3,omitempty"`
	Gater4    *string `json:"gater4,omitempty"`
	Gateptsr1 *string `json:"gateptsr1,omitempty" example:"0.0"`
	Gateptsr2 *string `json:"gateptsr2,omitempty" example:"0.0"`
	Gateptsr3 *string `json:"gateptsr3,omitempty"`
	Gateptsr4 *string `json:"gateptsr4,omitempty"`

	Windr1    *string `json:"windr1,omitempty" example:"-0.5"`
	Windr2    *string `json:"windr2,omitempty" example:"-0.3"`
	Windr3    *string `json:"windr3,omitempty"`
	Windr4    *string `json:"windr4,omitempty"`
	Windptsr1 *string `json:"windptsr1,omitempty" example:"1.2"`
	Windptsr2 *string `json:"windptsr2,omitempty" example:"0.8"`
	Windptsr3 *string `json:"windptsr3,omitempty"`
	Windptsr4 *string `json:"windptsr4,omitempty"`

	Reason     *string `json:"reason,omitempty" example:""`
	Totrun4    *string `json:"totrun4,omitempty"`
	Tot        *string `json:"tot" example:"269.9"`
	Valid      *int32  `json:"valid" example:"1"`
	Racepoints *string `json:"racepoints" example:"0.0"`
	Cuppoints  *string `json:"cuppoints" example:"100"`
	Version    *string `json:"version" example:"1"`
	Lastupdate *string `json:"lastupdate" example:"2025-01-25T18:30:00Z"`
	Posr4      *string `json:"posr4,omitempty"`
	Statusr4   *string `json:"statusr4,omitempty"`
}

type FISUpdateResultJPExample = FISInsertResultJPExample

type FISResultNK struct {
	Recid          int32   `json:"recid" example:"70001"`
	Raceid         *int32  `json:"raceid,omitempty" example:"8801"`
	Competitorid   *int32  `json:"competitorid,omitempty" example:"123456"`
	Status         *string `json:"status,omitempty" example:"OK"`
	Status2        *string `json:"status2,omitempty" example:""`
	Reason         *string `json:"reason,omitempty" example:""`
	Position       *int32  `json:"position,omitempty" example:"1"`
	Pf             *int32  `json:"pf,omitempty" example:"0"`
	Bib            *int32  `json:"bib,omitempty" example:"15"`
	Bibcolor       *string `json:"bibcolor,omitempty" example:"RED"`
	Fiscode        *int32  `json:"fiscode,omitempty" example:"3456789"`
	Competitorname *string `json:"competitorname,omitempty" example:"RIIBER Jarl Magnus"`
	Nationcode     *string `json:"nationcode,omitempty" example:"NOR"`
	Level          *string `json:"level,omitempty" example:"WC"`
	Heat           *string `json:"heat,omitempty" example:"1"`
	Stage          *string `json:"stage,omitempty" example:"F"`

	// Jump round 1
	J1r1      *string `json:"j1r1,omitempty" example:"18.5"`
	J2r1      *string `json:"j2r1,omitempty" example:"18.5"`
	J3r1      *string `json:"j3r1,omitempty" example:"18.5"`
	J4r1      *string `json:"j4r1,omitempty" example:"18.5"`
	J5r1      *string `json:"j5r1,omitempty" example:"18.5"`
	Speedr1   *string `json:"speedr1,omitempty" example:"89.5"`
	Distr1    *string `json:"distr1,omitempty" example:"135.0"`
	Disptsr1  *string `json:"disptsr1,omitempty" example:"76.5"`
	Judptsr1  *string `json:"judptsr1,omitempty" example:"56.0"`
	Gater1    *string `json:"gater1,omitempty" example:"10"`
	Gateptsr1 *string `json:"gateptsr1,omitempty" example:"-3.0"`
	Windr1    *string `json:"windr1,omitempty" example:"-0.8"`
	Windptsr1 *string `json:"windptsr1,omitempty" example:"2.5"`
	Totrun1   *string `json:"totrun1,omitempty" example:"134.0"`
	Posr1     *string `json:"posr1,omitempty" example:"1"`
	Statusr1  *string `json:"statusr1,omitempty" example:"OK"`

	// Jump round 2
	J1r2      *string `json:"j1r2,omitempty" example:"18.0"`
	J2r2      *string `json:"j2r2,omitempty" example:"18.0"`
	J3r2      *string `json:"j3r2,omitempty" example:"18.0"`
	J4r2      *string `json:"j4r2,omitempty" example:"18.0"`
	J5r2      *string `json:"j5r2,omitempty" example:"18.0"`
	Speedr2   *string `json:"speedr2,omitempty" example:"90.0"`
	Distr2    *string `json:"distr2,omitempty" example:"134.0"`
	Disptsr2  *string `json:"disptsr2,omitempty" example:"75.0"`
	Judptsr2  *string `json:"judptsr2,omitempty" example:"55.0"`
	Gater2    *string `json:"gater2,omitempty" example:"11"`
	Gateptsr2 *string `json:"gateptsr2,omitempty" example:"-3.5"`
	Windr2    *string `json:"windr2,omitempty" example:"-0.3"`
	Windptsr2 *string `json:"windptsr2,omitempty" example:"1.5"`
	Totrun2   *string `json:"totrun2,omitempty" example:"132.5"`
	Posr2     *string `json:"posr2,omitempty" example:"1"`
	Statusr2  *string `json:"statusr2,omitempty" example:"OK"`

	// Jump summary
	Pointsjump *string `json:"pointsjump,omitempty" example:"135.0"`
	Behindjump *string `json:"behindjump,omitempty" example:"0.0"`
	Posjump    *string `json:"posjump,omitempty" example:"1"`

	// Cross-country part
	Timecc    *string `json:"timecc,omitempty" example:"00:23:15.6"`
	Timeccint *string `json:"timeccint,omitempty" example:"1395"`
	Poscc     *string `json:"poscc,omitempty" example:"2"`
	Starttime *string `json:"starttime,omitempty" example:"12:30:00"`
	Statuscc  *string `json:"statuscc,omitempty" example:"OK"`
	Totbehind *string `json:"totbehind,omitempty" example:"5.6"`

	// Total result
	Timetot    *string `json:"timetot,omitempty" example:"00:24:10.2"`
	Timetotint *int32  `json:"timetotint,omitempty" example:"1450"`
	Valid      *int32  `json:"valid,omitempty" example:"1"`
	Racepoints *string `json:"racepoints,omitempty" example:"100.0"`
	Cuppoints  *string `json:"cuppoints,omitempty" example:"100"`
	Version    *string `json:"version,omitempty" example:"1"`
	Lastupdate *string `json:"lastupdate,omitempty" example:"2025-02-15T18:30:00Z"`
}

type FISAthleteResultNK struct {
	Recid          int32   `json:"recid" example:"70001"`
	Raceid         *int32  `json:"raceid,omitempty" example:"8801"`
	Position       *int32  `json:"position,omitempty" example:"1"`
	Racedate       *string `json:"racedate,omitempty" example:"2025-02-15"`
	Seasoncode     *int32  `json:"seasoncode,omitempty" example:"2025"`
	Distance       *string `json:"distance,omitempty" example:"10km"`
	Hill           *int32  `json:"hill,omitempty" example:"109"`
	Disciplinecode *string `json:"disciplinecode,omitempty" example:"NC"`
	Catcode        *string `json:"catcode,omitempty" example:"WC"`
	Place          *string `json:"place,omitempty" example:"Seefeld"`

	Posr1     *string `json:"posr1,omitempty" example:"1"`
	Speedr1   *string `json:"speedr1,omitempty" example:"89.5"`
	Distr1    *string `json:"distr1,omitempty" example:"135.0"`
	Judptsr1  *string `json:"judptsr1,omitempty" example:"56.0"`
	Windr1    *string `json:"windr1,omitempty" example:"-0.8"`
	Windptsr1 *string `json:"windptsr1,omitempty" example:"2.5"`
	Gater1    *string `json:"gater1,omitempty" example:"10"`
	Totrun1   *string `json:"totrun1,omitempty" example:"134.0"`

	Poscc      *string `json:"poscc,omitempty" example:"2"`
	Timetot    *string `json:"timetot,omitempty" example:"00:24:10.2"`
	Timetotint *int32  `json:"timetotint,omitempty" example:"1450"`
	Pointsjump *string `json:"pointsjump,omitempty" example:"135.0"`
}

type FISLastResultNKResponse struct {
	Result FISResultNK `json:"result"`
}

type FISRaceResultsNKResponse struct {
	Results []FISResultNK `json:"results"`
}

type FISAthleteResultsNKResponse struct {
	Results []FISAthleteResultNK `json:"results"`
}

type FISInsertResultNKExample struct {
	Recid          int32   `json:"recid" example:"70001"`
	Raceid         *int32  `json:"raceid" example:"8801"`
	Competitorid   *int32  `json:"competitorid" example:"123456"`
	Status         *string `json:"status" example:"OK"`
	Status2        *string `json:"status2" example:""`
	Reason         *string `json:"reason" example:""`
	Position       *int32  `json:"position" example:"1"`
	Pf             *int32  `json:"pf" example:"0"`
	Bib            *int32  `json:"bib" example:"15"`
	Bibcolor       *string `json:"bibcolor" example:"RED"`
	Fiscode        *int32  `json:"fiscode" example:"3456789"`
	Competitorname *string `json:"competitorname" example:"RIIBER Jarl Magnus"`
	Nationcode     *string `json:"nationcode" example:"NOR"`
	Level          *string `json:"level" example:"WC"`
	Heat           *string `json:"heat" example:"1"`
	Stage          *string `json:"stage" example:"F"`

	J1r1      *string `json:"j1r1" example:"18.5"`
	J2r1      *string `json:"j2r1" example:"18.5"`
	J3r1      *string `json:"j3r1" example:"18.5"`
	J4r1      *string `json:"j4r1" example:"18.5"`
	J5r1      *string `json:"j5r1" example:"18.5"`
	Speedr1   *string `json:"speedr1" example:"89.5"`
	Distr1    *string `json:"distr1" example:"135.0"`
	Disptsr1  *string `json:"disptsr1" example:"76.5"`
	Judptsr1  *string `json:"judptsr1" example:"56.0"`
	Gater1    *string `json:"gater1" example:"10"`
	Gateptsr1 *string `json:"gateptsr1" example:"-3.0"`
	Windr1    *string `json:"windr1" example:"-0.8"`
	Windptsr1 *string `json:"windptsr1" example:"2.5"`
	Totrun1   *string `json:"totrun1" example:"134.0"`
	Posr1     *string `json:"posr1" example:"1"`
	Statusr1  *string `json:"statusr1" example:"OK"`

	J1r2      *string `json:"j1r2" example:"18.0"`
	J2r2      *string `json:"j2r2" example:"18.0"`
	J3r2      *string `json:"j3r2" example:"18.0"`
	J4r2      *string `json:"j4r2" example:"18.0"`
	J5r2      *string `json:"j5r2" example:"18.0"`
	Speedr2   *string `json:"speedr2" example:"90.0"`
	Distr2    *string `json:"distr2" example:"134.0"`
	Disptsr2  *string `json:"disptsr2" example:"75.0"`
	Judptsr2  *string `json:"judptsr2" example:"55.0"`
	Gater2    *string `json:"gater2" example:"11"`
	Gateptsr2 *string `json:"gateptsr2" example:"-3.5"`
	Windr2    *string `json:"windr2" example:"-0.3"`
	Windptsr2 *string `json:"windptsr2" example:"1.5"`
	Totrun2   *string `json:"totrun2" example:"132.5"`
	Posr2     *string `json:"posr2" example:"1"`
	Statusr2  *string `json:"statusr2" example:"OK"`

	Pointsjump *string `json:"pointsjump" example:"135.0"`
	Behindjump *string `json:"behindjump" example:"0.0"`
	Posjump    *string `json:"posjump" example:"1"`

	Timecc     *string `json:"timecc" example:"00:23:15.6"`
	Timeccint  *string `json:"timeccint" example:"1395"`
	Poscc      *string `json:"poscc" example:"2"`
	Starttime  *string `json:"starttime" example:"12:30:00"`
	Statuscc   *string `json:"statuscc" example:"OK"`
	Totbehind  *string `json:"totbehind" example:"5.6"`
	Timetot    *string `json:"timetot" example:"00:24:10.2"`
	Timetotint *int32  `json:"timetotint" example:"1450"`
	Valid      *int32  `json:"valid" example:"1"`
	Racepoints *string `json:"racepoints" example:"100.0"`
	Cuppoints  *string `json:"cuppoints" example:"100"`
	Version    *string `json:"version" example:"1"`
	Lastupdate *string `json:"lastupdate" example:"2025-02-15T18:30:00Z"`
}

type FISUpdateResultNKExample = FISInsertResultNKExample

type FISAthlete struct {
	Fiscode   *int32  `json:"fiscode,omitempty" example:"342001"`
	Sporttiid *int32  `json:"sporttiid,omitempty" example:"123456"`
	Firstname *string `json:"firstname,omitempty" example:"Iivo"`
	Lastname  *string `json:"lastname,omitempty" example:"Niskanen"`
}

type FISInsertAthleteExample struct {
	Fiscode   int32   `json:"fiscode" example:"342001"`
	Sporttiid *int32  `json:"sporttiid,omitempty" example:"123456"`
	Firstname *string `json:"firstname,omitempty" example:"Iivo"`
	Lastname  *string `json:"lastname,omitempty" example:"Niskanen"`
}

type FISUpdateAthleteExample struct {
	Fiscode   int32   `json:"fiscode" example:"342001"`
	Sporttiid *int32  `json:"sporttiid,omitempty" example:"123456"`
	Firstname *string `json:"firstname,omitempty" example:"Iivo"`
	Lastname  *string `json:"lastname,omitempty" example:"Niskanen"`
}

type FISRacesSearchItem struct {
	Sectorcode     string  `json:"sectorcode" example:"CC"`
	Gender         *string `json:"gender,omitempty" example:"M"`
	Raceid         int32   `json:"raceid" example:"123456"`
	Racedate       *string `json:"racedate,omitempty" example:"2025-02-15"`
	Catcode        *string `json:"catcode,omitempty" example:"WC"`
	Description    *string `json:"description,omitempty" example:"World Cup Sprint"`
	Place          *string `json:"place,omitempty" example:"Lahti"`
	Nationcode     *string `json:"nationcode,omitempty" example:"FIN"`
	Disciplinecode *string `json:"disciplinecode,omitempty" example:"DSPR"`
}

type FISRacesSearchResponse struct {
	Races []FISRacesSearchItem `json:"races"`
}

type FISRaceDetailItem struct {
	Sectorcode string `json:"sectorcode" example:"CC"`

	Raceid         int32   `json:"raceid" example:"123456"`
	Eventid        *int32  `json:"eventid,omitempty" example:"98765"`
	Seasoncode     *int32  `json:"seasoncode,omitempty" example:"2025"`
	Racecodex      *int32  `json:"racecodex,omitempty" example:"1001"`
	Disciplineid   *string `json:"disciplineid,omitempty" example:"DSPR"`
	Disciplinecode *string `json:"disciplinecode,omitempty" example:"DSPR"`
	Catcode        *string `json:"catcode,omitempty" example:"WC"`
	Catcode2       *string `json:"catcode2,omitempty"`
	Catcode3       *string `json:"catcode3,omitempty"`
	Catcode4       *string `json:"catcode4,omitempty"`
	Gender         *string `json:"gender,omitempty" example:"M"`
	Racedate       *string `json:"racedate,omitempty" example:"2025-02-15"`       // YYYY-MM-DD
	Starteventdate *string `json:"starteventdate,omitempty" example:"2025-02-14"` // YYYY-MM-DD
	Description    *string `json:"description,omitempty" example:"World Cup 10 km"`
	Place          *string `json:"place,omitempty" example:"Lahti"`
	Nationcode     *string `json:"nationcode,omitempty" example:"FIN"`

	Td1id     *int32  `json:"td1id,omitempty" example:"1111"`
	Td1name   *string `json:"td1name,omitempty" example:"TD One"`
	Td1nation *string `json:"td1nation,omitempty" example:"NOR"`
	Td1code   *int32  `json:"td1code,omitempty" example:"2222"`
	Td2id     *int32  `json:"td2id,omitempty"`
	Td2name   *string `json:"td2name,omitempty"`
	Td2nation *string `json:"td2nation,omitempty"`
	Td2code   *int32  `json:"td2code,omitempty"`

	Calstatuscode  *string `json:"calstatuscode,omitempty" example:"O"`
	Procstatuscode *string `json:"procstatuscode,omitempty"`
	Receiveddate   *string `json:"receiveddate,omitempty" example:"2025-01-10"` // YYYY-MM-DD

	Pursuit  *string `json:"pursuit,omitempty"`
	Masse    *string `json:"masse,omitempty"`
	Relay    *string `json:"relay,omitempty"`
	Distance *string `json:"distance,omitempty" example:"10 km"` // CC: varchar, JP/NK: same column
	Hill     *string `json:"hill,omitempty"`                     // CC: varchar, JP/NK: integer -> string here
	Style    *string `json:"style,omitempty" example:"CL"`
	Qualif   *string `json:"qualif,omitempty"`
	Finale   *string `json:"finale,omitempty"`
	Homol    *string `json:"homol,omitempty"`

	Webcomment       *string `json:"webcomment,omitempty"`
	Displaystatus    *string `json:"displaystatus,omitempty"`
	Fisinterncomment *string `json:"fisinterncomment,omitempty"`

	Published         *int32  `json:"published,omitempty" example:"1"`
	Validforfispoints *int32  `json:"validforfispoints,omitempty" example:"1"`
	Usedfislist       *string `json:"usedfislist,omitempty"`
	Tolist            *string `json:"tolist,omitempty"`
	Discforlistcode   *string `json:"discforlistcode,omitempty"`
	Calculatedpenalty *string `json:"calculatedpenalty,omitempty"`
	Appliedpenalty    *string `json:"appliedpenalty,omitempty"`
	Appliedscala      *string `json:"appliedscala,omitempty"`
	Penscafixed       *string `json:"penscafixed,omitempty"`

	Version      *int32 `json:"version,omitempty" example:"1"`
	Nationraceid *int32 `json:"nationraceid,omitempty"`
	Provraceid   *int32 `json:"provraceid,omitempty"`
	Msql7evid    *int32 `json:"msql7evid,omitempty"`
	Mssql7id     *int32 `json:"mssql7id,omitempty"`
	Results      *int32 `json:"results,omitempty"`
	Pdf          *int32 `json:"pdf,omitempty"`

	Topbanner    *string `json:"topbanner,omitempty"`
	Bottombanner *string `json:"bottombanner,omitempty"`
	Toplogo      *string `json:"toplogo,omitempty"`
	Bottomlogo   *string `json:"bottomlogo,omitempty"`
	Gallery      *string `json:"gallery,omitempty"`

	Indi        *int32  `json:"indi,omitempty"`
	Team        *int32  `json:"team,omitempty"`
	Tabcount    *int32  `json:"tabcount,omitempty"`
	Columncount *int32  `json:"columncount,omitempty"`
	Level       *string `json:"level,omitempty"`

	Hloc1 *string `json:"hloc1,omitempty" example:"10:00:00"`
	Hloc2 *string `json:"hloc2,omitempty" example:"10:30:00"`
	Hloc3 *string `json:"hloc3,omitempty" example:"11:00:00"`
	Hcet1 *string `json:"hcet1,omitempty" example:"12:00:00"`
	Hcet2 *string `json:"hcet2,omitempty" example:"12:30:00"`
	Hcet3 *string `json:"hcet3,omitempty" example:"13:00:00"`

	Live        *int32  `json:"live,omitempty" example:"1"`
	Livestatus1 *string `json:"livestatus1,omitempty"`
	Livestatus2 *string `json:"livestatus2,omitempty"`
	Livestatus3 *string `json:"livestatus3,omitempty"`
	Liveinfo1   *string `json:"liveinfo1,omitempty"`
	Liveinfo2   *string `json:"liveinfo2,omitempty"`
	Liveinfo3   *string `json:"liveinfo3,omitempty"`

	Passwd     *string `json:"passwd,omitempty"`
	Timinglogo *string `json:"timinglogo,omitempty"`

	Validdate    *string `json:"validdate,omitempty" example:"2025-02-10"`
	Noepr        *int32  `json:"noepr,omitempty"`
	Tddoc        *int32  `json:"tddoc,omitempty"`
	Timingreport *int32  `json:"timingreport,omitempty"`

	SpecialCupPoints *int32  `json:"special_cup_points,omitempty"`
	SkipWcsl         *int32  `json:"skip_wcsl,omitempty"`
	Validforowg      *string `json:"validforowg,omitempty" example:"1"`

	Lastupdate *string `json:"lastupdate,omitempty" example:"2025-02-15T12:34:56Z"`
}

type FISRacesByIDsResponse struct {
	Sector string              `json:"sector" example:"CC"`
	Races  []FISRaceDetailItem `json:"races"`
}

type FISRacesCategoryCountItem struct {
	Catcode *string `json:"catcode,omitempty" example:"WC"`
	Total   int64   `json:"total" example:"42"`
}

type FISRacesCategoryCountsResponse struct {
	Seasoncode int32                       `json:"seasoncode" example:"2025"`
	Sectors    []string                    `json:"sectors" example:"CC,JP"`
	Nationcode *string                     `json:"nationcode,omitempty" example:"FIN"`
	Gender     *string                     `json:"gender,omitempty" example:"M"`
	Categories []FISRacesCategoryCountItem `json:"categories"`
}

type FISRacesNationCountItem struct {
	Nationcode string `json:"nationcode" example:"FIN"`
	Total      int64  `json:"total" example:"42"`
}

type FISRacesNationCountsResponse struct {
	Seasoncode int32                     `json:"seasoncode" example:"2025"`
	Sectors    []string                  `json:"sectors" example:"CC,JP"`
	Gender     *string                   `json:"gender,omitempty" example:"M"`
	Catcode    *string                   `json:"catcode,omitempty" example:"WC"`
	Nations    []FISRacesNationCountItem `json:"nations"`
}

type FISRacesTotalsResponse struct {
	Seasoncode int32    `json:"seasoncode" example:"2025"`
	Sectors    []string `json:"sectors" example:"CC,JP"`
	Catcode    *string  `json:"catcode,omitempty" example:"WC"`
	Gender     *string  `json:"gender,omitempty" example:"M"`
	Total      int64    `json:"total" example:"123"`
}

type FISCompetitorSeasonCatcodeItem struct {
	Seasoncode int32   `json:"seasoncode" example:"2025"`
	Catcode    *string `json:"catcode,omitempty" example:"WC"`
}

type FISCompetitorSeasonsCatcodesResponse struct {
	Competitorid int32                            `json:"competitorid" example:"123456"`
	Sector       string                           `json:"sector" example:"CC"`
	Items        []FISCompetitorSeasonCatcodeItem `json:"items"`
}

type FISLatestResultItem struct {
	Raceid         *int32  `json:"raceid,omitempty" example:"123456"`
	Racedate       *string `json:"racedate,omitempty" example:"2025-02-15"`
	Seasoncode     *int32  `json:"seasoncode,omitempty" example:"2025"`
	Disciplinecode *string `json:"disciplinecode,omitempty" example:"DSPR"`
	Catcode        *string `json:"catcode,omitempty" example:"WC"`
	Place          *string `json:"place,omitempty" example:"Lahti"`
	Nationcode     *string `json:"nationcode,omitempty" example:"FIN"`
	Position       *string `json:"position,omitempty" example:"1"`
	Timetot        *string `json:"timetot,omitempty" example:"25:32.4"`

	// Sector-specific examples (all optional)
	Distance   *string `json:"distance,omitempty" example:"10km"`
	Hill       *int32  `json:"hill,omitempty" example:"140"`
	Posr1      *string `json:"posr1,omitempty" example:"1"`
	Speedr1    *string `json:"speedr1,omitempty" example:"93.5"`
	Distr1     *string `json:"distr1,omitempty" example:"137.0"`
	Judptsr1   *string `json:"judptsr1,omitempty" example:"57.5"`
	Windr1     *string `json:"windr1,omitempty" example:"0.6"`
	Windptsr1  *string `json:"windptsr1,omitempty" example:"-3.0"`
	Gater1     *string `json:"gater1,omitempty" example:"11"`
	Poscc      *string `json:"poscc,omitempty" example:"3"`
	Timetotint *int32  `json:"timetotint,omitempty" example:"1534"`
	Pointsjump *string `json:"pointsjump,omitempty" example:"135.5"`
}

type FISLatestResultsResponse struct {
	Competitorid int32                 `json:"competitorid" example:"123456"`
	Sector       string                `json:"sector" example:"CC"`
	Seasoncode   *int32                `json:"seasoncode,omitempty" example:"2025"`
	Catcodes     []string              `json:"catcodes,omitempty" example:"WC,COC"`
	Limit        int32                 `json:"limit" example:"50"`
	Results      []FISLatestResultItem `json:"results"`
}
