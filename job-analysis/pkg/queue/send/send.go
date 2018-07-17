package queue

import (
	"log"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

const (
	//MqHost rabbit mq host url
	MqHost = "amqp://guest:guest@localhost:5672/"
)

func Send(msg []byte) {
	// Connect to rabbitmq
	conn, err := amqp.Dial(MqHost)

	// Handle error
	failOnError(err, "Failed to connect to RabbitMQ")
	// Close when function exits
	defer conn.Close()

	// Open a channel
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// Declare queue
	q, err := ch.QueueDeclare(
		"events", // name
		false,    // durable
		false,    // delete when unused
		false,    // exclusive
		false,    // no-wait
		nil,      // arguments
	)
	failOnError(err, "Failed to declare a queue")

	// Send message to queue
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        msg,
		})
	// log.Printf(" [x] Sent %s", msg)
	failOnError(err, "Failed to publish a message")
}
