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

type HandlerConfig interface{}

type HandlerFactory func(HandlerConfig) Handler

var SupportedHandler = map[HandlerType]HandlerFactory{}

func RegisterHandler(ht HandlerType, hf HandlerFactory) {
	SupportedHandler[ht] = hf
}

type Handler interface {
	Type() HandlerType
	IsOnline() bool
	Handle(s chan string)
}
