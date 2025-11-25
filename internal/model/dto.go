package model

type ShortenRequest struct {
	LongUrl string `json:"long_url" binding:"required,url"`
}

type ShortenResponse struct {
	ShortUrl string `json:"short_url"`
}
