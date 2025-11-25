package route

import (
	"url-shortener-t/internal/handler"

	"github.com/gin-gonic/gin"
)

func NewRouter(r *gin.Engine, urlHandler *handler.UrlHandler) {
	r.GET("/", handler.Index)
	r.POST("/shorten", urlHandler.Shorten)
}
