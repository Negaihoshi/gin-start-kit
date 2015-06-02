package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
)

var server *gin.Engine
var templates map[string]*template.Template

func main() {

	gin.SetMode(gin.DebugMode)
	// gin.SetMode(gin.TestMode)
	// gin.SetMode(gin.ReleaseMode)
	server = gin.Default()

	server.Static("/public/css/", "./public/css")
	server.Static("/public/js/", "./public/js/")
	server.Static("/public/font/", "./public/fonts/")
	server.Static("/public/img/", "./public/img/")
	server.Static("/public/template", "./tmp/public/template")

	server.Run(":3000")
}
