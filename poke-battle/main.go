package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
	"poke-battle/api"
	"poke-battle/db/sqlc"
	"poke-battle/event"
	"poke-battle/pb"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	dbUrl := os.Getenv("DATABASE_URL")
	connPool, err := pgxpool.New(context.Background(), dbUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	sqlStore := db.NewSQLStore(connPool)
	pokeApiClient := api.NewHTTPClient("https://pokeapi.co/api/v2/")

	pokeBattleStatsAddr := os.Getenv("POKE_BATTLE_STATS_ADDR")
	gRPCConn, err := grpc.Dial(pokeBattleStatsAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to gRPC server: %v\n", err)
		os.Exit(1)
	}
	gRPCClient := pb.NewPokeStatsClient(gRPCConn)

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to RabbitMQ: %v\n", err)
		os.Exit(1)
	}

	ch, err := conn.Channel()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fail to open a channel: %v\n", err)
		os.Exit(1)
	}

	q, err := ch.QueueDeclare(
		"longbattle",
		false,
		false,
		false,
		false,
		nil,
	)

	sender := event.NewRabbitMQSender(ch, q.Name)

	server := api.NewServer(sqlStore, pokeApiClient, gRPCClient, sender)
	serverAddr := os.Getenv("POKE_BATTLE_ADDR")
	fmt.Printf("Started poke battle on %s\n", serverAddr)
	err = server.Start(serverAddr)
	if err != nil {
		println("Fatal error")
	}
}
