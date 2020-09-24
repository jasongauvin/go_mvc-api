package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func WhatDoWeDo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": "I propose TEST",
	})
}
