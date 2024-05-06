package service

import (
	"geata/internal/app/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MappingRule model.MappingRule

// @summary List all MappingRule for a station
// @tags MappingRule
// @param stationID path int true "Station ID"
// @success 200 {array} MappingRule
// @router /api/v1/stations/{stationID}/mapping_rules [get]
func ListMappingRulesForStation(c *gin.Context) {
	ID := c.Param("stationID")
	stationID, err := strconv.Atoi(ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid station ID"})
	}
	rules, err := model.GetAllMappingRules(Engine, int64(stationID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, rules)
}

// @summary Get a MappingRule by ID
// @tags MappingRule
// @param id path int true "MappingRule ID"
// @success 200 {object} MappingRule
// @router /api/v1/mapping_rules/{id} [get]
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
// @param rule body MappingRule true "MappingRule"
// @success 200 {object} MappingRule
// @router /api/v1/mapping_rules/{id} [put]
func UpdateMappingRule(c *gin.Context) {
	ID := c.Param("id")
	id, err := strconv.Atoi(ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
	}
	var rule MappingRule
	if err := c.ShouldBindJSON(&rule); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	rule.ID = int64(id)
	err = model.UpdateMappingRule(Engine, (*model.MappingRule)(&rule))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, rule)
}

// @summary Delete a MappingRule
// @tags MappingRule
// @param id path int true "MappingRule ID"
// @success 204
// @router /api/v1/mapping_rules/{id} [delete]
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
