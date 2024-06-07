package model

import (
	"time"

	"xorm.io/xorm"
)

type Node struct {
	ID           int64     `xorm:"pk autoincr 'id'" json:"id"`
	ModelID      int64     `xorm:"'model_id' index" json:"modelId"`
	DataObjectID int64     `xorm:"'data_object_id' index" json:"-"`
	Name         string    `xorm:"'name'" json:"name"`
	Value        string    `xorm:"'value'" json:"value"`
	IEC61850Ref  string    `xorm:"'iec61850_ref'" json:"iec61850Ref"`
	DataSource   string    `xorm:"'data_source'" json:"dataSource"`
	CreatedAt    time.Time `xorm:"'created_at' created" json:"-"`
	UpdatedAt    time.Time `xorm:"'updated_at' updated" json:"-"`
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

func GetNodeByDataObjectID(engine *xorm.Engine, dataObjectID int64) ([]*Node, error) {
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

func UpdateNodeValueByModelIDAndDataSource(engine *xorm.Engine, modelID int64, ref, value, dataSource string) error {
	_, err := engine.Table(new(Node)).Where("model_id = ? AND iec61850_ref = ? AND data_source = ?", modelID, ref, dataSource).Update(map[string]interface{}{"value": value})
	return err
}

func UpdateNodeDataSourceByModelIDAndRef(engine *xorm.Engine, modelID int64, ref, dataSource string) error {
	_, err := engine.Table(new(Node)).Where("model_id = ? AND iec61850_ref = ?", modelID, ref).Update(map[string]interface{}{"data_source": dataSource})
	return err
}
