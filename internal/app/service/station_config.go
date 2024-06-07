package service

import (
	"fmt"
	"geata/internal/app/model"
)

type StationConfig model.StationConfig

func CreateStationConfigByStation(s *model.Station) *model.StationConfig {
	return &model.StationConfig{
		StationID:    s.ID,
		ModbusURL:    "tcp://localhost:502",
		IEC61850Host: s.Host,
		IEC61850Port: s.Port,
		MQTTBroker:   "tcp://localhost:1883",
		MQTTClientID: fmt.Sprintf("station-%d", s.ID),
		MQTTUsername: "geata",
		MQTTPassword: "geata",
		MQTTTopic:    fmt.Sprintf("station-%d", s.ID),
	}
}
