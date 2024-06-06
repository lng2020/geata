// Copyright 2023 lng2020. All rights reserved.
// SPDX-License-Identifier: MIT

package model

import (
	"time"

	"xorm.io/xorm"
)

type MqttDetail struct {
	ID          int64     `xorm:"pk autoincr 'id'" json:"id"`
	RuleID      int64     `xorm:"'rule_id' index" json:"ruleID"`
	Field       string    `xorm:"'field'" json:"field"`
	DataType    string    `xorm:"'data_type'" json:"dataType"`
	PayloadType string    `xorm:"'payload_type'" json:"payloadType"`
	CreatedAt   time.Time `xorm:"created" json:"createdAt"`
	UpdatedAt   time.Time `xorm:"updated" json:"updatedAt"`
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

func CreateOrUpdateMQTTDetail(engine *xorm.Engine, mqttDetail *MqttDetail) error {
	_, err := engine.Insert(mqttDetail)
	if err != nil {
		_, err = engine.ID(mqttDetail.ID).Update(mqttDetail)
	}
	return err
}
