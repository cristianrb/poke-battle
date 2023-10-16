package db

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"poke-stats/models"
)

type Store interface {
	//Querier
	GetPokeBattlesByPokemon(pokemon string) (models.PokeStats, error)
}

type SQLStore struct {
	*Queries
	connPool *pgxpool.Pool
}

func NewSQLStore(connPool *pgxpool.Pool) *SQLStore {
	return &SQLStore{
		connPool: connPool,
		Queries:  New(connPool),
	}
}

func (store *SQLStore) GetPokeBattlesByPokemon(pokemon string) (models.PokeStats, error) {

	battles, err := store.Queries.GetPokeBattlesByPokemon(context.Background(), GetPokeBattlesByPokemonParams{
		Pokemon1: pokemon,
		Pokemon2: pokemon,
	})
	if err != nil {
		return models.PokeStats{}, err
	}

	wins, loses := 0, 0
	for _, battle := range battles {
		if battle.Winner.String == pokemon {
			wins += 1
		} else {
			loses += 1
		}
	}
	winrate := float32(wins) / float32(wins+loses) * 100

	return models.PokeStats{
		Pokemon: pokemon,
		Wins:    int64(wins),
		Loses:   int64(loses),
		Winrate: winrate,
	}, nil
}
