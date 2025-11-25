package repository

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Repository[T interface{}] interface {
	Create(t T) T
	Find(id string) T
	FindAll() []T
	Update(id string, t T) T
	Delete(id string)
}

func ConnectDatabase(uri string) (*mongo.Client, error) {
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	return client, nil
}
