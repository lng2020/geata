package handler

import (
	"fmt"
	"log"

	"github.com/simonvetter/modbus"
)

// ModbusHandler represents your Modbus handler.
type ModbusHandler struct {
	client *modbus.ModbusClient
}

type ModbusHandlerConfig struct {
	URL string
}

func NewModbusHandler(hc HandlerConfig) Handler {
	config := hc.(ModbusHandlerConfig)
	client, err := modbus.NewClient(&modbus.ClientConfiguration{URL: config.URL})
	if err != nil {
		log.Fatal(err)
	}
	return &ModbusHandler{client: client}
}

func (h *ModbusHandler) Handle(s chan string) {
}

// ReadHoldingRegister reads a single holding register at the specified address.
func (h *ModbusHandler) ReadHoldingRegister(address uint16) (uint16, error) {
	reg, err := h.client.ReadRegister(address, modbus.HOLDING_REGISTER)
	if err != nil {
		return 0, fmt.Errorf("error reading holding register: %v", err)
	}
	return reg, nil
}

// WriteHoldingRegister writes a value to a holding register at the specified address.
func (h *ModbusHandler) WriteHoldingRegister(address, value uint16) error {
	err := h.client.WriteRegister(address, value)
	if err != nil {
		return fmt.Errorf("error writing holding register: %v", err)
	}
	return nil
}

// Close closes the Modbus connection.
func (h *ModbusHandler) Close() {
	if h.client != nil {
		h.client.Close()
	}
}

func (h *ModbusHandler) IsOnline() bool {
	return h.client != nil
}

func (h *ModbusHandler) Type() HandlerType {
	return ModbusHandlerType
}
