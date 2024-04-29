package service

import (
	"geata/internal/app/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MqttDetail model.MqttDetail

// @summary Get MqttDetail by RuleID
// @tags MqttDetail
// @param ruleID path int true "Rule ID"
// @success 200 {object} MqttDetail
// @router /mqtt_detail/{ruleID} [get]
func GetMqttDetailByRuleID(c *gin.Context) {
	ID := c.Param("ruleID")
	ruleID, err := strconv.Atoi(ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Rule ID"})
	}
	detail, err := model.GetMqttDetailByRuleID(Engine, int64(ruleID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, detail)
}
