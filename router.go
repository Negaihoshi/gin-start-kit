package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func routerService() {

	router.Static("/public/css/", "./public/css")
	router.Static("/public/js/", "./public/js/")
	router.Static("/public/font/", "./public/fonts/")
	router.Static("/public/img/", "./public/img/")
	router.Static("/public/template/", "./tmp/public/template")

	router.LoadHTMLGlob("templates/index.tmpl")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})
}
