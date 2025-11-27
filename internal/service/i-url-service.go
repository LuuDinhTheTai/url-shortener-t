package service

import (
	"url-shortener-t/internal/model"

	"github.com/gin-gonic/gin"
)

type IUrlService interface {
	Shorten(ctx *gin.Context, longUrl string) (*model.ShortenResponse, error)
	Redirect(ctx *gin.Context, pattern string) (string, error)
}
