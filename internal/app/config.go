// Copyright 2023 lng2020. All rights reserved.
// SPDX-License-Identifier: MIT

package app

import (
	"os"

	"gopkg.in/yaml.v2"
)

type AppConfig struct {
	Server     ServerConfig     `yaml:"server"`
	Logger     LoggerConfig     `yaml:"logger"`
	Database   DatabaseConfig   `yaml:"database"`
	MQTTBroker MQTTBrokerConfig `yaml:"mqtt_broker"`
}

type ServerConfig struct {
	Port int64  `yaml:"port"`
	Host string `yaml:"host"`
}

type LoggerConfig struct {
	LogLevel string `yaml:"log_level"`
	LogFile  string `yaml:"log_file"`
}

type IEC61850Config struct {
	ServerAddress string `yaml:"server_address"`
}

type DatabaseConfig struct {
	Type     string `yaml:"type"`
	Host     string `yaml:"host"`
	Port     int64  `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
}

type MQTTBrokerConfig struct {
	Port int64 `yaml:"port"`
}

func LoadConfig(configFilePath string) (*AppConfig, error) {
	configFile, err := os.ReadFile(configFilePath)
	if err != nil {
		return nil, err
	}

	config := &AppConfig{
		Server: ServerConfig{
			Port: 8080,
			Host: "0.0.0.0",
		},
		Logger: LoggerConfig{
			LogLevel: "info",
			LogFile:  "/path/to/default/logfile.log",
		},
		Database: DatabaseConfig{
			Type:     "postgresql",
			Host:     "db.example.com",
			Port:     5432,
			Username: "dbuser",
			Password: "dbpassword",
			DBName:   "mydatabase",
		},
		MQTTBroker: MQTTBrokerConfig{
			Port: 1883,
		},
	}

	if err := yaml.Unmarshal(configFile, config); err != nil {
		return nil, err
	}

	return config, nil
}
