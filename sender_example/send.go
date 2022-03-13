// JUST A EXAMPLE

package main

import (
	"encoding/json"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	sendMessageBroker()
}
func sendMessageBroker() {
	conn, _ := amqp.Dial("amqp://root:root@localhost:5672")
	defer conn.Close()
	ch, _ := conn.Channel()
	defer ch.Close()
	// q, _ := ch.QueueDeclare(
	// 	"mail",       // name
	// 	true,         // durable
	// 	false,        // delete when unused
	// 	false,        // exclusive
	// 	false,        // no-wait
	// 	nil,          // arguments
	// )
	b, _ := json.Marshal(map[string]interface{}{
		"to":      "[example@email.com]", // change this to test
		"subject": "Email confirmation!",
		"body": `<p>Please confirm your email by clicking the link below:</p>
				<p>https://example.com/</p>`,
	})
	ch.Publish(
		"",
		"mail",
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "application/json",
			Body:         b,
		},
	)
}
