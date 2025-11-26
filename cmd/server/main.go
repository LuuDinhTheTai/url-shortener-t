package main

import (
	"url-shortener-t/internal/config"
	"url-shortener-t/internal/handler"
	"url-shortener-t/internal/repository"
	"url-shortener-t/internal/route"
	"url-shortener-t/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	configuration := config.NewConfiguration()
	cfg, err := configuration.LoadConfig()
	if err != nil {
		panic(err)
	}

	client, err := repository.ConnectDatabase(cfg.Database.URI)
	if err != nil {
		panic(err)
	}

	urlRepository := repository.NewUrlRepository(cfg, client)
	urlService := service.NewUrlService(urlRepository)
	urlHandler := handler.NewUrlHandler(urlService)

	r := gin.Default()
	r.LoadHTMLGlob("web/template/*")
	r.Static("/css", "./web/css")
	r.Static("/js", "./web/js")

	route.NewRouter(r, urlHandler)

	err = r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
