package model

import (
	"time"

	"xorm.io/xorm"
)

type Node struct {
	ID           int64     `xorm:"pk autoincr 'id'" json:"ID"`
	DataObjectID int64     `xorm:"'data_object_id' index" json:"dataObjectID"`
	Name         string    `xorm:"'name'" json:"name"`
	Value        string    `xorm:"'value'" json:"value"`
	IEC61850Ref  string    `xorm:"'iec61850_ref'" json:"iec61850Ref"`
	DataSource   string    `xorm:"'data_source'" json:"dataSource"`
	CreatedAt    time.Time `xorm:"'created_at' created" json:"createdAt"`
	UpdatedAt    time.Time `xorm:"'updated_at' updated" json:"updatedAt"`
}

func (n *Node) TableName() string {
	return "node"
}

func CreateNode(engine *xorm.Engine, node *Node) error {
	_, err := engine.Insert(node)
	return err
}

func GetNodeByID(engine *xorm.Engine, id int64) (*Node, error) {
	node := new(Node)
	has, err := engine.ID(id).Get(node)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return node, nil
}

func GetNodeByRef(engine *xorm.Engine, ref string) (*Node, error) {
	node := new(Node)
	has, err := engine.Where("iec61850_ref = ?", ref).Get(node)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return node, nil
}

func GetNodesByDataObjectID(engine *xorm.Engine, dataObjectID int64) ([]*Node, error) {
	var nodes []*Node
	err := engine.Where("data_object_id = ?", dataObjectID).Find(&nodes)
	if err != nil {
		return nil, err
	}
	return nodes, nil
}

func UpdateNode(engine *xorm.Engine, node *Node) error {
	_, err := engine.ID(node.ID).Update(node)
	return err
}

func DeleteNode(engine *xorm.Engine, node *Node) error {
	_, err := engine.ID(node.ID).Delete(node)
	return err
}
