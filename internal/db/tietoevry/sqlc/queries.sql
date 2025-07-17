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

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;