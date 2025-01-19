-- name: InsertURL :one
INSERT INTO "urls" (id, url, created_at, count) VALUES ($1, $2, $3, $4) RETURNING *;

-- name: UpdateCountByID :one
UPDATE "urls"
SET count = urls.count + 1
WHERE "id" = ($1)
RETURNING *;

-- name: GetIDByURL :one
SELECT (id)
FROM "urls"
WHERE "url" = ($1);


