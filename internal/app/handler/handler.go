package handler

type HandlerType string

const (
	ModbusHandlerType   HandlerType = "Modbus"
	MQTTHandlerType     HandlerType = "MQTT"
	IEC61850HandlerType HandlerType = "IEC61850"
)

type HandlerConfig interface{}

type HandlerFactory func(HandlerConfig) Handler

var supportedHandler = map[HandlerType]HandlerFactory{}

func RegisterHandler(ht HandlerType, hf HandlerFactory) {
	supportedHandler[ht] = hf
}

// Handlers represents a collection of handlers
type Handlers struct {
	Handlers []*Handler
}

type Handler interface {
	Type() HandlerType
	IsOnline() bool
	Handle(s chan string)
}
