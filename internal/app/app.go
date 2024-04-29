// Copyright 2023 lng2020. All rights reserved.
// SPDX-License-Identifier: MIT

package app

import (
	"fmt"
	"geata/internal/app/handler"
	"geata/internal/app/model"
	"geata/internal/app/service"
	"geata/internal/app/web"
	"log"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
	"github.com/wind-c/comqtt/v2/mqtt"
	"github.com/wind-c/comqtt/v2/mqtt/hooks/auth"
	"github.com/wind-c/comqtt/v2/mqtt/listeners"
	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

type App struct {
	Config     *AppConfig
	Stations   []*service.Station
	db         *xorm.Engine
	mqttServer *mqtt.Server
}

var models = []any{
	new(model.Station),
	new(model.Node),
	new(model.MappingRule),
	new(model.MqttDetail),
	new(model.ModbusDetail),
	new(model.IEC61850Model),
	new(model.LogicalDevice),
	new(model.LogicalNode),
	new(model.DataObject),
}

func NewApp() *App {
	return &App{}
}

func (app *App) NewConfig(configFile string) error {
	if configFile == "" {
		configFile = "./internal/app/config.yaml"
	}
	config, err := LoadConfig(configFile)
	if err != nil {
		return err
	}
	app.Config = config
	return nil
}

func (app *App) InitDB() error {
	dbconf := app.Config.Database
	driverName := ""
	dataSourceName := ""

	switch dbconf.Type {
	case "mysql":
		driverName = "mysql"
		dataSourceName = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
			dbconf.Username,
			dbconf.Password,
			dbconf.Host,
			dbconf.Port,
			dbconf.DBName)
	case "sqlite3":
		driverName = "sqlite3"
		dataSourceName = dbconf.DBName
	case "postgres":
		driverName = "postgres"
		dataSourceName = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			dbconf.Host,
			dbconf.Port,
			dbconf.Username,
			dbconf.Password,
			dbconf.DBName)
	default:
		return fmt.Errorf("unsupported database type: %s", dbconf.Type)
	}

	Engine, err := xorm.NewEngine(driverName, dataSourceName)
	if err != nil {
		return err
	}

	err = Engine.Ping()
	if err != nil {
		return err
	}

	Engine.ShowSQL(true)
	Engine.SetMaxIdleConns(10)
	Engine.SetMaxOpenConns(100)
	Engine.SetMapper(names.GonicMapper{})

	err = Engine.Sync(models...)
	if err != nil {
		return err
	}

	app.db = Engine
	service.Engine = Engine
	return nil
}

func (app *App) RegisterHandlers() error {
	handler.RegisterHandler(handler.ModbusHandlerType, handler.NewModbusHandler)
	handler.RegisterHandler(handler.MQTTHandlerType, handler.NewMQTTHandler)
	handler.RegisterHandler(handler.IEC61850HandlerType, handler.NewIEC61850Handler)
	return nil
}

func (app *App) InitStations() error {
	stationsFromDB, err := model.GetAllStations(app.db)
	if err != nil {
		return err
	}

	for _, stationFromDB := range stationsFromDB {
		station := &service.Station{}
		err := station.InitFromDB(stationFromDB)
		if err != nil {
			return err
		}

		for _, handlerType := range handler.HandlerTypes {
			config := station.Configs[handlerType]
			station.Handlers[handlerType] = handler.SupportedHandler[handlerType](config)
		}

		app.Stations = append(app.Stations, station)
	}
	return nil
}

func (app *App) InitMQTTBroker() error {
	app.mqttServer = mqtt.New(nil)

	_ = app.mqttServer.AddHook(new(auth.AllowHook), nil)

	tcp := listeners.NewTCP("t1", fmt.Sprintf(":%s", strconv.Itoa(app.Config.MQTTBroker.Port)), nil)
	err := app.mqttServer.AddListener(tcp)
	if err != nil {
		return err
	}

	return nil
}

func (app *App) Init() error {
	err := app.InitDB()
	if err != nil {
		return err
	}

	err = app.InitMQTTBroker()
	if err != nil {
		return err
	}

	err = app.RegisterHandlers()
	if err != nil {
		return err
	}

	err = app.InitStations()
	if err != nil {
		return err
	}
	return nil
}

func (app *App) Start() error {
	go func() {
		err := app.mqttServer.Serve()
		if err != nil {
			log.Fatal("Failed to start MQTT server: ", err)
		}
		log.Println("MQTT server started")
	}()
	router := web.SetupRouter()
	return router.Run(fmt.Sprintf(":%d", app.Config.Server.Port))
}
