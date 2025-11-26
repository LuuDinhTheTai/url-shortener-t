package service

import (
	"url-shortener-t/internal/model"
	"url-shortener-t/internal/repository"

	"github.com/gin-gonic/gin"
)

type UrlService struct {
	urlRepository repository.IUrlRepository
}

func NewUrlService(urlRepository repository.IUrlRepository) IUrlService {
	return &UrlService{
		urlRepository: urlRepository,
	}
}

func (u *UrlService) Shorten(ctx *gin.Context, longUrl string) (model.ShortenResponse, error) {
	result := model.ShortenResponse{ShortUrl: "https://localhost:8080/shorten"}
	return result, nil
}
