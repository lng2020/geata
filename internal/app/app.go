// Copyright 2023 lng2020. All rights reserved.
// SPDX-License-Identifier: MIT

package app

import (
	"fmt"
	"geata/internal/app/handler"
)

// Handlers represents a collection of handlers
type Handlers struct {
	Modbus *handler.ModbusHandler
	// MQTT     *handler.MqttHandler
	// IEC61850 *handler.Iec61850Handler
}

// App represents the application.
type App struct {
	Config   *AppConfig
	Handlers *Handlers
}

// NewApp creates a new instance of App.
func NewApp() *App {
	return &App{}
}

// NewConfig creates a new instance of AppConfig.
func (app *App) NewConfig(configFile string) error {
	if configFile == "" {
		configFile = "config.yaml"
	}
	config, err := LoadConfig(configFile)
	if err != nil {
		return err
	}
	app.Config = config
	return nil
}

// NewHandlers creates a new instance of Handlers with initialized handlers.
func (app *App) NewHandlers() error {
	endpoint := fmt.Sprintf(app.Config.Plugins.Modbus.ServerAddress, app.Config.Plugins.Modbus.SlaveID)
	modbusHandler, err := handler.NewModbusHandler(endpoint)
	if err != nil {
		return err
	}

	// mqttHandler := NewMqttHandler()
	// if err != nil {
	// 	return nil, err
	// }

	// iec61850Handler := NewIec61850Handler()
	// if err != nil {
	// 	return nil, err
	// }

	app.Handlers = &Handlers{
		Modbus: modbusHandler,
		// MQTT:     mqttHandler,
		// IEC61850: iec61850Handler,
	}
	return nil
}

// Start starts application.
func (app *App) Start() error {
	return nil
}
