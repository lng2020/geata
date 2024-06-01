package model

import (
	"xorm.io/xorm"
)

type IEC61850Model struct {
	ID          int64  `xorm:"pk autoincr 'id'" json:"id"`
	Name        string `xorm:"'name'" json:"name"`
	Description string `xorm:"'description'" json:"description"`
}

type LogicalDevice struct {
	ID          int64  `xorm:"pk autoincr 'id'" json:"id"`
	ModelID     int64  `xorm:"'model_id' index" json:"modelID"`
	Name        string `xorm:"'name'" json:"name"`
	Description string `xorm:"'description'" json:"description"`
}

func (m *IEC61850Model) TableName() string {
	return "iec61850_model"
}

func (ld *LogicalDevice) TableName() string {
	return "logical_device"
}

type LogicalNode struct {
	ID              int64  `xorm:"pk autoincr 'id'" json:"id"`
	LogicalDeviceID int64  `xorm:"'logical_device_id' index" json:"logicalDeviceID"`
	Name            string `xorm:"'name'" json:"name"`
	Description     string `xorm:"'description'" json:"description"`
}

func (ln *LogicalNode) TableName() string {
	return "logical_node"
}

type DataObject struct {
	ID            int64  `xorm:"pk autoincr 'id'" json:"id"`
	LogicalNodeID int64  `xorm:"'logical_node_id' index" json:"logicalNodeID"`
	Name          string `xorm:"'name'" json:"name"`
	Description   string `xorm:"'description'" json:"description"`
}

func (do *DataObject) TableName() string {
	return "data_object"
}

func CreateIEC61850Model(engine *xorm.Engine, iec61850Model *IEC61850Model) error {
	_, err := engine.Insert(iec61850Model)
	return err
}

func GetIEC61850ModelByID(engine *xorm.Engine, id int64) (*IEC61850Model, error) {
	iec61850Model := new(IEC61850Model)
	has, err := engine.ID(id).Get(iec61850Model)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return iec61850Model, nil
}

func UpdateIEC61850Model(engine *xorm.Engine, iec61850Model *IEC61850Model) error {
	_, err := engine.ID(iec61850Model.ID).Update(iec61850Model)
	return err
}

func DeleteIEC61850Model(engine *xorm.Engine, iec61850Model *IEC61850Model) error {
	_, err := engine.ID(iec61850Model.ID).Delete(iec61850Model)
	return err
}

func CreateLogicalDevice(engine *xorm.Engine, logicalDevice *LogicalDevice) error {
	_, err := engine.Insert(logicalDevice)
	return err
}

func GetLogicalDeviceByID(engine *xorm.Engine, id int64) (*LogicalDevice, error) {
	logicalDevice := new(LogicalDevice)
	has, err := engine.ID(id).Get(logicalDevice)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return logicalDevice, nil
}

func UpdateLogicalDevice(engine *xorm.Engine, logicalDevice *LogicalDevice) error {
	_, err := engine.ID(logicalDevice.ID).Update(logicalDevice)
	return err
}

func DeleteLogicalDevice(engine *xorm.Engine, logicalDevice *LogicalDevice) error {
	_, err := engine.ID(logicalDevice.ID).Delete(logicalDevice)
	return err
}

func GetLogicalDeviceByModelID(engine *xorm.Engine, modelID int64) ([]LogicalDevice, error) {
	logicalDevices := make([]LogicalDevice, 0)
	err := engine.Where("model_id = ?", modelID).Find(&logicalDevices)
	if err != nil {
		return nil, err
	}
	return logicalDevices, nil
}

func CreateLogicalNode(engine *xorm.Engine, logicalNode *LogicalNode) error {
	_, err := engine.Insert(logicalNode)
	return err
}

func GetLogicalNodeByID(engine *xorm.Engine, id int64) (*LogicalNode, error) {
	logicalNode := new(LogicalNode)
	has, err := engine.ID(id).Get(logicalNode)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return logicalNode, nil
}

func GetLogicalNodeByLogicalDeviceID(engine *xorm.Engine, logicalDeviceID int64) ([]LogicalNode, error) {
	logicalNodes := make([]LogicalNode, 0)
	err := engine.Where("logical_device_id = ?", logicalDeviceID).Find(&logicalNodes)
	if err != nil {
		return nil, err
	}
	return logicalNodes, nil
}

func UpdateLogicalNode(engine *xorm.Engine, logicalNode *LogicalNode) error {
	_, err := engine.ID(logicalNode.ID).Update(logicalNode)
	return err
}

func DeleteLogicalNode(engine *xorm.Engine, logicalNode *LogicalNode) error {
	_, err := engine.ID(logicalNode.ID).Delete(logicalNode)
	return err
}

func CreateDataObject(engine *xorm.Engine, dataObject *DataObject) error {
	_, err := engine.Insert(dataObject)
	return err
}

func GetDataObjectByLogicalNodeID(engine *xorm.Engine, logicalNodeID int64) ([]DataObject, error) {
	dataObjects := make([]DataObject, 0)
	err := engine.Where("logical_node_id = ?", logicalNodeID).Find(&dataObjects)
	if err != nil {
		return nil, err
	}
	return dataObjects, nil
}

func GetDataObjectByID(engine *xorm.Engine, id int64) (*DataObject, error) {
	dataObject := new(DataObject)
	has, err := engine.ID(id).Get(dataObject)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return dataObject, nil
}

func UpdateDataObject(engine *xorm.Engine, dataObject *DataObject) error {
	_, err := engine.ID(dataObject.ID).Update(dataObject)
	return err
}

func DeleteDataObject(engine *xorm.Engine, dataObject *DataObject) error {
	_, err := engine.ID(dataObject.ID).Delete(dataObject)
	return err
}
