-- name: GetAthletesBySector :many
SELECT Firstname, Lastname, Fiscode
FROM A_competitor
WHERE SectorCode = $1
ORDER BY Fiscode;

-- name: GetNationsBySector :many
SELECT DISTINCT NationCode
FROM A_competitor
WHERE SectorCode = $1
ORDER BY NationCode ASC;

-- name: GetSkiJumpingSeasons :many
SELECT DISTINCT SeasonCode
FROM A_raceJP
ORDER BY SeasonCode DESC;

-- name: GetNordicCombinedSeasons :many
SELECT DISTINCT SeasonCode
FROM A_raceNK
ORDER BY SeasonCode DESC;

-- name: GetCrossCountrySeasons :many
SELECT DISTINCT SeasonCode
FROM A_raceCC
ORDER BY SeasonCode DESC;

-- name: GetSkiJumpingDisciplines :many
SELECT DISTINCT DisciplineCode
FROM A_raceJP
ORDER BY DisciplineCode ASC;

-- name: GetNordicCombinedDisciplines :many
SELECT DISTINCT DisciplineCode
FROM A_raceNK
ORDER BY DisciplineCode ASC;

-- name: GetCrossCountryDisciplines :many
SELECT DISTINCT DisciplineCode
FROM A_raceCC
ORDER BY DisciplineCode ASC;

-- name: GetSkiJumpingCategories :many
SELECT DISTINCT CatCode
FROM A_raceJP
ORDER BY CatCode ASC;

-- name: GetNordicCombinedCategories :many
SELECT DISTINCT CatCode
FROM A_raceNK
ORDER BY CatCode ASC;

-- name: GetCrossCountryCategories :many
SELECT DISTINCT CatCode
FROM A_raceCC
ORDER BY CatCode ASC;

-- name: GetCompetitorIDByFiscodeNK :one
SELECT CompetitorID
FROM A_competitor
WHERE Fiscode = $1 AND SectorCode = 'NK';

-- name: GetCompetitorIDByFiscodeJP :one
SELECT CompetitorID
FROM A_competitor
WHERE Fiscode = $1 AND SectorCode = 'JP';

-- name: GetCompetitorIDByFiscodeCC :one
SELECT CompetitorID
FROM A_competitor
WHERE Fiscode = $1 AND SectorCode = 'CC';


-- name: GetAthleteResultsNK :many
SELECT 
    rNK.RecID,
    rNK.RaceID,
    rNK.Position,
    aNK.RaceDate,
    aNK.SeasonCode,
    aNK.Distance,
    aNK.Hill,
    aNK.DisciplineCode,
    aNK.CatCode,
    aNK.Place,
    rNK.PosR1,
    rNK.SpeedR1,
    rNK.DistR1,
    rNK.JudPtsR1,
    rNK.WindR1,
    rNK.WindPtsR1,
    rNK.GateR1,
    rNK.TotRun1,
    rNK.PosCC,
    rNK.TimeTot,
    rNK.TimeTotInt,
    rNK.PointsJump
FROM A_resultNK rNK
JOIN A_raceNK aNK ON rNK.RaceID = aNK.RaceID
WHERE rNK.CompetitorID = $1
  AND ($2::int[]    IS NULL OR aNK.SeasonCode     = ANY($2))
  AND ($3::text[]   IS NULL OR aNK.DisciplineCode = ANY($3))
  AND ($4::text[]   IS NULL OR aNK.CatCode        = ANY($4))
ORDER BY aNK.RaceDate;

-- name: GetAthleteResultsJP :many
SELECT 
    rJP.RaceID,
    rJP.Position,
    aJP.RaceDate,
    aJP.SeasonCode,
    aJP.DisciplineCode,
    aJP.CatCode,
    aJP.Place,
    rJP.PosR1,
    rJP.SpeedR1,
    rJP.DistR1,
    rJP.JudPtsR1,
    rJP.WindR1,
    rJP.WindPtsR1,
    rJP.GateR1,
    rJP.PosR2,
    rJP.SpeedR2,
    rJP.DistR2,
    rJP.JudPtsR2,
    rJP.WindR2,
    rJP.WindPtsR2,
    rJP.GateR2,
    rJP.TotRun1,
    rJP.TotRun2
FROM A_resultJP rJP
JOIN A_raceJP aJP ON rJP.RaceID = aJP.RaceID
WHERE rJP.CompetitorID = $1
  AND ($2::int[]    IS NULL OR aJP.SeasonCode     = ANY($2))
  AND ($3::text[]   IS NULL OR aJP.DisciplineCode = ANY($3))
  AND ($4::text[]   IS NULL OR aJP.CatCode        = ANY($4))
ORDER BY aJP.RaceDate;

-- name: GetAthleteResultsCC :many
SELECT
    rCC.RecID,
    rCC.RaceID,
    rCC.Position,
    rCC.TimeTot,
    rCC.CompetitorID,
    aCC.RaceDate,
    aCC.SeasonCode,
    aCC.DisciplineCode,
    aCC.CatCode,
    aCC.Place
FROM A_resultCC rCC
JOIN A_raceCC aCC ON rCC.RaceID = aCC.RaceID
WHERE rCC.CompetitorID = $1
  AND ($2::int[]  IS NULL OR aCC.SeasonCode     = ANY($2))
  AND ($3::text[] IS NULL OR aCC.DisciplineCode = ANY($3))
  AND ($4::text[] IS NULL OR aCC.CatCode        = ANY($4))
ORDER BY aCC.RaceDate;

-- name: GetRacesNK :many
SELECT *
FROM A_raceNK
WHERE ($1::int[]  IS NULL OR SeasonCode     = ANY($1))
  AND ($2::text[] IS NULL OR DisciplineCode = ANY($2))
  AND ($3::text[] IS NULL OR CatCode        = ANY($3))
ORDER BY RaceID;

-- name: GetRacesJP :many
SELECT *
FROM A_raceJP
WHERE ($1::int[]  IS NULL OR SeasonCode     = ANY($1))
  AND ($2::text[] IS NULL OR DisciplineCode = ANY($2))
  AND ($3::text[] IS NULL OR CatCode        = ANY($3))
ORDER BY RaceID;

-- name: GetRacesCC :many
SELECT *
FROM A_raceCC
WHERE ($1::int[]  IS NULL OR SeasonCode     = ANY($1))
  AND ($2::text[] IS NULL OR DisciplineCode = ANY($2))
  AND ($3::text[] IS NULL OR CatCode        = ANY($3))
ORDER BY RaceID;


-- name: GetRaceResultsNKByRaceID :many
SELECT *
FROM public.a_resultnk
WHERE raceid = $1
ORDER BY recid;


-- name: GetRaceResultsJPByRaceID :many
SELECT *
FROM public.a_resultjp
WHERE raceid = $1
ORDER BY recid;


-- name: GetRaceResultsCCByRaceID :many
SELECT *
FROM public.a_resultcc
WHERE raceid = $1
ORDER BY recid;


-- name: GetLastRowCompetitor :one
SELECT *
FROM a_competitor
ORDER BY competitorid DESC
LIMIT 1;

-- name: GetLastRowRaceCC :one
SELECT *
FROM a_racecc
ORDER BY raceid DESC
LIMIT 1;

-- name: GetLastRowRaceJP :one
SELECT *
FROM a_racejp
ORDER BY raceid DESC
LIMIT 1;

-- name: GetLastRowRaceNK :one
SELECT *
FROM a_racenk
ORDER BY raceid DESC
LIMIT 1;

-- name: GetLastRowResultCC :one
SELECT *
FROM a_resultcc
ORDER BY recid DESC
LIMIT 1;

-- name: GetLastRowResultJP :one
SELECT *
FROM a_resultjp
ORDER BY recid DESC
LIMIT 1;

-- name: GetLastRowResultNK :one
SELECT *
FROM a_resultnk
ORDER BY recid DESC
LIMIT 1;


-- name: InsertCompetitor :exec
INSERT INTO public.a_competitor (
  competitorid, personid, ipcid, fiscode, birthdate, status_date,
  fee, dateofcreation, injury, version, compidmssql, carving, photo,
  notallowed, published, team, photo_big, lastupdate,
  statusnextlist, alternatenamecheck, deletedat, doped, createdby,
  categorycode, classname, data, lastupdateby, disciplines,
  type, sectorcode, classcode, lastname, firstname, gender, natteam,
  nationcode, nationalcode, skiclub, association, status,
  status_old, status_by, tragroup
) VALUES (
  $1,  $2,  $3,  $4,  $5,  $6,
  $7,  $8,  $9,  $10, $11, $12, $13,
  $14, $15, $16, $17, $18,
  $19, $20, $21, $22, $23,
  $24, $25, $26, $27, $28,
  $29, $30, $31, $32, $33, $34, $35,
  $36, $37, $38, $39, $40,
  $41, $42, $43
);


-- name: UpdateCompetitorByID :one
UPDATE public.a_competitor SET
  personid            = $2,
  ipcid               = $3,
  type                = $4,
  sectorcode          = $5,
  fiscode             = $6,
  lastname            = $7,
  firstname           = $8,
  gender              = $9,
  birthdate           = $10,
  nationcode          = $11,
  nationalcode        = $12,
  skiclub             = $13,
  association         = $14,
  status              = $15,
  status_old          = $16,
  status_by           = $17,
  status_date         = $18,
  statusnextlist      = $19,
  alternatenamecheck  = $20,
  fee                 = $21,
  dateofcreation      = $22,
  createdby           = $23,
  injury              = $24,
  version             = $25,
  compidmssql         = $26,
  carving             = $27,
  photo               = $28,
  notallowed          = $29,
  natteam             = $30,
  tragroup            = $31,
  published           = $32,
  doped               = $33,
  team                = $34,
  photo_big           = $35,
  data                = $36,
  lastupdateby        = $37,
  disciplines         = $38,
  lastupdate          = $39,
  deletedat           = $40,
  categorycode        = $41,
  classname           = $42,
  classcode           = $43
WHERE competitorid = $1
RETURNING competitorid;

-- name: DeleteCompetitorByID :one
DELETE FROM a_competitor
WHERE competitorid = $1
RETURNING competitorid;



-- name: InsertRaceCC :exec
INSERT INTO public.a_racecc (
  raceid, eventid, seasoncode, racecodex,
  disciplineid, disciplinecode, catcode, catcode2, catcode3, catcode4, gender,
  racedate, starteventdate, description, place, nationcode, receiveddate, validdate,
  td1id, td1name, td1nation, td1code,
  td2id, td2name, td2nation, td2code,
  calstatuscode, procstatuscode, displaystatus, fisinterncomment, webcomment,
  pursuit, masse, relay, distance, hill, style, qualif, finale, homol,
  published, validforfispoints, usedfislist, tolist, discforlistcode,
  calculatedpenalty, appliedpenalty, appliedscala, penscafixed,
  version, nationraceid, provraceid, msql7evid, mssql7id,
  topbanner, bottombanner, toplogo, bottomlogo, gallery,
  indi, team, tabcount, columncount, level,
  hloc1, hloc2, hloc3, hcet1, hcet2, hcet3,
  live, livestatus1, livestatus2, livestatus3, liveinfo1, liveinfo2, liveinfo3,
  passwd, timinglogo,
  results, pdf, noepr, tddoc, timingreport, special_cup_points, skip_wcsl,
  validforowg,
  lastupdate
) VALUES (
  $1,  $2,  $3,  $4,
  $5,  $6,  $7,  $8,  $9,  $10, $11,
  $12, $13, $14, $15, $16, $17, $18,
  $19, $20, $21, $22,
  $23, $24, $25, $26,
  $27, $28, $29, $30, $31,
  $32, $33, $34, $35, $36, $37, $38, $39, $40,
  $41, $42, $43, $44, $45,
  $46, $47, $48, $49,
  $50, $51, $52, $53, $54,
  $55, $56, $57, $58, $59,
  $60, $61, $62, $63, $64,
  $65, $66, $67, $68, $69, $70,
  $71, $72, $73, $74, $75, $76, $77,
  $78, $79,
  $80, $81, $82, $83, $84, $85, $86,
  $87,
  $88
);


-- name: UpdateRaceCCByID :one
UPDATE public.a_racecc SET
  eventid = $2,
  seasoncode = $3,
  racecodex = $4,
  disciplineid = $5,
  disciplinecode = $6,
  catcode = $7,
  catcode2 = $8,
  catcode3 = $9,
  catcode4 = $10,
  gender = $11,
  racedate = $12,
  starteventdate = $13,
  description = $14,
  place = $15,
  nationcode = $16,
  td1id = $17,
  td1name = $18,
  td1nation = $19,
  td1code = $20,
  td2id = $21,
  td2name = $22,
  td2nation = $23,
  td2code = $24,
  calstatuscode = $25,
  procstatuscode = $26,
  receiveddate = $27,
  pursuit = $28,
  masse = $29,
  relay = $30,
  distance = $31,
  hill = $32,
  style = $33,
  qualif = $34,
  finale = $35,
  homol = $36,
  webcomment = $37,
  displaystatus = $38,
  fisinterncomment = $39,
  published = $40,
  validforfispoints = $41,
  usedfislist = $42,
  tolist = $43,
  discforlistcode = $44,
  calculatedpenalty = $45,
  appliedpenalty = $46,
  appliedscala = $47,
  penscafixed = $48,
  version = $49,
  nationraceid = $50,
  provraceid = $51,
  msql7evid = $52,
  mssql7id = $53,
  results = $54,
  pdf = $55,
  topbanner = $56,
  bottombanner = $57,
  toplogo = $58,
  bottomlogo = $59,
  gallery = $60,
  indi = $61,
  team = $62,
  tabcount = $63,
  columncount = $64,
  level = $65,
  hloc1 = $66,
  hloc2 = $67,
  hloc3 = $68,
  hcet1 = $69,
  hcet2 = $70,
  hcet3 = $71,
  live = $72,
  livestatus1 = $73,
  livestatus2 = $74,
  livestatus3 = $75,
  liveinfo1 = $76,
  liveinfo2 = $77,
  liveinfo3 = $78,
  passwd = $79,
  timinglogo = $80,
  validdate = $81,
  noepr = $82,
  tddoc = $83,
  timingreport = $84,
  special_cup_points = $85,
  skip_wcsl = $86,
  validforowg = $87,
  lastupdate = $88
WHERE raceid = $1
RETURNING raceid;


-- name: DeleteRaceCCByID :one
DELETE FROM a_racecc
WHERE raceid = $1
RETURNING raceid;


-- name: InsertRaceJP :exec
INSERT INTO public.a_racejp (
  raceid, eventid, seasoncode, racecodex,
  disciplineid, disciplinecode, catcode, catcode2, catcode3, catcode4, gender,
  racedate, starteventdate, description, place, nationcode,
  td1id, td1name, td1nation, td1code,
  td2id, td2name, td2nation, td2code,
  calstatuscode, procstatuscode, receiveddate,
  pursuit, masse, relay, distance, hill, style, qualif, finale, homol,
  webcomment, displaystatus, fisinterncomment,
  published, validforfispoints, usedfislist, tolist, discforlistcode,
  calculatedpenalty, appliedpenalty, appliedscala, penscafixed,
  version, nationraceid, provraceid, msql7evid, mssql7id,
  results, pdf,
  topbanner, bottombanner, toplogo, bottomlogo, gallery,
  indi, team, tabcount, columncount, level,
  hloc1, hloc2, hloc3, hcet1, hcet2, hcet3,
  live, livestatus1, livestatus2, livestatus3,
  liveinfo1, liveinfo2, liveinfo3,
  passwd, timinglogo,
  validdate, noepr, tddoc, timingreport, special_cup_points, skip_wcsl,
  lastupdate, validforowg
) VALUES (
  $1,$2,$3,$4,
  $5,$6,$7,$8,$9,$10,$11,
  $12,$13,$14,$15,$16,
  $17,$18,$19,$20,
  $21,$22,$23,$24,
  $25,$26,$27,
  $28,$29,$30,$31,$32,$33,$34,$35,$36,
  $37,$38,$39,
  $40,$41,$42,$43,$44,
  $45,$46,$47,$48,
  $49,$50,$51,$52,$53,
  $54,$55,
  $56,$57,$58,$59,$60,
  $61,$62,$63,$64,$65,
  $66,$67,$68,$69,$70,$71,
  $72,$73,$74,$75,
  $76,$77,$78,
  $79,$80,
  $81,$82,$83,$84,$85,$86,
  $87,$88
);


-- name: UpdateRaceJPByID :one
UPDATE public.a_racejp SET
  eventid = $2,
  seasoncode = $3,
  racecodex = $4,
  disciplineid = $5,
  disciplinecode = $6,
  catcode = $7,
  catcode2 = $8,
  catcode3 = $9,
  catcode4 = $10,
  gender = $11,
  racedate = $12,
  starteventdate = $13,
  description = $14,
  place = $15,
  nationcode = $16,
  td1id = $17,
  td1name = $18,
  td1nation = $19,
  td1code = $20,
  td2id = $21,
  td2name = $22,
  td2nation = $23,
  td2code = $24,
  calstatuscode = $25,
  procstatuscode = $26,
  receiveddate = $27,
  pursuit = $28,
  masse = $29,
  relay = $30,
  distance = $31,
  hill = $32,
  style = $33,
  qualif = $34,
  finale = $35,
  homol = $36,
  webcomment = $37,
  displaystatus = $38,
  fisinterncomment = $39,
  published = $40,
  validforfispoints = $41,
  usedfislist = $42,
  tolist = $43,
  discforlistcode = $44,
  calculatedpenalty = $45,
  appliedpenalty = $46,
  appliedscala = $47,
  penscafixed = $48,
  version = $49,
  nationraceid = $50,
  provraceid = $51,
  msql7evid = $52,
  mssql7id = $53,
  results = $54,
  pdf = $55,
  topbanner = $56,
  bottombanner = $57,
  toplogo = $58,
  bottomlogo = $59,
  gallery = $60,
  indi = $61,
  team = $62,
  tabcount = $63,
  columncount = $64,
  level = $65,
  hloc1 = $66,
  hloc2 = $67,
  hloc3 = $68,
  hcet1 = $69,
  hcet2 = $70,
  hcet3 = $71,
  live = $72,
  livestatus1 = $73,
  livestatus2 = $74,
  livestatus3 = $75,
  liveinfo1 = $76,
  liveinfo2 = $77,
  liveinfo3 = $78,
  passwd = $79,
  timinglogo = $80,
  validdate = $81,
  noepr = $82,
  tddoc = $83,
  timingreport = $84,
  special_cup_points = $85,
  skip_wcsl = $86,
  lastupdate = $87,
  validforowg = $88
WHERE raceid = $1
RETURNING raceid;


-- name: DeleteRaceJPByID :one
DELETE FROM a_racejp
WHERE raceid = $1
RETURNING raceid;


-- name: InsertRaceNK :exec
INSERT INTO public.a_racenk (
  raceid, eventid, seasoncode, racecodex,
  disciplineid, disciplinecode, catcode, catcode2, catcode3, catcode4, gender,
  racedate, starteventdate, description, place, nationcode,
  td1id, td1name, td1nation, td1code,
  td2id, td2name, td2nation, td2code,
  calstatuscode, procstatuscode, receiveddate,
  pursuit, masse, relay, distance, hill, style, qualif, finale, homol,
  webcomment, displaystatus, fisinterncomment,
  published, validforfispoints, usedfislist, tolist, discforlistcode,
  calculatedpenalty, appliedpenalty, appliedscala, penscafixed,
  version, nationraceid, provraceid, msql7evid, mssql7id,
  results, pdf,
  topbanner, bottombanner, toplogo, bottomlogo, gallery,
  indi, team, tabcount, columncount, level,
  hloc1, hloc2, hloc3, hcet1, hcet2, hcet3,
  live, livestatus1, livestatus2, livestatus3,
  liveinfo1, liveinfo2, liveinfo3,
  passwd, timinglogo,
  validdate, noepr, tddoc, timingreport, special_cup_points, skip_wcsl,
  validforowg, lastupdate
) VALUES (
  $1,$2,$3,$4,
  $5,$6,$7,$8,$9,$10,$11,
  $12,$13,$14,$15,$16,
  $17,$18,$19,$20,
  $21,$22,$23,$24,
  $25,$26,$27,
  $28,$29,$30,$31,$32,$33,$34,$35,$36,
  $37,$38,$39,
  $40,$41,$42,$43,$44,
  $45,$46,$47,$48,
  $49,$50,$51,$52,$53,
  $54,$55,
  $56,$57,$58,$59,$60,
  $61,$62,$63,$64,$65,
  $66,$67,$68,$69,$70,$71,
  $72,$73,$74,$75,
  $76,$77,$78,
  $79,$80,
  $81,$82,$83,$84,$85,$86,
  $87,$88
);


-- name: UpdateRaceNKByID :one
UPDATE public.a_racenk SET
  eventid = $2,
  seasoncode = $3,
  racecodex = $4,
  disciplineid = $5,
  disciplinecode = $6,
  catcode = $7,
  catcode2 = $8,
  catcode3 = $9,
  catcode4 = $10,
  gender = $11,
  racedate = $12,
  starteventdate = $13,
  description = $14,
  place = $15,
  nationcode = $16,
  td1id = $17,
  td1name = $18,
  td1nation = $19,
  td1code = $20,
  td2id = $21,
  td2name = $22,
  td2nation = $23,
  td2code = $24,
  calstatuscode = $25,
  procstatuscode = $26,
  receiveddate = $27,
  pursuit = $28,
  masse = $29,
  relay = $30,
  distance = $31,
  hill = $32,
  style = $33,
  qualif = $34,
  finale = $35,
  homol = $36,
  webcomment = $37,
  displaystatus = $38,
  fisinterncomment = $39,
  published = $40,
  validforfispoints = $41,
  usedfislist = $42,
  tolist = $43,
  discforlistcode = $44,
  calculatedpenalty = $45,
  appliedpenalty = $46,
  appliedscala = $47,
  penscafixed = $48,
  version = $49,
  nationraceid = $50,
  provraceid = $51,
  msql7evid = $52,
  mssql7id = $53,
  results = $54,
  pdf = $55,
  topbanner = $56,
  bottombanner = $57,
  toplogo = $58,
  bottomlogo = $59,
  gallery = $60,
  indi = $61,
  team = $62,
  tabcount = $63,
  columncount = $64,
  level = $65,
  hloc1 = $66,
  hloc2 = $67,
  hloc3 = $68,
  hcet1 = $69,
  hcet2 = $70,
  hcet3 = $71,
  live = $72,
  livestatus1 = $73,
  livestatus2 = $74,
  livestatus3 = $75,
  liveinfo1 = $76,
  liveinfo2 = $77,
  liveinfo3 = $78,
  passwd = $79,
  timinglogo = $80,
  validdate = $81,
  noepr = $82,
  tddoc = $83,
  timingreport = $84,
  special_cup_points = $85,
  skip_wcsl = $86,
  validforowg = $87,
  lastupdate = $88
WHERE raceid = $1
RETURNING raceid;


-- name: DeleteRaceNKByID :one
DELETE FROM a_racenk
WHERE raceid = $1
RETURNING raceid;



-- name: InsertResultCC :exec
INSERT INTO a_resultcc (
  recid, raceid, competitorid, status, reason,
  "position", pf, status2, bib, bibcolor,
  fiscode, competitorname, nationcode, stage, level, heat,
  timer1, timer2, timer3, timetot, valid,
  racepoints, cuppoints, bonustime, bonuscuppoints, version,
  rg1, rg2, lastupdate
) VALUES (
  $1, $2, $3, $4, $5,
  $6, $7, $8, $9, $10,
  $11, $12, $13, $14, $15, $16,
  $17, $18, $19, $20, $21,
  $22, $23, $24, $25, $26,
  $27, $28, $29
);

-- name: UpdateResultCCByRecID :one
UPDATE a_resultcc SET
  raceid          = $2,
  competitorid    = $3,
  status          = $4,
  reason          = $5,
  "position"      = $6,
  pf              = $7,
  status2         = $8,
  bib             = $9,
  bibcolor        = $10,
  fiscode         = $11,
  competitorname  = $12,
  nationcode      = $13,
  stage           = $14,
  level           = $15,
  heat            = $16,
  timer1          = $17,
  timer2          = $18,
  timer3          = $19,
  timetot         = $20,
  valid           = $21,
  racepoints      = $22,
  cuppoints       = $23,
  bonustime       = $24,
  bonuscuppoints  = $25,
  version         = $26,
  rg1             = $27,
  rg2             = $28,
  lastupdate      = $29
WHERE recid = $1
RETURNING recid;

-- name: DeleteResultCCByRecID :one
DELETE FROM a_resultcc
WHERE recid = $1
RETURNING recid;


-- name: InsertResultJP :exec
INSERT INTO a_resultjp (
  recid, raceid, competitorid, status, status2, "position", bib,
  fiscode, competitorname, nationcode, level, heat, stage,
  j1r1, j2r1, j3r1, j4r1, j5r1, speedr1, distr1, disptsr1, judptsr1, totrun1, posr1, statusr1,
  j1r2, j2r2, j3r2, j4r2, j5r2, speedr2, distr2, disptsr2, judptsr2, totrun2, posr2, statusr2,
  j1r3, j2r3, j3r3, j4r3, j5r3, speedr3, distr3, disptsr3, judptsr3, totrun3, posr3, statusr3,
  j1r4, j2r4, j3r4, j4r4, j5r4, speedr4, distr4, disptsr4, judptsr4,
  gater1, gater2, gater3, gater4,
  gateptsr1, gateptsr2, gateptsr3, gateptsr4,
  windr1, windr2, windr3, windr4,
  windptsr1, windptsr2, windptsr3, windptsr4,
  reason, totrun4, tot, valid, racepoints, cuppoints, version, lastupdate, posr4, statusr4
) VALUES (
  $1,$2,$3,$4,$5,$6,$7,
  $8,$9,$10,$11,$12,$13,
  $14,$15,$16,$17,$18,$19,$20,$21,$22,$23,$24,$25,
  $26,$27,$28,$29,$30,$31,$32,$33,$34,$35,$36,
  $37,$38,$39,$40,$41,$42,$43,$44,$45,$46,$47,
  $48,$49,$50,$51,$52,$53,$54,$55,$56,$57,$58,
  $59,$60,$61,$62,
  $63,$64,$65,$66,
  $67,$68,$69,$70,
  $71,$72,$73,$74,
  $75,$76,$77,$78,$79,$80,$81,$82,$83,$84
);

-- name: UpdateResultJPByRecID :one
UPDATE a_resultjp SET
  raceid = $2, competitorid = $3, status = $4, status2 = $5, "position" = $6, bib = $7,
  fiscode = $8, competitorname = $9, nationcode = $10, level = $11, heat = $12, stage = $13,
  j1r1 = $14, j2r1 = $15, j3r1 = $16, j4r1 = $17, j5r1 = $18, speedr1 = $19, distr1 = $20, disptsr1 = $21, judptsr1 = $22, totrun1 = $23, posr1 = $24, statusr1 = $25,
  j1r2 = $26, j2r2 = $27, j3r2 = $28, j4r2 = $29, j5r2 = $30, speedr2 = $31, distr2 = $32, disptsr2 = $33, judptsr2 = $34, totrun2 = $35, posr2 = $36, statusr2 = $37,
  j1r3 = $38, j2r3 = $39, j3r3 = $40, j4r3 = $41, j5r3 = $42, speedr3 = $43, distr3 = $44, disptsr3 = $45, judptsr3 = $46, totrun3 = $47, posr3 = $48, statusr3 = $49,
  j1r4 = $50, j2r4 = $51, j3r4 = $52, j4r4 = $53, j5r4 = $54, speedr4 = $55, distr4 = $56, disptsr4 = $57, judptsr4 = $58,
  gater1 = $59, gater2 = $60, gater3 = $61, gater4 = $62,
  gateptsr1 = $63, gateptsr2 = $64, gateptsr3 = $65, gateptsr4 = $66,
  windr1 = $67, windr2 = $68, windr3 = $69, windr4 = $70,
  windptsr1 = $71, windptsr2 = $72, windptsr3 = $73, windptsr4 = $74,
  reason = $75, totrun4 = $76, tot = $77, valid = $78, racepoints = $79, cuppoints = $80, version = $81, lastupdate = $82, posr4 = $83, statusr4 = $84
WHERE recid = $1
RETURNING recid;

-- name: DeleteResultJPByRecID :one
DELETE FROM a_resultjp
WHERE recid = $1
RETURNING recid;


-- name: InsertResultNK :exec
INSERT INTO a_resultnk (
  recid, raceid, competitorid, status, status2, reason, "position", pf, bib, bibcolor,
  fiscode, competitorname, nationcode, level, heat, stage,
  j1r1, j2r1, j3r1, j4r1, j5r1, speedr1, distr1, disptsr1, judptsr1,
  gater1, gateptsr1, windr1, windptsr1, totrun1, posr1, statusr1,
  j1r2, j2r2, j3r2, j4r2, j5r2, speedr2, distr2, disptsr2, judptsr2,
  gater2, gateptsr2, windr2, windptsr2, totrun2, posr2, statusr2,
  pointsjump, behindjump, posjump, timecc, timeccint, poscc, starttime, statuscc, totbehind, timetot, timetotint,
  valid, racepoints, cuppoints, version, lastupdate
) VALUES (
  $1,$2,$3,$4,$5,$6,$7,$8,$9,$10,
  $11,$12,$13,$14,$15,$16,
  $17,$18,$19,$20,$21,$22,$23,$24,$25,
  $26,$27,$28,$29,$30,$31,$32,
  $33,$34,$35,$36,$37,$38,$39,$40,$41,
  $42,$43,$44,$45,$46,$47,$48,
  $49,$50,$51,$52,$53,$54,$55,$56,$57,$58,$59,
  $60,$61,$62,$63,$64
);

-- name: UpdateResultNKByRecID :one
UPDATE a_resultnk SET
  raceid=$2, competitorid=$3, status=$4, status2=$5, reason=$6, "position"=$7, pf=$8, bib=$9, bibcolor=$10,
  fiscode=$11, competitorname=$12, nationcode=$13, level=$14, heat=$15, stage=$16,
  j1r1=$17, j2r1=$18, j3r1=$19, j4r1=$20, j5r1=$21, speedr1=$22, distr1=$23, disptsr1=$24, judptsr1=$25,
  gater1=$26, gateptsr1=$27, windr1=$28, windptsr1=$29, totrun1=$30, posr1=$31, statusr1=$32,
  j1r2=$33, j2r2=$34, j3r2=$35, j4r2=$36, j5r2=$37, speedr2=$38, distr2=$39, disptsr2=$40, judptsr2=$41,
  gater2=$42, gateptsr2=$43, windr2=$44, windptsr2=$45, totrun2=$46, posr2=$47, statusr2=$48,
  pointsjump=$49, behindjump=$50, posjump=$51, timecc=$52, timeccint=$53, poscc=$54, starttime=$55, statuscc=$56, totbehind=$57, timetot=$58, timetotint=$59,
  valid=$60, racepoints=$61, cuppoints=$62, version=$63, lastupdate=$64
WHERE recid = $1
RETURNING recid;

-- name: DeleteResultNKByRecID :one
DELETE FROM a_resultnk
WHERE recid = $1
RETURNING recid;

-- name: GetAthletesBySporttiID :many
SELECT
  fiscode,
  sporttiid,
  firstname,
  lastname
FROM athlete
WHERE sporttiid = $1
ORDER BY fiscode;

-- name: InsertAthlete :exec
INSERT INTO public.athlete (
  fiscode,
  sporttiid,
  firstname,
  lastname
) VALUES (
  $1, $2, $3, $4
);

-- name: UpdateAthleteByFiscode :one
UPDATE public.athlete SET
  sporttiid = $2,
  firstname = $3,
  lastname  = $4
WHERE fiscode = $1
RETURNING fiscode;

-- name: DeleteAthleteByFiscode :one
DELETE FROM public.athlete
WHERE fiscode = $1
RETURNING fiscode;

-- name: SearchRacesCC :many
SELECT
  gender,
  raceid,
  racedate,
  catcode,
  description,
  place,
  nationcode,
  disciplinecode,
  'CC'::text AS sectorcode
FROM a_racecc
WHERE
  ($1::int4 = 0 OR seasoncode = $1::int4)
  AND ($2::text = '' OR nationcode = $2::text)
  AND ($3::text = '' OR gender = $3::text)
  AND ($4::text = '' OR catcode = $4::text)
  AND calstatuscode = 'O'
ORDER BY racedate;


-- name: SearchRacesJP :many
SELECT
  gender,
  raceid,
  racedate,
  catcode,
  description,
  place,
  nationcode,
  disciplinecode,
  'JP'::text AS sectorcode
FROM a_racejp
WHERE
  ($1::int4 = 0 OR seasoncode = $1::int4)
  AND ($2::text = '' OR nationcode = $2::text)
  AND ($3::text = '' OR gender = $3::text)
  AND ($4::text = '' OR catcode = $4::text)
  AND calstatuscode = 'O'
ORDER BY racedate;


-- name: SearchRacesNK :many
SELECT
  gender,
  raceid,
  racedate,
  catcode,
  description,
  place,
  nationcode,
  disciplinecode,
  'NK'::text AS sectorcode
FROM a_racenk
WHERE
  ($1::int4 = 0 OR seasoncode = $1::int4)
  AND ($2::text = '' OR nationcode = $2::text)
  AND ($3::text = '' OR gender = $3::text)
  AND ($4::text = '' OR catcode = $4::text)
  AND calstatuscode = 'O'
ORDER BY racedate;


-- name: GetLatestResultsCC :many
SELECT
  rcc.recid,
  rcc.raceid,
  rcc.position,
  rcc.timetot,
  rcc.competitorid,
  acc.racedate,
  acc.seasoncode,
  acc.disciplinecode,
  acc.catcode,
  acc.place,
  acc.nationcode
FROM a_resultcc AS rcc
JOIN a_racecc   AS acc
  ON rcc.raceid = acc.raceid
WHERE rcc.fiscode = $1::int4
  AND ($2::int4 = 0 OR acc.seasoncode = $2::int4)
  AND ($3::text[] IS NULL OR acc.catcode = ANY($3::text[]))
ORDER BY acc.racedate DESC
LIMIT $4::int4;


-- name: GetLatestResultsJP :many
SELECT 
  rjp.raceid,
  rjp.position,
  ajp.racedate,
  ajp.seasoncode,
  ajp.disciplinecode,
  ajp.catcode,
  ajp.place,
  ajp.nationcode,
  rjp.posr1,
  rjp.speedr1,
  rjp.distr1,
  rjp.judptsr1,
  rjp.windr1,
  rjp.windptsr1,
  rjp.gater1,
  rjp.posr2,
  rjp.speedr2,
  rjp.distr2,
  rjp.judptsr2,
  rjp.windr2,
  rjp.windptsr2,
  rjp.gater2,
  rjp.totrun1,
  rjp.totrun2
FROM a_resultjp AS rjp
JOIN a_racejp   AS ajp
  ON rjp.raceid = ajp.raceid
WHERE rjp.fiscode = $1::int4
  AND ($2::int4 = 0 OR ajp.seasoncode = $2::int4)
  AND ($3::text[] IS NULL OR ajp.catcode = ANY($3::text[]))
ORDER BY ajp.racedate DESC
LIMIT $4::int4;


-- name: GetLatestResultsNK :many
SELECT 
  rnk.recid,
  rnk.raceid,
  rnk.position,
  ank.racedate,
  ank.seasoncode,
  ank.distance,
  ank.hill,
  ank.disciplinecode,
  ank.catcode,
  ank.place,
  ank.nationcode,
  rnk.posr1,
  rnk.speedr1,
  rnk.distr1,
  rnk.judptsr1,
  rnk.windr1,
  rnk.windptsr1,
  rnk.gater1,
  rnk.totrun1,
  rnk.poscc,
  rnk.timetot,
  rnk.timetotint,
  rnk.pointsjump
FROM a_resultnk AS rnk
JOIN a_racenk   AS ank
  ON rnk.raceid = ank.raceid
WHERE rnk.fiscode = $1::int4
  AND ($2::int4 = 0 OR ank.seasoncode = $2::int4)
  AND ($3::text[] IS NULL OR ank.catcode = ANY($3::text[]))
ORDER BY ank.racedate DESC
LIMIT $4::int4;


-- name: SearchCompetitors :many
SELECT *
FROM a_competitor
WHERE ($1::text = '' OR nationcode  = $1::text)
  AND ($2::text = '' OR sectorcode  = $2::text)
  AND ($3::text = '' OR gender      = $3::text)
  AND ($4::date = '0001-01-01' OR birthdate >= $4::date)
  AND ($5::date = '0001-01-01' OR birthdate <= $5::date)
ORDER BY competitorid;


-- name: GetRacesByIDsCC :many
SELECT *
FROM A_raceCC
WHERE raceid = ANY($1::int[])
ORDER BY raceid;

-- name: GetRacesByIDsJP :many
SELECT *
FROM A_raceJP
WHERE raceid = ANY($1::int[])
ORDER BY raceid;

-- name: GetRacesByIDsNK :many
SELECT *
FROM A_raceNK
WHERE raceid = ANY($1::int[])
ORDER BY raceid;

-- name: GetSeasonsCatcodesCCByCompetitor :many
SELECT DISTINCT
  rcc.seasoncode,
  rcc.catcode
FROM a_racecc   AS rcc
JOIN a_resultcc AS res
  ON rcc.raceid = res.raceid
WHERE res.fiscode = $1::int4
ORDER BY rcc.seasoncode DESC, rcc.catcode ASC;


-- name: GetSeasonsCatcodesJPByCompetitor :many
SELECT DISTINCT
  rjp.seasoncode,
  rjp.catcode
FROM a_racejp   AS rjp
JOIN a_resultjp AS res
  ON rjp.raceid = res.raceid
WHERE res.fiscode = $1::int4
ORDER BY rjp.seasoncode DESC, rjp.catcode ASC;


-- name: GetSeasonsCatcodesNKByCompetitor :many
SELECT DISTINCT
  rnk.seasoncode,
  rnk.catcode
FROM a_racenk   AS rnk
JOIN a_resultnk AS res
  ON rnk.raceid = res.raceid
WHERE res.fiscode = $1::int4
ORDER BY rnk.seasoncode DESC, rnk.catcode ASC;

-- name: GetRaceCountsByCategoryCC :many
SELECT
  catcode,
  COUNT(*) AS total
FROM a_racecc
WHERE seasoncode = $1::int4
  AND ($2::text = '' OR nationcode = $2::text)
  AND ($3::text = '' OR gender     = $3::text)
GROUP BY catcode
ORDER BY total DESC;


-- name: GetRaceCountsByCategoryJP :many
SELECT
  catcode,
  COUNT(*) AS total
FROM a_racejp
WHERE seasoncode = $1::int4
  AND ($2::text = '' OR nationcode = $2::text)
  AND ($3::text = '' OR gender     = $3::text)
GROUP BY catcode
ORDER BY total DESC;


-- name: GetRaceCountsByCategoryNK :many
SELECT
  catcode,
  COUNT(*) AS total
FROM a_racenk
WHERE seasoncode = $1::int4
  AND ($2::text = '' OR nationcode = $2::text)
  AND ($3::text = '' OR gender     = $3::text)
GROUP BY catcode
ORDER BY total DESC;

-- name: GetRaceTotalCC :one
SELECT COUNT(*) AS total
FROM a_racecc
WHERE seasoncode = $1::int4
  AND ($2::text = '' OR catcode = $2::text)
  AND ($3::text = '' OR gender  = $3::text);


-- name: GetRaceTotalJP :one
SELECT COUNT(*) AS total
FROM a_racejp
WHERE seasoncode = $1::int4
  AND ($2::text = '' OR catcode = $2::text)
  AND ($3::text = '' OR gender  = $3::text);


-- name: GetRaceTotalNK :one
SELECT COUNT(*) AS total
FROM a_racenk
WHERE seasoncode = $1::int4
  AND ($2::text = '' OR catcode = $2::text)
  AND ($3::text = '' OR gender  = $3::text);

-- name: GetCompetitorCountsByNation :many
SELECT
  nationcode,
  COUNT(*) AS competitors
FROM a_competitor
WHERE ($1::text = ''          OR sectorcode = $1::text)
  AND ($2::text = ''          OR gender     = $2::text)
  AND ($3::date = '0001-01-01' OR birthdate >= $3::date)
  AND ($4::date = '0001-01-01' OR birthdate <= $4::date)
GROUP BY nationcode
ORDER BY competitors DESC;


-- name: GetRaceCountsByNationCC :many
SELECT
  nationcode,
  COUNT(*) AS total
FROM a_racecc
WHERE seasoncode = $1::int4
  AND ($2::text = '' OR catcode = $2::text)
  AND ($3::text = '' OR gender  = $3::text)
GROUP BY nationcode
ORDER BY total DESC;


-- name: GetRaceCountsByNationJP :many
SELECT
  nationcode,
  COUNT(*) AS total
FROM a_racejp
WHERE seasoncode = $1::int4
  AND ($2::text = '' OR catcode = $2::text)
  AND ($3::text = '' OR gender  = $3::text)
GROUP BY nationcode
ORDER BY total DESC;


-- name: GetRaceCountsByNationNK :many
SELECT
  nationcode,
  COUNT(*) AS total
FROM a_racenk
WHERE seasoncode = $1::int4
  AND ($2::text = '' OR catcode = $2::text)
  AND ($3::text = '' OR gender  = $3::text)
GROUP BY nationcode
ORDER BY total DESC;
