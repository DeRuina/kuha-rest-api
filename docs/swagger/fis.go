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
	Competitorid int32   `json:"competitorid" example:"123456"`
	Fiscode      *int32  `json:"fiscode,omitempty" example:"342001"`
	Lastname     *string `json:"lastname,omitempty" example:"Niskanen"`
	Firstname    *string `json:"firstname,omitempty" example:"Iivo"`
	Gender       *string `json:"gender,omitempty" example:"M"`
	Birthdate    *string `json:"birthdate,omitempty" example:"1992-01-12"`
	Nationcode   *string `json:"nationcode,omitempty" example:"FIN"`
	Status       *string `json:"status,omitempty" example:"ACTIVE"`
	Sectorcode   *string `json:"sectorcode,omitempty" example:"CC"`
	StatusDate   *string `json:"status_date,omitempty" example:"2025-10-28T09:45:00Z"`
	Lastupdate   *string `json:"lastupdate,omitempty" example:"2025-10-28T09:45:00Z"`
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
