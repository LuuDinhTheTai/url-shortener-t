package repository

import (
	"url-shortener-t/internal/model"

	"github.com/gin-gonic/gin"
)

type IUrlRepository interface {
	Create(ctx *gin.Context, url model.Url) (*model.Url, error)
	FindById(ctx *gin.Context, id string) (*model.Url, error)
	FindByPattern(ctx *gin.Context, pattern string) (*model.Url, error)
	Update(ctx *gin.Context, id, string, url model.Url) (*model.Url, error)
	Delete(ctx *gin.Context, id string) error
}
