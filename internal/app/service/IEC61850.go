package service

import (
	"crypto/sha256"
	"fmt"
	"geata/internal/app/model"
	"geata/internal/app/parser"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type IEC61850Model struct {
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
	Name        string `json:"name"`
	Description string `json:"description"`
}

type DataObject struct {
	Name          string          `json:"name"`
	Description   string          `json:"description"`
	DataAttribute []DataAttribute `json:"dataAttribute"`
}

type DataAttribute model.Node

type IEC61850ModelFileParsedResult struct {
	IEC61850Model IEC61850Model `json:"IEC61850Model"`
	FileHashName  string        `json:"fileHashName"`
}

// @summary Get IEC61850 model by ID
// @tags IEC61850
// @produce json
// @param id path int true "IEC61850 model ID"
// @success 200 {object} IEC61850Model
// @router /api/v1/iec61850/model/{id} [get]
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
// @success 200 IE61850ModelFileParsedResult

// @router /api/v1/iec61850/upload [post]
func UploadIEC61850ModleFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fileHashName, err := hashFile(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	tempFile := fmt.Sprintf("tmp/%s", fileHashName)
	err = c.SaveUploadedFile(file, tempFile)
	defer os.Remove(tempFile)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	scl, err := parser.ParseIEC61850Model(tempFile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	parsedIEC61850Model, parsedDataObjects := pruneResult(scl)
	// using transaction to insert to db first
	session := Engine.NewSession()
	defer session.Close()
	err = session.Begin()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// insert IEC61850Model
	iec61850Model := &model.IEC61850Model{
		Name:         parsedIEC61850Model.Name,
		Description:  parsedIEC61850Model.Description,
		FileHashName: fileHashName,
	}
	_, err = session.Insert(iec61850Model)
	if err != nil {
		session.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// insert LogicalDevice
	for _, logicalDevice := range parsedIEC61850Model.LogicalDevice {
		ld := &model.LogicalDevice{
			ModelID:     iec61850Model.ID,
			Name:        logicalDevice.Name,
			Description: logicalDevice.Description,
		}
		_, err = session.Insert(ld)
		if err != nil {
			session.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		// insert LogicalNode
		for _, logicalNode := range logicalDevice.LogicalNode {
			ln := &model.LogicalNode{
				LogicalDeviceID: ld.ID,
				Name:            logicalNode.Name,
			}
			_, err = session.Insert(ln)
			if err != nil {
				session.Rollback()
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			for _, dataObject := range parsedDataObjects {
				do := &model.DataObject{
					LogicalNodeID: ln.ID,
					Name:          dataObject.Name,
					Description:   dataObject.Description,
				}
				_, err = session.Insert(do)
				if err != nil {
					session.Rollback()
					c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
					return
				}
				for _, dataAttribute := range dataObject.DataAttribute {
					node := &model.Node{
						DataObjectID: do.ID,
						Name:         dataAttribute.Name,
						Value:        dataAttribute.Value,
						IEC61850Ref:  dataAttribute.IEC61850Ref,
						DataSource:   dataAttribute.DataSource,
					}
					_, err = session.Insert(node)
					if err != nil {
						session.Rollback()
						c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
						return
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
	c.JSON(http.StatusOK, IEC61850ModelFileParsedResult{
		IEC61850Model: *parsedIEC61850Model,
		FileHashName:  fileHashName,
	})
}

func pruneResult(scl *parser.SCL) (*IEC61850Model, []DataObject) {
	model := &IEC61850Model{
		Name:          scl.Header.ID,
		Description:   scl.Header.NameStructure,
		LogicalDevice: make([]LogicalDevice, 0),
	}

	var dataObjects []DataObject

	for _, ied := range scl.IED {
		for _, accessPoint := range ied.AccessPoint {
			for _, lDevice := range accessPoint.Server.LDevice {
				logicalDevice := LogicalDevice{
					Name:        lDevice.Inst,
					LogicalNode: make([]LogicalNode, 0),
				}

				ln0 := LogicalNode{
					Name: lDevice.LN0.LnClass + lDevice.LN0.Inst,
				}

				for _, doi := range lDevice.LN0.DOI {
					dataObject := DataObject{
						Name:          doi.Name,
						Description:   "",
						DataAttribute: make([]DataAttribute, 0),
					}

					for _, dai := range doi.DAI {
						dataAttribute := DataAttribute{
							Name:  dai.Name,
							Value: dai.Val,
						}
						dataObject.DataAttribute = append(dataObject.DataAttribute, dataAttribute)
					}

					dataObjects = append(dataObjects, dataObject)
				}

				logicalDevice.LogicalNode = append(logicalDevice.LogicalNode, ln0)

				for _, ln := range lDevice.LN {
					logicalNode := LogicalNode{
						Name: ln.LnClass + ln.Inst,
					}

					for _, doi := range ln.DOI {
						dataObject := DataObject{
							Name:          doi.Name,
							Description:   "",
							DataAttribute: make([]DataAttribute, 0),
						}

						for _, dai := range doi.DAI {
							dataAttribute := DataAttribute{
								Name:  dai.Name,
								Value: dai.Val,
							}
							dataObject.DataAttribute = append(dataObject.DataAttribute, dataAttribute)
						}

						dataObjects = append(dataObjects, dataObject)
					}

					logicalDevice.LogicalNode = append(logicalDevice.LogicalNode, logicalNode)
				}

				model.LogicalDevice = append(model.LogicalDevice, logicalDevice)
			}
		}
	}
	return model, dataObjects
}

func hashFile(file *multipart.FileHeader) (string, error) {
	f, err := file.Open()
	if err != nil {
		return "", err
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
