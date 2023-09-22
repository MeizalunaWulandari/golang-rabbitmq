package main

import (
	"github.com/pkg/errors"
	"fmt"
	"os"

	"github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672")
	if err != nil {
		panic(errors.Wrap(err, "Failed to connect"))
	}
	defer conn.Close()

	ch, err :=  conn.Channel()
	if err != nil {
	 panic(errors.Wrap(err, "Failed to get Channel"))
	}
	defer ch.Close()

	q, err := ch.QueueDeclare("hello", false, false, false,false, nil)
	if err != nil {
	  panic(errors.Wrap(err, "Failed to declare Channel"))
	}

	err = ch.Publish("", q.Name, false, false, amqp091.Publishing{
	  ContentType: "text/plain",
	  Body: []byte(os.Args[1]),
	  })
	if err != nil {
	 panic(errors.Wrap(err, "Failed to publish messages"))
	}

	fmt.Println("Send message:", os.Args[1])

}
