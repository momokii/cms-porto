package controller

import (
	"cms-porto/middleware"
	"cms-porto/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

	var InputLogin models.LoginData

	if err := c.ShouldBindJSON(&InputLogin); err != nil {
		middleware.ErrorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"errors":  false,
		"message": "success login",
		"data": gin.H{
			"token": "ini token",
			"type":  "JWT",
		},
	})
}

func CheckToken(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"errors":  false,
		"message": "token valid",
	})
}
