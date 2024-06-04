// Copyright 2023 lng2020. All rights reserved.
// SPDX-License-Identifier: MIT

package app

import (
	"context"
	"fmt"
	"geata/internal/app/logger"
	"geata/internal/app/model"
	"geata/internal/app/service"
	"geata/internal/app/web"
	"log/slog"
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
	mqttServer *mqtt.Server
	ctx        context.Context
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
	new(model.AuditLog),
	new(model.User),
	new(model.StationConfig),
}

func NewApp(ctx context.Context) *App {
	return &App{ctx: ctx}
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
	Engine.SetLogger(logger.NewXORMLogger())

	err = Engine.Sync(models...)
	if err != nil {
		return err
	}

	service.Engine = Engine
	return nil
}

func (app *App) InitStations() error {
	stationsInDB, err := model.GetAllStations(service.Engine)
	if err != nil {
		return err
	}

	for _, stationInDB := range stationsInDB {
		station := service.StationInitFromDB(stationInDB)
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

	go func() {
		err := app.mqttServer.Serve()
		if err != nil {
			slog.Error("Failed to start the MQTT server: ", err)
		}
		slog.Info("MQTT server started on", "port", strconv.Itoa(app.Config.MQTTBroker.Port))
	}()

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

	err = app.InitStations()
	if err != nil {
		return err
	}

	return nil
}

func (app *App) Start() error {
	for _, station := range app.Stations {
		for _, handler := range station.Handlers {
			go handler.Handle(app.ctx, station.Datas[handler.Type()])
		}
	}
	router := web.SetupRouter()
	return router.Run(fmt.Sprintf(":%d", app.Config.Server.Port))
}
