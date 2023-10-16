package db

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"math/rand"
	"poke-battle/models"
)

type Store interface {
	//Querier
	SaveBattle(pokes []string) (*models.PokemonBattle, error)
	ListBattles(limit, offset int32) ([]PokemonBattle, error)
	GetBattle(id string) (PokemonBattle, error)
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

func (store *SQLStore) SaveBattle(pokes []string) (*models.PokemonBattle, error) {
	args := CreatePokeBattleParams{
		Pokemon1: pokes[0],
		Pokemon2: pokes[1],
		Winner: pgtype.Text{
			String: pokes[rand.Intn(2)],
			Valid:  true,
		},
	}

	battle, err := store.Queries.CreatePokeBattle(context.Background(), args)
	if err != nil {
		return nil, err
	}
	return &models.PokemonBattle{
		ID:       battle.ID,
		Pokemon1: battle.Pokemon1,
		Pokemon2: battle.Pokemon2,
	}, nil
}

func (store *SQLStore) ListBattles(limit, offset int32) ([]PokemonBattle, error) {
	args := ListPokeBattlesParams{
		Limit:  limit,
		Offset: offset,
	}
	return store.Queries.ListPokeBattles(context.Background(), args)
}

func (store *SQLStore) GetBattle(id string) (PokemonBattle, error) {
	uuid, err := uuid.Parse(id)
	if err != nil {
		return PokemonBattle{}, err
	}
	return store.Queries.GetPokeBattle(context.Background(), uuid)
}
