package models

import "github.com/google/uuid"

type Pokemon struct {
	Name string `json:"name"`
}

type PokemonBattle struct {
	ID       uuid.UUID
	Pokemon1 string `json:"pokemon1"`
	Pokemon2 string `json:"pokemon2"`
}
