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
