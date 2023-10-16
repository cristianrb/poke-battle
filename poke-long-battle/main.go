package main

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"poke-long-battle/event"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("failed to connect to RabbitMQ: %s", err.Error())
	}
	defer conn.Close()

	consumer := event.NewConsumer(conn)
	if err := consumer.Listen(); err != nil {
		log.Fatalf("failed to listen to RabbitMQ: %s", err.Error())
	}
}
