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
	FC_ST FunctionalConstraint = iota
	FC_MX
	FC_SP
	FC_SV
	FC_CF
	FC_DC
	FC_SG
	FC_SE
	FC_SR
	FC_OR
	FC_BL
	FC_EX
	FC_CO
	FC_US
	FC_MS
	FC_RP
	FC_BR
	FC_LG
	FC_GO

	FC_ALL  = 0xFF
	FC_None = -1
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
	triggerOptions uint8
	mmsValue       *mms.MmsValue
	sAddr          uint32
}
