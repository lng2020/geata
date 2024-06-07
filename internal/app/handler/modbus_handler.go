package handler

import (
	"context"
	"fmt"
	"geata/internal/app/logger"
	"geata/internal/app/model"
	"log/slog"

	"github.com/simonvetter/modbus"
	"xorm.io/xorm"
)

type ModbusHandler struct {
	client  *modbus.ModbusClient
	ModelID int64
	Engine  *xorm.Engine
}

type ModbusHandlerConfig struct {
	URL     string
	ModelID int64
	Engine  *xorm.Engine
}

func (c ModbusHandlerConfig) Type() HandlerType {
	return ModbusHandlerType
}

func NewModbusHandlerConfig(url string, id int64, e *xorm.Engine) HandlerConfig {
	return &ModbusHandlerConfig{
		URL:     url,
		ModelID: id,
		Engine:  e,
	}
}

func (hc *ModbusHandlerConfig) NewHandler() Handler {
	client, err := modbus.NewClient(&modbus.ClientConfiguration{URL: hc.URL})
	if err != nil {
		slog.Error("Failed to create modbus client", logger.ErrAttr(err))
		return nil
	}
	return &ModbusHandler{client: client, ModelID: hc.ModelID, Engine: hc.Engine}
}

func (h *ModbusHandler) Handle(ctx context.Context, s chan Data) {
	for {
		err := h.client.Open()
		if err != nil {
			slog.Error("Failed to open modbus connection", logger.ErrAttr(err))
			break
		}
		details := make(map[string]*model.ModbusDetail)
		res, err := model.GetAllModbusRulesByModelID(h.Engine, h.ModelID)
		if err != nil {
			slog.Error("Failed to get modbus rules", logger.ErrAttr(err))
			return
		}
		for _, rule := range res {
			detail, err := model.GetModbusDetailByRuleID(h.Engine, rule.ID)
			if err != nil || detail == nil {
				slog.Error("Failed to get modbus detail", logger.ErrAttr(err))
				return
			}
			details[rule.IEC61850Ref] = detail
		}
		for ref, detail := range details {
			reg, err := h.ReadHoldingRegister(uint16(detail.StartAddress))
			if err != nil {
				slog.Error("Failed to read holding register", logger.ErrAttr(err))
				continue
			}
			data := Data{
				IEC61850Ref: ref,
				Value:       fmt.Sprintf("%d", reg),
				DataSource:  "Modbus",
			}
			s <- data
		}
		err = h.client.Close()
		if err != nil {
			slog.Error("Failed to close modbus connection", logger.ErrAttr(err))
		}
	}
}

func (h *ModbusHandler) ReadHoldingRegister(address uint16) (uint16, error) {
	reg, err := h.client.ReadRegister(address, modbus.HOLDING_REGISTER)
	if err != nil {
		return 0, fmt.Errorf("error reading holding register: %v", err)
	}
	return reg, nil
}

func (h *ModbusHandler) IsOnline() bool {
	return h.client != nil
}

func (h *ModbusHandler) Type() HandlerType {
	return ModbusHandlerType
}
