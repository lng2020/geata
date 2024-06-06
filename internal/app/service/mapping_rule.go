package service

import (
	"geata/internal/app/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MappingRule model.MappingRule

type UpdateMappingRuleRequest struct {
	MappingRule  MappingRule        `json:"mappingRule"`
	ModbusDetail model.ModbusDetail `json:"modbusDetail"`
	MQTTDetail   model.MqttDetail   `json:"mqttDetail"`
}

// @summary List all MappingRule for a station
// @tags MappingRule
// @param stationID path int true "Station ID"
// @success 200 {array} MappingRule
// @router /api/v1/station/{model_id}/mapping_rule [get]
func ListMappingRulesByModelID(c *gin.Context) {
	ID := c.Param("model_id")
	model_id, err := strconv.Atoi(ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Model ID"})
	}
	rules, err := model.GetMappingByModelID(Engine, int64(model_id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, rules)
}

// @summary Get a MappingRule by ID
// @tags MappingRule
// @param id path int true "MappingRule ID"
// @success 200 {object} MappingRule
// @router /api/v1/mapping_rule/{id} [get]
func GetMappingRuleByID(c *gin.Context) {
	ID := c.Param("id")
	id, err := strconv.Atoi(ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
	}
	rule, err := model.GetMappingRuleByID(Engine, int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, rule)
}

// @summary Update a MappingRule
// @tags MappingRule
// @param id path int true "MappingRule ID"
// @param rule body UpdateMappingRuleRequest true "MappingRule"
// @success 200 {object} MappingRule
// @router /api/v1/mapping_rule/{rule_id} [put]
func UpdateMappingRule(c *gin.Context) {
	ID := c.Param("rule_id")
	id, err := strconv.Atoi(ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
	}
	var req UpdateMappingRuleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	rule := req.MappingRule
	rule.ID = int64(id)
	session := Engine.NewSession()
	defer session.Close()
	err = session.Begin()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	err = model.UpdateMappingRule(Engine, (*model.MappingRule)(&rule))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	typ := rule.Type
	switch typ {
	case "Modbus":
		req.ModbusDetail.RuleID = rule.ID
		err = model.CreateOrUpdateModbusDetail(Engine, &req.ModbusDetail)
		if err != nil {
			session.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	case "MQTT":
		req.MQTTDetail.RuleID = rule.ID
		err = model.CreateOrUpdateMQTTDetail(Engine, &req.MQTTDetail)
		if err != nil {
			session.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	}

	c.JSON(http.StatusOK, rule)
}

// @summary Delete a MappingRule
// @tags MappingRule
// @param id path int true "MappingRule ID"
// @success 204
// @router /api/v1/mapping_rule/{id} [delete]
func DeleteMappingRule(c *gin.Context) {
	ID := c.Param("id")
	id, err := strconv.Atoi(ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
	}
	rule, err := model.GetMappingRuleByID(Engine, int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	err = model.DeleteMappingRule(Engine, rule)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.Status(http.StatusNoContent)
}
