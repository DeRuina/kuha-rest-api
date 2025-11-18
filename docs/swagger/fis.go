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
