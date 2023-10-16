package models

type PokeStats struct {
	Pokemon string  `json:"pokemon"`
	Wins    int64   `json:"wins"`
	Loses   int64   `json:"loses"`
	Winrate float32 `json:"winrate"`
}
