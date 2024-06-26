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

	tcp := listeners.NewTCP("t1", fmt.Sprintf(":%s", strconv.Itoa(int(app.Config.MQTTBroker.Port))), nil)
	err := app.mqttServer.AddListener(tcp)
	if err != nil {
		return err
	}

	go func() {
		err := app.mqttServer.Serve()
		if err != nil {
			slog.Error("Failed to start the MQTT server: ", err)
		}
		slog.Info("MQTT server started on", "port", strconv.Itoa(int(app.Config.MQTTBroker.Port)))
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
	stationDataQueue := make(chan service.StationData, len(app.Stations))
	stationIDCancelMap := make(map[int64]context.CancelFunc)
	for _, station := range app.Stations {
		ctx, cancel := context.WithCancel(app.ctx)
		go station.Start(ctx, stationDataQueue)
		stationIDCancelMap[station.ID] = cancel
	}
	go HandleStationData(app.ctx, stationDataQueue)
	go HandleUpdateStation(app.ctx, stationIDCancelMap, stationDataQueue)
	router := web.SetupRouter()
	return router.Run(fmt.Sprintf(":%d", app.Config.Server.Port))
}

func HandleStationData(ctx context.Context, stationDataQueue chan service.StationData) {
	for {
		select {
		case stationData := <-stationDataQueue:
			err := model.UpdateNodeValueByModelIDAndDataSource(service.Engine, stationData.ModelID, stationData.Data.IEC61850Ref, stationData.Data.Value, stationData.Data.DataSource)
			if err != nil {
				slog.Error("Failed to update node value", logger.ErrAttr(err))
			}
		case <-ctx.Done():
			slog.Info("HandleStationData stopped")
			return
		}
	}
}

func HandleUpdateStation(ctx context.Context, m map[int64]context.CancelFunc, stationDataQueue chan service.StationData) {
	for {
		select {
		case <-ctx.Done():
			slog.Info("HandleUpdateStation stopped")
			return
		default:
			stationInDB, err := model.GetAllStations(service.Engine)
			if err != nil {
				slog.Error("Failed to get all stations", logger.ErrAttr(err))
				continue
			}
			for _, station := range stationInDB {
				if _, ok := m[station.ID]; !ok {
					slog.Info("New station found", "station_id", station.ID)
					newStation := service.StationInitFromDB(station)
					newCtx, cancel := context.WithCancel(ctx)
					m[station.ID] = cancel
					go newStation.Start(newCtx, stationDataQueue)
				}
			}
			for id, cancel := range m {
				found := false
				for _, station := range stationInDB {
					if station.ID == id {
						found = true
						break
					}
				}
				if !found {
					slog.Info("Station removed", "station_id", id)
					cancel()
					delete(m, id)
				}
			}
		}
	}
}
