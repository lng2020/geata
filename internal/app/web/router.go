// Copyright 2023 lng2020. All rights reserved.
// SPDX-License-Identifier: MIT

package web

import (
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

	// @Summary Get a list of clients
	// @Description Get a list of clients
	// @ID get-client-list
	// @Produce  json
	// @Success 200 {array} Client
	// @Router /clients [get]
	r.GET("/clients", getClientList)

	// Serve Swagger UI
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
