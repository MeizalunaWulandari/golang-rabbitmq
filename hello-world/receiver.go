package main

import (
	"log"

	"github.com/pkg/errors"

	"github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672")
	if err != nil {
		panic(errors.Wrap(err, "Failed to connect"))
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		panic(errors.Wrap(err, "Failed to get Channel"))
	}
	defer ch.Close()

	q, err := ch.QueueDeclare("hello", false, false, false, false, nil)
	if err != nil {
		panic(errors.Wrap(err, "Failed to declare Channel"))
	}

	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	forever := make(chan struct{})

	log.Printf("[*] Waiting for messages, To exit press CTRL+C")
	<-forever
}
