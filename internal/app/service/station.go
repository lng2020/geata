package service

import (
	"geata/internal/app/handler"
	"geata/internal/app/model"
)

type Station struct {
	ID   int64
	Name string
	Host string
	Port int64
	*handler.Handlers
}

func (s *Station) InitFromDB(stationFromDB *model.Station) error {
	s.ID = stationFromDB.ID
	s.Name = stationFromDB.Name
	s.Host = stationFromDB.Host
	s.Port = stationFromDB.Port

	return nil
}
