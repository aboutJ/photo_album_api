package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func JSONStatusOk(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": message,
		"data":    data,
	})
}

func JSONStatusBadRequest(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{
		"success": false,
		"error":   err.Error(),
	})
}

func JSONStatusInternalServerError(c *gin.Context, err error, message string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"success": false,
		"message": message,
		"reason":  err.Error(),
	})
}
