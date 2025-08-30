-- name: GetAllSporttiIDs :many
SELECT sportti_id FROM sportti_id_list;


-- Prefer updating customer metadata if it already exists.
-- name: UpsertCustomer :exec
INSERT INTO customer (
    idcustomer, firstname, lastname, idgroups, dob, sex, dob_year, dob_month, dob_day,
    pid_number, company, occupation, education, address, phone_home, phone_work, phone_mobile,
    faxno, email, username, password, readonly, warnings, allow_to_save, allow_to_cloud, flag2,
    idsport, medication, addinfo, team_name, add1, athlete, add10, add20, updatemode, weight_kg,
    height_cm, date_modified, recom_testlevel, created_by, mod_by, mod_date, deleted,
    created_date, modded, allow_anonymous_data, locked, allow_to_sprintai, tosprintai_from,
    stat_sent, sportti_id
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9,
    $10, $11, $12, $13, $14, $15, $16, $17,
    $18, $19, $20, $21, $22, $23, $24, $25, $26,
    $27, $28, $29, $30, $31, $32, $33, $34, $35, $36,
    $37, $38, $39, $40, $41, $42, $43, $44,
    $45, $46, $47, $48, $49, $50, $51
)
ON CONFLICT (idcustomer) DO UPDATE SET
    firstname = EXCLUDED.firstname,
    lastname = EXCLUDED.lastname,
    idgroups = EXCLUDED.idgroups,
    dob = EXCLUDED.dob,
    sex = EXCLUDED.sex,
    dob_year = EXCLUDED.dob_year,
    dob_month = EXCLUDED.dob_month,
    dob_day = EXCLUDED.dob_day,
    pid_number = EXCLUDED.pid_number,
    company = EXCLUDED.company,
    occupation = EXCLUDED.occupation,
    education = EXCLUDED.education,
    address = EXCLUDED.address,
    phone_home = EXCLUDED.phone_home,
    phone_work = EXCLUDED.phone_work,
    phone_mobile = EXCLUDED.phone_mobile,
    faxno = EXCLUDED.faxno,
    email = EXCLUDED.email,
    username = EXCLUDED.username,
    password = EXCLUDED.password,
    readonly = EXCLUDED.readonly,
    warnings = EXCLUDED.warnings,
    allow_to_save = EXCLUDED.allow_to_save,
    allow_to_cloud = EXCLUDED.allow_to_cloud,
    flag2 = EXCLUDED.flag2,
    idsport = EXCLUDED.idsport,
    medication = EXCLUDED.medication,
    addinfo = EXCLUDED.addinfo,
    team_name = EXCLUDED.team_name,
    add1 = EXCLUDED.add1,
    athlete = EXCLUDED.athlete,
    add10 = EXCLUDED.add10,
    add20 = EXCLUDED.add20,
    updatemode = EXCLUDED.updatemode,
    weight_kg = EXCLUDED.weight_kg,
    height_cm = EXCLUDED.height_cm,
    date_modified = EXCLUDED.date_modified,
    recom_testlevel = EXCLUDED.recom_testlevel,
    created_by = EXCLUDED.created_by,
    mod_by = EXCLUDED.mod_by,
    mod_date = EXCLUDED.mod_date,
    deleted = EXCLUDED.deleted,
    created_date = EXCLUDED.created_date,
    modded = EXCLUDED.modded,
    allow_anonymous_data = EXCLUDED.allow_anonymous_data,
    locked = EXCLUDED.locked,
    allow_to_sprintai = EXCLUDED.allow_to_sprintai,
    tosprintai_from = EXCLUDED.tosprintai_from,
    stat_sent = EXCLUDED.stat_sent,
    sportti_id = EXCLUDED.sportti_id;


-- name: InsertMeasurement :exec
INSERT INTO measurement_list (
    idmeasurement, measname, idcustomer, tablename, idpatterndef,
    do_year, do_month, do_day, do_hour, do_min, sessionno, info,
    measurements, groupnotes, cbcharts, cbcomments,
    created_by, mod_by, mod_date, deleted, created_date, modded,
    test_location, keywords, tester_name, modder_name, meastype,
    sent_to_sprintai
) VALUES (
    $1, $2, $3, $4, $5,
    $6, $7, $8, $9, $10, $11, $12,
    $13, $14, $15, $16,
    $17, $18, $19, $20, $21, $22,
    $23, $24, $25, $26, $27,
    $28
)
ON CONFLICT (idmeasurement) DO NOTHING;


-- name: InsertDirTest :exec
INSERT INTO dirtest (
    iddirtest, idmeasurement, meascols, weightkg, heightcm, bmi,
    fat_pr, fat_p1, fat_p2, fat_p3, fat_p4, fat_style, fat_equip,
    fvc, fev1, air_press, air_temp, air_humid, testprotocol, air_press_unit,
    settingslist, lt1_x, lt1_y, lt2_x, lt2_y, vt1_x, vt2_x, vt1_y, vt2_y,
    lt1_calc_x, lt1_calc_y, lt2_calc_x, lt2_calc_y, protocolmodel, testtype,
    protocolxval, steptime, w_rest, created_by, mod_by, mod_date, deleted,
    created_date, modded, norawdata
) VALUES (
    $1, $2, $3, $4, $5, $6,
    $7, $8, $9, $10, $11, $12, $13,
    $14, $15, $16, $17, $18, $19, $20,
    $21, $22, $23, $24, $25, $26, $27, $28, $29,
    $30, $31, $32, $33, $34, $35,
    $36, $37, $38, $39, $40, $41, $42,
    $43, $44, $45
)
ON CONFLICT (iddirtest) DO NOTHING;


-- name: InsertDirTestStep :exec
INSERT INTO dirteststeps (
    iddirteststeps, idmeasurement, stepno, ana_time, timestop, speed, pace,
    angle, elev, vo2calc, t_tot, t_ex, fico2, fio2, feco2, feo2, vde, vco2,
    vo2, bf, ve, petco2, peto2, vo2kg, re, hr, la, rer, ve_stpd, veo2,
    veco2, tv, ee_ae, la_vo2, o2pulse, vde_tv, va, o2sa, rpe,
    bp_sys, bp_dia, own1, own2, own3, own4, own5,
    step_is_rest, step_is_30max, step_is_60max, step_is_rec, calc_start,
    calc_end, comments, timestart, duration, eco, p, wkg,
    vo2_30s, vo2_pr, step_is_last, deleted, created_by,
    mod_by, mod_date, created_date, modded, own6, own7, own8, own9, own10,
    to2, tco2
) VALUES (
    $1, $2, $3, $4, $5, $6, $7,
    $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18,
    $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29,
    $30, $31, $32, $33, $34, $35, $36, $37, $38, $39,
    $40, $41, $42, $43, $44, $45, $46,
    $47, $48, $49, $50, $51,
    $52, $53, $54, $55, $56, $57, $58,
    $59, $60, $61, $62, $63,
    $64, $65, $66, $67, $68, $69, $70, $71, $72,
    $73, $74
)
ON CONFLICT (iddirteststeps) DO NOTHING;


-- name: InsertDirRawData :exec
INSERT INTO dirrawdata (
    iddirrawdata, idmeasurement, rawdata, columndata, info, unitsdata,
    created_by, mod_by, mod_date, deleted, created_date, modded
) VALUES (
    $1, $2, $3, $4, $5, $6,
    $7, $8, $9, $10, $11, $12
)
ON CONFLICT (iddirrawdata) DO NOTHING;


-- name: InsertDirReport :exec
INSERT INTO dirreport (
    iddirreport, page_instructions, idmeasurement, template_rec, librec_name,
    created_by, mod_by, mod_date, deleted, created_date, modded
) VALUES (
    $1, $2, $3, $4, $5,
    $6, $7, $8, $9, $10, $11
)
ON CONFLICT (iddirreport) DO NOTHING;


-- name: InsertDirResults :exec
INSERT INTO dirresults (
    iddirresults, idmeasurement, max_vo2mlkgmin, max_vo2mlmin, max_vo2,
    max_hr, max_speed, max_pace, max_p, max_pkg, max_angle, max_lac,
    max_add1, max_add2, max_add3,
    lac_ank_vo2mlkgmin, lac_ank_vo2mlmin, lac_ank_vo2, lac_ank_vo2pr,
    lac_ank_hr, lac_ank_speed, lac_ank_pace, lac_ank_p, lac_ank_pkg,
    lac_ank_angle, lac_ank_lac, lac_ank_add1, lac_ank_add2, lac_ank_add3,
    lac_aerk_vo2mlkgmin, lac_aerk_vo2mlmin, lac_aerk_vo2, lac_aerk_vo2pr,
    lac_aerk_hr, lac_aerk_speed, lac_aerk_pace, lac_aerk_p, lac_aerk_pkg,
    lac_aerk_angle, lac_aerk_lac, lac_aerk_add1, lac_aerk_add2, lac_aerk_add3,
    vent_ank_vo2mlkgmin, vent_ank_vo2mlmin, vent_ank_vo2, vent_ank_vo2pr,
    vent_ank_hr, vent_ank_speed, vent_ank_pace, vent_ank_p, vent_ank_pkg,
    vent_ank_angle, vent_ank_lac, vent_ank_add1, vent_ank_add2, vent_ank_add3,
    vent_aerk_vo2mlkgmin, vent_aerk_vo2mlmin, vent_aerk_vo2, vent_aerk_vo2pr,
    vent_aerk_hr, vent_aerk_speed, vent_aerk_pace, vent_aerk_p, vent_aerk_pkg,
    vent_aerk_angle, vent_aerk_lac, vent_aerk_add1, vent_aerk_add2, vent_aerk_add3,
    created_by, mod_by, mod_date, deleted, created_date, modded
) VALUES (
    $1, $2, $3, $4, $5,
    $6, $7, $8, $9, $10, $11, $12,
    $13, $14, $15,
    $16, $17, $18, $19,
    $20, $21, $22, $23, $24,
    $25, $26, $27, $28, $29,
    $30, $31, $32, $33,
    $34, $35, $36, $37, $38,
    $39, $40, $41, $42, $43,
    $44, $45, $46, $47,
    $48, $49, $50, $51, $52,
    $53, $54, $55, $56, $57,
    $58, $59, $60,
    $61, $62, $63, $64, $65, $66,
    $67, $68, $69, $70, $71, $72,
    $73, $74, $75,
    $76, $77
)
ON CONFLICT (iddirresults) DO NOTHING;

-- name: GetCustomerByID :one
SELECT * FROM customer
WHERE idcustomer = $1;

-- name: GetMeasurementsByCustomer :many
SELECT *
FROM measurement_list
WHERE idcustomer = $1
ORDER BY idmeasurement;

-- name: GetDirTestsByMeasurementIDs :many
SELECT *
FROM dirtest
WHERE idmeasurement = ANY($1::int[])
ORDER BY idmeasurement;


-- name: GetDirTestStepsByMeasurementIDs :many
SELECT *
FROM dirteststeps
WHERE idmeasurement = ANY($1::int[])
ORDER BY idmeasurement, stepno;

-- name: GetDirReportsByMeasurementIDs :many
SELECT *
FROM dirreport
WHERE idmeasurement = ANY($1::int[])
ORDER BY idmeasurement;


-- name: GetDirRawDataByMeasurementIDs :many
SELECT *
FROM dirrawdata
WHERE idmeasurement = ANY($1::int[])
ORDER BY idmeasurement;


-- name: GetDirResultsByMeasurementIDs :many
SELECT *
FROM dirresults
WHERE idmeasurement = ANY($1::int[])
ORDER BY idmeasurement;
