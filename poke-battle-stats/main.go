package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"net"
	"os"
	db "poke-stats/db/sqlc"
	server "poke-stats/server"
)

var addr = "0.0.0.0:50050"

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}

	dbUrl := os.Getenv("DATABASE_URL")
	connPool, err := pgxpool.New(context.Background(), dbUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	sqlStore := db.NewSQLStore(connPool)

	s := server.NewServer(sqlStore)
	err = s.Start(lis)
	if err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
