-- name: UpsertUser :exec
INSERT INTO users (
    id, sportti_id, profile_gender, profile_birthdate, profile_weight,
    profile_height, profile_resting_heart_rate, profile_maximum_heart_rate,
    profile_aerobic_threshold, profile_anaerobic_threshold, profile_vo2max
) VALUES (
    $1, $2, $3, $4, $5,
    $6, $7, $8,
    $9, $10, $11
)
ON CONFLICT (id) DO UPDATE SET
    sportti_id = EXCLUDED.sportti_id,
    profile_gender = EXCLUDED.profile_gender,
    profile_birthdate = EXCLUDED.profile_birthdate,
    profile_weight = EXCLUDED.profile_weight,
    profile_height = EXCLUDED.profile_height,
    profile_resting_heart_rate = EXCLUDED.profile_resting_heart_rate,
    profile_maximum_heart_rate = EXCLUDED.profile_maximum_heart_rate,
    profile_aerobic_threshold = EXCLUDED.profile_aerobic_threshold,
    profile_anaerobic_threshold = EXCLUDED.profile_anaerobic_threshold,
    profile_vo2max = EXCLUDED.profile_vo2max;

-- name: DeleteUser :execrows
DELETE FROM users
WHERE id = $1;

-- name: InsertExercise :exec
INSERT INTO exercises (
    id, created_at, updated_at, user_id, start_time, duration,
    comment, sport_type, detailed_sport_type, distance, avg_heart_rate,
    max_heart_rate, trimp, sprint_count, avg_speed, max_speed,
    source, status, calories, training_load, raw_id,
    feeling, recovery, rpe, raw_data
) VALUES (
    $1, $2, $3, $4, $5, $6,
    $7, $8, $9, $10, $11,
    $12, $13, $14, $15, $16,
    $17, $18, $19, $20, $21,
    $22, $23, $24, $25
)
ON CONFLICT (id) DO UPDATE SET
  updated_at          = GREATEST(exercises.updated_at, EXCLUDED.updated_at),
  start_time          = EXCLUDED.start_time,
  duration            = EXCLUDED.duration,
  comment             = EXCLUDED.comment,
  sport_type          = EXCLUDED.sport_type,
  detailed_sport_type = EXCLUDED.detailed_sport_type,
  distance            = EXCLUDED.distance,
  avg_heart_rate      = EXCLUDED.avg_heart_rate,
  max_heart_rate      = EXCLUDED.max_heart_rate,
  trimp               = EXCLUDED.trimp,
  sprint_count        = EXCLUDED.sprint_count,
  avg_speed           = EXCLUDED.avg_speed,
  max_speed           = EXCLUDED.max_speed,
  status              = EXCLUDED.status,
  calories            = EXCLUDED.calories,
  training_load       = EXCLUDED.training_load,
  feeling             = EXCLUDED.feeling,
  recovery            = EXCLUDED.recovery,
  rpe                 = EXCLUDED.rpe,
  raw_data            = EXCLUDED.raw_data;


-- name: InsertExerciseHRZone :exec
INSERT INTO exercise_hr_zones (
    exercise_id, zone_index, seconds_in_zone,
    lower_limit, upper_limit, created_at, updated_at
) VALUES (
    $1, $2, $3,
    $4, $5, $6, $7
)
ON CONFLICT (exercise_id, zone_index) DO UPDATE SET
  seconds_in_zone = EXCLUDED.seconds_in_zone,
  lower_limit     = EXCLUDED.lower_limit,
  upper_limit     = EXCLUDED.upper_limit,
  updated_at      = GREATEST(exercise_hr_zones.updated_at, EXCLUDED.updated_at);

-- name: InsertExerciseSample :exec
INSERT INTO exercise_samples (
    id, user_id, exercise_id,
    sample_type, recording_rate, samples, source
) VALUES (
    $1, $2, $3,
    $4, $5, $6, $7
)
ON CONFLICT (exercise_id, sample_type) DO UPDATE SET
    recording_rate = EXCLUDED.recording_rate,
    samples = EXCLUDED.samples;

-- name: InsertExerciseSection :exec
INSERT INTO exercise_sections (
    id, user_id, exercise_id,
    created_at, updated_at, start_time, end_time,
    section_type, name, comment, source, raw_id, raw_data
) VALUES (
    $1, $2, $3,
    $4, $5, $6, $7,
    $8, $9, $10, $11, $12, $13
)
ON CONFLICT (id) DO UPDATE SET
  exercise_id = EXCLUDED.exercise_id,
  updated_at  = GREATEST(exercise_sections.updated_at, EXCLUDED.updated_at),
  start_time  = EXCLUDED.start_time,
  end_time    = EXCLUDED.end_time,
  section_type= EXCLUDED.section_type,
  name        = EXCLUDED.name,
  comment     = EXCLUDED.comment,
  raw_data    = EXCLUDED.raw_data;


-- name: InsertSymptom :exec
INSERT INTO symptoms (
    id, user_id, date, symptom, severity, comment, source,
    created_at, updated_at, raw_id, original_id, recovered,
    pain_index, side, category, additional_data
) VALUES (
    $1, $2, $3, $4, $5, $6, $7,
    $8, $9, $10, $11, $12,
    $13, $14, $15, $16
)
ON CONFLICT (id) DO UPDATE SET
  symptom         = EXCLUDED.symptom,
  severity        = EXCLUDED.severity,
  comment         = EXCLUDED.comment,
  updated_at      = GREATEST(symptoms.updated_at, EXCLUDED.updated_at),
  original_id     = EXCLUDED.original_id,
  recovered       = EXCLUDED.recovered,
  pain_index      = EXCLUDED.pain_index,
  side            = EXCLUDED.side,
  category        = EXCLUDED.category,
  additional_data = EXCLUDED.additional_data;


-- name: InsertMeasurement :exec
INSERT INTO measurements (
    id, created_at, updated_at, user_id, date, name, name_type,
    source, value, value_numeric, comment, raw_id, raw_data, additional_info
) VALUES (
    $1, $2, $3, $4, $5, $6, $7,
    $8, $9, $10, $11, $12, $13, $14
)
ON CONFLICT (id) DO UPDATE SET
  updated_at     = GREATEST(measurements.updated_at, EXCLUDED.updated_at),
  name_type      = EXCLUDED.name_type,
  value          = EXCLUDED.value,
  value_numeric  = EXCLUDED.value_numeric,
  comment        = EXCLUDED.comment,
  raw_data       = EXCLUDED.raw_data,
  additional_info= EXCLUDED.additional_info;


-- name: InsertTestResult :exec
INSERT INTO test_results (
    id, user_id, type_id, type_type, type_result_type, type_name,
    timestamp, name, comment, data, created_at, updated_at,
    test_event_id, test_event_name, test_event_date,
    test_event_template_test_id, test_event_template_test_name,
    test_event_template_test_limits
) VALUES (
    $1, $2, $3, $4, $5, $6,
    $7, $8, $9, $10, $11, $12,
    $13, $14, $15, $16, $17, $18
)
ON CONFLICT (id) DO UPDATE SET
  type_id                        = EXCLUDED.type_id,
  type_type                      = EXCLUDED.type_type,
  type_result_type               = EXCLUDED.type_result_type,
  type_name                      = EXCLUDED.type_name,
  timestamp                      = EXCLUDED.timestamp,
  name                           = EXCLUDED.name,
  comment                        = EXCLUDED.comment,
  data                           = EXCLUDED.data,
  updated_at                     = GREATEST(test_results.updated_at, EXCLUDED.updated_at),
  test_event_id                  = EXCLUDED.test_event_id,
  test_event_name                = EXCLUDED.test_event_name,
  test_event_date                = EXCLUDED.test_event_date,
  test_event_template_test_id    = EXCLUDED.test_event_template_test_id,
  test_event_template_test_name  = EXCLUDED.test_event_template_test_name,
  test_event_template_test_limits= EXCLUDED.test_event_template_test_limits;


-- name: InsertQuestionnaireAnswer :exec
INSERT INTO question_answers (
    user_id, questionnaire_instance_id, questionnaire_name_fi,
    questionnaire_name_en, questionnaire_key, question_id, question_label_fi,
    question_label_en, question_type, option_id, option_value,
    option_label_fi, option_label_en, free_text, created_at, updated_at, value
) VALUES (
    $1, $2, $3,
    $4, $5, $6, $7,
    $8, $9, $10, $11,
    $12, $13, $14, $15, $16, $17
)
ON CONFLICT (questionnaire_instance_id, question_id, user_id) DO UPDATE SET
  questionnaire_name_fi = EXCLUDED.questionnaire_name_fi,
  questionnaire_name_en = EXCLUDED.questionnaire_name_en,
  questionnaire_key     = EXCLUDED.questionnaire_key,
  question_label_fi     = EXCLUDED.question_label_fi,
  question_label_en     = EXCLUDED.question_label_en,
  question_type         = EXCLUDED.question_type,
  option_id             = EXCLUDED.option_id,
  option_value          = EXCLUDED.option_value,
  option_label_fi       = EXCLUDED.option_label_fi,
  option_label_en       = EXCLUDED.option_label_en,
  free_text             = EXCLUDED.free_text,
  updated_at            = GREATEST(question_answers.updated_at, EXCLUDED.updated_at),
  value                 = EXCLUDED.value;



-- name: InsertActivityZone :exec
INSERT INTO activity_zones (
    user_id, date, created_at, updated_at,
    seconds_in_zone_0, seconds_in_zone_1, seconds_in_zone_2,
    seconds_in_zone_3, seconds_in_zone_4, seconds_in_zone_5,
    source, raw_data
) VALUES (
    $1, $2, $3, $4,
    $5, $6, $7,
    $8, $9, $10,
    $11, $12
)
ON CONFLICT (user_id, date, source) DO UPDATE SET
  updated_at       = GREATEST(activity_zones.updated_at, EXCLUDED.updated_at),
  seconds_in_zone_0= EXCLUDED.seconds_in_zone_0,
  seconds_in_zone_1= EXCLUDED.seconds_in_zone_1,
  seconds_in_zone_2= EXCLUDED.seconds_in_zone_2,
  seconds_in_zone_3= EXCLUDED.seconds_in_zone_3,
  seconds_in_zone_4= EXCLUDED.seconds_in_zone_4,
  seconds_in_zone_5= EXCLUDED.seconds_in_zone_5,
  raw_data         = EXCLUDED.raw_data;

-- name: GetUser :one
SELECT * FROM users WHERE id = $1;

-- name: LogDeletedUser :exec
INSERT INTO deleted_users_log (user_id, sportti_id)
SELECT users.id, users.sportti_id 
FROM users 
WHERE users.id = $1;

-- name: GetDeletedUsers :many
SELECT id, user_id, sportti_id, deleted_at
FROM deleted_users_log
ORDER BY deleted_at DESC;

-- name: GetExercisesByUser :many
SELECT * FROM exercises
WHERE user_id = $1
ORDER BY start_time DESC;

-- name: GetExerciseHRZones :many
SELECT * FROM exercise_hr_zones
WHERE exercise_id = $1
ORDER BY zone_index;

-- name: GetExerciseSamples :many
SELECT * FROM exercise_samples
WHERE exercise_id = $1;

-- name: GetExerciseSections :many
SELECT * FROM exercise_sections
WHERE exercise_id = $1
ORDER BY start_time;

-- name: GetSymptomsByUser :many
SELECT * FROM symptoms
WHERE user_id = $1
ORDER BY date DESC, created_at DESC;

-- name: GetMeasurementsByUser :many
SELECT * FROM measurements
WHERE user_id = $1
ORDER BY date DESC, created_at DESC;

-- name: GetTestResultsByUser :many
SELECT * FROM test_results
WHERE user_id = $1
ORDER BY timestamp DESC, created_at DESC;

-- name: GetQuestionnairesByUser :many
SELECT * FROM question_answers
WHERE user_id = $1
ORDER BY created_at DESC;

-- name: GetActivityZonesByUser :many
SELECT * FROM activity_zones
WHERE user_id = $1
ORDER BY date DESC, created_at DESC;