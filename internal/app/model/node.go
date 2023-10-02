package model

import (
	"time"

	"xorm.io/xorm"
)

type node struct {
	ID          int64     `xorm:"pk autoincr 'id'" json:"id"`
	StationID   int64     `xorm:"'station_id' index" json:"station_id"`
	IEC61850Ref string    `xorm:"'iec61850_ref'" json:"iec61850_ref"`
	DataSource  string    `xorm:"'data_source'" json:"data_source"`
	CreatedAt   time.Time `xorm:"'created_at' created" json:"created_at"`
	UpdatedAt   time.Time `xorm:"'updated_at' updated" json:"updated_at"`
}

func (n *node) TableName() string {
	return "node"
}

func Createnode(engine *xorm.Engine, node *node) error {
	_, err := engine.Insert(node)
	return err
}

func GetnodeByID(engine *xorm.Engine, id int64) (*node, error) {
	node := new(node)
	has, err := engine.ID(id).Get(node)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return node, nil
}

func GetnodeByRef(engine *xorm.Engine, ref string) (*node, error) {
	node := new(node)
	has, err := engine.Where("iec61850_ref = ?", ref).Get(node)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return node, nil
}

func GetAllnodes(engine *xorm.Engine, stationID int64) ([]*node, error) {
	var nodes []*node
	err := engine.Where("station_id = ?", stationID).Find(&nodes)
	if err != nil {
		return nil, err
	}
	return nodes, nil
}

func Updatenode(engine *xorm.Engine, node *node) error {
	_, err := engine.ID(node.ID).Update(node)
	return err
}

func Deletenode(engine *xorm.Engine, node *node) error {
	_, err := engine.ID(node.ID).Delete(node)
	return err
}
