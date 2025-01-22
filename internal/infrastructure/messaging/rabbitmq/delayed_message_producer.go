package rabbitmq

import (
	"encoding/json"

	"github.com/streadway/amqp"
)

func NewSchedulingMessage(client *RabbitMQClient) *SchedulingMessageProducer {
	return &SchedulingMessageProducer{
		Channel: client.Channel,
		Config:  client.Config,
	}
}

type SchedulingMessageProducer struct {
	Channel *amqp.Channel
	Config  RabbitMQConfig
}

func (p *SchedulingMessageProducer) PublishMessage(message MessageData, delay int) error {
	jsonData, err := json.Marshal(message)
	if err != nil {
		return err
	}

	messageProperties := amqp.Table{
		"x-delay": int32(delay),
	}

	// Publish the delayed message
	return p.Channel.Publish(
		p.Config.Exchange, // Exchange name
		"",                // Routing key
		false,             // Mandatory
		false,             // Immediate
		amqp.Publishing{
			ContentType: "application/json", // Content type is JSON
			Body:        jsonData,           // Serialized JSON data
			Headers:     messageProperties,  // Delayed message properties
		},
	)
}
