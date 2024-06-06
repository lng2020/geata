package handler

import (
	"context"
	"encoding/json"
	"geata/internal/app/logger"
	"log/slog"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Message struct {
	Ref   string `json:"ref"`
	Value string `json:"value"`
}

type MQTTHandler struct {
	client  mqtt.Client
	topic   string
	ModelID int64
}

type MQTTHandlerConfig struct {
	Broker   string
	ClientID string
	Username string
	Password string
	Topic    string
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
		Topic:    topic,
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
		return nil
	}
	return &MQTTHandler{
		client: client,
		topic:  hc.Topic,
	}
}

func (h *MQTTHandler) Handle(ctx context.Context, s chan Data) {
	token := h.client.Subscribe(h.topic, 0, func(client mqtt.Client, msg mqtt.Message) {
		jsonMsg := Message{}
		if err := json.Unmarshal(msg.Payload(), &jsonMsg); err != nil {
			slog.Error("Failed to unmarshal message", logger.ErrAttr(err))
			return
		}
		ref := jsonMsg.Ref
		value := jsonMsg.Value
		slog.Info("Received message", logger.StringAttr("topic", ref), logger.StringAttr("value", value))
		s <- Data{IEC61850Ref: ref, Value: value, DataSource: "MQTT"}
	})
	token.Wait()
	if token.Error() != nil {
		slog.Error("Failed to subscribe to topic", logger.ErrAttr(token.Error()))
		return
	}

	<-ctx.Done()

	if token := h.client.Unsubscribe(h.topic); token.Wait() && token.Error() != nil {
		slog.Error("Failed to unsubscribe from topic", logger.ErrAttr(token.Error()))
	}
	h.client.Disconnect(250)
}

func (h *MQTTHandler) IsOnline() bool {
	return h.client.IsConnected()
}

func (h *MQTTHandler) Type() HandlerType {
	return MQTTHandlerType
}
