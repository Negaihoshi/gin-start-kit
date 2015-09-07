package main

import "github.com/gin-gonic/gin"

var router *gin.Engine

func main() {

	gin.SetMode(gin.DebugMode)
	// gin.SetMode(gin.TestMode)
	// gin.SetMode(gin.ReleaseMode)
	router = gin.Default()
	routerService()
	router.Run(":3000")
}
