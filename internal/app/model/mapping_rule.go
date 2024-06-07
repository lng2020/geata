// Copyright 2023 lng2020. All rights reserved.
// SPDX-License-Identifier: MIT

package model

import (
	"time"

	"xorm.io/xorm"
)

type MappingRule struct {
	ID          int64     `xorm:"pk autoincr 'id'" json:"id"`
	ModelID     int64     `xorm:"'model_id' index" json:"modelId"`
	IEC61850Ref string    `xorm:"'iec61850_ref'" json:"iec61850Ref"`
	Type        string    `xorm:"'type'" json:"type"`
	CreateTime  time.Time `xorm:"'create_time' created" json:"createTime"`
	UpdateTime  time.Time `xorm:"'update_time' updated" json:"updateTime"`
}

func (r *MappingRule) TableName() string {
	return "mapping_rule"
}

func CreateMappingRule(engine *xorm.Engine, rule *MappingRule) error {
	_, err := engine.Insert(rule)
	return err
}

func GetMappingRuleByID(engine *xorm.Engine, id int64) (*MappingRule, error) {
	rule := new(MappingRule)
	has, err := engine.ID(id).Get(rule)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return rule, nil
}

func GetMappingByModelID(engine *xorm.Engine, stationID int64) ([]*MappingRule, error) {
	var rules []*MappingRule
	err := engine.Where("model_id = ?", stationID).Find(&rules)
	if err != nil {
		return nil, err
	}
	return rules, nil
}

func UpdateMappingRule(engine *xorm.Engine, rule *MappingRule) error {
	_, err := engine.ID(rule.ID).Update(rule)
	return err
}

func DeleteMappingRule(engine *xorm.Engine, rule *MappingRule) error {
	_, err := engine.ID(rule.ID).Delete(rule)
	return err
}

func GetAllModbusRulesByModelID(engine *xorm.Engine, modelID int64) ([]*MappingRule, error) {
	var rules []*MappingRule
	err := engine.Where("model_id = ? AND type = ?", modelID, "Modbus").Find(&rules)
	if err != nil {
		return nil, err
	}
	return rules, nil
}

func GetAllMqttRulesByModelID(engine *xorm.Engine, modelID int64) ([]*MappingRule, error) {
	var rules []*MappingRule
	err := engine.Where("model_id = ? AND type = ?", modelID, "MQTT").Find(&rules)
	if err != nil {
		return nil, err
	}
	return rules, nil
}
