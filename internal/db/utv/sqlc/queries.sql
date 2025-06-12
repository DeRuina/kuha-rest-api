-- Common sql queries

-- user_db.py

-- name: AddUser :exec
INSERT INTO user_data(user_id, data)
VALUES ($1, $2);

-- name: RetrieveUser :one
SELECT data
FROM user_data
WHERE user_id = $1;

-- name: AddOrUpdateUser :exec
INSERT INTO user_data(user_id, data)
VALUES ($1, $2)
ON CONFLICT (user_id)
    DO UPDATE SET data = EXCLUDED.data;

-- name: DeleteUser :exec
DELETE
FROM user_data
WHERE user_id = $1;

-- name: CreateGroup :exec
INSERT INTO utv_groups(id, group_name, created, active, deleted)
VALUES($1, $2, $3, $4, $5);

-- name: ListGroups :many
SELECT * FROM utv_groups;

-- name: ListGroupMembers :many
SELECT user_id, added
FROM utv_group_members
WHERE group_id = $1;

-- name: ListGroupsForUser :many
SELECT utv_groups.id, utv_groups.group_name, utv_groups.created, utv_groups.active, utv_groups.deleted FROM utv_groups, utv_group_members
WHERE utv_groups.id = utv_group_members.group_id
    AND utv_group_members.user_id = $1;

-- name: DeleteGroup :exec
DELETE FROM utv_groups
WHERE id = $1;

-- name: AddUserToGroup :exec
INSERT INTO utv_group_members(group_id, user_id, added)
VALUES ($1, $2, $3);

-- name: RemoveUserFromGroup :exec
DELETE from utv_group_members
WHERE user_id = $1 AND group_id = $2;

-- notifications.py

-- name: AddNotification :one
INSERT INTO notifications(id, to_id, from_id, status, expires, notification)
VALUES(uuid_generate_v4(), $1, $2, $3, $4, $5)
RETURNING id;

-- name: ListNotifications :many
SELECT id, to_id, from_id, status, expires, notification
FROM notifications
WHERE
    ($1 IS NULL OR to_id = $1) AND
    ($2 IS NULL OR from_id = $2) AND
    ($3 IS NULL OR expires >= $3) AND
    ($4 IS NULL OR $4);

-- name: GetNotification :one
SELECT id, to_id, from_id, status, expires, notification
FROM notifications
WHERE id = @id AND ($1 IS NULL OR expires >= $1);

-- name: ToggleNotificationExpiration :one
UPDATE notifications
SET expires = -(SELECT expires FROM notifications WHERE notifications.id = $1)
WHERE notifications.id = $1
RETURNING expires;

-- name: SetNotificationStatus :one
UPDATE notifications
SET status = $1 
WHERE id = $2
RETURNING status;

-- api_db.py

-- name: GetAllDataTypes :many
SELECT data
FROM source_cache
WHERE ($1 IS NULL OR source = $1);

-- name: SetResourceMetadata :exec
INSERT INTO resource_data(resource_id, data)
VALUES($1, $2)
ON CONFLICT(resource_id)
DO UPDATE SET data = EXCLUDED.data;

-- name: GetResourceMetadata :many
SELECT data
FROM resource_data
WHERE resource_id = $1;

-- name: SetAppData :exec
INSERT INTO app_data(app_id, field_name, data)
VALUES($1, $2, $3)
ON CONFLICT(app_id, field_name)
DO UPDATE SET data = EXCLUDED.data;

-- name: GetAppData :many
SELECT data
FROM app_data
WHERE app_id = $1
AND field_name = $2;

-- name: SetPersonalInformation :exec
INSERT INTO user_data(user_id, data)
VALUES($1, $2)
ON CONFLICT(user_id)
DO UPDATE SET DATA = $3;

-- Wearable related functions

-- api_db.py

-- name: GetDatesFromCoachtechData :many
WITH cte AS (
    SELECT coachtech_id
    FROM coachtech_ids
    WHERE user_id = $1
)
SELECT DISTINCT summary_date
FROM (
    SELECT summary_date
    FROM coachtech_data
    WHERE coachtech_id = (SELECT coachtech_id FROM cte)
    AND
    summary_date BETWEEN $2 AND $3
    AND
    data ->> 'testType' LIKE $4
) AS acceptables
ORDER BY summary_date DESC;

-- name: GetDatesFromOuraData :many
SELECT DISTINCT summary_date
FROM oura_data
WHERE user_id = @user_id
AND (@after_date::date IS NULL OR summary_date >= @after_date)
AND (@before_date::date IS NULL OR summary_date <= @before_date)
ORDER BY summary_date DESC;



-- name: GetDatesFromPolarData :many
SELECT DISTINCT summary_date
FROM polar_data
WHERE user_id = @user_id
AND (@after_date::date IS NULL OR summary_date >= @after_date)
AND (@before_date::date IS NULL OR summary_date <= @before_date)
ORDER BY summary_date DESC;

-- name: GetDatesFromSuuntoData :many
SELECT DISTINCT summary_date
FROM suunto_data
WHERE user_id = @user_id
AND (@after_date::date IS NULL OR summary_date >= @after_date)
AND (@before_date::date IS NULL OR summary_date <= @before_date)
ORDER BY summary_date DESC;

-- name: GetTypesFromSuuntoData :many
SELECT DISTINCT jsonb_object_keys(data)
FROM suunto_data
WHERE user_id = @user_id
AND summary_date = @date;

-- name: GetAllDataForDateSuunto :one
SELECT data
FROM suunto_data
WHERE user_id = $1
AND summary_date = $2;

-- name: GetSpecificDataForDateSuunto :one
SELECT data->$3::text
FROM suunto_data
WHERE user_id = $1
AND summary_date = $2;

-- name: GetTypesFromCoachtechData :many
WITH cte AS (
    SELECT coachtech_id
    FROM coachtech_ids
    WHERE user_id = $1
)
SELECT DISTINCT data->>'testType', data->>'time', data->>'id'
FROM coachtech_data
WHERE summary_date = $2
  AND coachtech_id = (SELECT coachtech_id FROM cte)
  AND (
      ($3 IS NULL AND data->>'testType' IS NOT NULL) OR
      (data->>'testType' = $3)
  );

-- name: GetTypesFromOuraData :many
SELECT DISTINCT jsonb_object_keys(data)
FROM oura_data
WHERE user_id = @user_id
AND summary_date = @date;


-- name: GetTypesFromPolarData :many
SELECT DISTINCT jsonb_object_keys(data)
FROM polar_data
WHERE user_id = @user_id
AND summary_date = @date;

-- name: GetDataPointFromCoachtechData :many
WITH cte AS (
    SELECT coachtech_id
    FROM coachtech_ids
    WHERE user_id = $1
)
SELECT data
FROM coachtech_data
WHERE summary_date=$2
AND coachtech_id = (SELECT coachtech_id FROM cte);

-- name: GetAllDataForDateOura :one
SELECT data
FROM oura_data
WHERE user_id = $1
AND summary_date = $2;

-- name: GetSpecificDataForDateOura :one
SELECT data->$3::text
FROM oura_data
WHERE user_id = $1
AND summary_date = $2;


-- name: GetAllDataForDatePolar :one
SELECT data
FROM polar_data
WHERE user_id = $1
AND summary_date = $2;

-- name: GetUniqueCoachtechDataTypes :many
WITH cte AS (
    SELECT coachtech_id
    FROM coachtech_ids
    WHERE user_id = $1
)
SELECT DISTINCT data->>'testType'
FROM coachtech_data
WHERE coachtech_id = (SELECT coachtech_id FROM cte)
AND summary_date BETWEEN to_timestamp($2)::date AND to_timestamp($3)::date;

-- name: GetSpecificDataForDatePolar :one
SELECT data->$3::text
FROM polar_data
WHERE user_id = $1
AND summary_date = $2;

-- name: GetDatesFromGarminData :many
SELECT DISTINCT summary_date
FROM garmin_data
WHERE user_id = @user_id
AND (@after_date::date IS NULL OR summary_date >= @after_date)
AND (@before_date::date IS NULL OR summary_date <= @before_date)
ORDER BY summary_date DESC;

-- name: GetAllDataForDateGarmin :one
SELECT data
FROM garmin_data
WHERE user_id = $1
AND summary_date = $2;

-- name: GetSpecificDataForDateGarmin :one
SELECT data->$3::text
FROM garmin_data
WHERE user_id = $1
AND summary_date = $2;

-- name: GetTypesFromGarminData :many
SELECT DISTINCT jsonb_object_keys(data)
FROM garmin_data
WHERE user_id = @user_id
AND summary_date = @date;

-- name: InsertOuraData :exec
INSERT INTO oura_data (user_id, summary_date, data)
VALUES ($1, $2, $3)
ON CONFLICT (user_id, summary_date)
DO UPDATE SET data = EXCLUDED.data;

-- name: InsertPolarData :exec
INSERT INTO polar_data (user_id, summary_date, data)
VALUES ($1, $2, $3)
ON CONFLICT (user_id, summary_date)
DO UPDATE SET data = EXCLUDED.data;

-- name: InsertSuuntoData :exec
INSERT INTO suunto_data (user_id, summary_date, data)
VALUES ($1, $2, $3)
ON CONFLICT (user_id, summary_date)
DO UPDATE SET data = EXCLUDED.data;

-- name: InsertGarminData :exec
INSERT INTO garmin_data (user_id, summary_date, data)
VALUES ($1, $2, $3)
ON CONFLICT (user_id, summary_date)
DO UPDATE SET data = EXCLUDED.data;

-- name: DeleteGarminData :exec
DELETE FROM garmin_data
WHERE user_id = $1 AND summary_date = $2;

-- name: DeleteOuraData :exec
DELETE FROM oura_data
WHERE user_id = $1 AND summary_date = $2;

-- name: DeletePolarData :exec
DELETE FROM polar_data
WHERE user_id = $1 AND summary_date = $2;

-- name: DeleteSuuntoData :exec
DELETE FROM suunto_data
WHERE user_id = $1 AND summary_date = $2;
