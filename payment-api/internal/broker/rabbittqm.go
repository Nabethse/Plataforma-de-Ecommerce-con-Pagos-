package broker

import (
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func NewRabbitMQ() (*RabbitMQ, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &RabbitMQ{conn, ch}, nil
}

func (r *RabbitMQ) Publish(queue string, message interface{}) error {
	body, _ := json.Marshal(message)

	err := r.channel.Publish(
		"",
		queue,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		log.Println("Error al publicar mensaje:", err)
	}
	return err
}

func (r *RabbitMQ) Close() {
	r.channel.Close()
	r.conn.Close()
}