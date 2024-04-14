package handler

import (
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MQTTHandler struct {
	client mqtt.Client
	topic  string
}

func NewMQTTHandler(broker, clientID, username, password, topic string) *MQTTHandler {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetClientID(clientID)
	opts.SetUsername(username)
	opts.SetPassword(password)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("Failed to connect to MQTT broker: %v", token.Error())
	}

	return &MQTTHandler{
		client: client,
		topic:  topic,
	}
}

func (h *MQTTHandler) Handle(s chan string) {
	token := h.client.Subscribe(h.topic, 0, func(client mqtt.Client, msg mqtt.Message) {
		s <- string(msg.Payload())
	})
	token.Wait()
	if token.Error() != nil {
		log.Printf("Failed to subscribe to topic: %v", token.Error())
		return
	}

	select {}
}

func (h *MQTTHandler) Close() {
	h.client.Disconnect(250)
}

func (h *MQTTHandler) IsOnline() bool {
	return h.client.IsConnected()
}

func (h *MQTTHandler) Type() HandlerType {
	return MQTTHandlerType
}