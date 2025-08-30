package klab

import (
	"time"

	klabsqlc "github.com/DeRuina/KUHA-REST-API/internal/db/klab"
)

// Clean response types (what you want in JSON)
type KlabMeasurementResponse struct {
	IdMeasurement  int32      `json:"idMeasurement"`
	MeasName       *string    `json:"MeasName,omitempty"`
	IdCustomer     int32      `json:"idCustomer"`
	TableName      *string    `json:"tableName,omitempty"`
	IdPatternDef   *string    `json:"idpatterndef,omitempty"`
	DoYear         *int16     `json:"do_year,omitempty"`
	DoMonth        *int16     `json:"do_month,omitempty"`
	DoDay          *int16     `json:"do_day,omitempty"`
	DoHour         *int16     `json:"do_hour,omitempty"`
	DoMin          *int16     `json:"do_min,omitempty"`
	SessionNo      *int32     `json:"sessionno,omitempty"`
	Info           *string    `json:"info,omitempty"`
	Measurements   *string    `json:"measurements,omitempty"`
	GroupNotes     *string    `json:"groupnotes,omitempty"`
	CbCharts       *string    `json:"cbcharts,omitempty"`
	CbComments     *string    `json:"cbcomments,omitempty"`
	CreatedBy      *int64     `json:"created_by,omitempty"`
	ModBy          *int64     `json:"mod_by,omitempty"`
	ModDate        *time.Time `json:"mod_date,omitempty"`
	Deleted        *int16     `json:"deleted,omitempty"`
	CreatedDate    *time.Time `json:"created_date,omitempty"`
	Modded         *int16     `json:"modded,omitempty"`
	TestLocation   *string    `json:"test_location,omitempty"`
	Keywords       *string    `json:"keywords,omitempty"`
	TesterName     *string    `json:"tester_name,omitempty"`
	ModderName     *string    `json:"modder_name,omitempty"`
	MeasType       *int32     `json:"meastype,omitempty"`
	SentToSprintAI *time.Time `json:"sent_to_sprintai,omitempty"`
}

type KlabDirTestResponse struct {
	IdDirTest     int32      `json:"idDirTest"`
	IdMeasurement int32      `json:"idMeasurement"`
	MeasCols      *string    `json:"MeasCols,omitempty"`
	WeightKg      *float64   `json:"weightkg,omitempty"`
	HeightCm      *float64   `json:"heightcm,omitempty"`
	Bmi           *float64   `json:"bmi,omitempty"`
	FatPr         *float64   `json:"fat_pr,omitempty"`
	FatP1         *float64   `json:"fat_p1,omitempty"`
	FatP2         *float64   `json:"fat_p2,omitempty"`
	FatP3         *float64   `json:"fat_p3,omitempty"`
	FatP4         *float64   `json:"fat_p4,omitempty"`
	FatStyle      *int32     `json:"fat_style,omitempty"`
	FatEquip      *string    `json:"fat_equip,omitempty"`
	Fvc           *float64   `json:"fvc,omitempty"`
	Fev1          *float64   `json:"fev1,omitempty"`
	AirPress      *float64   `json:"air_press,omitempty"`
	AirTemp       *float64   `json:"air_temp,omitempty"`
	AirHumid      *float64   `json:"air_humid,omitempty"`
	TestProtocol  *string    `json:"testprotocol,omitempty"`
	AirPressUnit  *int32     `json:"air_press_unit,omitempty"`
	SettingsList  *string    `json:"settingslist,omitempty"`
	Lt1X          *float64   `json:"lt1_x,omitempty"`
	Lt1Y          *float64   `json:"lt1_y,omitempty"`
	Lt2X          *float64   `json:"lt2_x,omitempty"`
	Lt2Y          *float64   `json:"lt2_y,omitempty"`
	Vt1X          *float64   `json:"vt1_x,omitempty"`
	Vt2X          *float64   `json:"vt2_x,omitempty"`
	Vt1Y          *float64   `json:"vt1_y,omitempty"`
	Vt2Y          *float64   `json:"vt2_y,omitempty"`
	Lt1CalcX      *float64   `json:"lt1_calc_x,omitempty"`
	Lt1CalcY      *float64   `json:"lt1_calc_y,omitempty"`
	Lt2CalcX      *float64   `json:"lt2_calc_x,omitempty"`
	Lt2CalcY      *float64   `json:"lt2_calc_y,omitempty"`
	ProtocolModel *int16     `json:"protocolmodel,omitempty"`
	TestType      *int16     `json:"testtype,omitempty"`
	ProtocolXVal  *int16     `json:"protocolxval,omitempty"`
	StepTime      *int32     `json:"steptime,omitempty"`
	WRest         *int16     `json:"w_rest,omitempty"`
	CreatedBy     *int64     `json:"created_by,omitempty"`
	ModBy         *int64     `json:"mod_by,omitempty"`
	ModDate       *time.Time `json:"mod_date,omitempty"`
	Deleted       *int16     `json:"deleted,omitempty"`
	CreatedDate   *time.Time `json:"created_date,omitempty"`
	Modded        *int16     `json:"modded,omitempty"`
	NoRawData     *int16     `json:"norawdata,omitempty"`
}

type KlabDirTestStepResponse struct {
	IdDirTestSteps int32      `json:"iddirteststeps"`
	IdMeasurement  int32      `json:"idmeasurement"`
	StepNo         *int32     `json:"stepno,omitempty"`
	AnaTime        *int32     `json:"ana_time,omitempty"`
	TimeStop       *float64   `json:"timestop,omitempty"`
	Speed          *float64   `json:"speed,omitempty"`
	Pace           *float64   `json:"pace,omitempty"`
	Angle          *float64   `json:"angle,omitempty"`
	Elev           *float64   `json:"elev,omitempty"`
	Vo2Calc        *float64   `json:"vo2calc,omitempty"`
	TTot           *float64   `json:"t_tot,omitempty"`
	TEx            *float64   `json:"t_ex,omitempty"`
	Fico2          *float64   `json:"fico2,omitempty"`
	Fio2           *float64   `json:"fio2,omitempty"`
	Feco2          *float64   `json:"feco2,omitempty"`
	Feo2           *float64   `json:"feo2,omitempty"`
	Vde            *float64   `json:"vde,omitempty"`
	Vco2           *float64   `json:"vco2,omitempty"`
	Vo2            *float64   `json:"vo2,omitempty"`
	Bf             *float64   `json:"bf,omitempty"`
	Ve             *float64   `json:"ve,omitempty"`
	Petco2         *float64   `json:"petco2,omitempty"`
	Peto2          *float64   `json:"peto2,omitempty"`
	Vo2Kg          *float64   `json:"vo2kg,omitempty"`
	Re             *float64   `json:"re,omitempty"`
	Hr             *float64   `json:"hr,omitempty"`
	La             *float64   `json:"la,omitempty"`
	Rer            *float64   `json:"rer,omitempty"`
	VeStpd         *float64   `json:"ve_stpd,omitempty"`
	Veo2           *float64   `json:"veo2,omitempty"`
	Veco2          *float64   `json:"veco2,omitempty"`
	Tv             *float64   `json:"tv,omitempty"`
	EeAe           *float64   `json:"ee_ae,omitempty"`
	LaVo2          *float64   `json:"la_vo2,omitempty"`
	O2Pulse        *float64   `json:"o2pulse,omitempty"`
	VdeTv          *float64   `json:"vde_tv,omitempty"`
	Va             *float64   `json:"va,omitempty"`
	O2Sa           *float64   `json:"o2sa,omitempty"`
	Rpe            *float64   `json:"rpe,omitempty"`
	BpSys          *float64   `json:"bp_sys,omitempty"`
	BpDia          *float64   `json:"bp_dia,omitempty"`
	Own1           *float64   `json:"own1,omitempty"`
	Own2           *float64   `json:"own2,omitempty"`
	Own3           *float64   `json:"own3,omitempty"`
	Own4           *float64   `json:"own4,omitempty"`
	Own5           *float64   `json:"own5,omitempty"`
	StepIsRest     *int32     `json:"step_is_rest,omitempty"`
	StepIs30Max    *int32     `json:"step_is_30max,omitempty"`
	StepIs60Max    *int32     `json:"step_is_60max,omitempty"`
	StepIsRec      *int32     `json:"step_is_rec,omitempty"`
	CalcStart      *int32     `json:"calc_start,omitempty"`
	CalcEnd        *int32     `json:"calc_end,omitempty"`
	Comments       *string    `json:"comments,omitempty"`
	TimeStart      *float64   `json:"timestart,omitempty"`
	Duration       *float64   `json:"duration,omitempty"`
	Eco            *float64   `json:"eco,omitempty"`
	P              *float64   `json:"p,omitempty"`
	Wkg            *float64   `json:"wkg,omitempty"`
	Vo230s         *float64   `json:"vo2_30s,omitempty"`
	Vo2Pr          *float64   `json:"vo2_pr,omitempty"`
	StepIsLast     *int32     `json:"step_is_last,omitempty"`
	Deleted        *int16     `json:"deleted,omitempty"`
	CreatedBy      *int64     `json:"created_by,omitempty"`
	ModBy          *int64     `json:"mod_by,omitempty"`
	ModDate        *time.Time `json:"mod_date,omitempty"`
	CreatedDate    *time.Time `json:"created_date,omitempty"`
	Modded         *int16     `json:"modded,omitempty"`
	Own6           *float64   `json:"own6,omitempty"`
	Own7           *float64   `json:"own7,omitempty"`
	Own8           *float64   `json:"own8,omitempty"`
	Own9           *float64   `json:"own9,omitempty"`
	Own10          *float64   `json:"own10,omitempty"`
	To2            *float64   `json:"to2,omitempty"`
	Tco2           *float64   `json:"tco2,omitempty"`
}

type KlabDirReportResponse struct {
	IdDirReport      int32      `json:"iddirreport"`
	PageInstructions *string    `json:"page_instructions,omitempty"`
	IdMeasurement    int32      `json:"idmeasurement"`
	TemplateRec      *int32     `json:"template_rec,omitempty"`
	LibrecName       *string    `json:"librec_name,omitempty"`
	CreatedBy        *int64     `json:"created_by,omitempty"`
	ModBy            *int64     `json:"mod_by,omitempty"`
	ModDate          *time.Time `json:"mod_date,omitempty"`
	Deleted          *int16     `json:"deleted,omitempty"`
	CreatedDate      *time.Time `json:"created_date,omitempty"`
	Modded           *int16     `json:"modded,omitempty"`
}

type KlabDirRawDataResponse struct {
	IdDirRawData  int32      `json:"idDirRawData"`
	IdMeasurement int32      `json:"idMeasurement"`
	RawData       *string    `json:"rawdata,omitempty"`
	ColumnData    *string    `json:"columndata,omitempty"`
	Info          *string    `json:"info,omitempty"`
	UnitsData     *string    `json:"unitsdata,omitempty"`
	CreatedBy     *int64     `json:"created_by,omitempty"`
	ModBy         *int64     `json:"mod_by,omitempty"`
	ModDate       *time.Time `json:"mod_date,omitempty"`
	Deleted       *int16     `json:"deleted,omitempty"`
	CreatedDate   *time.Time `json:"created_date,omitempty"`
	Modded        *int16     `json:"modded,omitempty"`
}

type KlabDirResultsResponse struct {
	IdDirResults       int32      `json:"iddirresults"`
	IdMeasurement      int32      `json:"idmeasurement"`
	MaxVo2MlKgMin      *float64   `json:"max_vo2mlkgmin,omitempty"`
	MaxVo2MlMin        *float64   `json:"max_vo2mlmin,omitempty"`
	MaxVo2             *float64   `json:"max_vo2,omitempty"`
	MaxHr              *float64   `json:"max_hr,omitempty"`
	MaxSpeed           *float64   `json:"max_speed,omitempty"`
	MaxPace            *float64   `json:"max_pace,omitempty"`
	MaxP               *float64   `json:"max_p,omitempty"`
	MaxPkg             *float64   `json:"max_pkg,omitempty"`
	MaxAngle           *float64   `json:"max_angle,omitempty"`
	MaxLac             *float64   `json:"max_lac,omitempty"`
	MaxAdd1            *float64   `json:"max_add1,omitempty"`
	MaxAdd2            *float64   `json:"max_add2,omitempty"`
	MaxAdd3            *float64   `json:"max_add3,omitempty"`
	LacAnkVo2MlKgMin   *float64   `json:"lac_ank_vo2mlkgmin,omitempty"`
	LacAnkVo2MlMin     *float64   `json:"lac_ank_vo2mlmin,omitempty"`
	LacAnkVo2          *float64   `json:"lac_ank_vo2,omitempty"`
	LacAnkVo2Pr        *float64   `json:"lac_ank_vo2pr,omitempty"`
	LacAnkHr           *float64   `json:"lac_ank_hr,omitempty"`
	LacAnkSpeed        *float64   `json:"lac_ank_speed,omitempty"`
	LacAnkPace         *float64   `json:"lac_ank_pace,omitempty"`
	LacAnkP            *float64   `json:"lac_ank_p,omitempty"`
	LacAnkPkg          *float64   `json:"lac_ank_pkg,omitempty"`
	LacAnkAngle        *float64   `json:"lac_ank_angle,omitempty"`
	LacAnkLac          *float64   `json:"lac_ank_lac,omitempty"`
	LacAnkAdd1         *float64   `json:"lac_ank_add1,omitempty"`
	LacAnkAdd2         *float64   `json:"lac_ank_add2,omitempty"`
	LacAnkAdd3         *float64   `json:"lac_ank_add3,omitempty"`
	LacAerkVo2MlKgMin  *float64   `json:"lac_aerk_vo2mlkgmin,omitempty"`
	LacAerkVo2MlMin    *float64   `json:"lac_aerk_vo2mlmin,omitempty"`
	LacAerkVo2         *float64   `json:"lac_aerk_vo2,omitempty"`
	LacAerkVo2Pr       *float64   `json:"lac_aerk_vo2pr,omitempty"`
	LacAerkHr          *float64   `json:"lac_aerk_hr,omitempty"`
	LacAerkSpeed       *float64   `json:"lac_aerk_speed,omitempty"`
	LacAerkPace        *float64   `json:"lac_aerk_pace,omitempty"`
	LacAerkP           *float64   `json:"lac_aerk_p,omitempty"`
	LacAerkPkg         *float64   `json:"lac_aerk_pkg,omitempty"`
	LacAerkAngle       *float64   `json:"lac_aerk_angle,omitempty"`
	LacAerkLac         *float64   `json:"lac_aerk_lac,omitempty"`
	LacAerkAdd1        *float64   `json:"lac_aerk_add1,omitempty"`
	LacAerkAdd2        *float64   `json:"lac_aerk_add2,omitempty"`
	LacAerkAdd3        *float64   `json:"lac_aerk_add3,omitempty"`
	VentAnkVo2MlKgMin  *float64   `json:"vent_ank_vo2mlkgmin,omitempty"`
	VentAnkVo2MlMin    *float64   `json:"vent_ank_vo2mlmin,omitempty"`
	VentAnkVo2         *float64   `json:"vent_ank_vo2,omitempty"`
	VentAnkVo2Pr       *float64   `json:"vent_ank_vo2pr,omitempty"`
	VentAnkHr          *float64   `json:"vent_ank_hr,omitempty"`
	VentAnkSpeed       *float64   `json:"vent_ank_speed,omitempty"`
	VentAnkPace        *float64   `json:"vent_ank_pace,omitempty"`
	VentAnkP           *float64   `json:"vent_ank_p,omitempty"`
	VentAnkPkg         *float64   `json:"vent_ank_pkg,omitempty"`
	VentAnkAngle       *float64   `json:"vent_ank_angle,omitempty"`
	VentAnkLac         *float64   `json:"vent_ank_lac,omitempty"`
	VentAnkAdd1        *float64   `json:"vent_ank_add1,omitempty"`
	VentAnkAdd2        *float64   `json:"vent_ank_add2,omitempty"`
	VentAnkAdd3        *float64   `json:"vent_ank_add3,omitempty"`
	VentAerkVo2MlKgMin *float64   `json:"vent_aerk_vo2mlkgmin,omitempty"`
	VentAerkVo2MlMin   *float64   `json:"vent_aerk_vo2mlmin,omitempty"`
	VentAerkVo2        *float64   `json:"vent_aerk_vo2,omitempty"`
	VentAerkVo2Pr      *float64   `json:"vent_aerk_vo2pr,omitempty"`
	VentAerkHr         *float64   `json:"vent_aerk_hr,omitempty"`
	VentAerkSpeed      *float64   `json:"vent_aerk_speed,omitempty"`
	VentAerkPace       *float64   `json:"vent_aerk_pace,omitempty"`
	VentAerkP          *float64   `json:"vent_aerk_p,omitempty"`
	VentAerkPkg        *float64   `json:"vent_aerk_pkg,omitempty"`
	VentAerkAngle      *float64   `json:"vent_aerk_angle,omitempty"`
	VentAerkLac        *float64   `json:"vent_aerk_lac,omitempty"`
	VentAerkAdd1       *float64   `json:"vent_aerk_add1,omitempty"`
	VentAerkAdd2       *float64   `json:"vent_aerk_add2,omitempty"`
	VentAerkAdd3       *float64   `json:"vent_aerk_add3,omitempty"`
	CreatedBy          *int64     `json:"created_by,omitempty"`
	ModBy              *int64     `json:"mod_by,omitempty"`
	ModDate            *time.Time `json:"mod_date,omitempty"`
	Deleted            *int16     `json:"deleted,omitempty"`
	CreatedDate        *time.Time `json:"created_date,omitempty"`
	Modded             *int16     `json:"modded,omitempty"`
}

// Main response struct
type KlabDataNoCustomerResponse struct {
	CustomerID   int32                     `json:"customer_id"`
	Measurements []KlabMeasurementResponse `json:"measurements"`
	DirTests     []KlabDirTestResponse     `json:"dirtest"`
	DirTestSteps []KlabDirTestStepResponse `json:"dirteststeps"`
	DirReports   []KlabDirReportResponse   `json:"dirreport"`
	DirRawData   []KlabDirRawDataResponse  `json:"dirrawdata"`
	DirResults   []KlabDirResultsResponse  `json:"dirresults"`
}

func convertMeasurement(m klabsqlc.MeasurementList) KlabMeasurementResponse {
	resp := KlabMeasurementResponse{
		IdMeasurement: m.Idmeasurement,
		IdCustomer:    m.Idcustomer,
	}

	if m.Measname.Valid {
		resp.MeasName = &m.Measname.String
	}
	if m.Tablename.Valid {
		resp.TableName = &m.Tablename.String
	}
	if m.Idpatterndef.Valid {
		resp.IdPatternDef = &m.Idpatterndef.String
	}
	if m.DoYear.Valid {
		resp.DoYear = &m.DoYear.Int16
	}
	if m.DoMonth.Valid {
		resp.DoMonth = &m.DoMonth.Int16
	}
	if m.DoDay.Valid {
		resp.DoDay = &m.DoDay.Int16
	}
	if m.DoHour.Valid {
		resp.DoHour = &m.DoHour.Int16
	}
	if m.DoMin.Valid {
		resp.DoMin = &m.DoMin.Int16
	}
	if m.Sessionno.Valid {
		resp.SessionNo = &m.Sessionno.Int32
	}
	if m.Info.Valid {
		resp.Info = &m.Info.String
	}
	if m.Measurements.Valid {
		resp.Measurements = &m.Measurements.String
	}
	if m.Groupnotes.Valid {
		resp.GroupNotes = &m.Groupnotes.String
	}
	if m.Cbcharts.Valid {
		resp.CbCharts = &m.Cbcharts.String
	}
	if m.Cbcomments.Valid {
		resp.CbComments = &m.Cbcomments.String
	}
	if m.CreatedBy.Valid {
		resp.CreatedBy = &m.CreatedBy.Int64
	}
	if m.ModBy.Valid {
		resp.ModBy = &m.ModBy.Int64
	}
	if m.ModDate.Valid {
		resp.ModDate = &m.ModDate.Time
	}
	if m.Deleted.Valid {
		resp.Deleted = &m.Deleted.Int16
	}
	if m.CreatedDate.Valid {
		resp.CreatedDate = &m.CreatedDate.Time
	}
	if m.Modded.Valid {
		resp.Modded = &m.Modded.Int16
	}
	if m.TestLocation.Valid {
		resp.TestLocation = &m.TestLocation.String
	}
	if m.Keywords.Valid {
		resp.Keywords = &m.Keywords.String
	}
	if m.TesterName.Valid {
		resp.TesterName = &m.TesterName.String
	}
	if m.ModderName.Valid {
		resp.ModderName = &m.ModderName.String
	}
	if m.Meastype.Valid {
		resp.MeasType = &m.Meastype.Int32
	}
	if m.SentToSprintai.Valid {
		resp.SentToSprintAI = &m.SentToSprintai.Time
	}

	return resp
}

func convertDirTest(d klabsqlc.Dirtest) KlabDirTestResponse {
	resp := KlabDirTestResponse{
		IdDirTest:     d.Iddirtest,
		IdMeasurement: d.Idmeasurement,
	}

	if d.Meascols.Valid {
		resp.MeasCols = &d.Meascols.String
	}
	if d.Weightkg.Valid {
		resp.WeightKg = &d.Weightkg.Float64
	}
	if d.Heightcm.Valid {
		resp.HeightCm = &d.Heightcm.Float64
	}
	if d.Bmi.Valid {
		resp.Bmi = &d.Bmi.Float64
	}
	if d.FatPr.Valid {
		resp.FatPr = &d.FatPr.Float64
	}
	if d.FatP1.Valid {
		resp.FatP1 = &d.FatP1.Float64
	}
	if d.FatP2.Valid {
		resp.FatP2 = &d.FatP2.Float64
	}
	if d.FatP3.Valid {
		resp.FatP3 = &d.FatP3.Float64
	}
	if d.FatP4.Valid {
		resp.FatP4 = &d.FatP4.Float64
	}
	if d.FatStyle.Valid {
		resp.FatStyle = &d.FatStyle.Int32
	}
	if d.FatEquip.Valid {
		resp.FatEquip = &d.FatEquip.String
	}
	if d.Fvc.Valid {
		resp.Fvc = &d.Fvc.Float64
	}
	if d.Fev1.Valid {
		resp.Fev1 = &d.Fev1.Float64
	}
	if d.AirPress.Valid {
		resp.AirPress = &d.AirPress.Float64
	}
	if d.AirTemp.Valid {
		resp.AirTemp = &d.AirTemp.Float64
	}
	if d.AirHumid.Valid {
		resp.AirHumid = &d.AirHumid.Float64
	}
	if d.Testprotocol.Valid {
		resp.TestProtocol = &d.Testprotocol.String
	}
	if d.AirPressUnit.Valid {
		resp.AirPressUnit = &d.AirPressUnit.Int32
	}
	if d.Settingslist.Valid {
		resp.SettingsList = &d.Settingslist.String
	}
	if d.Lt1X.Valid {
		resp.Lt1X = &d.Lt1X.Float64
	}
	if d.Lt1Y.Valid {
		resp.Lt1Y = &d.Lt1Y.Float64
	}
	if d.Lt2X.Valid {
		resp.Lt2X = &d.Lt2X.Float64
	}
	if d.Lt2Y.Valid {
		resp.Lt2Y = &d.Lt2Y.Float64
	}
	if d.Vt1X.Valid {
		resp.Vt1X = &d.Vt1X.Float64
	}
	if d.Vt2X.Valid {
		resp.Vt2X = &d.Vt2X.Float64
	}
	if d.Vt1Y.Valid {
		resp.Vt1Y = &d.Vt1Y.Float64
	}
	if d.Vt2Y.Valid {
		resp.Vt2Y = &d.Vt2Y.Float64
	}
	if d.Lt1CalcX.Valid {
		resp.Lt1CalcX = &d.Lt1CalcX.Float64
	}
	if d.Lt1CalcY.Valid {
		resp.Lt1CalcY = &d.Lt1CalcY.Float64
	}
	if d.Lt2CalcX.Valid {
		resp.Lt2CalcX = &d.Lt2CalcX.Float64
	}
	if d.Lt2CalcY.Valid {
		resp.Lt2CalcY = &d.Lt2CalcY.Float64
	}
	if d.Protocolmodel.Valid {
		resp.ProtocolModel = &d.Protocolmodel.Int16
	}
	if d.Testtype.Valid {
		resp.TestType = &d.Testtype.Int16
	}
	if d.Protocolxval.Valid {
		resp.ProtocolXVal = &d.Protocolxval.Int16
	}
	if d.Steptime.Valid {
		resp.StepTime = &d.Steptime.Int32
	}
	if d.WRest.Valid {
		resp.WRest = &d.WRest.Int16
	}
	if d.CreatedBy.Valid {
		resp.CreatedBy = &d.CreatedBy.Int64
	}
	if d.ModBy.Valid {
		resp.ModBy = &d.ModBy.Int64
	}
	if d.ModDate.Valid {
		resp.ModDate = &d.ModDate.Time
	}
	if d.Deleted.Valid {
		resp.Deleted = &d.Deleted.Int16
	}
	if d.CreatedDate.Valid {
		resp.CreatedDate = &d.CreatedDate.Time
	}
	if d.Modded.Valid {
		resp.Modded = &d.Modded.Int16
	}
	if d.Norawdata.Valid {
		resp.NoRawData = &d.Norawdata.Int16
	}

	return resp
}

func convertDirTestStep(d klabsqlc.Dirteststep) KlabDirTestStepResponse {
	resp := KlabDirTestStepResponse{
		IdDirTestSteps: d.Iddirteststeps,
		IdMeasurement:  d.Idmeasurement,
	}

	if d.Stepno.Valid {
		resp.StepNo = &d.Stepno.Int32
	}
	if d.AnaTime.Valid {
		resp.AnaTime = &d.AnaTime.Int32
	}
	if d.Timestop.Valid {
		resp.TimeStop = &d.Timestop.Float64
	}
	if d.Speed.Valid {
		resp.Speed = &d.Speed.Float64
	}
	if d.Pace.Valid {
		resp.Pace = &d.Pace.Float64
	}
	if d.Angle.Valid {
		resp.Angle = &d.Angle.Float64
	}
	if d.Elev.Valid {
		resp.Elev = &d.Elev.Float64
	}
	if d.Vo2calc.Valid {
		resp.Vo2Calc = &d.Vo2calc.Float64
	}
	if d.TTot.Valid {
		resp.TTot = &d.TTot.Float64
	}
	if d.TEx.Valid {
		resp.TEx = &d.TEx.Float64
	}
	if d.Fico2.Valid {
		resp.Fico2 = &d.Fico2.Float64
	}
	if d.Fio2.Valid {
		resp.Fio2 = &d.Fio2.Float64
	}
	if d.Feco2.Valid {
		resp.Feco2 = &d.Feco2.Float64
	}
	if d.Feo2.Valid {
		resp.Feo2 = &d.Feo2.Float64
	}
	if d.Vde.Valid {
		resp.Vde = &d.Vde.Float64
	}
	if d.Vco2.Valid {
		resp.Vco2 = &d.Vco2.Float64
	}
	if d.Vo2.Valid {
		resp.Vo2 = &d.Vo2.Float64
	}
	if d.Bf.Valid {
		resp.Bf = &d.Bf.Float64
	}
	if d.Ve.Valid {
		resp.Ve = &d.Ve.Float64
	}
	if d.Petco2.Valid {
		resp.Petco2 = &d.Petco2.Float64
	}
	if d.Peto2.Valid {
		resp.Peto2 = &d.Peto2.Float64
	}
	if d.Vo2kg.Valid {
		resp.Vo2Kg = &d.Vo2kg.Float64
	}
	if d.Re.Valid {
		resp.Re = &d.Re.Float64
	}
	if d.Hr.Valid {
		resp.Hr = &d.Hr.Float64
	}
	if d.La.Valid {
		resp.La = &d.La.Float64
	}
	if d.Rer.Valid {
		resp.Rer = &d.Rer.Float64
	}
	if d.VeStpd.Valid {
		resp.VeStpd = &d.VeStpd.Float64
	}
	if d.Veo2.Valid {
		resp.Veo2 = &d.Veo2.Float64
	}
	if d.Veco2.Valid {
		resp.Veco2 = &d.Veco2.Float64
	}
	if d.Tv.Valid {
		resp.Tv = &d.Tv.Float64
	}
	if d.EeAe.Valid {
		resp.EeAe = &d.EeAe.Float64
	}
	if d.LaVo2.Valid {
		resp.LaVo2 = &d.LaVo2.Float64
	}
	if d.O2pulse.Valid {
		resp.O2Pulse = &d.O2pulse.Float64
	}
	if d.VdeTv.Valid {
		resp.VdeTv = &d.VdeTv.Float64
	}
	if d.Va.Valid {
		resp.Va = &d.Va.Float64
	}
	if d.O2sa.Valid {
		resp.O2Sa = &d.O2sa.Float64
	}
	if d.Rpe.Valid {
		resp.Rpe = &d.Rpe.Float64
	}
	if d.BpSys.Valid {
		resp.BpSys = &d.BpSys.Float64
	}
	if d.BpDia.Valid {
		resp.BpDia = &d.BpDia.Float64
	}
	if d.Own1.Valid {
		resp.Own1 = &d.Own1.Float64
	}
	if d.Own2.Valid {
		resp.Own2 = &d.Own2.Float64
	}
	if d.Own3.Valid {
		resp.Own3 = &d.Own3.Float64
	}
	if d.Own4.Valid {
		resp.Own4 = &d.Own4.Float64
	}
	if d.Own5.Valid {
		resp.Own5 = &d.Own5.Float64
	}
	if d.StepIsRest.Valid {
		resp.StepIsRest = &d.StepIsRest.Int32
	}
	if d.StepIs30max.Valid {
		resp.StepIs30Max = &d.StepIs30max.Int32
	}
	if d.StepIs60max.Valid {
		resp.StepIs60Max = &d.StepIs60max.Int32
	}
	if d.StepIsRec.Valid {
		resp.StepIsRec = &d.StepIsRec.Int32
	}
	if d.CalcStart.Valid {
		resp.CalcStart = &d.CalcStart.Int32
	}
	if d.CalcEnd.Valid {
		resp.CalcEnd = &d.CalcEnd.Int32
	}
	if d.Comments.Valid {
		resp.Comments = &d.Comments.String
	}
	if d.Timestart.Valid {
		resp.TimeStart = &d.Timestart.Float64
	}
	if d.Duration.Valid {
		resp.Duration = &d.Duration.Float64
	}
	if d.Eco.Valid {
		resp.Eco = &d.Eco.Float64
	}
	if d.P.Valid {
		resp.P = &d.P.Float64
	}
	if d.Wkg.Valid {
		resp.Wkg = &d.Wkg.Float64
	}
	if d.Vo230s.Valid {
		resp.Vo230s = &d.Vo230s.Float64
	}
	if d.Vo2Pr.Valid {
		resp.Vo2Pr = &d.Vo2Pr.Float64
	}
	if d.StepIsLast.Valid {
		resp.StepIsLast = &d.StepIsLast.Int32
	}
	if d.Deleted.Valid {
		resp.Deleted = &d.Deleted.Int16
	}
	if d.CreatedBy.Valid {
		resp.CreatedBy = &d.CreatedBy.Int64
	}
	if d.ModBy.Valid {
		resp.ModBy = &d.ModBy.Int64
	}
	if d.ModDate.Valid {
		resp.ModDate = &d.ModDate.Time
	}
	if d.CreatedDate.Valid {
		resp.CreatedDate = &d.CreatedDate.Time
	}
	if d.Modded.Valid {
		resp.Modded = &d.Modded.Int16
	}
	if d.Own6.Valid {
		resp.Own6 = &d.Own6.Float64
	}
	if d.Own7.Valid {
		resp.Own7 = &d.Own7.Float64
	}
	if d.Own8.Valid {
		resp.Own8 = &d.Own8.Float64
	}
	if d.Own9.Valid {
		resp.Own9 = &d.Own9.Float64
	}
	if d.Own10.Valid {
		resp.Own10 = &d.Own10.Float64
	}
	if d.To2.Valid {
		resp.To2 = &d.To2.Float64
	}
	if d.Tco2.Valid {
		resp.Tco2 = &d.Tco2.Float64
	}

	return resp
}

func convertDirReport(d klabsqlc.Dirreport) KlabDirReportResponse {
	resp := KlabDirReportResponse{
		IdDirReport:   d.Iddirreport,
		IdMeasurement: d.Idmeasurement,
	}

	if d.PageInstructions.Valid {
		resp.PageInstructions = &d.PageInstructions.String
	}
	if d.TemplateRec.Valid {
		resp.TemplateRec = &d.TemplateRec.Int32
	}
	if d.LibrecName.Valid {
		resp.LibrecName = &d.LibrecName.String
	}
	if d.CreatedBy.Valid {
		resp.CreatedBy = &d.CreatedBy.Int64
	}
	if d.ModBy.Valid {
		resp.ModBy = &d.ModBy.Int64
	}
	if d.ModDate.Valid {
		resp.ModDate = &d.ModDate.Time
	}
	if d.Deleted.Valid {
		resp.Deleted = &d.Deleted.Int16
	}
	if d.CreatedDate.Valid {
		resp.CreatedDate = &d.CreatedDate.Time
	}
	if d.Modded.Valid {
		resp.Modded = &d.Modded.Int16
	}

	return resp
}

func convertDirRawData(d klabsqlc.Dirrawdatum) KlabDirRawDataResponse {
	resp := KlabDirRawDataResponse{
		IdDirRawData:  d.Iddirrawdata,
		IdMeasurement: d.Idmeasurement,
	}

	if d.Rawdata.Valid {
		resp.RawData = &d.Rawdata.String
	}
	if d.Columndata.Valid {
		resp.ColumnData = &d.Columndata.String
	}
	if d.Info.Valid {
		resp.Info = &d.Info.String
	}
	if d.Unitsdata.Valid {
		resp.UnitsData = &d.Unitsdata.String
	}
	if d.CreatedBy.Valid {
		resp.CreatedBy = &d.CreatedBy.Int64
	}
	if d.ModBy.Valid {
		resp.ModBy = &d.ModBy.Int64
	}
	if d.ModDate.Valid {
		resp.ModDate = &d.ModDate.Time
	}
	if d.Deleted.Valid {
		resp.Deleted = &d.Deleted.Int16
	}
	if d.CreatedDate.Valid {
		resp.CreatedDate = &d.CreatedDate.Time
	}
	if d.Modded.Valid {
		resp.Modded = &d.Modded.Int16
	}

	return resp
}

func convertDirResults(d klabsqlc.Dirresult) KlabDirResultsResponse {
	resp := KlabDirResultsResponse{
		IdDirResults:  d.Iddirresults,
		IdMeasurement: d.Idmeasurement,
	}

	if d.MaxVo2mlkgmin.Valid {
		resp.MaxVo2MlKgMin = &d.MaxVo2mlkgmin.Float64
	}
	if d.MaxVo2mlmin.Valid {
		resp.MaxVo2MlMin = &d.MaxVo2mlmin.Float64
	}
	if d.MaxVo2.Valid {
		resp.MaxVo2 = &d.MaxVo2.Float64
	}
	if d.MaxHr.Valid {
		resp.MaxHr = &d.MaxHr.Float64
	}
	if d.MaxSpeed.Valid {
		resp.MaxSpeed = &d.MaxSpeed.Float64
	}
	if d.MaxPace.Valid {
		resp.MaxPace = &d.MaxPace.Float64
	}
	if d.MaxP.Valid {
		resp.MaxP = &d.MaxP.Float64
	}
	if d.MaxPkg.Valid {
		resp.MaxPkg = &d.MaxPkg.Float64
	}
	if d.MaxAngle.Valid {
		resp.MaxAngle = &d.MaxAngle.Float64
	}
	if d.MaxLac.Valid {
		resp.MaxLac = &d.MaxLac.Float64
	}
	if d.MaxAdd1.Valid {
		resp.MaxAdd1 = &d.MaxAdd1.Float64
	}
	if d.MaxAdd2.Valid {
		resp.MaxAdd2 = &d.MaxAdd2.Float64
	}
	if d.MaxAdd3.Valid {
		resp.MaxAdd3 = &d.MaxAdd3.Float64
	}
	if d.LacAnkVo2mlkgmin.Valid {
		resp.LacAnkVo2MlKgMin = &d.LacAnkVo2mlkgmin.Float64
	}
	if d.LacAnkVo2mlmin.Valid {
		resp.LacAnkVo2MlMin = &d.LacAnkVo2mlmin.Float64
	}
	if d.LacAnkVo2.Valid {
		resp.LacAnkVo2 = &d.LacAnkVo2.Float64
	}
	if d.LacAnkVo2pr.Valid {
		resp.LacAnkVo2Pr = &d.LacAnkVo2pr.Float64
	}
	if d.LacAnkHr.Valid {
		resp.LacAnkHr = &d.LacAnkHr.Float64
	}
	if d.LacAnkSpeed.Valid {
		resp.LacAnkSpeed = &d.LacAnkSpeed.Float64
	}
	if d.LacAnkPace.Valid {
		resp.LacAnkPace = &d.LacAnkPace.Float64
	}
	if d.LacAnkP.Valid {
		resp.LacAnkP = &d.LacAnkP.Float64
	}
	if d.LacAnkPkg.Valid {
		resp.LacAnkPkg = &d.LacAnkPkg.Float64
	}
	if d.LacAnkAngle.Valid {
		resp.LacAnkAngle = &d.LacAnkAngle.Float64
	}
	if d.LacAnkLac.Valid {
		resp.LacAnkLac = &d.LacAnkLac.Float64
	}
	if d.LacAnkAdd1.Valid {
		resp.LacAnkAdd1 = &d.LacAnkAdd1.Float64
	}
	if d.LacAnkAdd2.Valid {
		resp.LacAnkAdd2 = &d.LacAnkAdd2.Float64
	}
	if d.LacAnkAdd3.Valid {
		resp.LacAnkAdd3 = &d.LacAnkAdd3.Float64
	}
	if d.LacAerkVo2mlkgmin.Valid {
		resp.LacAerkVo2MlKgMin = &d.LacAerkVo2mlkgmin.Float64
	}
	if d.LacAerkVo2mlmin.Valid {
		resp.LacAerkVo2MlMin = &d.LacAerkVo2mlmin.Float64
	}
	if d.LacAerkVo2.Valid {
		resp.LacAerkVo2 = &d.LacAerkVo2.Float64
	}
	if d.LacAerkVo2pr.Valid {
		resp.LacAerkVo2Pr = &d.LacAerkVo2pr.Float64
	}
	if d.LacAerkHr.Valid {
		resp.LacAerkHr = &d.LacAerkHr.Float64
	}
	if d.LacAerkSpeed.Valid {
		resp.LacAerkSpeed = &d.LacAerkSpeed.Float64
	}
	if d.LacAerkPace.Valid {
		resp.LacAerkPace = &d.LacAerkPace.Float64
	}
	if d.LacAerkP.Valid {
		resp.LacAerkP = &d.LacAerkP.Float64
	}
	if d.LacAerkPkg.Valid {
		resp.LacAerkPkg = &d.LacAerkPkg.Float64
	}
	if d.LacAerkAngle.Valid {
		resp.LacAerkAngle = &d.LacAerkAngle.Float64
	}
	if d.LacAerkLac.Valid {
		resp.LacAerkLac = &d.LacAerkLac.Float64
	}
	if d.LacAerkAdd1.Valid {
		resp.LacAerkAdd1 = &d.LacAerkAdd1.Float64
	}
	if d.LacAerkAdd2.Valid {
		resp.LacAerkAdd2 = &d.LacAerkAdd2.Float64
	}
	if d.LacAerkAdd3.Valid {
		resp.LacAerkAdd3 = &d.LacAerkAdd3.Float64
	}
	if d.VentAnkVo2mlkgmin.Valid {
		resp.VentAnkVo2MlKgMin = &d.VentAnkVo2mlkgmin.Float64
	}
	if d.VentAnkVo2mlmin.Valid {
		resp.VentAnkVo2MlMin = &d.VentAnkVo2mlmin.Float64
	}
	if d.VentAnkVo2.Valid {
		resp.VentAnkVo2 = &d.VentAnkVo2.Float64
	}
	if d.VentAnkVo2pr.Valid {
		resp.VentAnkVo2Pr = &d.VentAnkVo2pr.Float64
	}
	if d.VentAnkHr.Valid {
		resp.VentAnkHr = &d.VentAnkHr.Float64
	}
	if d.VentAnkSpeed.Valid {
		resp.VentAnkSpeed = &d.VentAnkSpeed.Float64
	}
	if d.VentAnkPace.Valid {
		resp.VentAnkPace = &d.VentAnkPace.Float64
	}
	if d.VentAnkP.Valid {
		resp.VentAnkP = &d.VentAnkP.Float64
	}
	if d.VentAnkPkg.Valid {
		resp.VentAnkPkg = &d.VentAnkPkg.Float64
	}
	if d.VentAnkAngle.Valid {
		resp.VentAnkAngle = &d.VentAnkAngle.Float64
	}
	if d.VentAnkLac.Valid {
		resp.VentAnkLac = &d.VentAnkLac.Float64
	}
	if d.VentAnkAdd1.Valid {
		resp.VentAnkAdd1 = &d.VentAnkAdd1.Float64
	}
	if d.VentAnkAdd2.Valid {
		resp.VentAnkAdd2 = &d.VentAnkAdd2.Float64
	}
	if d.VentAnkAdd3.Valid {
		resp.VentAnkAdd3 = &d.VentAnkAdd3.Float64
	}
	if d.VentAerkVo2mlkgmin.Valid {
		resp.VentAerkVo2MlKgMin = &d.VentAerkVo2mlkgmin.Float64
	}
	if d.VentAerkVo2mlmin.Valid {
		resp.VentAerkVo2MlMin = &d.VentAerkVo2mlmin.Float64
	}
	if d.VentAerkVo2.Valid {
		resp.VentAerkVo2 = &d.VentAerkVo2.Float64
	}
	if d.VentAerkVo2pr.Valid {
		resp.VentAerkVo2Pr = &d.VentAerkVo2pr.Float64
	}
	if d.VentAerkHr.Valid {
		resp.VentAerkHr = &d.VentAerkHr.Float64
	}
	if d.VentAerkSpeed.Valid {
		resp.VentAerkSpeed = &d.VentAerkSpeed.Float64
	}
	if d.VentAerkPace.Valid {
		resp.VentAerkPace = &d.VentAerkPace.Float64
	}
	if d.VentAerkP.Valid {
		resp.VentAerkP = &d.VentAerkP.Float64
	}
	if d.VentAerkPkg.Valid {
		resp.VentAerkPkg = &d.VentAerkPkg.Float64
	}
	if d.VentAerkAngle.Valid {
		resp.VentAerkAngle = &d.VentAerkAngle.Float64
	}
	if d.VentAerkLac.Valid {
		resp.VentAerkLac = &d.VentAerkLac.Float64
	}
	if d.VentAerkAdd1.Valid {
		resp.VentAerkAdd1 = &d.VentAerkAdd1.Float64
	}
	if d.VentAerkAdd2.Valid {
		resp.VentAerkAdd2 = &d.VentAerkAdd2.Float64
	}
	if d.VentAerkAdd3.Valid {
		resp.VentAerkAdd3 = &d.VentAerkAdd3.Float64
	}
	if d.CreatedBy.Valid {
		resp.CreatedBy = &d.CreatedBy.Int64
	}
	if d.ModBy.Valid {
		resp.ModBy = &d.ModBy.Int64
	}
	if d.ModDate.Valid {
		resp.ModDate = &d.ModDate.Time
	}
	if d.Deleted.Valid {
		resp.Deleted = &d.Deleted.Int16
	}
	if d.CreatedDate.Valid {
		resp.CreatedDate = &d.CreatedDate.Time
	}
	if d.Modded.Valid {
		resp.Modded = &d.Modded.Int16
	}

	return resp
}
