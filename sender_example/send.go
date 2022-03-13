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
	q, _ := ch.QueueDeclare(
		"mail",
		false,
		false,
		false,
		false,
		nil,
	)
	b, _ := json.Marshal(map[string]interface{}{
		"to":      "example@email.com", // change this to test
		"subject": "Email confirmation!",
		"body": `<p>Please confirm your email by clicking the link below:</p>
				<p>https://example.com/</p>`,
	})
	ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        b,
		},
	)
}
