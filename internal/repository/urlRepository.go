package repository

import (
	"url-shortener-t/internal/config"
	"url-shortener-t/internal/model"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type UrlRepository struct {
	collection *mongo.Collection
}

func NewUrlRepository(cfg config.Config, client *mongo.Client) IUrlRepository {
	return &UrlRepository{
		collection: client.Database(cfg.Database.Name).Collection("urls"),
	}
}

func (u *UrlRepository) Create(ctx *gin.Context, url model.Url) (*model.Url, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UrlRepository) FindById(ctx *gin.Context, id string) (*model.Url, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UrlRepository) FindByPattern(ctx *gin.Context, pattern string) (*model.Url, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UrlRepository) Update(ctx *gin.Context, id, string, url model.Url) (*model.Url, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UrlRepository) Delete(ctx *gin.Context, id string) error {
	//TODO implement me
	panic("implement me")
}
