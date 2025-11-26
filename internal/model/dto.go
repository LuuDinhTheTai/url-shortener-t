package model

import "time"

type ShortenRequest struct {
	LongUrl string `json:"long_url" binding:"required,url"`
}

type ShortenResponse struct {
	LongUrl    string    `json:"long_url"`
	ShortUrl   string    `json:"short_url"`
	TotalClick int       `json:"total_click"`
	CreatedAt  time.Time `json:"created_at"`
	UpdateAt   time.Time `json:"update_at"`
	ExpiredAt  time.Time `json:"expired_at"`
}
