package service

import (
	"geata/internal/app/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type IEC61850Model struct {
	ID            int64
	Name          string
	Description   string
	LogicalDevice []LogicalDevice
}

type LogicalDevice struct {
	ID          int64
	ModelID     int64
	Name        string
	Description string
	LogicalNode []LogicalNode
}

type LogicalNode struct {
	ID              int64
	LogicalDeviceID int64
	Name            string
}

type DataObject struct {
	ID            int64
	LogicalNodeID int64
	Name          string
	Description   string
	DataAttribute []DataAttribute
}

type DataAttribute model.Node

// @summary Get IEC61850 model by ID
// @tags IEC61850
// @produce json
// @param id path int true "IEC61850 model ID"
// @success 200 {object} IEC61850Model
// @router /iec61850/model/{id} [get]
func GetIEC61850ModelByID(c *gin.Context) {
	ID := c.Param("id")
	modelID, err := strconv.Atoi(ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	model, err := model.GetIEC61850ModelByID(Engine, int64(modelID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, model)
}

// @summary Get DataObject by ID
// @tags IEC61850
// @produce json
// @param id path int true "DataObject ID"
// @success 200 {object} DataObject
// @router /iec61850/data_object/{id} [get]
func GetDataObjectByID(c *gin.Context) {
	ID := c.Param("id")
	dataObjectID, err := strconv.Atoi(ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	dataObject, err := model.GetDataObjectByID(Engine, int64(dataObjectID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, dataObject)
}

// @summary Get Node by ID
// @tags IEC61850
// @produce json
// @param id path int true "Node ID"
// @success 200 {object} DataAttribute
// @router /iec61850/node/{id} [get]
func GetNodeByID(c *gin.Context) {
	ID := c.Param("id")
	nodeID, err := strconv.Atoi(ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	node, err := model.GetNodeByID(Engine, int64(nodeID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, node)
}

// @summary Get Node by IEC61850 reference
// @tags IEC61850
// @produce json
// @param ref path string true "IEC61850 reference"
// @success 200 {object} DataAttribute
// @router /iec61850/node/ref/{ref} [get]
func GetNodeByRef(c *gin.Context) {
	ref := c.Param("ref")
	node, err := model.GetNodeByRef(Engine, ref)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, node)
}

// @summary Get Nodes by DataObject ID
// @tags IEC61850
// @produce json
// @param id path int true "DataObject ID"
// @success 200 {array} DataAttribute
// @router /iec61850/nodes/data_object/{id} [get]
func GetNodesByDataObjectID(c *gin.Context) {
	ID := c.Param("id")
	dataObjectID, err := strconv.Atoi(ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	nodes, err := model.GetNodesByDataObjectID(Engine, int64(dataObjectID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, nodes)
}

// @summary Update DataSource of a Node
// @tags IEC61850
// @accept json
// @produce json
// @param id path int true "Node ID"
// @param dataSource body string true "DataSource"
// @success 200 {string} string "OK"
// @router /iec61850/node/{id}/data_source [put]
func UpdateNodeDataSource(c *gin.Context) {
	ID := c.Param("id")
	nodeID, err := strconv.Atoi(ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var dataSource string
	err = c.BindJSON(&dataSource)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	node, err := model.GetNodeByID(Engine, int64(nodeID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	node.DataSource = dataSource
	err = model.UpdateNode(Engine, node)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "OK")
}
