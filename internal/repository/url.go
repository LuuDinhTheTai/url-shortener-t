package repository

import (
	"fmt"
	"url-shortener-t/internal/config"
	"url-shortener-t/internal/model"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type UrlRepository struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewUrlRepository(cfg *config.Configuration, client *mongo.Client) *UrlRepository {
	return &UrlRepository{
		client:     client,
		collection: client.Database(cfg.Database.Name).Collection("url"),
	}
}

func (u UrlRepository) Create(ctx *gin.Context, url model.Url) model.Url {
	result, err := u.collection.InsertOne(ctx, url)
	if err != nil {
		fmt.Println(err)
	}

	return result.InsertedID.(model.Url)
}

func (u UrlRepository) Find(id string) model.Url {
	//TODO implement me
	panic("implement me")
}

func (u UrlRepository) FindAll() []model.Url {
	//TODO implement me
	panic("implement me")
}

func (u UrlRepository) Update(ctx *gin.Context, id string, url model.Url) model.Url {
	result, err := u.collection.UpdateByID(ctx, id, url)
	if err != nil {
		fmt.Println(err)
	}

	return result.UpsertedID.(model.Url)
}

func (u UrlRepository) Delete(id string) {
	//TODO implement me
	panic("implement me")
}
