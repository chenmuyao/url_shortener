-- name: InsertURL :one
INSERT INTO "urls" (url) VALUES ($1) RETURNING *;

-- name: GetByID :one
SELECT (url)
FROM "urls"
WHERE "id" = ($1);

-- name: GetIDByURL :one
SELECT (id)
FROM "urls"
WHERE "url" = ($1);


