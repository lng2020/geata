package service

import (
	"geata/internal/app/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ModbusDetail model.ModbusDetail

// @summary Get ModbusDetail by RuleID
// @tags ModbusDetail
// @param ruleID path int true "Rule ID"
// @success 200 {object} ModbusDetail
// @router /modbus_detail/{ruleID} [get]
func GetModbusDetailByRuleID(c *gin.Context) {
	ID := c.Param("ruleID")
	id, err := strconv.Atoi(ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Rule ID"})
	}
	detail, err := model.GetModbusDetailByRuleID(Engine, int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, detail)
}
