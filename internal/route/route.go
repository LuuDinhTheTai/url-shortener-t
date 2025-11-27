package route

import (
	"url-shortener-t/internal/handler"

	"github.com/gin-gonic/gin"
)

func NewRouter(r *gin.Engine, indexHanlder *handler.IndexHandler, urlHandler *handler.UrlHandler) {
	r.LoadHTMLGlob("web/template/*")
	r.Static("/css", "./web/css")
	r.Static("/js", "./web/js")

	r.GET("/", indexHanlder.Index)
	r.POST("/shorten", urlHandler.Shorten)
	r.GET("/:pattern", urlHandler.Redirect)
}
