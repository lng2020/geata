package iec61850

import "geata/internal/app/iec61850/mms"

type ModelNodeType int

const (
	LogicalDeviceModelType ModelNodeType = iota
	LogicalNodeModelType
	DataObjectModelType
	DataAttributeModelType
)

type FunctionalConstraint int

const (
	IEC61850_FC_ST FunctionalConstraint = iota
	IEC61850_FC_MX
	IEC61850_FC_SP
	IEC61850_FC_SV
	IEC61850_FC_CF
	IEC61850_FC_DC
	IEC61850_FC_SG
	IEC61850_FC_SE
	IEC61850_FC_SR
	IEC61850_FC_OR
	IEC61850_FC_BL
	IEC61850_FC_EX
	IEC61850_FC_CO
	IEC61850_FC_US
	IEC61850_FC_MS
	IEC61850_FC_RP
	IEC61850_FC_BR
	IEC61850_FC_LG
	IEC61850_FC_GO

	IEC61850_FC_ALL  = 0xFF
	IEC61850_FC_None = -1
)

type DataAttributeType int

const (
	BOOLEAN DataAttributeType = iota
	INT8
	INT16
	INT32
	INT64
	UINT8
	UINT16
	UINT32
	UINT64
	FLOAT32
	FLOAT64
	STRING
)

type ModelNode struct {
	modelType ModelNodeType
	name      string
	children  []*ModelNode
	parent    *ModelNode
}

type LogicalDevice struct {
	modelType ModelNodeType
	name      string
	children  []*ModelNode
	parent    *ModelNode
}

type LogicalNode struct {
	modelType ModelNodeType
	name      string
	children  []*ModelNode
	parent    *ModelNode
}

type DataObject struct {
	modelType    ModelNodeType
	name         string
	children     []*ModelNode
	parent       *ModelNode
	elementCount int
}

type DataAttribute struct {
	modelType      ModelNodeType
	name           string
	children       []*ModelNode
	parent         *ModelNode
	elementCount   int
	fc             FunctionalConstraint
	datype         DataAttributeType
	triggerOptions byte
	mmsValue       *mms.MmsValue
	sAddr          uint32
}
