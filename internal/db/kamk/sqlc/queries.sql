-- name: InsertInjury :exec
INSERT INTO public.injuries (
  user_id, injury_type, severity, pain_level, description, date_start, status, injury_id, meta
) VALUES (
  $1, $2, $3, $4, $5, NOW(), 0, $6, $7
);


-- name: MarkInjuryRecoveredByID :exec
UPDATE public.injuries
SET status = 1,
    date_end = NOW()
WHERE injury_id = $1
  AND user_id = $2;

-- name: GetActiveInjuriesByUser :many
SELECT
  user_id,
  injury_type,
  severity,
  pain_level,
  description,
  date_start,
  status,
  date_end,
  injury_id,
  meta
FROM public.injuries
WHERE user_id = $1
  AND status = 0
ORDER BY date_start DESC;

-- name: GetMaxInjuryIDForUser :one
SELECT COALESCE(MAX(injury_id), 0)::int4 AS id
FROM public.injuries
WHERE user_id = $1;


-- name: InsertQuestionnaire :one
INSERT INTO public.querys (
  user_id, query_type, answers, comment, "timestamp", meta
) VALUES (
  $1, $2, $3, $4, NOW(), $5
)
RETURNING id;

-- name: GetQuestionnairesByUser :many
SELECT
  id,
  user_id,
  query_type,
  answers,
  comment,
  "timestamp",
  meta
FROM public.querys
WHERE user_id = $1
ORDER BY "timestamp" DESC;

-- name: IsQuizDoneToday :many
SELECT
  id,
  user_id,
  query_type,
  answers,
  comment,
  "timestamp",
  meta
FROM public.querys
WHERE user_id = $1
  AND query_type    = $2
  AND "timestamp"  >= $3
  AND "timestamp"  <  $4
ORDER BY "timestamp" DESC;

-- name: UpdateQuestionnaireByID :execrows
UPDATE public.querys
SET answers = $3,
    comment = $4
WHERE user_id = $1
  AND id           = $2;

-- name: DeleteInjuryByID :execrows
DELETE FROM public.injuries
WHERE user_id = $1
  AND injury_id     = $2;

-- name: DeleteQuestionnaireByID :execrows
DELETE FROM public.querys
WHERE user_id = $1
  AND id           = $2;