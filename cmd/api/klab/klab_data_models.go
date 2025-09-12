package klabapi

import (
	"time"

	klabsqlc "github.com/DeRuina/KUHA-REST-API/internal/db/klab"
	"github.com/DeRuina/KUHA-REST-API/internal/utils"
)

// Models
type KlabCustomerInput struct {
	IdCustomer         *int32   `json:"idCustomer" validate:"required"`
	FirstName          *string  `json:"FirstName" validate:"required"`
	LastName           *string  `json:"LastName" validate:"required"`
	IdGroups           *int32   `json:"idGroups"`
	DOB                *string  `json:"DOB"`
	SEX                *int32   `json:"SEX"`
	DobYear            *int32   `json:"dob_year"`
	DobMonth           *int32   `json:"dob_month"`
	DobDay             *int32   `json:"dob_day"`
	PidNumber          *string  `json:"pid_number"`
	Company            *string  `json:"company"`
	Occupation         *string  `json:"occupation"`
	Education          *string  `json:"education"`
	Address            *string  `json:"address"`
	PhoneHome          *string  `json:"phone_home"`
	PhoneWork          *string  `json:"phone_work"`
	PhoneMobile        *string  `json:"phone_mobile"`
	Faxno              *string  `json:"faxno"`
	Email              *string  `json:"email"`
	Username           *string  `json:"username"`
	Password           *string  `json:"password"`
	Readonly           *int32   `json:"readonly"`
	Warnings           *int32   `json:"warnings"`
	AllowToSave        *int32   `json:"allow_to_save"`
	AllowToCloud       *int32   `json:"allow_to_cloud"`
	Flag2              *int32   `json:"flag2"`
	Idsport            *int32   `json:"idsport"`
	Medication         *string  `json:"medication"`
	Addinfo            *string  `json:"addinfo"`
	TeamName           *string  `json:"team_name"`
	Add1               *int32   `json:"add1"`
	Athlete            *int32   `json:"athlete"`
	Add10              *string  `json:"add10"`
	Add20              *string  `json:"add20"`
	Updatemode         *int32   `json:"updatemode"`
	WeightKg           *float64 `json:"weight_kg"`
	HeightCm           *float64 `json:"height_cm"`
	DateModified       *float64 `json:"date_modified"`
	RecomTestlevel     *int32   `json:"recom_testlevel"`
	CreatedBy          *int64   `json:"created_by"`
	ModBy              *int64   `json:"mod_by"`
	ModDate            *string  `json:"mod_date"`
	Deleted            *int32   `json:"deleted"`
	CreatedDate        *string  `json:"created_date"`
	Modded             *int32   `json:"modded"`
	AllowAnonymousData *string  `json:"allow_anonymous_data" validate:"omitempty,oneof=0 1"`
	Locked             *int32   `json:"locked"`
	AllowToSprintai    *int32   `json:"allow_to_sprintai"`
	TosprintaiFrom     *string  `json:"tosprintai_from"`
	StatSent           *string  `json:"stat_sent"`
	SporttiID          *string  `json:"sportti_id" validate:"required"`
}

type KlabMeasurementInput struct {
	IdMeasurement  *int32  `json:"idMeasurement" validate:"required"`
	MeasName       *string `json:"MeasName"`
	IdCustomer     *int32  `json:"idCustomer"`
	TableName      *string `json:"tableName"`
	IdPatternDef   *string `json:"idpatterndef"`
	DoYear         *int32  `json:"do_year"`
	DoMonth        *int32  `json:"do_month"`
	DoDay          *int32  `json:"do_day"`
	DoHour         *int32  `json:"do_hour"`
	DoMin          *int32  `json:"do_min"`
	SessionNo      *int32  `json:"sessionno"`
	Info           *string `json:"info"`
	Measurements   *string `json:"measurements"`
	GroupNotes     *string `json:"groupnotes"`
	CbCharts       *string `json:"cbcharts"`
	CbComments     *string `json:"cbcomments"`
	CreatedBy      *int64  `json:"created_by"`
	ModBy          *int64  `json:"mod_by"`
	ModDate        *string `json:"mod_date"`
	Deleted        *int32  `json:"deleted"`
	CreatedDate    *string `json:"created_date"`
	Modded         *int32  `json:"modded"`
	TestLocation   *string `json:"test_location"`
	Keywords       *string `json:"keywords"`
	TesterName     *string `json:"tester_name"`
	ModderName     *string `json:"modder_name"`
	MeasType       *int32  `json:"meastype"`
	SentToSprintAI *string `json:"sent_to_sprintai"`
}

type KlabDirTestInput struct {
	IdDirTest     *int32   `json:"idDirTest" validate:"required"`
	IdMeasurement *int32   `json:"idMeasurement" validate:"required"`
	MeasCols      *string  `json:"MeasCols"`
	WeightKg      *float64 `json:"weightkg"`
	HeightCm      *float64 `json:"heightcm"`
	BMI           *float64 `json:"bmi"`
	FatPr         *float64 `json:"fat_pr"`
	FatP1         *float64 `json:"fat_p1"`
	FatP2         *float64 `json:"fat_p2"`
	FatP3         *float64 `json:"fat_p3"`
	FatP4         *float64 `json:"fat_p4"`
	FatStyle      *int32   `json:"fat_style"`
	FatEquip      *string  `json:"fat_equip"`
	FVC           *float64 `json:"fvc"`
	FEV1          *float64 `json:"fev1"`
	AirPress      *float64 `json:"air_press"`
	AirTemp       *float64 `json:"air_temp"`
	AirHumid      *float64 `json:"air_humid"`
	TestProtocol  *string  `json:"testprotocol"`
	AirPressUnit  *int32   `json:"air_press_unit"`
	SettingsList  *string  `json:"settingslist"`
	Lt1X          *float64 `json:"lt1_x"`
	Lt1Y          *float64 `json:"lt1_y"`
	Lt2X          *float64 `json:"lt2_x"`
	Lt2Y          *float64 `json:"lt2_y"`
	Vt1X          *float64 `json:"vt1_x"`
	Vt2X          *float64 `json:"vt2_x"`
	Vt1Y          *float64 `json:"vt1_y"`
	Vt2Y          *float64 `json:"vt2_y"`
	Lt1CalcX      *float64 `json:"lt1_calc_x"`
	Lt1CalcY      *float64 `json:"lt1_calc_y"`
	Lt2CalcX      *float64 `json:"lt2_calc_x"`
	Lt2CalcY      *float64 `json:"lt2_calc_y"`
	ProtocolModel *int32   `json:"protocolmodel"`
	TestType      *int32   `json:"testtype"`
	ProtocolXVal  *int32   `json:"protocolxval"`
	StepTime      *int32   `json:"steptime"`
	WRest         *int32   `json:"w_rest"`
	CreatedBy     *int64   `json:"created_by"`
	ModBy         *int64   `json:"mod_by"`
	ModDate       *string  `json:"mod_date"`
	Deleted       *int32   `json:"deleted"`
	CreatedDate   *string  `json:"created_date"`
	Modded        *int32   `json:"modded"`
	NoRawData     *int32   `json:"norawdata"`
}

type KlabDirTestStepInput struct {
	Iddirteststeps *int32   `json:"iddirteststeps" validate:"required"`
	Idmeasurement  *int32   `json:"idmeasurement" validate:"required"`
	Stepno         *int32   `json:"stepno"`
	AnaTime        *int32   `json:"ana_time"`
	TimeStop       *float64 `json:"timestop"`
	Speed          *float64 `json:"speed"`
	Pace           *float64 `json:"pace"`
	Angle          *float64 `json:"angle"`
	Elev           *float64 `json:"elev"`
	Vo2calc        *float64 `json:"vo2calc"`
	TTot           *float64 `json:"t_tot"`
	TEx            *float64 `json:"t_ex"`
	Fico2          *float64 `json:"fico2"`
	Fio2           *float64 `json:"fio2"`
	Feco2          *float64 `json:"feco2"`
	Feo2           *float64 `json:"feo2"`
	Vde            *float64 `json:"vde"`
	Vco2           *float64 `json:"vco2"`
	Vo2            *float64 `json:"vo2"`
	Bf             *float64 `json:"bf"`
	Ve             *float64 `json:"ve"`
	Petco2         *float64 `json:"petco2"`
	Peto2          *float64 `json:"peto2"`
	Vo2kg          *float64 `json:"vo2kg"`
	Re             *float64 `json:"re"`
	Hr             *float64 `json:"hr"`
	La             *float64 `json:"la"`
	Rer            *float64 `json:"rer"`
	VeStpd         *float64 `json:"ve_stpd"`
	Veo2           *float64 `json:"veo2"`
	Veco2          *float64 `json:"veco2"`
	Tv             *float64 `json:"tv"`
	EeAe           *float64 `json:"ee_ae"`
	LaVo2          *float64 `json:"la_vo2"`
	O2pulse        *float64 `json:"o2pulse"`
	VdeTv          *float64 `json:"vde_tv"`
	Va             *float64 `json:"va"`
	O2sa           *float64 `json:"o2sa"`
	Rpe            *float64 `json:"rpe"`
	BpSys          *float64 `json:"bp_sys"`
	BpDia          *float64 `json:"bp_dia"`
	Own1           *float64 `json:"own1"`
	Own2           *float64 `json:"own2"`
	Own3           *float64 `json:"own3"`
	Own4           *float64 `json:"own4"`
	Own5           *float64 `json:"own5"`
	StepIsRest     *int32   `json:"step_is_rest"`
	StepIs30max    *int32   `json:"step_is_30max"`
	StepIs60max    *int32   `json:"step_is_60max"`
	StepIsRec      *int32   `json:"step_is_rec"`
	CalcStart      *int32   `json:"calc_start"`
	CalcEnd        *int32   `json:"calc_end"`
	Comments       *string  `json:"comments"`
	TimeStart      *float64 `json:"timestart"`
	Duration       *float64 `json:"duration"`
	Eco            *float64 `json:"eco"`
	P              *float64 `json:"p"`
	Wkg            *float64 `json:"wkg"`
	Vo230s         *float64 `json:"vo2_30s"`
	Vo2Pr          *float64 `json:"vo2_pr"`
	StepIsLast     *int32   `json:"step_is_last"`
	Deleted        *int32   `json:"deleted"`
	CreatedBy      *int64   `json:"created_by"`
	ModBy          *int64   `json:"mod_by"`
	ModDate        *string  `json:"mod_date"`
	CreatedDate    *string  `json:"created_date"`
	Modded         *int32   `json:"modded"`
	Own6           *float64 `json:"own6"`
	Own7           *float64 `json:"own7"`
	Own8           *float64 `json:"own8"`
	Own9           *float64 `json:"own9"`
	Own10          *float64 `json:"own10"`
	To2            *float64 `json:"to2"`
	Tco2           *float64 `json:"tco2"`
}

type KlabDirReportInput struct {
	Iddirreport   *int32  `json:"iddirreport" validate:"required"`
	PageInstr     *string `json:"page_instructions"`
	Idmeasurement *int32  `json:"idmeasurement" validate:"required"`
	TemplateRec   *int32  `json:"template_rec"`
	LibrecName    *string `json:"librec_name"`
	CreatedBy     *int64  `json:"created_by"`
	ModBy         *int64  `json:"mod_by"`
	ModDate       *string `json:"mod_date"`
	Deleted       *int32  `json:"deleted"`
	CreatedDate   *string `json:"created_date"`
	Modded        *int32  `json:"modded"`
}

type KlabDirRawDataInput struct {
	IdDirRawData  *int32  `json:"idDirRawData" validate:"required"`
	IdMeasurement *int32  `json:"idMeasurement" validate:"required"`
	RawData       *string `json:"rawdata"`
	ColumnData    *string `json:"columndata"`
	Info          *string `json:"info"`
	UnitsData     *string `json:"unitsdata"`
	CreatedBy     *int64  `json:"created_by"`
	ModBy         *int64  `json:"mod_by"`
	ModDate       *string `json:"mod_date"`
	Deleted       *int32  `json:"deleted"`
	CreatedDate   *string `json:"created_date"`
	Modded        *int32  `json:"modded"`
}

type KlabDirResultsInput struct {
	Iddirresults       *int32   `json:"iddirresults" validate:"required"`
	Idmeasurement      *int32   `json:"idmeasurement" validate:"required"`
	MaxVo2mlkgmin      *float64 `json:"max_vo2mlkgmin"`
	MaxVo2mlmin        *float64 `json:"max_vo2mlmin"`
	MaxVo2             *float64 `json:"max_vo2"`
	MaxHr              *float64 `json:"max_hr"`
	MaxSpeed           *float64 `json:"max_speed"`
	MaxPace            *float64 `json:"max_pace"`
	MaxP               *float64 `json:"max_p"`
	MaxPkg             *float64 `json:"max_pkg"`
	MaxAngle           *float64 `json:"max_angle"`
	MaxLac             *float64 `json:"max_lac"`
	MaxAdd1            *float64 `json:"max_add1"`
	MaxAdd2            *float64 `json:"max_add2"`
	MaxAdd3            *float64 `json:"max_add3"`
	LacAnkVo2mlkgmin   *float64 `json:"lac_ank_vo2mlkgmin"`
	LacAnkVo2mlmin     *float64 `json:"lac_ank_vo2mlmin"`
	LacAnkVo2          *float64 `json:"lac_ank_vo2"`
	LacAnkVo2pr        *float64 `json:"lac_ank_vo2pr"`
	LacAnkHr           *float64 `json:"lac_ank_hr"`
	LacAnkSpeed        *float64 `json:"lac_ank_speed"`
	LacAnkPace         *float64 `json:"lac_ank_pace"`
	LacAnkP            *float64 `json:"lac_ank_p"`
	LacAnkPkg          *float64 `json:"lac_ank_pkg"`
	LacAnkAngle        *float64 `json:"lac_ank_angle"`
	LacAnkLac          *float64 `json:"lac_ank_lac"`
	LacAnkAdd1         *float64 `json:"lac_ank_add1"`
	LacAnkAdd2         *float64 `json:"lac_ank_add2"`
	LacAnkAdd3         *float64 `json:"lac_ank_add3"`
	LacAerkVo2mlkgmin  *float64 `json:"lac_aerk_vo2mlkgmin"`
	LacAerkVo2mlmin    *float64 `json:"lac_aerk_vo2mlmin"`
	LacAerkVo2         *float64 `json:"lac_aerk_vo2"`
	LacAerkVo2pr       *float64 `json:"lac_aerk_vo2pr"`
	LacAerkHr          *float64 `json:"lac_aerk_hr"`
	LacAerkSpeed       *float64 `json:"lac_aerk_speed"`
	LacAerkPace        *float64 `json:"lac_aerk_pace"`
	LacAerkP           *float64 `json:"lac_aerk_p"`
	LacAerkPkg         *float64 `json:"lac_aerk_pkg"`
	LacAerkAngle       *float64 `json:"lac_aerk_angle"`
	LacAerkLac         *float64 `json:"lac_aerk_lac"`
	LacAerkAdd1        *float64 `json:"lac_aerk_add1"`
	LacAerkAdd2        *float64 `json:"lac_aerk_add2"`
	LacAerkAdd3        *float64 `json:"lac_aerk_add3"`
	VentAnkVo2mlkgmin  *float64 `json:"vent_ank_vo2mlkgmin"`
	VentAnkVo2mlmin    *float64 `json:"vent_ank_vo2mlmin"`
	VentAnkVo2         *float64 `json:"vent_ank_vo2"`
	VentAnkVo2pr       *float64 `json:"vent_ank_vo2pr"`
	VentAnkHr          *float64 `json:"vent_ank_hr"`
	VentAnkSpeed       *float64 `json:"vent_ank_speed"`
	VentAnkPace        *float64 `json:"vent_ank_pace"`
	VentAnkP           *float64 `json:"vent_ank_p"`
	VentAnkPkg         *float64 `json:"vent_ank_pkg"`
	VentAnkAngle       *float64 `json:"vent_ank_angle"`
	VentAnkLac         *float64 `json:"vent_ank_lac"`
	VentAnkAdd1        *float64 `json:"vent_ank_add1"`
	VentAnkAdd2        *float64 `json:"vent_ank_add2"`
	VentAnkAdd3        *float64 `json:"vent_ank_add3"`
	VentAerkVo2mlkgmin *float64 `json:"vent_aerk_vo2mlkgmin"`
	VentAerkVo2mlmin   *float64 `json:"vent_aerk_vo2mlmin"`
	VentAerkVo2        *float64 `json:"vent_aerk_vo2"`
	VentAerkVo2pr      *float64 `json:"vent_aerk_vo2pr"`
	VentAerkHr         *float64 `json:"vent_aerk_hr"`
	VentAerkSpeed      *float64 `json:"vent_aerk_speed"`
	VentAerkPace       *float64 `json:"vent_aerk_pace"`
	VentAerkP          *float64 `json:"vent_aerk_p"`
	VentAerkPkg        *float64 `json:"vent_aerk_pkg"`
	VentAerkAngle      *float64 `json:"vent_aerk_angle"`
	VentAerkLac        *float64 `json:"vent_aerk_lac"`
	VentAerkAdd1       *float64 `json:"vent_aerk_add1"`
	VentAerkAdd2       *float64 `json:"vent_aerk_add2"`
	VentAerkAdd3       *float64 `json:"vent_aerk_add3"`
	CreatedBy          *int64   `json:"created_by"`
	ModBy              *int64   `json:"mod_by"`
	ModDate            *string  `json:"mod_date"`
	Deleted            *int32   `json:"deleted"`
	CreatedDate        *string  `json:"created_date"`
	Modded             *int32   `json:"modded"`
}

// Customer -> UpsertCustomerParams
func mapCustomerToParams(in KlabCustomerInput, sporttiID string) (klabsqlc.UpsertCustomerParams, error) {
	id := utils.DerefInt32(in.IdCustomer)

	var (
		dob, modDate, createdDate, toSprintFrom, statSent *time.Time
		err                                               error
	)
	if in.DOB != nil {
		dob, err = utils.ParseDatePtr(in.DOB)
		if err != nil {
			return klabsqlc.UpsertCustomerParams{}, err
		}
	}
	if in.ModDate != nil {
		modDate, err = utils.ParseTimestampPtr(in.ModDate)
		if err != nil {
			return klabsqlc.UpsertCustomerParams{}, err
		}
	}
	if in.CreatedDate != nil {
		createdDate, err = utils.ParseTimestampPtr(in.CreatedDate)
		if err != nil {
			return klabsqlc.UpsertCustomerParams{}, err
		}
	}
	if in.TosprintaiFrom != nil {
		toSprintFrom, err = utils.ParseDatePtr(in.TosprintaiFrom)
		if err != nil {
			return klabsqlc.UpsertCustomerParams{}, err
		}
	}
	if in.StatSent != nil {
		statSent, err = utils.ParseDatePtr(in.StatSent)
		if err != nil {
			return klabsqlc.UpsertCustomerParams{}, err
		}
	}

	return klabsqlc.UpsertCustomerParams{
		Idcustomer:         id,
		Firstname:          utils.DerefString(in.FirstName),
		Lastname:           utils.DerefString(in.LastName),
		Idgroups:           utils.NullInt32Ptr(in.IdGroups),
		Dob:                utils.NullTimePtr(dob),
		Sex:                utils.NullInt32Ptr(in.SEX),
		DobYear:            utils.NullInt32Ptr(in.DobYear),
		DobMonth:           utils.NullInt32Ptr(in.DobMonth),
		DobDay:             utils.NullInt32Ptr(in.DobDay),
		PidNumber:          utils.NullStringPtr(in.PidNumber),
		Company:            utils.NullStringPtr(in.Company),
		Occupation:         utils.NullStringPtr(in.Occupation),
		Education:          utils.NullStringPtr(in.Education),
		Address:            utils.NullStringPtr(in.Address),
		PhoneHome:          utils.NullStringPtr(in.PhoneHome),
		PhoneWork:          utils.NullStringPtr(in.PhoneWork),
		PhoneMobile:        utils.NullStringPtr(in.PhoneMobile),
		Faxno:              utils.NullStringPtr(in.Faxno),
		Email:              utils.NullStringPtr(in.Email),
		Username:           utils.NullStringPtr(in.Username),
		Password:           utils.NullStringPtr(in.Password),
		Readonly:           utils.NullInt32Ptr(in.Readonly),
		Warnings:           utils.NullInt32Ptr(in.Warnings),
		AllowToSave:        utils.NullInt32Ptr(in.AllowToSave),
		AllowToCloud:       utils.NullInt32Ptr(in.AllowToCloud),
		Flag2:              utils.NullInt32Ptr(in.Flag2),
		Idsport:            utils.NullInt32Ptr(in.Idsport),
		Medication:         utils.NullStringPtr(in.Medication),
		Addinfo:            utils.NullStringPtr(in.Addinfo),
		TeamName:           utils.NullStringPtr(in.TeamName),
		Add1:               utils.NullInt32Ptr(in.Add1),
		Athlete:            utils.NullInt32Ptr(in.Athlete),
		Add10:              utils.NullStringPtr(in.Add10),
		Add20:              utils.NullStringPtr(in.Add20),
		Updatemode:         utils.NullInt32Ptr(in.Updatemode),
		WeightKg:           utils.NullFloat64Ptr(in.WeightKg),
		HeightCm:           utils.NullFloat64Ptr(in.HeightCm),
		DateModified:       utils.NullFloat64Ptr(in.DateModified),
		RecomTestlevel:     utils.NullInt32Ptr(in.RecomTestlevel),
		CreatedBy:          utils.NullInt64Ptr(in.CreatedBy),
		ModBy:              utils.NullInt64Ptr(in.ModBy),
		ModDate:            utils.NullTimePtr(modDate),
		Deleted:            utils.NullInt16FromInt32Ptr(in.Deleted),
		CreatedDate:        utils.NullTimePtr(createdDate),
		Modded:             utils.NullInt16FromInt32Ptr(in.Modded),
		AllowAnonymousData: utils.NullStringPtr(in.AllowAnonymousData),
		Locked:             utils.NullInt16FromInt32Ptr(in.Locked),
		AllowToSprintai:    utils.NullInt32Ptr(in.AllowToSprintai),
		TosprintaiFrom:     utils.NullTimePtr(toSprintFrom),
		StatSent:           utils.NullTimePtr(statSent),
		SporttiID:          utils.NullStringPtr(&sporttiID),
	}, nil
}

// Measurement -> InsertMeasurementParams
func mapMeasurementToParams(in KlabMeasurementInput) (klabsqlc.InsertMeasurementParams, error) {
	var (
		modDate, createdDate, sentTo *time.Time
		err                          error
	)
	if in.ModDate != nil {
		modDate, err = utils.ParseTimestampPtr(in.ModDate)
		if err != nil {
			return klabsqlc.InsertMeasurementParams{}, err
		}
	}
	if in.CreatedDate != nil {
		createdDate, err = utils.ParseTimestampPtr(in.CreatedDate)
		if err != nil {
			return klabsqlc.InsertMeasurementParams{}, err
		}
	}
	if in.SentToSprintAI != nil {
		sentTo, err = utils.ParseTimestampPtr(in.SentToSprintAI)
		if err != nil {
			return klabsqlc.InsertMeasurementParams{}, err
		}
	}

	return klabsqlc.InsertMeasurementParams{
		Idmeasurement:  utils.DerefInt32(in.IdMeasurement),
		Measname:       utils.NullStringPtr(in.MeasName),
		Idcustomer:     utils.DerefInt32(in.IdCustomer),
		Tablename:      utils.NullStringPtr(in.TableName),
		Idpatterndef:   utils.NullStringPtr(in.IdPatternDef),
		DoYear:         utils.NullInt16FromInt32Ptr(in.DoYear),
		DoMonth:        utils.NullInt16FromInt32Ptr(in.DoMonth),
		DoDay:          utils.NullInt16FromInt32Ptr(in.DoDay),
		DoHour:         utils.NullInt16FromInt32Ptr(in.DoHour),
		DoMin:          utils.NullInt16FromInt32Ptr(in.DoMin),
		Sessionno:      utils.NullInt32Ptr(in.SessionNo),
		Info:           utils.NullStringPtr(in.Info),
		Measurements:   utils.NullStringPtr(in.Measurements),
		Groupnotes:     utils.NullStringPtr(in.GroupNotes),
		Cbcharts:       utils.NullStringPtr(in.CbCharts),
		Cbcomments:     utils.NullStringPtr(in.CbComments),
		CreatedBy:      utils.NullInt64Ptr(in.CreatedBy),
		ModBy:          utils.NullInt64Ptr(in.ModBy),
		ModDate:        utils.NullTimePtr(modDate),
		Deleted:        utils.NullInt16FromInt32Ptr(in.Deleted),
		CreatedDate:    utils.NullTimePtr(createdDate),
		Modded:         utils.NullInt16FromInt32Ptr(in.Modded),
		TestLocation:   utils.NullStringPtr(in.TestLocation),
		Keywords:       utils.NullStringPtr(in.Keywords),
		TesterName:     utils.NullStringPtr(in.TesterName),
		ModderName:     utils.NullStringPtr(in.ModderName),
		Meastype:       utils.NullInt32Ptr(in.MeasType),
		SentToSprintai: utils.NullTimePtr(sentTo),
	}, nil
}

// DirTest -> InsertDirTestParams
func mapDirTestToParams(in KlabDirTestInput) (klabsqlc.InsertDirTestParams, error) {
	var (
		modDate, createdDate *time.Time
		err                  error
	)
	if in.ModDate != nil {
		modDate, err = utils.ParseTimestampPtr(in.ModDate)
		if err != nil {
			return klabsqlc.InsertDirTestParams{}, err
		}
	}
	if in.CreatedDate != nil {
		createdDate, err = utils.ParseTimestampPtr(in.CreatedDate)
		if err != nil {
			return klabsqlc.InsertDirTestParams{}, err
		}
	}

	return klabsqlc.InsertDirTestParams{
		Iddirtest:     utils.DerefInt32(in.IdDirTest),
		Idmeasurement: utils.DerefInt32(in.IdMeasurement),
		Meascols:      utils.NullStringPtr(in.MeasCols),
		Weightkg:      utils.NullFloat64Ptr(in.WeightKg),
		Heightcm:      utils.NullFloat64Ptr(in.HeightCm),
		Bmi:           utils.NullFloat64Ptr(in.BMI),
		FatPr:         utils.NullFloat64Ptr(in.FatPr),
		FatP1:         utils.NullFloat64Ptr(in.FatP1),
		FatP2:         utils.NullFloat64Ptr(in.FatP2),
		FatP3:         utils.NullFloat64Ptr(in.FatP3),
		FatP4:         utils.NullFloat64Ptr(in.FatP4),
		FatStyle:      utils.NullInt32Ptr(in.FatStyle),
		FatEquip:      utils.NullStringPtr(in.FatEquip),
		Fvc:           utils.NullFloat64Ptr(in.FVC),
		Fev1:          utils.NullFloat64Ptr(in.FEV1),
		AirPress:      utils.NullFloat64Ptr(in.AirPress),
		AirTemp:       utils.NullFloat64Ptr(in.AirTemp),
		AirHumid:      utils.NullFloat64Ptr(in.AirHumid),
		Testprotocol:  utils.NullStringPtr(in.TestProtocol),
		AirPressUnit:  utils.NullInt32Ptr(in.AirPressUnit),
		Settingslist:  utils.NullStringPtr(in.SettingsList),
		Lt1X:          utils.NullFloat64Ptr(in.Lt1X),
		Lt1Y:          utils.NullFloat64Ptr(in.Lt1Y),
		Lt2X:          utils.NullFloat64Ptr(in.Lt2X),
		Lt2Y:          utils.NullFloat64Ptr(in.Lt2Y),
		Vt1X:          utils.NullFloat64Ptr(in.Vt1X),
		Vt2X:          utils.NullFloat64Ptr(in.Vt2X),
		Vt1Y:          utils.NullFloat64Ptr(in.Vt1Y),
		Vt2Y:          utils.NullFloat64Ptr(in.Vt2Y),
		Lt1CalcX:      utils.NullFloat64Ptr(in.Lt1CalcX),
		Lt1CalcY:      utils.NullFloat64Ptr(in.Lt1CalcY),
		Lt2CalcX:      utils.NullFloat64Ptr(in.Lt2CalcX),
		Lt2CalcY:      utils.NullFloat64Ptr(in.Lt2CalcY),
		Protocolmodel: utils.NullInt16FromInt32Ptr(in.ProtocolModel),
		Testtype:      utils.NullInt16FromInt32Ptr(in.TestType),
		Protocolxval:  utils.NullInt16FromInt32Ptr(in.ProtocolXVal),
		Steptime:      utils.NullInt32Ptr(in.StepTime),
		WRest:         utils.NullInt16FromInt32Ptr(in.WRest),
		CreatedBy:     utils.NullInt64Ptr(in.CreatedBy),
		ModBy:         utils.NullInt64Ptr(in.ModBy),
		ModDate:       utils.NullTimePtr(modDate),
		Deleted:       utils.NullInt16FromInt32Ptr(in.Deleted),
		CreatedDate:   utils.NullTimePtr(createdDate),
		Modded:        utils.NullInt16FromInt32Ptr(in.Modded),
		Norawdata:     utils.NullInt16FromInt32Ptr(in.NoRawData),
	}, nil
}

// DirTestStep -> InsertDirTestStepParams
func mapDirTestStepToParams(in KlabDirTestStepInput) (klabsqlc.InsertDirTestStepParams, error) {
	var (
		modDate, createdDate *time.Time
		err                  error
	)
	if in.ModDate != nil {
		modDate, err = utils.ParseTimestampPtr(in.ModDate)
		if err != nil {
			return klabsqlc.InsertDirTestStepParams{}, err
		}
	}
	if in.CreatedDate != nil {
		createdDate, err = utils.ParseTimestampPtr(in.CreatedDate)
		if err != nil {
			return klabsqlc.InsertDirTestStepParams{}, err
		}
	}

	return klabsqlc.InsertDirTestStepParams{
		Iddirteststeps: utils.DerefInt32(in.Iddirteststeps),
		Idmeasurement:  utils.DerefInt32(in.Idmeasurement),
		Stepno:         utils.NullInt32Ptr(in.Stepno),
		AnaTime:        utils.NullInt32Ptr(in.AnaTime),
		Timestop:       utils.NullFloat64Ptr(in.TimeStop),
		Speed:          utils.NullFloat64Ptr(in.Speed),
		Pace:           utils.NullFloat64Ptr(in.Pace),
		Angle:          utils.NullFloat64Ptr(in.Angle),
		Elev:           utils.NullFloat64Ptr(in.Elev),
		Vo2calc:        utils.NullFloat64Ptr(in.Vo2calc),
		TTot:           utils.NullFloat64Ptr(in.TTot),
		TEx:            utils.NullFloat64Ptr(in.TEx),
		Fico2:          utils.NullFloat64Ptr(in.Fico2),
		Fio2:           utils.NullFloat64Ptr(in.Fio2),
		Feco2:          utils.NullFloat64Ptr(in.Feco2),
		Feo2:           utils.NullFloat64Ptr(in.Feo2),
		Vde:            utils.NullFloat64Ptr(in.Vde),
		Vco2:           utils.NullFloat64Ptr(in.Vco2),
		Vo2:            utils.NullFloat64Ptr(in.Vo2),
		Bf:             utils.NullFloat64Ptr(in.Bf),
		Ve:             utils.NullFloat64Ptr(in.Ve),
		Petco2:         utils.NullFloat64Ptr(in.Petco2),
		Peto2:          utils.NullFloat64Ptr(in.Peto2),
		Vo2kg:          utils.NullFloat64Ptr(in.Vo2kg),
		Re:             utils.NullFloat64Ptr(in.Re),
		Hr:             utils.NullFloat64Ptr(in.Hr),
		La:             utils.NullFloat64Ptr(in.La),
		Rer:            utils.NullFloat64Ptr(in.Rer),
		VeStpd:         utils.NullFloat64Ptr(in.VeStpd),
		Veo2:           utils.NullFloat64Ptr(in.Veo2),
		Veco2:          utils.NullFloat64Ptr(in.Veco2),
		Tv:             utils.NullFloat64Ptr(in.Tv),
		EeAe:           utils.NullFloat64Ptr(in.EeAe),
		LaVo2:          utils.NullFloat64Ptr(in.LaVo2),
		O2pulse:        utils.NullFloat64Ptr(in.O2pulse),
		VdeTv:          utils.NullFloat64Ptr(in.VdeTv),
		Va:             utils.NullFloat64Ptr(in.Va),
		O2sa:           utils.NullFloat64Ptr(in.O2sa),
		Rpe:            utils.NullFloat64Ptr(in.Rpe),
		BpSys:          utils.NullFloat64Ptr(in.BpSys),
		BpDia:          utils.NullFloat64Ptr(in.BpDia),
		Own1:           utils.NullFloat64Ptr(in.Own1),
		Own2:           utils.NullFloat64Ptr(in.Own2),
		Own3:           utils.NullFloat64Ptr(in.Own3),
		Own4:           utils.NullFloat64Ptr(in.Own4),
		Own5:           utils.NullFloat64Ptr(in.Own5),
		StepIsRest:     utils.NullInt32Ptr(in.StepIsRest),
		StepIs30max:    utils.NullInt32Ptr(in.StepIs30max),
		StepIs60max:    utils.NullInt32Ptr(in.StepIs60max),
		StepIsRec:      utils.NullInt32Ptr(in.StepIsRec),
		CalcStart:      utils.NullInt32Ptr(in.CalcStart),
		CalcEnd:        utils.NullInt32Ptr(in.CalcEnd),
		Comments:       utils.NullStringPtr(in.Comments),
		Timestart:      utils.NullFloat64Ptr(in.TimeStart),
		Duration:       utils.NullFloat64Ptr(in.Duration),
		Eco:            utils.NullFloat64Ptr(in.Eco),
		P:              utils.NullFloat64Ptr(in.P),
		Wkg:            utils.NullFloat64Ptr(in.Wkg),
		Vo230s:         utils.NullFloat64Ptr(in.Vo230s),
		Vo2Pr:          utils.NullFloat64Ptr(in.Vo2Pr),
		StepIsLast:     utils.NullInt32Ptr(in.StepIsLast),
		Deleted:        utils.NullInt16FromInt32Ptr(in.Deleted),
		CreatedBy:      utils.NullInt64Ptr(in.CreatedBy),
		ModBy:          utils.NullInt64Ptr(in.ModBy),
		ModDate:        utils.NullTimePtr(modDate),
		CreatedDate:    utils.NullTimePtr(createdDate),
		Modded:         utils.NullInt16FromInt32Ptr(in.Modded),
		Own6:           utils.NullFloat64Ptr(in.Own6),
		Own7:           utils.NullFloat64Ptr(in.Own7),
		Own8:           utils.NullFloat64Ptr(in.Own8),
		Own9:           utils.NullFloat64Ptr(in.Own9),
		Own10:          utils.NullFloat64Ptr(in.Own10),
		To2:            utils.NullFloat64Ptr(in.To2),
		Tco2:           utils.NullFloat64Ptr(in.Tco2),
	}, nil
}

// DirReport -> InsertDirReportParams
func mapDirReportToParams(in KlabDirReportInput) (klabsqlc.InsertDirReportParams, error) {
	var (
		modDate, createdDate *time.Time
		err                  error
	)
	if in.ModDate != nil {
		modDate, err = utils.ParseTimestampPtr(in.ModDate)
		if err != nil {
			return klabsqlc.InsertDirReportParams{}, err
		}
	}
	if in.CreatedDate != nil {
		createdDate, err = utils.ParseTimestampPtr(in.CreatedDate)
		if err != nil {
			return klabsqlc.InsertDirReportParams{}, err
		}
	}

	return klabsqlc.InsertDirReportParams{
		Iddirreport:      utils.DerefInt32(in.Iddirreport),
		PageInstructions: utils.NullStringPtr(in.PageInstr),
		Idmeasurement:    utils.DerefInt32(in.Idmeasurement),
		TemplateRec:      utils.NullInt32Ptr(in.TemplateRec),
		LibrecName:       utils.NullStringPtr(in.LibrecName),
		CreatedBy:        utils.NullInt64Ptr(in.CreatedBy),
		ModBy:            utils.NullInt64Ptr(in.ModBy),
		ModDate:          utils.NullTimePtr(modDate),
		Deleted:          utils.NullInt16FromInt32Ptr(in.Deleted),
		CreatedDate:      utils.NullTimePtr(createdDate),
		Modded:           utils.NullInt16FromInt32Ptr(in.Modded),
	}, nil
}

// DirRawData -> InsertDirRawDataParams
func mapDirRawDataToParams(in KlabDirRawDataInput) (klabsqlc.InsertDirRawDataParams, error) {
	var (
		modDate, createdDate *time.Time
		err                  error
	)
	if in.ModDate != nil {
		modDate, err = utils.ParseTimestampPtr(in.ModDate)
		if err != nil {
			return klabsqlc.InsertDirRawDataParams{}, err
		}
	}
	if in.CreatedDate != nil {
		createdDate, err = utils.ParseTimestampPtr(in.CreatedDate)
		if err != nil {
			return klabsqlc.InsertDirRawDataParams{}, err
		}
	}

	return klabsqlc.InsertDirRawDataParams{
		Iddirrawdata:  utils.DerefInt32(in.IdDirRawData),
		Idmeasurement: utils.DerefInt32(in.IdMeasurement),
		Rawdata:       utils.NullStringPtr(in.RawData),
		Columndata:    utils.NullStringPtr(in.ColumnData),
		Info:          utils.NullStringPtr(in.Info),
		Unitsdata:     utils.NullStringPtr(in.UnitsData),
		CreatedBy:     utils.NullInt64Ptr(in.CreatedBy),
		ModBy:         utils.NullInt64Ptr(in.ModBy),
		ModDate:       utils.NullTimePtr(modDate),
		Deleted:       utils.NullInt16FromInt32Ptr(in.Deleted),
		CreatedDate:   utils.NullTimePtr(createdDate),
		Modded:        utils.NullInt16FromInt32Ptr(in.Modded),
	}, nil
}

// DirResults -> InsertDirResultsParams
func mapDirResultsToParams(in KlabDirResultsInput) (klabsqlc.InsertDirResultsParams, error) {
	var (
		modDate, createdDate *time.Time
		err                  error
	)
	if in.ModDate != nil {
		modDate, err = utils.ParseTimestampPtr(in.ModDate)
		if err != nil {
			return klabsqlc.InsertDirResultsParams{}, err
		}
	}
	if in.CreatedDate != nil {
		createdDate, err = utils.ParseTimestampPtr(in.CreatedDate)
		if err != nil {
			return klabsqlc.InsertDirResultsParams{}, err
		}
	}

	return klabsqlc.InsertDirResultsParams{
		Iddirresults:       utils.DerefInt32(in.Iddirresults),
		Idmeasurement:      utils.DerefInt32(in.Idmeasurement),
		MaxVo2mlkgmin:      utils.NullFloat64Ptr(in.MaxVo2mlkgmin),
		MaxVo2mlmin:        utils.NullFloat64Ptr(in.MaxVo2mlmin),
		MaxVo2:             utils.NullFloat64Ptr(in.MaxVo2),
		MaxHr:              utils.NullFloat64Ptr(in.MaxHr),
		MaxSpeed:           utils.NullFloat64Ptr(in.MaxSpeed),
		MaxPace:            utils.NullFloat64Ptr(in.MaxPace),
		MaxP:               utils.NullFloat64Ptr(in.MaxP),
		MaxPkg:             utils.NullFloat64Ptr(in.MaxPkg),
		MaxAngle:           utils.NullFloat64Ptr(in.MaxAngle),
		MaxLac:             utils.NullFloat64Ptr(in.MaxLac),
		MaxAdd1:            utils.NullFloat64Ptr(in.MaxAdd1),
		MaxAdd2:            utils.NullFloat64Ptr(in.MaxAdd2),
		MaxAdd3:            utils.NullFloat64Ptr(in.MaxAdd3),
		LacAnkVo2mlkgmin:   utils.NullFloat64Ptr(in.LacAnkVo2mlkgmin),
		LacAnkVo2mlmin:     utils.NullFloat64Ptr(in.LacAnkVo2mlmin),
		LacAnkVo2:          utils.NullFloat64Ptr(in.LacAnkVo2),
		LacAnkVo2pr:        utils.NullFloat64Ptr(in.LacAnkVo2pr),
		LacAnkHr:           utils.NullFloat64Ptr(in.LacAnkHr),
		LacAnkSpeed:        utils.NullFloat64Ptr(in.LacAnkSpeed),
		LacAnkPace:         utils.NullFloat64Ptr(in.LacAnkPace),
		LacAnkP:            utils.NullFloat64Ptr(in.LacAnkP),
		LacAnkPkg:          utils.NullFloat64Ptr(in.LacAnkPkg),
		LacAnkAngle:        utils.NullFloat64Ptr(in.LacAnkAngle),
		LacAnkLac:          utils.NullFloat64Ptr(in.LacAnkLac),
		LacAnkAdd1:         utils.NullFloat64Ptr(in.LacAnkAdd1),
		LacAnkAdd2:         utils.NullFloat64Ptr(in.LacAnkAdd2),
		LacAnkAdd3:         utils.NullFloat64Ptr(in.LacAnkAdd3),
		LacAerkVo2mlkgmin:  utils.NullFloat64Ptr(in.LacAerkVo2mlkgmin),
		LacAerkVo2mlmin:    utils.NullFloat64Ptr(in.LacAerkVo2mlmin),
		LacAerkVo2:         utils.NullFloat64Ptr(in.LacAerkVo2),
		LacAerkVo2pr:       utils.NullFloat64Ptr(in.LacAerkVo2pr),
		LacAerkHr:          utils.NullFloat64Ptr(in.LacAerkHr),
		LacAerkSpeed:       utils.NullFloat64Ptr(in.LacAerkSpeed),
		LacAerkPace:        utils.NullFloat64Ptr(in.LacAerkPace),
		LacAerkP:           utils.NullFloat64Ptr(in.LacAerkP),
		LacAerkPkg:         utils.NullFloat64Ptr(in.LacAerkPkg),
		LacAerkAngle:       utils.NullFloat64Ptr(in.LacAerkAngle),
		LacAerkLac:         utils.NullFloat64Ptr(in.LacAerkLac),
		LacAerkAdd1:        utils.NullFloat64Ptr(in.LacAerkAdd1),
		LacAerkAdd2:        utils.NullFloat64Ptr(in.LacAerkAdd2),
		LacAerkAdd3:        utils.NullFloat64Ptr(in.LacAerkAdd3),
		VentAnkVo2mlkgmin:  utils.NullFloat64Ptr(in.VentAnkVo2mlkgmin),
		VentAnkVo2mlmin:    utils.NullFloat64Ptr(in.VentAnkVo2mlmin),
		VentAnkVo2:         utils.NullFloat64Ptr(in.VentAnkVo2),
		VentAnkVo2pr:       utils.NullFloat64Ptr(in.VentAnkVo2pr),
		VentAnkHr:          utils.NullFloat64Ptr(in.VentAnkHr),
		VentAnkSpeed:       utils.NullFloat64Ptr(in.VentAnkSpeed),
		VentAnkPace:        utils.NullFloat64Ptr(in.VentAnkPace),
		VentAnkP:           utils.NullFloat64Ptr(in.VentAnkP),
		VentAnkPkg:         utils.NullFloat64Ptr(in.VentAnkPkg),
		VentAnkAngle:       utils.NullFloat64Ptr(in.VentAnkAngle),
		VentAnkLac:         utils.NullFloat64Ptr(in.VentAnkLac),
		VentAnkAdd1:        utils.NullFloat64Ptr(in.VentAnkAdd1),
		VentAnkAdd2:        utils.NullFloat64Ptr(in.VentAnkAdd2),
		VentAnkAdd3:        utils.NullFloat64Ptr(in.VentAnkAdd3),
		VentAerkVo2mlkgmin: utils.NullFloat64Ptr(in.VentAerkVo2mlkgmin),
		VentAerkVo2mlmin:   utils.NullFloat64Ptr(in.VentAerkVo2mlmin),
		VentAerkVo2:        utils.NullFloat64Ptr(in.VentAerkVo2),
		VentAerkVo2pr:      utils.NullFloat64Ptr(in.VentAerkVo2pr),
		VentAerkHr:         utils.NullFloat64Ptr(in.VentAerkHr),
		VentAerkSpeed:      utils.NullFloat64Ptr(in.VentAerkSpeed),
		VentAerkPace:       utils.NullFloat64Ptr(in.VentAerkPace),
		VentAerkP:          utils.NullFloat64Ptr(in.VentAerkP),
		VentAerkPkg:        utils.NullFloat64Ptr(in.VentAerkPkg),
		VentAerkAngle:      utils.NullFloat64Ptr(in.VentAerkAngle),
		VentAerkLac:        utils.NullFloat64Ptr(in.VentAerkLac),
		VentAerkAdd1:       utils.NullFloat64Ptr(in.VentAerkAdd1),
		VentAerkAdd2:       utils.NullFloat64Ptr(in.VentAerkAdd2),
		VentAerkAdd3:       utils.NullFloat64Ptr(in.VentAerkAdd3),
		CreatedBy:          utils.NullInt64Ptr(in.CreatedBy),
		ModBy:              utils.NullInt64Ptr(in.ModBy),
		ModDate:            utils.NullTimePtr(modDate),
		Deleted:            utils.NullInt16FromInt32Ptr(in.Deleted),
		CreatedDate:        utils.NullTimePtr(createdDate),
		Modded:             utils.NullInt16FromInt32Ptr(in.Modded),
	}, nil
}
