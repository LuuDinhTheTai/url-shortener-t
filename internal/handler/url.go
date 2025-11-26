package handler

import (
	"net/http"
	"url-shortener-t/internal/model"
	"url-shortener-t/internal/service"

	"github.com/gin-gonic/gin"
)

type UrlHandler struct {
	urlService service.IUrlService
}

func NewUrlHandler(urlService service.IUrlService) *UrlHandler {
	return &UrlHandler{
		urlService: urlService,
	}
}

func (h *UrlHandler) Shorten(ctx *gin.Context) {
	var request model.ShortenRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := h.urlService.Shorten(ctx, request.LongUrl)
	if err != nil {
		// TODO log err
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response)
}
