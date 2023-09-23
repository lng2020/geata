// Copyright 2023 lng2020. All rights reserved.
// SPDX-License-Identifier: MIT

package app

import (
	"os"

	"gopkg.in/yaml.v2"
)

// AppConfig contains the configuration settings for your application.
type AppConfig struct {
	Server   ServerConfig   `yaml:"server"`
	Logger   LoggerConfig   `yaml:"logger"`
	Plugins  PluginConfig   `yaml:"plugins"`
	Database DatabaseConfig `yaml:"database"`
}

// ServerConfig contains server-related configuration.
type ServerConfig struct {
	Port int    `yaml:"port"`
	Host string `yaml:"host"`
}

// LoggerConfig contains logger-related configuration.
type LoggerConfig struct {
	LogLevel string `yaml:"log_level"`
	LogFile  string `yaml:"log_file"`
}

// PluginConfig contains plugin-related configuration.
type PluginConfig struct {
	MQTT     MQTTConfig     `yaml:"mqtt"`
	Modbus   ModbusConfig   `yaml:"modbus"`
	IEC61850 IEC61850Config `yaml:"iec61850"`
}

// MQTTConfig contains MQTT plugin configuration.
type MQTTConfig struct {
	BrokerAddress string `yaml:"broker_address"`
	Username      string `yaml:"username"`
	Password      string `yaml:"password"`
}

// ModbusConfig contains Modbus plugin configuration.
type ModbusConfig struct {
	ServerAddress string `yaml:"server_address"`
	SlaveID       int    `yaml:"slave_id"`
}

// IEC61850Config contains IEC 61850 plugin configuration.
type IEC61850Config struct {
	ServerAddress string `yaml:"server_address"`
}

// DatabaseConfig contains database-related configuration.
type DatabaseConfig struct {
	Type     string `yaml:"type"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
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
		Plugins: PluginConfig{
			MQTT: MQTTConfig{
				BrokerAddress: "tcp://mqtt.example.com:1883",
				Username:      "",
				Password:      "",
			},
			Modbus: ModbusConfig{
				ServerAddress: "tcp://modbus.example.com:502",
				SlaveID:       1,
			},
			IEC61850: IEC61850Config{
				ServerAddress: "iec61850.example.com",
			},
		},
		Database: DatabaseConfig{
			Type:     "postgresql",
			Host:     "db.example.com",
			Port:     5432,
			Username: "dbuser",
			Password: "dbpassword",
			DBName:   "mydatabase",
		},
	}

	if err := yaml.Unmarshal(configFile, config); err != nil {
		return nil, err
	}

	return config, nil
}
