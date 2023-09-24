// Copyright 2023 lng2020. All rights reserved.
// SPDX-License-Identifier: MIT

package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getClientList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
