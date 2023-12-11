package middleware

import "github.com/gin-gonic/gin"

func ErrorResponse(c *gin.Context, statusCode int, errorMessage string) {
	c.JSON(statusCode, gin.H{
		"errors":  true,
		"message": errorMessage,
	})
	c.Abort()
}
