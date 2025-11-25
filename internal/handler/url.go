package handler

import (
	"url-shortener-t/internal/model"
	"url-shortener-t/internal/service"

	"github.com/gin-gonic/gin"
)

type UrlHandler struct {
	urlService service.UrlService
}

func NewUrlHandler(urlService service.UrlService) *UrlHandler {
	return &UrlHandler{
		urlService: urlService,
	}
}

func (h UrlHandler) Shorten(ctx *gin.Context) {
	var request model.ShortenRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	response, err := h.urlService.Shorten(request.LongUrl)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, response)
}
