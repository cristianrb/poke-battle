-- name: GetPokeBattlesByPokemon :many
SELECT * FROM pokemon_battle
WHERE pokemon1 = $1 OR pokemon2 = $2;
