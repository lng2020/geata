package handler

import (
	"context"
	"geata/internal/app/logger"
	"log/slog"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MQTTHandler struct {
	client mqtt.Client
	topic  string
}

type MQTTHandlerConfig struct {
	Broker   string
	ClientID string
	Username string
	Password string
}

func (c MQTTHandlerConfig) Type() HandlerType {
	return MQTTHandlerType
}

func NewMQTTHandlerConfig(broker, clientID, username, password, topic string) HandlerConfig {
	return &MQTTHandlerConfig{
		Broker:   broker,
		ClientID: clientID,
		Username: username,
		Password: password,
	}
}

func (hc *MQTTHandlerConfig) NewHandler() Handler {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(hc.Broker)
	opts.SetClientID(hc.ClientID)
	opts.SetUsername(hc.Username)
	opts.SetPassword(hc.Password)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		slog.Error("Failed to connect to MQTT broker", logger.ErrAttr(token.Error()))
	}

	return &MQTTHandler{
		client: client,
	}
}

func (h *MQTTHandler) Handle(ctx context.Context, s chan Data) {
	token := h.client.Subscribe(h.topic, 0, func(client mqtt.Client, msg mqtt.Message) {
		ref := msg.Topic()
		value := string(msg.Payload())
		s <- Data{IEC61850Ref: ref, Value: value}
	})
	token.Wait()
	if token.Error() != nil {
		slog.Error("Failed to subscribe to topic", logger.ErrAttr(token.Error()))
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
