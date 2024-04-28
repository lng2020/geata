// Copyright 2023 lng2020. All rights reserved.
// SPDX-License-Identifier: MIT

package web

import (
	_ "geata/internal/docs"

	"geata/internal/app/service"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Geata API v1
// @version 1.0
// @description Geata API v1
// @host localhost:8080
// @BasePath /api/v1
func SetupRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")

	{
		v1.GET("/stations", service.ListStations)
		{
			v1.GET("/stations/:station_id", service.GetStation)
			v1.POST("/stations", service.CreateStation)
			v1.PUT("/stations/:station_id", service.UpdateStation)
			v1.DELETE("/stations/:station_id", service.DeleteStation)
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
