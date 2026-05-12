-- name: GetHero :one
SELECT * FROM heroes
WHERE id = ? LIMIT 1;

-- name: ListHeroes :many
SELECT * FROM heroes
ORDER BY created_at DESC;

-- name: CreateHero :one
INSERT INTO heroes (
    id, ancestry_name, class_name, motivation, origin, background_name, quirks
) VALUES (
    ?, ?, ?, ?, ?, ?, ?
)
RETURNING *;

-- name: DeleteHero :exec
DELETE FROM heroes
WHERE id = ?;
