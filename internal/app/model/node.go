// Copyright 2023 lng2020. All rights reserved.
// SPDX-License-Identifier: MIT

package model

import (
	"time"

	"xorm.io/xorm"
)

type Node struct {
	ID          int64     `xorm:"pk autoincr 'id'" json:"id"`
	StationID   int64     `xorm:"'station_id' index" json:"station_id"`
	IEC61850Ref string    `xorm:"'iec61850_ref'" json:"iec61850_ref"`
	DataSource  string    `xorm:"'data_source'" json:"data_source"`
	CreatedAt   time.Time `xorm:"'created_at' created" json:"created_at"`
	UpdatedAt   time.Time `xorm:"'updated_at' updated" json:"updated_at"`
}

func (n *Node) TableName() string {
	return "Node"
}

func CreateNode(engine *xorm.Engine, Node *Node) error {
	_, err := engine.Insert(Node)
	return err
}

func GetNodeByID(engine *xorm.Engine, id int64) (*Node, error) {
	Node := new(Node)
	has, err := engine.ID(id).Get(Node)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return Node, nil
}

func GetNodeByRef(engine *xorm.Engine, ref string) (*Node, error) {
	Node := new(Node)
	has, err := engine.Where("iec61850_ref = ?", ref).Get(Node)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return Node, nil
}

func GetAllNodes(engine *xorm.Engine, stationID int64) ([]*Node, error) {
	var Nodes []*Node
	err := engine.Where("station_id = ?", stationID).Find(&Nodes)
	if err != nil {
		return nil, err
	}
	return Nodes, nil
}

func UpdateNode(engine *xorm.Engine, Node *Node) error {
	_, err := engine.ID(Node.ID).Update(Node)
	return err
}

func DeleteNode(engine *xorm.Engine, Node *Node) error {
	_, err := engine.ID(Node.ID).Delete(Node)
	return err
}
