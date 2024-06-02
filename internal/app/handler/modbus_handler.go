package handler

import (
	"context"
	"fmt"
	"geata/internal/app/logger"
	"log/slog"

	"github.com/simonvetter/modbus"
)

type ModbusHandler struct {
	client *modbus.ModbusClient
}

type ModbusHandlerConfig struct {
	URL string
}

func (c ModbusHandlerConfig) Type() HandlerType {
	return ModbusHandlerType
}

func NewModbusHandlerConfig(url string) HandlerConfig {
	return &ModbusHandlerConfig{
		URL: url,
	}
}

func (hc *ModbusHandlerConfig) NewHandler() Handler {
	client, err := modbus.NewClient(&modbus.ClientConfiguration{URL: hc.URL})
	if err != nil {
		slog.Error("Failed to create modbus client", logger.ErrAttr(err))
	}
	return &ModbusHandler{client: client}
}

func (h *ModbusHandler) Handle(ctx context.Context, s chan Data) {
	slog.Info("ModbusHandler is not implemented")
}

func (h *ModbusHandler) ReadHoldingRegister(address uint16) (uint16, error) {
	reg, err := h.client.ReadRegister(address, modbus.HOLDING_REGISTER)
	if err != nil {
		return 0, fmt.Errorf("error reading holding register: %v", err)
	}
	return reg, nil
}

func (h *ModbusHandler) WriteHoldingRegister(address, value uint16) error {
	err := h.client.WriteRegister(address, value)
	if err != nil {
		return fmt.Errorf("error writing holding register: %v", err)
	}
	return nil
}

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
