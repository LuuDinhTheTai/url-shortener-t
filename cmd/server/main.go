package main

import (
	"context"
	"fmt"
	"log"
	"time"
	"url-shortener-t/internal/config"
	"url-shortener-t/internal/model"
	"url-shortener-t/internal/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	cfg := config.LoadConfig()

	_, collection := repository.ConnectDatabase(cfg.Database.URI, cfg.Database.Name)

	newUrl := model.Url{
		Domain:     "facebook.com",
		Alias:      "fb-profile",
		LongUrl:    "https://www.facebook.com/profile.php?id=1000...",
		TotalClick: 0,
		CreatedAt:  time.Now(),
		UpdateAt:   time.Now(),
		ExpiredAt:  time.Now().Add(7 * 24 * time.Hour),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, newUrl)
	if err != nil {
		log.Fatal("Lỗi khi thêm dữ liệu: ", err)
	}

	oid, _ := result.InsertedID.(primitive.ObjectID)
	fmt.Printf("Đã thêm thành công! ID của bản ghi: %s\n", oid.Hex())

}
