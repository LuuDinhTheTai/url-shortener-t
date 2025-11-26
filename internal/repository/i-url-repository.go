package repository

import (
	"url-shortener-t/internal/model"

	"github.com/gin-gonic/gin"
)

type IUrlRepository interface {
	Save(ctx *gin.Context, url model.Url) (*model.Url, error)
	FindByPattern(ctx *gin.Context, pattern string) (*model.Url, error)
}
