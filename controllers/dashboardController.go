package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetDataDashboard(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"errors":  false,
		"message": "success get data dashboard",
		"data": gin.H{
			"data": "ini data",
		},
	})
}
