package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jasongauvin/go_mvc-api/api/controllers"
)

func InitializeRouter(r *gin.Engine) {
	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.GET("/doyouloveme", controllers.LoveMe)
			v1.GET("/whatdowedo", controllers.WhatDoWeDo)
		}
	}

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data": "Hello Léo",
		})
	})

}
