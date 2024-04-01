package server

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	db "poke-stats/db/sqlc"
	"poke-stats/pb"
)

type Server struct {
	pb.PokeStatsServer
	gRPCServer *grpc.Server
	store      db.Store
}

func NewServer(store db.Store) *Server {
	server := &Server{}
	s := grpc.NewServer()
	pb.RegisterPokeStatsServer(s, server)
	reflection.Register(s)

	server.gRPCServer = s
	server.store = store
	return server
}

func (s *Server) Start(listener net.Listener) error {
	return s.gRPCServer.Serve(listener)
}

func (s *Server) GetPokeStats(ctx context.Context, in *pb.PokeStatsRequest) (*pb.PokeStatsResponse, error) {
	pokeStats, err := s.store.GetPokeBattlesByPokemon(in.Pokemon)
	if err != nil {
		return nil, err
	}

	println(pokeStats.Winrate)
	return &pb.PokeStatsResponse{
		Pokemon: pokeStats.Pokemon,
		Wins:    pokeStats.Wins,
		Loses:   pokeStats.Loses,
		Winrate: pokeStats.Winrate,
	}, nil
}
