// Copyright 2023 lng2020. All rights reserved.
// SPDX-License-Identifier: MIT

package web

import (
	_ "geata/internal/docs"
	"log/slog"
	"net/http"

	"geata/internal/app/logger"
	"geata/internal/app/model"
	"geata/internal/app/service"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RoleMiddleware(rt model.RoleType) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte("your_secret_key"), nil
		})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if claims["role"] == float64(rt) {
				c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
				return
			}
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		c.Next()
	}
}

func TokenAuthMiddleware(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("your_secret_key"), nil
	})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	if _, ok := token.Claims.(jwt.MapClaims); !ok || !token.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	c.Next()
}

// @title Geata API v1
// @version 1.0
// @description Geata API v1
// @host localhost:8080
// @BasePath /api/v1
func SetupRouter() *gin.Engine {
	gin.SetMode("release")
	r := gin.New()
	// suppress gin's debug console log
	r.Use(logger.NewGinLoggerMiddleware(slog.Default()))
	r.Use(gin.Recovery())

	v1 := r.Group("/api/v1")

	{
		v1.POST("/login", service.Login)
		v1.POST("/register", service.Register)

		v1.GET("/stations", service.ListStations)
		{
			v1.GET("/stations/:station_id", service.GetStation)
			v1.POST("/stations", service.CreateStation)
			v1.PUT("/stations/:station_id", service.UpdateStation)
			v1.DELETE("/stations/:station_id", service.DeleteStation)
		}

		v1.GET("/stations/:station_id/mapping_rules", service.ListMappingRulesForStation)

		{
			v1.GET("/mapping_rules/:rule_id", service.GetMappingRuleByID)
			v1.PUT("/mapping_rules/:rule_id", service.UpdateMappingRule)
		}

		v1.GET("/mqtt_detail/:rule_id", service.GetMqttDetailByRuleID)
		v1.GET("/modbus_detail/:rule_id", service.GetModbusDetailByRuleID)

		v1.GET("/iec61850/model/:model_id", service.GetIEC61850ModelByID)
		v1.GET("/iec61850/data_object/:object_id", service.GetDataObjectByID)
		v1.GET("/iec61850/data_object/:object_id/node", service.GetNodesByDataObjectID)
		v1.GET("/iec61850/node/:id", service.GetNodeByID)
		v1.GET("/iec61850/node/ref/:ref", service.GetNodeByRef)
		v1.PUT("/iec61850/node/:id/data_source", service.UpdateNodeDataSource)

		v1.GET("/audit_log/:log_id", service.GetAuditLogByID)
		v1.GET("/audit_log", service.GetAllAuditLogs)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
