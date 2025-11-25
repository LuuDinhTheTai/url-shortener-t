package service

import (
	"url-shortener-t/internal/model"
	"url-shortener-t/internal/repository"
)

type urlService struct {
	urlRepository repository.UrlRepository
}

func (u urlService) Shorten(longUrl string) (model.ShortenResponse, error) {
	var response = model.ShortenResponse{
		ShortUrl: "https://localhost:8080/shortened",
	}

	return response, nil
}

func NewUrlService(urlRepository *repository.UrlRepository) UrlService {
	return &urlService{
		urlRepository: *urlRepository,
	}
}
