// Copyright 2023 lng2020. All rights reserved.
// SPDX-License-Identifier: MIT

package web

import (
	_ "geata/internal/docs"
	"log/slog"
	"net/http"
	"strings"
	"time"

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
		tokenString := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte("your_secret_key"), nil
		})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		mapClaims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		role := mapClaims["role"].(float64)
		if model.RoleType(role) != rt {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
			return
		}
		c.Next()
	}
}

func TokenAuthMiddleware(c *gin.Context) {
	tokenString := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("your_secret_key"), nil
	})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	mapClaims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	exp := mapClaims["exp"].(float64)
	if exp < float64(time.Now().Unix()) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token expired"})
		return
	}

	user_id := mapClaims["user_id"].(float64)
	c.Set("user_id", int64(user_id))
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
		user := v1.Group("/user")
		user.Use(TokenAuthMiddleware)
		{
			user.PATCH("/lang", service.UpdateUserLang)
		}

		v1.GET("/station", service.ListStations)
		{
			v1.GET("/station/:station_id", service.GetStation)
			v1.POST("/station", service.CreateStation)
			v1.PUT("/station/:station_id", service.UpdateStation)
			v1.DELETE("/station/:station_id", service.DeleteStation)
		}

		v1.GET("/station/:station_id/config/iec61850", service.GetIEC61850ConfigForStation)
		v1.GET("/station/:station_id/config/mqtt", service.GetMqttConfigForStation)
		v1.GET("/station/:station_id/config/modbus", service.GetModbusConfigForStation)

		{
			v1.GET("/mapping_rule/:rule_id", service.GetMappingRuleByID)
			v1.PUT("/mapping_rule/:rule_id", service.UpdateMappingRule)
		}

		v1.GET("/mqtt_detail/:rule_id", service.GetMqttDetailByRuleID)
		v1.GET("/modbus_detail/:rule_id", service.GetModbusDetailByRuleID)

		v1.GET("/iec61850/model/:model_id", service.GetIEC61850ModelByID)
		v1.GET("/iec61850/model/:model_id/mapping_rule", service.ListMappingRulesByModelID)
		v1.GET("/iec61850/logical_node/:logical_node_id/data_object", service.GetDataObjectByLogicalNodeID)
		v1.GET("/iec61850/node/:id", service.GetNodeByID)
		v1.GET("/iec61850/node/ref/:ref", service.GetNodeByRef)
		v1.PUT("/iec61850/node/:id/data_source", service.UpdateNodeDataSource)
		v1.POST("/iec61850/upload", service.UploadIEC61850ModelFile)

		v1.GET("/audit_log/:log_id", service.GetAuditLogByID)
		v1.GET("/audit_log", service.GetAllAuditLogs)

		managment := v1.Group("/management")
		managment.Use(RoleMiddleware(model.RoleAdmin))
		{
			managment.GET("/user", service.ListUsers)
		}

	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
