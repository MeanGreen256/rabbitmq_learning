package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"strings"
	"time"

	"rabbitmq_learning/pkg/messaging"
	"rabbitmq_learning/pkg/types"

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

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	messageContent := "Hello World from Go!"
	if len(os.Args) > 1 {
		messageContent = strings.Join(os.Args[1:], " ")
	}

	msg := types.Message{
		Sender:    "ClientApp",
		Content:   messageContent,
		Timestamp: time.Now(),
	}

	body, err := json.Marshal(msg)
	failOnError(err, "Failed to marshal JSON")

	err = ch.PublishWithContext(ctx,
		"",     // exchange (use default)
		q.Name, // routing key (queue name)
		false,  // mandatory
		false,  // immediate
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	failOnError(err, "Failed to publish a message")

	log.Printf(" [x] Sent message: %s", msg.Content)
}
