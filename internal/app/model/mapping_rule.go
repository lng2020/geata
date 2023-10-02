package model

import (
	"time"

	"xorm.io/xorm"
)

type MappingRule struct {
	ID          int64     `xorm:"pk autoincr 'id'" json:"id"`
	StationID   int64     `xorm:"'station_id' index" json:"station_id"`
	IEC61850Ref string    `xorm:"'iec61850_ref'" json:"iec61850_ref"`
	Type        string    `xorm:"'type'" json:"type"`
	CreateTime  time.Time `xorm:"'create_time' created" json:"create_time"`
	UpdateTime  time.Time `xorm:"'update_time' updated" json:"update_time"`
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

func GetAllMappingRules(engine *xorm.Engine, stationID int64) ([]*MappingRule, error) {
	var rules []*MappingRule
	err := engine.Where("station_id = ?", stationID).Find(&rules)
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
