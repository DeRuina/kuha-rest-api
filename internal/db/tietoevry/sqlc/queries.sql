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
ON CONFLICT (source, raw_id) DO NOTHING;


-- name: InsertExerciseHRZone :exec
INSERT INTO exercise_hr_zones (
    exercise_id, zone_index, seconds_in_zone,
    lower_limit, upper_limit, created_at, updated_at
) VALUES (
    $1, $2, $3,
    $4, $5, $6, $7
)
ON CONFLICT (exercise_id, zone_index) DO NOTHING;

-- name: InsertExerciseSample :exec
INSERT INTO exercise_samples (
    id, user_id, exercise_id,
    sample_type, recording_rate, samples, source
) VALUES (
    $1, $2, $3,
    $4, $5, $6, $7
)
ON CONFLICT (exercise_id, sample_type) DO NOTHING;


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
ON CONFLICT (id) DO NOTHING;


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
ON CONFLICT (source, user_id, date, raw_id) DO NOTHING;


-- name: InsertMeasurement :exec
INSERT INTO measurements (
    id, created_at, updated_at, user_id, date, name, name_type,
    source, value, value_numeric, comment, raw_id, raw_data, additional_info
) VALUES (
    $1, $2, $3, $4, $5, $6, $7,
    $8, $9, $10, $11, $12, $13, $14
)
ON CONFLICT (source, user_id, date, name, raw_id) DO NOTHING;

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
ON CONFLICT (id) DO NOTHING;


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
ON CONFLICT (user_id, questionnaire_instance_id, question_id) DO NOTHING;


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
ON CONFLICT (user_id, date, source) DO NOTHING;

-- name: GetUser :one
SELECT * FROM users WHERE id = $1;

-- name: LogDeletedUser :exec
INSERT INTO deleted_users_log (user_id, sportti_id)
SELECT id, sportti_id 
FROM users 
WHERE id = $1;