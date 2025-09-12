package swagger

type UserDataKlabResponse struct {
	Users []string `json:"sportti_ids" example:"12345,67890,54321"`
}

// KlabCustomerResponse mirrors public.customer (51 columns)
type KlabCustomerResponse struct {
	Idcustomer         int32    `json:"idcustomer" example:"7842"`
	Firstname          string   `json:"firstname" example:"Mikael"`
	Lastname           string   `json:"lastname" example:"Järvinen"`
	Idgroups           *int32   `json:"idgroups,omitempty" example:"1205"`
	Dob                *string  `json:"dob,omitempty" example:"1992-08-15"` // date (YYYY-MM-DD)
	Sex                *int32   `json:"sex,omitempty" example:"1"`
	DobYear            *int32   `json:"dob_year,omitempty" example:"1992"`
	DobMonth           *int32   `json:"dob_month,omitempty" example:"8"`
	DobDay             *int32   `json:"dob_day,omitempty" example:"15"`
	PidNumber          *string  `json:"pid_number,omitempty" example:"150892-123A"`
	Company            *string  `json:"company,omitempty" example:"Kuopio Sports Institute"`
	Occupation         *string  `json:"occupation,omitempty" example:"Athlete"`
	Education          *string  `json:"education,omitempty" example:"Bachelor"`
	Address            *string  `json:"address,omitempty" example:"Tulliportinkatu 12, Kuopio"`
	PhoneHome          *string  `json:"phone_home,omitempty" example:"+358-17-1234567"`
	PhoneWork          *string  `json:"phone_work,omitempty" example:"+358-17-7654321"`
	PhoneMobile        *string  `json:"phone_mobile,omitempty" example:"+358-40-1234567"`
	Faxno              *string  `json:"faxno,omitempty"`
	Email              *string  `json:"email,omitempty" example:"mikael.jarvinen@example.com"`
	Username           *string  `json:"username,omitempty" example:"mjarvinen"`
	Password           *string  `json:"password,omitempty" example:"securepass123"`
	Readonly           *int32   `json:"readonly,omitempty" example:"0"`
	Warnings           *int32   `json:"warnings,omitempty" example:"0"`
	AllowToSave        *int32   `json:"allow_to_save,omitempty" example:"1"`
	AllowToCloud       *int32   `json:"allow_to_cloud,omitempty" example:"1"`
	Flag2              *int32   `json:"flag2,omitempty" example:"0"`
	Idsport            *int32   `json:"idsport,omitempty" example:"42"`
	Medication         *string  `json:"medication,omitempty" example:"None"`
	Addinfo            *string  `json:"addinfo,omitempty" example:"Elite cross-country skier"`
	TeamName           *string  `json:"team_name,omitempty" example:"Finnish National Ski Team"`
	Add1               *int32   `json:"add1,omitempty" example:"1"`
	Athlete            *int32   `json:"athlete,omitempty" example:"1"`
	Add10              *string  `json:"add10,omitempty" example:"Field 10"`
	Add20              *string  `json:"add20,omitempty" example:"Field 20"`
	Updatemode         *int32   `json:"updatemode,omitempty" example:"1"`
	WeightKg           *float64 `json:"weight_kg,omitempty" example:"72.5"`
	HeightCm           *float64 `json:"height_cm,omitempty" example:"178.2"`
	DateModified       *float64 `json:"date_modified,omitempty" example:"1693478400.0"` // double precision
	RecomTestlevel     *int32   `json:"recom_testlevel,omitempty" example:"3"`
	CreatedBy          *int64   `json:"created_by,omitempty" example:"1001"`
	ModBy              *int64   `json:"mod_by,omitempty" example:"1001"`
	ModDate            *string  `json:"mod_date,omitempty" example:"2024-08-29T10:30:00Z"`     // timestamp (RFC3339)
	Deleted            *int32   `json:"deleted,omitempty" example:"0"`                         // smallint -> int32 in JSON
	CreatedDate        *string  `json:"created_date,omitempty" example:"2024-08-15T14:22:00Z"` // timestamp (RFC3339)
	Modded             *int32   `json:"modded,omitempty" example:"0"`                          // smallint -> int32 in JSON
	AllowAnonymousData *string  `json:"allow_anonymous_data,omitempty" example:"0"`
	Locked             *int32   `json:"locked,omitempty" example:"0"` // smallint -> int32
	AllowToSprintai    *int32   `json:"allow_to_sprintai,omitempty" example:"1"`
	TosprintaiFrom     *string  `json:"tosprintai_from,omitempty" example:"2024-08-01"` // date (YYYY-MM-DD)
	StatSent           *string  `json:"stat_sent,omitempty" example:"2024-08-20"`       // date (YYYY-MM-DD)
	SporttiID          *string  `json:"sportti_id,omitempty" example:"27353728"`        // varchar(64)
}

type UserKlabResponse struct {
	Customer []KlabCustomerResponse `json:"customer"`
}

type KlabCustomer struct {
	IdCustomer         *int32   `json:"idCustomer" example:"7842"`
	FirstName          *string  `json:"FirstName" example:"Mikael"`
	LastName           *string  `json:"LastName" example:"Järvinen"`
	IdGroups           *int32   `json:"idGroups,omitempty" example:"1205"`
	DOB                *string  `json:"DOB,omitempty" example:"1992-08-15"` // YYYY-MM-DD
	SEX                *int32   `json:"SEX,omitempty" example:"1"`
	DobYear            *int32   `json:"dob_year,omitempty" example:"1992"`
	DobMonth           *int32   `json:"dob_month,omitempty" example:"8"`
	DobDay             *int32   `json:"dob_day,omitempty" example:"15"`
	PidNumber          *string  `json:"pid_number,omitempty" example:"150892-123A"`
	Company            *string  `json:"company,omitempty" example:"Kuopio Sports Institute"`
	Occupation         *string  `json:"occupation,omitempty" example:"Athlete"`
	Education          *string  `json:"education,omitempty" example:"Bachelor"`
	Address            *string  `json:"address,omitempty" example:"Tulliportinkatu 12, Kuopio"`
	PhoneHome          *string  `json:"phone_home,omitempty" example:"+358-17-1234567"`
	PhoneWork          *string  `json:"phone_work,omitempty" example:"+358-17-7654321"`
	PhoneMobile        *string  `json:"phone_mobile,omitempty" example:"+358-40-1234567"`
	Faxno              *string  `json:"faxno,omitempty"`
	Email              *string  `json:"email,omitempty" example:"mikael.jarvinen@example.com"`
	Username           *string  `json:"username,omitempty" example:"mjarvinen"`
	Password           *string  `json:"password,omitempty" example:"securepass123"`
	Readonly           *int32   `json:"readonly,omitempty" example:"0"`
	Warnings           *int32   `json:"warnings,omitempty" example:"0"`
	AllowToSave        *int32   `json:"allow_to_save,omitempty" example:"1"`
	AllowToCloud       *int32   `json:"allow_to_cloud,omitempty" example:"1"`
	Flag2              *int32   `json:"flag2,omitempty" example:"0"`
	Idsport            *int32   `json:"idsport,omitempty" example:"42"`
	Medication         *string  `json:"medication,omitempty" example:"None"`
	Addinfo            *string  `json:"addinfo,omitempty" example:"Elite cross-country skier"`
	TeamName           *string  `json:"team_name,omitempty" example:"Finnish National Ski Team"`
	Add1               *int32   `json:"add1,omitempty" example:"1"`
	Athlete            *int32   `json:"athlete,omitempty" example:"1"`
	Add10              *string  `json:"add10,omitempty" example:"Field 10"`
	Add20              *string  `json:"add20,omitempty" example:"Field 20"`
	Updatemode         *int32   `json:"updatemode,omitempty" example:"1"`
	WeightKg           *float64 `json:"weight_kg,omitempty" example:"72.5"`
	HeightCm           *float64 `json:"height_cm,omitempty" example:"178.2"`
	DateModified       *float64 `json:"date_modified,omitempty" example:"1693478400.0"`
	RecomTestlevel     *int32   `json:"recom_testlevel,omitempty" example:"3"`
	CreatedBy          *int64   `json:"created_by,omitempty" example:"1001"`
	ModBy              *int64   `json:"mod_by,omitempty" example:"1001"`
	ModDate            *string  `json:"mod_date,omitempty" example:"2024-08-29T10:30:00Z"`     // RFC3339
	Deleted            *int32   `json:"deleted,omitempty" example:"0"`                         // smallint
	CreatedDate        *string  `json:"created_date,omitempty" example:"2024-08-15T14:22:00Z"` // RFC3339
	Modded             *int32   `json:"modded,omitempty" example:"0"`                          // smallint
	AllowAnonymousData *string  `json:"allow_anonymous_data,omitempty" example:"0"`
	Locked             *int32   `json:"locked,omitempty" example:"0"` // smallint
	AllowToSprintai    *int32   `json:"allow_to_sprintai,omitempty" example:"1"`
	TosprintaiFrom     *string  `json:"tosprintai_from,omitempty" example:"2024-08-01"` // YYYY-MM-DD
	StatSent           *string  `json:"stat_sent,omitempty" example:"2024-08-20"`       // YYYY-MM-DD
	SporttiID          *string  `json:"sportti_id,omitempty" example:"27353728"`
}

type KlabMeasurement struct {
	IdMeasurement  *int32  `json:"idMeasurement" example:"888423"`
	MeasName       *string `json:"MeasName,omitempty" example:"VO2 Max Test"`
	IdCustomer     *int32  `json:"idCustomer" example:"7842"`
	TableName      *string `json:"tableName,omitempty" example:"dirtest"`
	IdPatternDef   *string `json:"idpatterndef,omitempty" example:"pattern_vo2"`
	DoYear         *int32  `json:"do_year,omitempty" example:"2024"`
	DoMonth        *int32  `json:"do_month,omitempty" example:"8"`
	DoDay          *int32  `json:"do_day,omitempty" example:"29"`
	DoHour         *int32  `json:"do_hour,omitempty" example:"10"`
	DoMin          *int32  `json:"do_min,omitempty" example:"30"`
	SessionNo      *int32  `json:"sessionno,omitempty" example:"1"`
	Info           *string `json:"info,omitempty" example:"Pre-season assessment"`
	Measurements   *string `json:"measurements,omitempty" example:"Standard protocol"`
	GroupNotes     *string `json:"groupnotes,omitempty" example:"Group Alpha"`
	CbCharts       *string `json:"cbcharts,omitempty" example:"chart_1"`
	CbComments     *string `json:"cbcomments,omitempty" example:"Good performance"`
	CreatedBy      *int64  `json:"created_by,omitempty" example:"1001"`
	ModBy          *int64  `json:"mod_by,omitempty" example:"1001"`
	ModDate        *string `json:"mod_date,omitempty" example:"2024-08-29T10:30:00Z"`     // RFC3339
	Deleted        *int32  `json:"deleted,omitempty" example:"0"`                         // smallint
	CreatedDate    *string `json:"created_date,omitempty" example:"2024-08-29T10:30:00Z"` // RFC3339
	Modded         *int32  `json:"modded,omitempty" example:"0"`                          // smallint
	TestLocation   *string `json:"test_location,omitempty" example:"Lab A"`
	Keywords       *string `json:"keywords,omitempty" example:"VO2max, endurance"`
	TesterName     *string `json:"tester_name,omitempty" example:"Dr. Korhonen"`
	ModderName     *string `json:"modder_name,omitempty" example:"Dr. Korhonen"`
	MeasType       *int32  `json:"meastype,omitempty" example:"1"`
	SentToSprintAI *string `json:"sent_to_sprintai,omitempty" example:"2024-08-29T11:45:00Z"` // RFC3339
}

type KlabDirTest struct {
	IdDirTest     *int32   `json:"idDirTest" example:"777456"`
	IdMeasurement *int32   `json:"idMeasurement" example:"888423"`
	MeasCols      *string  `json:"MeasCols,omitempty" example:"col2measindex=1\\r\\ncol2name=VO2"`
	WeightKg      *float64 `json:"weightkg,omitempty" example:"72.5"`
	HeightCm      *float64 `json:"heightcm,omitempty" example:"178.2"`
	BMI           *float64 `json:"bmi,omitempty" example:"22.8"`
	FatPr         *float64 `json:"fat_pr,omitempty" example:"8.5"`
	FatP1         *float64 `json:"fat_p1,omitempty" example:"2.1"`
	FatP2         *float64 `json:"fat_p2,omitempty" example:"2.3"`
	FatP3         *float64 `json:"fat_p3,omitempty" example:"2.0"`
	FatP4         *float64 `json:"fat_p4,omitempty" example:"2.1"`
	FatStyle      *int32   `json:"fat_style,omitempty" example:"1"`
	FatEquip      *string  `json:"fat_equip,omitempty" example:"Calipers"`
	FVC           *float64 `json:"fvc,omitempty" example:"5.85"`
	FEV1          *float64 `json:"fev1,omitempty" example:"4.92"`
	AirPress      *float64 `json:"air_press,omitempty" example:"1013.2"`
	AirTemp       *float64 `json:"air_temp,omitempty" example:"21.5"`
	AirHumid      *float64 `json:"air_humid,omitempty" example:"45.2"`
	TestProtocol  *string  `json:"testprotocol,omitempty" example:"Incremental"`
	AirPressUnit  *int32   `json:"air_press_unit,omitempty" example:"1"`
	SettingsList  *string  `json:"settingslist,omitempty" example:"protocol=inc;start=100"`
	Lt1X          *float64 `json:"lt1_x,omitempty" example:"185.5"`
	Lt1Y          *float64 `json:"lt1_y,omitempty" example:"2.8"`
	Lt2X          *float64 `json:"lt2_x,omitempty" example:"245.8"`
	Lt2Y          *float64 `json:"lt2_y,omitempty" example:"4.2"`
	Vt1X          *float64 `json:"vt1_x,omitempty" example:"190.2"`
	Vt2X          *float64 `json:"vt2_x,omitempty" example:"250.1"`
	Vt1Y          *float64 `json:"vt1_y,omitempty" example:"42.5"`
	Vt2Y          *float64 `json:"vt2_y,omitempty" example:"58.9"`
	Lt1CalcX      *float64 `json:"lt1_calc_x,omitempty" example:"186.1"`
	Lt1CalcY      *float64 `json:"lt1_calc_y,omitempty" example:"2.82"`
	Lt2CalcX      *float64 `json:"lt2_calc_x,omitempty" example:"246.3"`
	Lt2CalcY      *float64 `json:"lt2_calc_y,omitempty" example:"4.18"`
	ProtocolModel *int32   `json:"protocolmodel,omitempty" example:"1"` // smallint
	TestType      *int32   `json:"testtype,omitempty" example:"1"`      // smallint
	ProtocolXVal  *int32   `json:"protocolxval,omitempty" example:"1"`  // smallint
	StepTime      *int32   `json:"steptime,omitempty" example:"180"`
	WRest         *int32   `json:"w_rest,omitempty" example:"50"` // smallint
	CreatedBy     *int64   `json:"created_by,omitempty" example:"1001"`
	ModBy         *int64   `json:"mod_by,omitempty" example:"1001"`
	ModDate       *string  `json:"mod_date,omitempty" example:"2024-08-29T10:30:00Z"`     // RFC3339
	Deleted       *int32   `json:"deleted,omitempty" example:"0"`                         // smallint
	CreatedDate   *string  `json:"created_date,omitempty" example:"2024-08-29T10:30:00Z"` // RFC3339
	Modded        *int32   `json:"modded,omitempty" example:"0"`                          // smallint
	NoRawData     *int32   `json:"norawdata,omitempty" example:"0"`                       // smallint
}

type KlabDirTestStep struct {
	Iddirteststeps *int32   `json:"iddirteststeps" example:"666621"`
	Idmeasurement  *int32   `json:"idmeasurement" example:"888423"`
	Stepno         *int32   `json:"stepno,omitempty" example:"0"`
	AnaTime        *int32   `json:"ana_time,omitempty" example:"0"`
	TimeStop       *float64 `json:"timestop,omitempty" example:"180.0"`
	Speed          *float64 `json:"speed,omitempty" example:"0.0"`
	Pace           *float64 `json:"pace,omitempty" example:"0.0"`
	Angle          *float64 `json:"angle,omitempty" example:"0.0"`
	Elev           *float64 `json:"elev,omitempty" example:"100.0"`
	Vo2calc        *float64 `json:"vo2calc,omitempty" example:"12.5"`
	TTot           *float64 `json:"t_tot,omitempty" example:"180.0"`
	TEx            *float64 `json:"t_ex,omitempty" example:"175.2"`
	Fico2          *float64 `json:"fico2,omitempty" example:"0.04"`
	Fio2           *float64 `json:"fio2,omitempty" example:"20.93"`
	Feco2          *float64 `json:"feco2,omitempty" example:"3.2"`
	Feo2           *float64 `json:"feo2,omitempty" example:"17.8"`
	Vde            *float64 `json:"vde,omitempty" example:"28.5"`
	Vco2           *float64 `json:"vco2,omitempty" example:"1.02"`
	Vo2            *float64 `json:"vo2,omitempty" example:"1.25"`
	Bf             *float64 `json:"bf,omitempty" example:"18.5"`
	Ve             *float64 `json:"ve,omitempty" example:"32.1"`
	Petco2         *float64 `json:"petco2,omitempty" example:"38.2"`
	Peto2          *float64 `json:"peto2,omitempty" example:"102.5"`
	Vo2kg          *float64 `json:"vo2kg,omitempty" example:"17.2"`
	Re             *float64 `json:"re,omitempty" example:"0.82"`
	Hr             *float64 `json:"hr,omitempty" example:"95.0"`
	La             *float64 `json:"la,omitempty" example:"1.2"`
	Rer            *float64 `json:"rer,omitempty" example:"0.82"`
	VeStpd         *float64 `json:"ve_stpd,omitempty" example:"30.8"`
	Veo2           *float64 `json:"veo2,omitempty" example:"25.6"`
	Veco2          *float64 `json:"veco2,omitempty" example:"31.2"`
	Tv             *float64 `json:"tv,omitempty" example:"1.74"`
	EeAe           *float64 `json:"ee_ae,omitempty" example:"5.8"`
	LaVo2          *float64 `json:"la_vo2,omitempty" example:"0.96"`
	O2pulse        *float64 `json:"o2pulse,omitempty" example:"13.2"`
	VdeTv          *float64 `json:"vde_tv,omitempty" example:"16.4"`
	Va             *float64 `json:"va,omitempty" example:"28.1"`
	O2sa           *float64 `json:"o2sa,omitempty" example:"98.5"`
	Rpe            *float64 `json:"rpe,omitempty" example:"6.0"`
	BpSys          *float64 `json:"bp_sys,omitempty" example:"120.0"`
	BpDia          *float64 `json:"bp_dia,omitempty" example:"80.0"`
	Own1           *float64 `json:"own1,omitempty"`
	Own2           *float64 `json:"own2,omitempty"`
	Own3           *float64 `json:"own3,omitempty"`
	Own4           *float64 `json:"own4,omitempty"`
	Own5           *float64 `json:"own5,omitempty"`
	StepIsRest     *int32   `json:"step_is_rest,omitempty" example:"1"`
	StepIs30max    *int32   `json:"step_is_30max,omitempty" example:"0"`
	StepIs60max    *int32   `json:"step_is_60max,omitempty" example:"0"`
	StepIsRec      *int32   `json:"step_is_rec,omitempty" example:"0"`
	CalcStart      *int32   `json:"calc_start,omitempty" example:"60"`
	CalcEnd        *int32   `json:"calc_end,omitempty" example:"180"`
	Comments       *string  `json:"comments,omitempty" example:"Rest phase"`
	TimeStart      *float64 `json:"timestart,omitempty" example:"0.0"`
	Duration       *float64 `json:"duration,omitempty" example:"180.0"`
	Eco            *float64 `json:"eco,omitempty" example:"4.2"`
	P              *float64 `json:"p,omitempty" example:"100.0"`
	Wkg            *float64 `json:"wkg,omitempty" example:"1.38"`
	Vo230s         *float64 `json:"vo2_30s,omitempty" example:"17.1"`
	Vo2Pr          *float64 `json:"vo2_pr,omitempty" example:"25.2"`
	StepIsLast     *int32   `json:"step_is_last,omitempty" example:"0"`
	Deleted        *int32   `json:"deleted,omitempty" example:"0"` // smallint
	CreatedBy      *int64   `json:"created_by,omitempty" example:"1001"`
	ModBy          *int64   `json:"mod_by,omitempty" example:"1001"`
	ModDate        *string  `json:"mod_date,omitempty" example:"2024-08-29T10:30:00Z"`     // RFC3339
	CreatedDate    *string  `json:"created_date,omitempty" example:"2024-08-29T10:30:00Z"` // RFC3339
	Modded         *int32   `json:"modded,omitempty" example:"0"`                          // smallint
	Own6           *float64 `json:"own6,omitempty"`
	Own7           *float64 `json:"own7,omitempty"`
	Own8           *float64 `json:"own8,omitempty"`
	Own9           *float64 `json:"own9,omitempty"`
	Own10          *float64 `json:"own10,omitempty"`
	To2            *float64 `json:"to2,omitempty" example:"0.85"`
	Tco2           *float64 `json:"tco2,omitempty" example:"1.02"`
}

type KlabDirReport struct {
	Iddirreport   *int32  `json:"iddirreport" example:"555852"`
	PageInstr     *string `json:"page_instructions,omitempty" example:"otsake=Training\\r\\nlabel_vo2max=VO2 Max"`
	Idmeasurement *int32  `json:"idmeasurement,omitempty" example:"888423"`
	TemplateRec   *int32  `json:"template_rec,omitempty" example:"1"`
	LibrecName    *string `json:"librec_name,omitempty" example:"VO2Max_Template"`
	CreatedBy     *int64  `json:"created_by,omitempty" example:"1001"`
	ModBy         *int64  `json:"mod_by,omitempty" example:"1001"`
	ModDate       *string `json:"mod_date,omitempty" example:"2024-08-29T10:30:00Z"`     // RFC3339
	Deleted       *int32  `json:"deleted,omitempty" example:"0"`                         // smallint
	CreatedDate   *string `json:"created_date,omitempty" example:"2024-08-29T10:30:00Z"` // RFC3339
	Modded        *int32  `json:"modded,omitempty" example:"0"`                          // smallint
}

type KlabDirRawData struct {
	IdDirRawData  *int32  `json:"idDirRawData" example:"444856"`
	IdMeasurement *int32  `json:"idMeasurement" example:"888423"`
	RawData       *string `json:"rawdata,omitempty" example:"180;95;1250;17200;1020"`
	ColumnData    *string `json:"columndata,omitempty" example:"Time;HR;VO2;VO2kg;VCO2"`
	Info          *string `json:"info,omitempty" example:"VO2 Max Raw Data"`
	UnitsData     *string `json:"unitsdata,omitempty" example:"s;bpm;ml/min;ml/kg/min;ml/min"`
	CreatedBy     *int64  `json:"created_by,omitempty" example:"1001"`
	ModBy         *int64  `json:"mod_by,omitempty" example:"1001"`
	ModDate       *string `json:"mod_date,omitempty" example:"2024-08-29T10:30:00Z"`     // RFC3339
	Deleted       *int32  `json:"deleted,omitempty" example:"0"`                         // smallint
	CreatedDate   *string `json:"created_date,omitempty" example:"2024-08-29T10:30:00Z"` // RFC3339
	Modded        *int32  `json:"modded,omitempty" example:"0"`                          // smallint
}

type KlabDirResults struct {
	Iddirresults       *int32   `json:"iddirresults" example:"333421"`
	Idmeasurement      *int32   `json:"idmeasurement" example:"888423"`
	MaxVo2mlkgmin      *float64 `json:"max_vo2mlkgmin,omitempty" example:"68.5"`
	MaxVo2mlmin        *float64 `json:"max_vo2mlmin,omitempty" example:"4965.25"`
	MaxVo2             *float64 `json:"max_vo2,omitempty" example:"68.5"`
	MaxHr              *float64 `json:"max_hr,omitempty" example:"192.0"`
	MaxSpeed           *float64 `json:"max_speed,omitempty" example:"0.0"`
	MaxPace            *float64 `json:"max_pace,omitempty" example:"0.0"`
	MaxP               *float64 `json:"max_p,omitempty" example:"375.0"`
	MaxPkg             *float64 `json:"max_pkg,omitempty" example:"5.17"`
	MaxAngle           *float64 `json:"max_angle,omitempty" example:"0.0"`
	MaxLac             *float64 `json:"max_lac,omitempty" example:"8.9"`
	MaxAdd1            *float64 `json:"max_add1,omitempty"`
	MaxAdd2            *float64 `json:"max_add2,omitempty"`
	MaxAdd3            *float64 `json:"max_add3,omitempty"`
	LacAnkVo2mlkgmin   *float64 `json:"lac_ank_vo2mlkgmin,omitempty" example:"45.2"`
	LacAnkVo2mlmin     *float64 `json:"lac_ank_vo2mlmin,omitempty" example:"3277.0"`
	LacAnkVo2          *float64 `json:"lac_ank_vo2,omitempty" example:"45.2"`
	LacAnkVo2pr        *float64 `json:"lac_ank_vo2pr,omitempty" example:"66.0"`
	LacAnkHr           *float64 `json:"lac_ank_hr,omitempty" example:"150.0"`
	LacAnkSpeed        *float64 `json:"lac_ank_speed,omitempty" example:"0.0"`
	LacAnkPace         *float64 `json:"lac_ank_pace,omitempty" example:"0.0"`
	LacAnkP            *float64 `json:"lac_ank_p,omitempty" example:"185.5"`
	LacAnkPkg          *float64 `json:"lac_ank_pkg,omitempty" example:"2.56"`
	LacAnkAngle        *float64 `json:"lac_ank_angle,omitempty" example:"0.0"`
	LacAnkLac          *float64 `json:"lac_ank_lac,omitempty" example:"2.0"`
	LacAnkAdd1         *float64 `json:"lac_ank_add1,omitempty"`
	LacAnkAdd2         *float64 `json:"lac_ank_add2,omitempty"`
	LacAnkAdd3         *float64 `json:"lac_ank_add3,omitempty"`
	LacAerkVo2mlkgmin  *float64 `json:"lac_aerk_vo2mlkgmin,omitempty" example:"58.8"`
	LacAerkVo2mlmin    *float64 `json:"lac_aerk_vo2mlmin,omitempty" example:"4263.0"`
	LacAerkVo2         *float64 `json:"lac_aerk_vo2,omitempty" example:"58.8"`
	LacAerkVo2pr       *float64 `json:"lac_aerk_vo2pr,omitempty" example:"85.8"`
	LacAerkHr          *float64 `json:"lac_aerk_hr,omitempty" example:"175.0"`
	LacAerkSpeed       *float64 `json:"lac_aerk_speed,omitempty" example:"0.0"`
	LacAerkPace        *float64 `json:"lac_aerk_pace,omitempty" example:"0.0"`
	LacAerkP           *float64 `json:"lac_aerk_p,omitempty" example:"245.8"`
	LacAerkPkg         *float64 `json:"lac_aerk_pkg,omitempty" example:"3.39"`
	LacAerkAngle       *float64 `json:"lac_aerk_angle,omitempty" example:"0.0"`
	LacAerkLac         *float64 `json:"lac_aerk_lac,omitempty" example:"4.0"`
	LacAerkAdd1        *float64 `json:"lac_aerk_add1,omitempty"`
	LacAerkAdd2        *float64 `json:"lac_aerk_add2,omitempty"`
	LacAerkAdd3        *float64 `json:"lac_aerk_add3,omitempty"`
	VentAnkVo2mlkgmin  *float64 `json:"vent_ank_vo2mlkgmin,omitempty" example:"46.1"`
	VentAnkVo2mlmin    *float64 `json:"vent_ank_vo2mlmin,omitempty" example:"3342.25"`
	VentAnkVo2         *float64 `json:"vent_ank_vo2,omitempty" example:"46.1"`
	VentAnkVo2pr       *float64 `json:"vent_ank_vo2pr,omitempty" example:"67.3"`
	VentAnkHr          *float64 `json:"vent_ank_hr,omitempty" example:"152.0"`
	VentAnkSpeed       *float64 `json:"vent_ank_speed,omitempty" example:"0.0"`
	VentAnkPace        *float64 `json:"vent_ank_pace,omitempty" example:"0.0"`
	VentAnkP           *float64 `json:"vent_ank_p,omitempty" example:"190.2"`
	VentAnkPkg         *float64 `json:"vent_ank_pkg,omitempty" example:"2.62"`
	VentAnkAngle       *float64 `json:"vent_ank_angle,omitempty" example:"0.0"`
	VentAnkLac         *float64 `json:"vent_ank_lac,omitempty" example:"2.1"`
	VentAnkAdd1        *float64 `json:"vent_ank_add1,omitempty"`
	VentAnkAdd2        *float64 `json:"vent_ank_add2,omitempty"`
	VentAnkAdd3        *float64 `json:"vent_ank_add3,omitempty"`
	VentAerkVo2mlkgmin *float64 `json:"vent_aerk_vo2mlkgmin,omitempty" example:"59.5"`
	VentAerkVo2mlmin   *float64 `json:"vent_aerk_vo2mlmin,omitempty" example:"4313.75"`
	VentAerkVo2        *float64 `json:"vent_aerk_vo2,omitempty" example:"59.5"`
	VentAerkVo2pr      *float64 `json:"vent_aerk_vo2pr,omitempty" example:"86.9"`
	VentAerkHr         *float64 `json:"vent_aerk_hr,omitempty" example:"177.0"`
	VentAerkSpeed      *float64 `json:"vent_aerk_speed,omitempty" example:"0.0"`
	VentAerkPace       *float64 `json:"vent_aerk_pace,omitempty" example:"0.0"`
	VentAerkP          *float64 `json:"vent_aerk_p,omitempty" example:"250.1"`
	VentAerkPkg        *float64 `json:"vent_aerk_pkg,omitempty" example:"3.45"`
	VentAerkAngle      *float64 `json:"vent_aerk_angle,omitempty" example:"0.0"`
	VentAerkLac        *float64 `json:"vent_aerk_lac,omitempty" example:"4.2"`
	VentAerkAdd1       *float64 `json:"vent_aerk_add1,omitempty"`
	VentAerkAdd2       *float64 `json:"vent_aerk_add2,omitempty"`
	VentAerkAdd3       *float64 `json:"vent_aerk_add3,omitempty"`
	CreatedBy          *int64   `json:"created_by,omitempty" example:"1001"`
	ModBy              *int64   `json:"mod_by,omitempty" example:"1001"`
	ModDate            *string  `json:"mod_date,omitempty" example:"2024-08-29T10:30:00Z"`     // RFC3339
	Deleted            *int32   `json:"deleted,omitempty" example:"0"`                         // smallint
	CreatedDate        *string  `json:"created_date,omitempty" example:"2024-08-29T10:30:00Z"` // RFC3339
	Modded             *int32   `json:"modded,omitempty" example:"0"`                          // smallint
}

type KlabDataResponse struct {
	SporttiID    int32             `json:"sportti_id" example:"27353728"`
	Measurements []KlabMeasurement `json:"measurements"`
	DirTest      []KlabDirTest     `json:"dirtest"`
	DirTestSteps []KlabDirTestStep `json:"dirteststeps"`
	DirReport    []KlabDirReport   `json:"dirreport"`
	DirRawData   []KlabDirRawData  `json:"dirrawdata"`
	DirResults   []KlabDirResults  `json:"dirresults"`
}

// Updated to show multiple entries in arrays
type KlabDataBulkDoc struct {
	Customer27353728 KlabDataBundle `json:"27353728"`
}

type KlabDataBundle struct {
	Customer        []KlabCustomer    `json:"customer,omitempty"`
	MeasurementList []KlabMeasurement `json:"measurement_list,omitempty"`
	DirTest         []KlabDirTest     `json:"dirtest,omitempty"`
	DirTestSteps    []KlabDirTestStep `json:"dirteststeps,omitempty"`
	DirReport       []KlabDirReport   `json:"dirreport,omitempty"`
	DirRawData      []KlabDirRawData  `json:"dirrawdata,omitempty"`
	DirResults      []KlabDirResults  `json:"dirresults,omitempty"`
}
