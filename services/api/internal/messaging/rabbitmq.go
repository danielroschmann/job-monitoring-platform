package messaging

import (
	"encoding/json"
	"fmt"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	JobsExchange = "jobs"

	RemoteOkQueue  = "remoteok.queue"
	JobIndexQueue  = "jobindex.queue"
	ItJobbankQueue = "itjobbank.queue"

	RemoteOkRoutingKey  = "collect.remoteok"
	JobIndexRoutingKey  = "collect.jobindex"
	ItJobbankRoutingKey = "collect.itjobbank"
)

type RabbitClient struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

type CollectorsConfig struct {
	Name       string
	Queue      string
	RoutingKey string
}

var Collectors = []CollectorsConfig{
	{
		Name:       "RemoteOk",
		Queue:      RemoteOkQueue,
		RoutingKey: RemoteOkRoutingKey,
	},
	{
		Name:       "Jobindex",
		Queue:      JobIndexQueue,
		RoutingKey: JobIndexRoutingKey,
	},
	{
		Name:       "It-jobbank",
		Queue:      ItJobbankQueue,
		RoutingKey: ItJobbankRoutingKey,
	},
}

func Connect() (*RabbitClient, error) {
	url := fmt.Sprintf(
		"amqp://%s:%s@%s:%s/",
		os.Getenv("RABBITMQ_USER"),
		os.Getenv("RABBITMQ_PASSWORD"),
		os.Getenv("RABBITMQ_HOST"),
		os.Getenv("RABBITMQ_PORT"),
	)

	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, err
	}

	if err := ch.Confirm(false); err != nil {
		ch.Close()
		conn.Close()
		return nil, err
	}

	return &RabbitClient{
		conn: conn,
		ch:   ch,
	}, nil
}

func (rc *RabbitClient) Close() error {
	if err := rc.ch.Close(); err != nil {
		return err
	}
	return rc.conn.Close()
}

func (rc *RabbitClient) DeclareExchange(name string, exchangeType string, durable bool) error {
	return rc.ch.ExchangeDeclare(name, exchangeType, durable, false, false, false, nil)
}
func (rc *RabbitClient) SetupInfrastructure() error {
	if err := rc.DeclareExchange(JobsExchange, "topic", true); err != nil {
		return err
	}

	for _, collector := range Collectors {
		if err := rc.DeclareQueue(collector.Queue, true, false); err != nil {
			return err
		}

		if err := rc.DeclareBinding(
			collector.Queue,
			collector.RoutingKey,
			JobsExchange,
		); err != nil {
			return err
		}
	}

	return nil

}

func (rc *RabbitClient) DeclareQueue(queueName string, durable bool, autoDelete bool) error {
	_, err := rc.ch.QueueDeclare(queueName, durable, autoDelete, false, false, nil)
	return err
}

func (rc *RabbitClient) DeclareBinding(queueName string, routingKey string, exchange string) error {
	return rc.ch.QueueBind(queueName, routingKey, exchange, false, nil)
}
