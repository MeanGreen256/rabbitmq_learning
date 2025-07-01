package main

import (
	"encoding/json"
	"log"
	"time"

	"rabbitmq-json-example/pkg/messaging"
	"rabbitmq-json-example/pkg/types"

	"github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp091.Dial(messaging.RabbitMQURL)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		messaging.QueueName, // name
		true,                // durable
		false,               // delete when unused
		false,               // exclusive
		false,               // no-wait
		nil,                 // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack (acknowledge message immediately)
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			var msg types.Message
			err := json.Unmarshal(d.Body, &msg)
			if err != nil {
				log.Printf("Error decoding JSON: %s", err)
				continue
			}
			log.Printf("Received a message from '%s': '%s' at %s", msg.Sender, msg.Content, msg.Timestamp.Format(time.RFC3339))
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

