package main

import (
	"github.com/gin-gonic/gin"
	. "raspberry/apis"
	"raspberry/middleware"
)

func initRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.Logger())
	router.GET("/", IndexAPI)
	return router
}
