package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoveMe(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "I love you",
	})
}
