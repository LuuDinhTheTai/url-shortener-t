package service

import (
	"url-shortener-t/internal/model"

	"github.com/gin-gonic/gin"
)

type IUrlService interface {
	Shorten(ctx *gin.Context, longUrl string) (model.ShortenResponse, error)
}
