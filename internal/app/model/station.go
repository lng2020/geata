// Copyright 2023 lng2020. All rights reserved.
// SPDX-License-Identifier: MIT

package model

import (
	"time"

	"xorm.io/xorm"
)

type Station struct {
	ID             int64     `xorm:"pk autoincr 'id'" json:"id"`
	Name           string    `xorm:"'name'" json:"name"`
	Host           string    `xorm:"'host'" json:"host"`
	Port           int       `xorm:"'port'" json:"port"`
	IsOnline       bool      `xorm:"'is_online'" json:"is_online"`
	LastOnlineTime time.Time `xorm:"'last_online_time'" json:"last_online_time"`
	CreatedAt      time.Time `xorm:"'created_at' created" json:"created_at"`
	UpdatedAt      time.Time `xorm:"'updated_at' updated" json:"updated_at"`
}

func (s *Station) TableName() string {
	return "station"
}

func CreateStation(engine *xorm.Engine, station *Station) error {
	_, err := engine.Insert(station)
	return err
}

func GetStationByID(engine *xorm.Engine, id int64) (*Station, error) {
	station := new(Station)
	has, err := engine.ID(id).Get(station)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return station, nil
}

func GetAllStations(engine *xorm.Engine) ([]*Station, error) {
	var stations []*Station
	err := engine.Find(&stations)
	if err != nil {
		return nil, err
	}
	return stations, nil
}

func UpdateStation(engine *xorm.Engine, station *Station) error {
	_, err := engine.ID(station.ID).Update(station)
	return err
}

func DeleteStation(engine *xorm.Engine, station *Station) error {
	_, err := engine.ID(station.ID).Delete(station)
	return err
}
