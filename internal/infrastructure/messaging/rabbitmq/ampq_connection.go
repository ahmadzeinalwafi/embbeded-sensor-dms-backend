package rabbitmq

import (
	"log"

	"github.com/streadway/amqp"
)

type RabbitMQConfig struct {
	DSN      string
	Exchange string
	Queue    string
}

type RabbitMQClient struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
	Config     RabbitMQConfig
}

type MessageData struct {
	DelayTime int `json:"DelayTime"`
	Device_Id string `json:"Device_Id"`
}

func NewRabbitMQConnection(config RabbitMQConfig) (*RabbitMQClient, error) {
	conn, err := amqp.Dial(config.DSN)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %s", err)
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %s", err)
		return nil, err
	}

	return &RabbitMQClient{
		Connection: conn,
		Channel:    ch,
		Config:     config,
	}, nil
}

func (r *RabbitMQClient) SetupExchange() error {
	return r.Channel.ExchangeDeclare(
		r.Config.Exchange,   // Exchange name
		"x-delayed-message", // Exchange type
		true,                // Durable
		false,               // Auto-delete
		false,               // Internal
		false,               // No-wait
		amqp.Table{
			"x-delayed-type": "direct", // The underlying exchange type
		},
	)
}

// SetupQueue declares the queue in RabbitMQ.
func (r *RabbitMQClient) SetupQueue() error {
	_, err := r.Channel.QueueDeclare(
		r.Config.Queue, // Queue name
		true,           // Durable
		false,          // Auto-delete
		false,          // Exclusive
		false,          // No-wait
		nil,            // Arguments
	)
	return err
}

// BindQueue binds the queue to the exchange.
func (r *RabbitMQClient) BindQueue() error {
	return r.Channel.QueueBind(
		r.Config.Queue,    // Queue name
		"",                // Routing key
		r.Config.Exchange, // Exchange name
		false,             // No-wait
		nil,               // Arguments
	)
}
