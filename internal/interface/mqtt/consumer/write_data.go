package mqtt

import (
	"context"
	"encoding/json"
	"log"

	usecase "dms/internal/application/usecases"
	entity "dms/internal/domain/entities"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// MQTTConsumer defines the interface for consuming MQTT messages.
type MQTTConsumer struct {
	Client    mqtt.Client
	Topic     string
	Processor usecase.MessageProcessorUseCase
}

// NewMQTTConsumer initializes a new MQTT consumer.
func NewMQTTConsumer(client mqtt.Client, topic string, processor usecase.MessageProcessorUseCase) *MQTTConsumer {
	return &MQTTConsumer{
		Client:    client,
		Topic:     topic,
		Processor: processor,
	}
}

// Start starts the MQTT consumer and subscribes to the topic.
func (m *MQTTConsumer) Start() error {
	// Subscribe to the topic
	if token := m.Client.Subscribe(m.Topic, 2, m.messageHandler()); token.Wait() && token.Error() != nil {
		return token.Error()
	}

	log.Println("Subscribed to topic:", m.Topic)
	return nil
}

// messageHandler processes the incoming messages.
func (m *MQTTConsumer) messageHandler() func(mqtt.Client, mqtt.Message) {
	return func(client mqtt.Client, msg mqtt.Message) {
		var messageData entity.HistoricalDeviceRecords
		err := json.Unmarshal(msg.Payload(), &messageData)
		if err != nil {
			log.Printf("Error unmarshalling message: %v", err)
			return
		}

		err = m.Processor.ProcessMessage(context.Background(), messageData)
		if err != nil {
			log.Printf("Error processing message: %v", err)
		}
	}
}
