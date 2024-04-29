package service

import (
	"geata/internal/app/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AuditLog model.AuditLog

// @summary Get audit log by ID
// @tags AuditLog
// @produce json
// @param id path int true "Audit log ID"
// @success 200 {object} AuditLog
// @router /audit_log/{id} [get]
func GetAuditLogByID(c *gin.Context) {
	ID := c.Param("id")
	logID, err := strconv.Atoi(ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log, err := model.GetAuditLogByID(Engine, int64(logID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, log)
}

// @summary Get all audit logs
// @tags AuditLog
// @produce json
// @success 200 {array} AuditLog
// @router /audit_log [get]
func GetAllAuditLogs(c *gin.Context) {
	logs, err := model.GetAllAuditLogs(Engine)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, logs)
}
