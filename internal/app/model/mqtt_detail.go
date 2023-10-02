// Copyright 2023 lng2020. All rights reserved.
// SPDX-License-Identifier: MIT

package model

import "xorm.io/xorm"

type MqttDetail struct {
	ID     int64 `xorm:"pk autoincr 'id'" json:"id"`
	RuleID int64 `xorm:"'rule_id' index" json:"rule_id"`
	// other fields specific to MqttDetail as needed.
}

func (m *MqttDetail) TableName() string {
	return "mqtt_detail"
}

func CreateMqttDetail(engine *xorm.Engine, mqttDetail *MqttDetail) error {
	_, err := engine.Insert(mqttDetail)
	return err
}

func GetMqttDetailByRuleID(engine *xorm.Engine, ruleID int64) (*MqttDetail, error) {
	mqttDetail := new(MqttDetail)
	has, err := engine.ID(ruleID).Get(mqttDetail)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return mqttDetail, nil
}

func UpdateMqttDetail(engine *xorm.Engine, mqttDetail *MqttDetail) error {
	_, err := engine.ID(mqttDetail.ID).Update(mqttDetail)
	return err
}

func DeleteMqttDetail(engine *xorm.Engine, mqttDetail *MqttDetail) error {
	_, err := engine.ID(mqttDetail.ID).Delete(mqttDetail)
	return err
}
