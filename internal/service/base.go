package service

import "url-shortener-t/internal/model"

type UrlService interface {
	Shorten(longUrl string) (model.ShortenResponse, error)
}
