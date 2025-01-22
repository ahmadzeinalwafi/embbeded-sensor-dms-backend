package mqtt

import (
	"fmt"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
)

type MQTTConnection struct {
	Client mqtt.Client
}

func NewMQTTConnection(broker string) (*MQTTConnection, error) {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetClientID(fmt.Sprintf("go_mqtt_client_%s", uuid.New().String())) // Unique client ID
	opts.SetUsername("user")
	opts.SetPassword("password")
	opts.SetCleanSession(true)
	opts.SetKeepAlive(120 * time.Second)
	opts.SetAutoReconnect(true)
	opts.SetOnConnectHandler(func(client mqtt.Client) {
		log.Println("Successfully connected to MQTT broker")
	})
	opts.SetConnectionLostHandler(func(client mqtt.Client, err error) {
		log.Printf("Connection lost: %v", err)
	})

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return nil, fmt.Errorf("failed to connect to MQTT broker: %v", token.Error())
	}

	log.Println("Successfully connected to MQTT broker")
	return &MQTTConnection{
		Client: client,
	}, nil
}
