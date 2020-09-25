package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


func Gojira(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"data": "Sauce",
	})
}