package handler

import (
	"fmt"
	"time"

	"github.com/simonvetter/modbus"
)

// ModbusHandler represents your Modbus handler.
type ModbusHandler struct {
	client *modbus.ModbusClient
}

// NewModbusHandler creates a new instance of ModbusHandler.
func NewModbusHandler(endpoint string) (*ModbusHandler, error) {
	var err error

	// Create a Modbus client for the specified endpoint
	client, err := modbus.NewClient(&modbus.ClientConfiguration{
		URL:     endpoint,
		Timeout: 1 * time.Second,
	})
	// Note: Use udp:// for Modbus TCP over UDP

	if err != nil {
		return nil, fmt.Errorf("error creating Modbus client: %v", err)
	}

	// Attempt to connect
	err = client.Open()
	if err != nil {
		return nil, fmt.Errorf("error opening Modbus connection: %v", err)
	}

	return &ModbusHandler{
		client: client,
	}, nil
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
