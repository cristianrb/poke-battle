package api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	param "github.com/oceanicdev/chi-param"
	"math/rand"
	"net/http"
	"poke-battle/db/sqlc"
	"poke-battle/event"
	"poke-battle/models"
	"poke-battle/pb"
)

type Server struct {
	router        *chi.Mux
	store         db.Store
	pokeApiClient HTTPClient
	gRPCClient    pb.PokeStatsClient
	sender        event.Sender
}

type PokemonWithError struct {
	Pokemon string
	Error   error
}

func NewServer(store db.Store, pokeApiClient HTTPClient, gRPCClient pb.PokeStatsClient, sender event.Sender) *Server {
	return &Server{
		store:         store,
		pokeApiClient: pokeApiClient,
		gRPCClient:    gRPCClient,
		sender:        sender,
	}
}

func (server *Server) setupRouter() {
	r := chi.NewRouter()
	r.Route("/api/pokebattle", func(r chi.Router) {
		r.Post("/", server.createBattle)
		r.Post("/long", server.createLongBattle)
		r.Get("/{id}", server.getBattle)
		r.Get("/", server.listBattles)
		r.Get("/stats", server.getStats)
	})

	server.router = r
}

func (server *Server) Start(address string) error {
	server.setupRouter()
	return http.ListenAndServe(address, server.router)
}

func (server *Server) createBattle(w http.ResponseWriter, r *http.Request) {
	pokesChannel := make(chan string, 2)

	for i := 0; i < 2; i++ {
		go func() {
			poke, err := server.getRandomPokemon()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			pokesChannel <- poke
		}()
	}

	var pokes []string
	for i := 0; i < 2; i++ {
		pokes = append(pokes, <-pokesChannel)
	}

	pokemonBattle, err := server.store.SaveBattle(pokes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	err = json.NewEncoder(w).Encode(*pokemonBattle)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (server *Server) getRandomPokemon() (string, error) {
	pokeApiResponse := models.Pokemon{}
	err := server.pokeApiClient.Do(fmt.Sprintf("pokemon/%d", rand.Intn(151)), &pokeApiResponse)
	if err != nil {
		return "", err
	}

	return pokeApiResponse.Name, nil
}

func (server *Server) getBattle(w http.ResponseWriter, r *http.Request) {
	battleId := chi.URLParam(r, "id")
	battle, err := server.store.GetBattle(battleId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	err = json.NewEncoder(w).Encode(battle)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (server *Server) listBattles(w http.ResponseWriter, r *http.Request) {
	limit, err := param.QueryInt32(r, "limit")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	offset, err := param.QueryInt32(r, "offset")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	battles, err := server.store.ListBattles(limit, offset*limit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(battles)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (server *Server) getStats(w http.ResponseWriter, r *http.Request) {
	pokemon, err := param.QueryString(r, "pokemon")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	req := pb.PokeStatsRequest{
		Pokemon: pokemon,
	}

	stats, err := server.gRPCClient.GetPokeStats(context.Background(), &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	err = json.NewEncoder(w).Encode(stats)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (server *Server) createLongBattle(w http.ResponseWriter, r *http.Request) {
	pokesChannel := make(chan string, 8)

	for i := 0; i < 8; i++ {
		go func() {
			poke, err := server.getRandomPokemon()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			pokesChannel <- poke
		}()
	}

	var pokes string
	for i := 0; i < 8; i++ {
		pokes = fmt.Sprintf("%s,%s", pokes, <-pokesChannel)
	}

	err := server.sender.Send(context.Background(), []byte(pokes))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
