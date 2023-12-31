// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: pokemon_battle.sql

package db

import (
	"context"
)

const getPokeBattlesByPokemon = `-- name: GetPokeBattlesByPokemon :many
SELECT id, pokemon1, pokemon2, winner FROM pokemon_battle
WHERE pokemon1 = $1 OR pokemon2 = $2
`

type GetPokeBattlesByPokemonParams struct {
	Pokemon1 string
	Pokemon2 string
}

func (q *Queries) GetPokeBattlesByPokemon(ctx context.Context, arg GetPokeBattlesByPokemonParams) ([]PokemonBattle, error) {
	rows, err := q.db.Query(ctx, getPokeBattlesByPokemon, arg.Pokemon1, arg.Pokemon2)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []PokemonBattle
	for rows.Next() {
		var i PokemonBattle
		if err := rows.Scan(
			&i.ID,
			&i.Pokemon1,
			&i.Pokemon2,
			&i.Winner,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
