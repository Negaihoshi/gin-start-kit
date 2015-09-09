package main

import (
	"net/http"

	"github.com/gin-gonic/contrib/renders/multitemplate"
	"github.com/gin-gonic/gin"
)

func routerService() {
	router.Static("/public/css/", "./public/css")
	router.Static("/public/js/", "./public/js/")
	router.Static("/public/font/", "./public/fonts/")
	router.Static("/public/img/", "./public/img/")
	router.Static("/public/template/", "./tmp/public/template")

	templates := multitemplate.New()
	templates.AddFromFiles("index",
		"templates/layout.tmpl",
		"templates/index.tmpl")

	router.HTMLRender = templates

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index", gin.H{
			"title": "Main website",
		})
	})
}
