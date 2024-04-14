package handler

type HandlerType string

const (
	ModbusHandlerType   HandlerType = "Modbus"
	MQTTHandlerType     HandlerType = "MQTT"
	IEC61850HandlerType HandlerType = "IEC61850"
)

type HandlerFactory func() Handler

var supportedHandler = map[HandlerType]HandlerFactory{}

func RegisterHandler(ht HandlerType, hf HandlerFactory) {
	supportedHandler[ht] = hf
}

// Handlers represents a collection of handlers
type Handlers struct {
	Handlers []*Handler
}

type Handler interface {
	// TODO: add more methods
	Type() HandlerType
	IsOnline() bool
	Handle(s chan string)
}
