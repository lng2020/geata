// Copyright 2023 lng2020. All rights reserved.
// SPDX-License-Identifier: MIT

package app

import (
	"fmt"
	"geata/internal/app/handler"
	"geata/internal/app/model"
	"geata/internal/app/web"

	"xorm.io/xorm"
)

// App represents the application.
type App struct {
	Config   *AppConfig
	Stations []*handler.Station
	db       *xorm.Engine
}

// NewApp creates a new instance of App.
func NewApp() *App {
	return &App{}
}

// NewConfig creates a new instance of AppConfig.
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

// InitDB initializes database.
func (app *App) InitDB() error {
	dbconf := app.Config.Database
	driverName := ""

	switch dbconf.Type {
	case "mysql":
		driverName = "mysql"
	case "sqlite3":
		driverName = "sqlite3"
	default:
		return fmt.Errorf("unsupported database type: %s", dbconf.Type)
	}

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbconf.Username,
		dbconf.Password,
		dbconf.Host,
		dbconf.Port,
		dbconf.DBName)

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

	app.db = Engine
	return nil
}

func (app *App) Init() error {
	err := app.InitDB()
	if err != nil {
		return err
	}

	app.Stations = make([]*handler.Station, 0)

	stationsFromDB, err := model.GetAllStations(app.db)
	if err != nil {
		return err
	}

	for _, station := range stationsFromDB {
		stationHandlers := &handler.Handlers{
			// TODO: add other handlers
		}
		station := &handler.Station{
			Name:     station.Name,
			Handlers: stationHandlers,
		}
		app.Stations = append(app.Stations, station)
	}
	return nil
}

// Start starts application.
func (app *App) Start() error {
	router := web.SetupRouter()
	return router.Run(fmt.Sprintf(":%d", app.Config.Server.Port))
}
