// Copyright 2023 lng2020. All rights reserved.
// SPDX-License-Identifier: MIT

package model

import (
	"time"

	"xorm.io/xorm"
)

type ModbusDetail struct {
	ID           int64     `xorm:"pk autoincr 'id'" json:"id"`
	RuleID       int64     `xorm:"'rule_id' index" json:"ruleID"`
	FunctionCode int64     `xorm:"'function_code'" json:"function_code"`
	StartAddress int64     `xorm:"'start_address'" json:"start_address"`
	Length       int64     `xorm:"'length'" json:"length"`
	DataType     int64     `xorm:"'data_type'" json:"data_type"`
	CreateTime   time.Time `xorm:"'create_time' created" json:"createTime"`
	UpdateTime   time.Time `xorm:"'update_time' updated" json:"updateTime"`
}

func (m *ModbusDetail) TableName() string {
	return "modbus_detail"
}

func CreateModbusDetail(engine *xorm.Engine, modbusDetail *ModbusDetail) error {
	_, err := engine.Insert(modbusDetail)
	return err
}

func GetModbusDetailByRuleID(engine *xorm.Engine, ruleID int64) (*ModbusDetail, error) {
	modbusDetail := new(ModbusDetail)
	has, err := engine.ID(ruleID).Get(modbusDetail)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return modbusDetail, nil
}

func UpdateModbusDetail(engine *xorm.Engine, modbusDetail *ModbusDetail) error {
	_, err := engine.ID(modbusDetail.ID).Update(modbusDetail)
	return err
}

func DeleteModbusDetail(engine *xorm.Engine, modbusDetail *ModbusDetail) error {
	_, err := engine.ID(modbusDetail.ID).Delete(modbusDetail)
	return err
}
