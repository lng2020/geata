package handler

// Handlers represents a collection of handlers
type Handlers struct {
	Modbus *ModbusHandler
	// MQTT     *handler.MqttHandler
	// IEC61850 *handler.Iec61850Handler
}
