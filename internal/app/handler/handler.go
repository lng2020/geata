package handler

type HandlerType string

const (
	ModbusHandlerType   HandlerType = "Modbus"
	MQTTHandlerType     HandlerType = "MQTT"
	IEC61850HandlerType HandlerType = "IEC61850"
)

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
	Handle(s chan string)
}
