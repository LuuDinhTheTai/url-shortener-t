package handler

import (
	"github.com/gin-gonic/gin"
)

type IndexHandler struct {
}

func NewIndexHandler() *IndexHandler {
	return &IndexHandler{}
}

func (h *IndexHandler) Index(ctx *gin.Context) {
	ctx.HTML(200, "index.html", nil)
}
