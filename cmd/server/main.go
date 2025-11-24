package main

import (
	"url-shortener-t/internal/route"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("web/template/*")
	r.Static("/css", "./web/css")

	route.NewRouter(r)

	r.Run(":8080")
}
