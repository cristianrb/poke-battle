package main

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"os"
	"poke-long-battle/event"
)

func main() {
	rabbitMQUrl := os.Getenv("RABBITMQ_URL")
	conn, err := amqp.Dial(rabbitMQUrl)
	if err != nil {
		log.Fatalf("failed to connect to RabbitMQ: %s", err.Error())
	}
	defer conn.Close()

	consumer := event.NewConsumer(conn)
	fmt.Printf("Started poke long battle\n")
	if err := consumer.Listen(); err != nil {
		log.Fatalf("failed to listen to RabbitMQ: %s", err.Error())
	}
}
