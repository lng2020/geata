package service

import (
	"geata/internal/app/model"
	"geata/internal/app/parser"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type IEC61850Model struct {
	ID            int64           `json:"id"`
	Name          string          `json:"name"`
	Description   string          `json:"description"`
	LogicalDevice []LogicalDevice `json:"logicalDevice"`
}

type LogicalDevice struct {
	Name        string        `json:"name"`
	Description string        `json:"description"`
	LogicalNode []LogicalNode `json:"logicalNode"`
}

type LogicalNode struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type DataObject struct {
	Name          string          `json:"name"`
	Description   string          `json:"description"`
	DataAttribute []DataAttribute `json:"dataAttribute"`
}

type DataAttribute model.Node

// @summary Get IEC61850 model by ID
// @tags IEC61850
// @produce json
// @param id path int true "IEC61850 model ID"
// @success 200 {object} IEC61850Model
// @router /api/v1/iec61850/model/{model_id} [get]
func GetIEC61850ModelByID(c *gin.Context) {
	ID := c.Param("model_id")
	modelID, err := strconv.Atoi(ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	iec61850Model, err := model.GetIEC61850ModelByID(Engine, int64(modelID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	LDs, err := model.GetLogicalDeviceByModelID(Engine, iec61850Model.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	resp := IEC61850Model{
		Name:          iec61850Model.Name,
		Description:   iec61850Model.Description,
		LogicalDevice: make([]LogicalDevice, 0),
	}
	for _, ld := range LDs {
		logicalDevice := LogicalDevice{
			Name:        ld.Name,
			Description: ld.Description,
			LogicalNode: make([]LogicalNode, 0),
		}
		LNs, err := model.GetLogicalNodeByLogicalDeviceID(Engine, ld.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		for _, ln := range LNs {
			logicalNode := LogicalNode{
				ID:          ln.ID,
				Name:        ln.Name,
				Description: ln.Description,
			}
			logicalDevice.LogicalNode = append(logicalDevice.LogicalNode, logicalNode)
		}
		resp.LogicalDevice = append(resp.LogicalDevice, logicalDevice)
	}
	c.JSON(http.StatusOK, resp)
}

// @summary Get DataObject by ID
// @tags IEC61850
// @produce json
// @param id path int true "DataObject ID"
// @success 200 {object} DataObject
// @router /api/v1/iec61850/data_object/{id} [get]
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

// @summary Get DataObject by LogicalNode ID
// @tags IEC61850
// @produce json
// @param id path int true "LogicalNode ID"
// @success 200 {array} DataObject
// @router /api/v1/iec61850/logical_node/{logical_node_id}/data_object [get]
func GetDataObjectByLogicalNodeID(c *gin.Context) {
	ID := c.Param("logical_node_id")
	logicalNodeID, err := strconv.Atoi(ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	dataObjects, err := model.GetDataObjectByLogicalNodeID(Engine, int64(logicalNodeID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	resp := make([]DataObject, 0)
	for _, dataObject := range dataObjects {
		do := DataObject{
			Name:          dataObject.Name,
			Description:   dataObject.Description,
			DataAttribute: make([]DataAttribute, 0),
		}
		DAs, err := model.GetNodeByDataObjectID(Engine, dataObject.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		for _, da := range DAs {
			dataAttribute := DataAttribute(*da)
			do.DataAttribute = append(do.DataAttribute, dataAttribute)
		}
		resp = append(resp, do)
	}
	c.JSON(http.StatusOK, resp)
}

// @summary Get Node by ID
// @tags IEC61850
// @produce json
// @param id path int true "Node ID"
// @success 200 {object} DataAttribute
// @router /api/v1/iec61850/node/{id} [get]
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
// @router /api/v1/iec61850/node/ref/{ref} [get]
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
// @param object_id path int true "DataObject ID"
// @success 200 {array} DataAttribute
// @router /api/v1/iec61850/data_object/{object_id}/node [get]
func GetNodeByDataObjectID(c *gin.Context) {
	ID := c.Param("id")
	dataObjectID, err := strconv.Atoi(ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	nodes, err := model.GetNodeByDataObjectID(Engine, int64(dataObjectID))
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
// @router /api/v1/iec61850/node/{id}/data_source [put]
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

// @summary upload IEC61850 Model file
// @tags IEC61850
// @accept multipart/form-data
// @produce json
// @param file formData file true "ICD file"
// @success 200 IEC61850Model

// @router /api/v1/iec61850/upload [post]
func UploadIEC61850ModelFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tempFile, err := os.CreateTemp("", "tempfile-")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer os.Remove(tempFile.Name())

	err = c.SaveUploadedFile(file, tempFile.Name())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	scl, err := parser.ParseIEC61850Model(tempFile.Name())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// resp is the response that will be returned to the client
	var resp IEC61850Model

	// using transaction to insert to db first
	session := Engine.NewSession()
	defer session.Close()
	err = session.Begin()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// those are the data objects that will be inserted to db
	iec61850Model := &model.IEC61850Model{
		Name:        scl.Header.ID,
		Description: scl.Header.NameStructure,
	}
	resp.Name = iec61850Model.Name
	resp.Description = iec61850Model.Description

	_, err = session.Insert(iec61850Model)
	if err != nil {
		session.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	for _, ied := range scl.IED {
		for _, accessPoint := range ied.AccessPoint {
			for _, lDevice := range accessPoint.Server.LDevice {
				logicalDevice := model.LogicalDevice{
					Name:        lDevice.Inst,
					ModelID:     iec61850Model.ID,
					Description: "",
				}
				_, err = session.Insert(logicalDevice)
				if err != nil {
					session.Rollback()
					c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
					return
				}

				respLogicalDevice := LogicalDevice{
					Name:        lDevice.Inst,
					Description: "",
					LogicalNode: make([]LogicalNode, 0),
				}

				ln0 := model.LogicalNode{
					Name:            lDevice.LN0.LnClass + lDevice.LN0.Inst,
					LogicalDeviceID: logicalDevice.ID,
					Description:     "",
				}

				_, err = session.Insert(ln0)
				if err != nil {
					session.Rollback()
					c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
					return
				}

				respLN0 := LogicalNode{
					Name: lDevice.LN0.LnClass + lDevice.LN0.Inst,
				}

				respLogicalDevice.LogicalNode = append(respLogicalDevice.LogicalNode, respLN0)

				for _, doi := range lDevice.LN0.DOI {
					dataObject := model.DataObject{
						Name:          doi.Name,
						Description:   "",
						LogicalNodeID: ln0.ID,
					}

					_, err = session.Insert(dataObject)
					if err != nil {
						session.Rollback()
						c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
						return
					}

					for _, dai := range doi.DAI {
						dataAttribute := model.Node{
							DataObjectID: dataObject.ID,
							Name:         dai.Name,
							Value:        dai.Val,
							IEC61850Ref:  logicalDevice.Name + "/" + ln0.Name + "$" + dataObject.Name + "$" + dai.Name,
							DataSource:   "IEC61850",
						}
						_, err = session.Insert(dataAttribute)
						if err != nil {
							session.Rollback()
							c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
							return
						}
					}
				}

				for _, ln := range lDevice.LN {
					logicalNode := model.LogicalNode{
						Name:            ln.LnClass + ln.Inst,
						LogicalDeviceID: logicalDevice.ID,
						Description:     "",
					}

					_, err = session.Insert(logicalNode)
					if err != nil {
						session.Rollback()
						c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
						return
					}

					respLogicalNode := LogicalNode{
						Name: ln.LnClass + ln.Inst,
					}
					respLogicalDevice.LogicalNode = append(respLogicalDevice.LogicalNode, respLogicalNode)

					for _, doi := range ln.DOI {
						dataObject := model.DataObject{
							Name:          doi.Name,
							Description:   "",
							LogicalNodeID: logicalNode.ID,
						}

						_, err = session.Insert(dataObject)
						if err != nil {
							session.Rollback()
							c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
							return
						}

						for _, dai := range doi.DAI {
							dataAttribute := model.Node{
								DataObjectID: dataObject.ID,
								Name:         dai.Name,
								Value:        dai.Val,
								IEC61850Ref:  logicalDevice.Name + "/" + logicalNode.Name + "$" + dataObject.Name + "$" + dai.Name,
								DataSource:   "IEC61850",
							}
							_, err = session.Insert(dataAttribute)
							if err != nil {
								session.Rollback()
								c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
								return
							}
						}
					}
				}
			}
		}
	}
	err = session.Commit()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
