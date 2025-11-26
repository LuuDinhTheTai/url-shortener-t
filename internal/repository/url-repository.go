package repository

import (
	"errors"
	"url-shortener-t/internal/config"
	"url-shortener-t/internal/model"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
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

func (u *UrlRepository) Save(ctx *gin.Context, url model.Url) (*model.Url, error) {
	_, err := u.collection.InsertOne(ctx, url)
	if err != nil {
		return nil, err
	}
	return &url, nil
}

func (u *UrlRepository) FindByPattern(ctx *gin.Context, pattern string) (*model.Url, error) {
	var url model.Url

	err := u.collection.FindOne(ctx, bson.M{"pattern": pattern}).Decode(&url)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}

	return &url, nil
}
