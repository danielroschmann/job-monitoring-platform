package messaging

import (
	"encoding/json"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type CollectJobMessage struct {
	Source      string    `json:"source"`
	RequestedBy uint      `json:"requested_by"`
	RequestedAt time.Time `json:"requested_at"`
}

func (rc *RabbitClient) PublishMessage(exchange string, routingKey string, payload any) error {
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	return rc.ch.Publish(exchange, routingKey, false, false, amqp.Publishing{

		ContentType: "application/json",
		Body:        []byte(payloadBytes),
	})

}
