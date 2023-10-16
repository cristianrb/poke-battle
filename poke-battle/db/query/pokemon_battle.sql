-- name: CreatePokeBattle :one
INSERT INTO pokemon_battle (pokemon1, pokemon2, winner) VALUES ($1, $2, $3) RETURNING *;

-- name: ListPokeBattles :many
SELECT * FROM pokemon_battle LIMIT $1 OFFSET $2;

-- name: GetPokeBattle :one
SELECT * FROM pokemon_battle WHERE id = $1;