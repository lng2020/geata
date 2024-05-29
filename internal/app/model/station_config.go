package model

import "xorm.io/xorm"

type StationConfig struct {
	ID           int64  `xorm:"pk autoincr 'id'"`
	StationID    int64  `xorm:"'station_id'"`
	ModbusURL    string `xorm:"'modbus_url'"`
	IEC61850Host string `xorm:"'iec61850_host'"`
	IEC61850Port int64  `xorm:"'iec61850_port'"`
	MQTTBroker   string `xorm:"'mqtt_broker'"`
	MQTTClientID string `xorm:"'mqtt_client_id'"`
	MQTTUsername string `xorm:"'mqtt_username'"`
	MQTTPassword string `xorm:"'mqtt_password'"`
	MQTTTopic    string `xorm:"'mqtt_topic'"`
}

func (s *StationConfig) TableName() string {
	return "station_config"
}

func CreateStationConfig(engine *xorm.Engine, stationConfig *StationConfig) error {
	_, err := engine.Insert(stationConfig)
	return err
}

func GetStationConfigByStationID(engine *xorm.Engine, stationID int64) (*StationConfig, error) {
	stationConfig := new(StationConfig)
	has, err := engine.Where("station_id = ?", stationID).Get(stationConfig)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return stationConfig, nil
}
