package service

import "geata/internal/app/model"

type StationConfig model.StationConfig

func CreateDefaultStationConfig(stationID int64) *model.StationConfig {
	return &model.StationConfig{
		StationID:    stationID,
		ModbusURL:    "tcp://localhost:502",
		IEC61850Host: "localhost",
		IEC61850Port: 102,
		MQTTBroker:   "tcp://localhost:1883",
		MQTTClientID: "geata",
		MQTTUsername: "geata",
		MQTTPassword: "geata",
		MQTTTopic:    "geata",
	}
}
