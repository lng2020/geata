package handler

import "context"

type HandlerType string

const (
	ModbusHandlerType   HandlerType = "Modbus"
	MQTTHandlerType     HandlerType = "MQTT"
	IEC61850HandlerType HandlerType = "IEC61850"
)

type Data struct {
	IEC61850Ref string
	Value       string
}

var HandlerTypes = []HandlerType{
	ModbusHandlerType,
	MQTTHandlerType,
	IEC61850HandlerType,
}

type HandlerConfig interface {
	NewHandler() Handler
	Type() HandlerType
}

type Handler interface {
	Type() HandlerType
	IsOnline() bool
	Handle(ctx context.Context, s chan Data)
}
