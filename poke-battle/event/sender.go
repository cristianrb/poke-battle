package event

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Sender interface {
	Send(ctx context.Context, data []byte) error
}

type RabbitMQSender struct {
	channel   *amqp.Channel
	queueName string
}

func NewRabbitMQSender(ch *amqp.Channel, queueName string) *RabbitMQSender {
	return &RabbitMQSender{
		channel:   ch,
		queueName: queueName,
	}
}

func (s *RabbitMQSender) Send(ctx context.Context, data []byte) error {
	return s.channel.PublishWithContext(
		ctx,
		"",
		s.queueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        data,
		},
	)
}
