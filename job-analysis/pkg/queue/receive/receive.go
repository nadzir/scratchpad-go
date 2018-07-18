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

func StartReceiver(channel chan []byte) {
	// Connect
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	// Close when function ends
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// redeclare queue to ensure it exists
	// It is declared in sender.go
	q, err := ch.QueueDeclare(
		"events", // name
		true,     // durable
		false,    // delete when unused
		false,    // exclusive
		false,    // no-wait
		nil,      // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			channel <- d.Body
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

// StartReceiver(channel chan []byte) {
// 	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
// 	failOnError(err, "Failed to connect to RabbitMQ")
// 	defer conn.Close()

// 	ch, err := conn.Channel()
// 	failOnError(err, "Failed to open a channel")
// 	defer ch.Close()

// 	err = ch.ExchangeDeclare(
// 		"events", // name
// 		"fanout", // type
// 		true,     // durable
// 		false,    // auto-deleted
// 		false,    // internal
// 		false,    // no-wait
// 		nil,      // arguments
// 	)
// 	failOnError(err, "Failed to declare an exchange")

// 	q, err := ch.QueueDeclare(
// 		"stats", // name
// 		true,    // durable
// 		false,   // delete when usused
// 		false,   // exclusive
// 		false,   // no-wait
// 		nil,     // arguments
// 	)
// 	failOnError(err, "Failed to declare a queue")

// 	err = ch.QueueBind(
// 		q.Name,   // queue name
// 		"",       // routing key
// 		"events", // exchange
// 		false,
// 		nil)
// 	failOnError(err, "Failed to bind a queue")

// 	msgs, err := ch.Consume(
// 		q.Name, // queue
// 		"",     // consumer
// 		true,   // auto-ack
// 		true,   // exclusive
// 		false,  // no-local
// 		false,  // no-wait
// 		nil,    // args
// 	)
// 	failOnError(err, "Failed to register a consumer")

// 	forever := make(chan bool)

// 	go func() {
// 		for d := range msgs {
// 			channel <- d.Body
// 		}
// 	}()

// 	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
// 	<-forever
// }
