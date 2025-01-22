package main

import (
	rabbitmq "dms/internal/infrastructure/messaging/rabbitmq" // Import the rabbitmq package
	"log"
	"os"
)

func main() {
	rabbitMQConfig := rabbitmq.RabbitMQConfig{
		DSN:      "amqp://guest:guest@localhost:5672/",
		Exchange: "delayed_exchange",
		Queue:    "delayed_queue",
	}

	rabbitMQClient, err := rabbitmq.NewRabbitMQConnection(rabbitMQConfig)
	if err != nil {
		log.Fatalf("Failed to create RabbitMQ client: %s", err)
		os.Exit(1)
	}

	delayedMessageProducer := rabbitmq.NewSchedulingMessage(rabbitMQClient)

	if err := rabbitMQClient.SetupExchange(); err != nil {
		log.Fatalf("Failed to declare an exchange: %s", err)
	}
	if err := rabbitMQClient.SetupQueue(); err != nil {
		log.Fatalf("Failed to declare a queue: %s", err)
	}
	if err := rabbitMQClient.BindQueue(); err != nil {
		log.Fatalf("Failed to bind the queue: %s", err)
	}

	messageData := rabbitmq.MessageData{
		DelayTime: 8,
		Device_Id: "This is a delayed JSON message!",
	}

	if err := delayedMessageProducer.PublishMessage(messageData, 5000); err != nil {
		log.Fatalf("Failed to publish a message: %s", err)
	} else {
		log.Println("JSON message sent successfully!")
	}
}
