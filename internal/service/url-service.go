package service

import (
	"fmt"
	"time"
	"url-shortener-t/internal/model"
	"url-shortener-t/internal/repository"
	"url-shortener-t/pkg"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type UrlService struct {
	urlRepository repository.IUrlRepository
}

func NewUrlService(urlRepository repository.IUrlRepository) IUrlService {
	return &UrlService{
		urlRepository: urlRepository,
	}
}

func (u *UrlService) Shorten(ctx *gin.Context, longUrl string) (*model.ShortenResponse, error) {
	newID := bson.NewObjectID()
	pattern := pkg.Encode([]byte(newID.Hex()))

	url := model.Url{
		Id:         newID,
		Domain:     "localhost:8080",
		Pattern:    pattern,
		LongUrl:    longUrl,
		TotalClick: 0,
		CreatedAt:  time.Now(),
		UpdateAt:   time.Now(),
		ExpiredAt:  time.Now().Add(time.Hour * 24 * 7),
	}

	saved, err := u.urlRepository.Save(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("failed to save url: %w", err)
	}

	response := &model.ShortenResponse{
		LongUrl:    saved.LongUrl,
		ShortUrl:   saved.Domain + "/" + saved.Pattern,
		TotalClick: saved.TotalClick,
		CreatedAt:  saved.CreatedAt,
		UpdateAt:   saved.UpdateAt,
		ExpiredAt:  saved.ExpiredAt,
	}

	return response, nil
}

func (u *UrlService) Redirect(ctx *gin.Context, pattern string) (string, error) {
	url, err := u.urlRepository.FindByPattern(ctx, pattern)
	if err != nil {
		return "", fmt.Errorf("failed to find url by pattern: %w", err)
	}

	//url.TotalClick++
	//url.UpdateAt = time.Now()

	//_, err = u.urlRepository.Update(ctx, url)
	//if err != nil {
	//	return "", fmt.Errorf("failed to update url: %w", err)
	//}

	return url.LongUrl, nil
}
