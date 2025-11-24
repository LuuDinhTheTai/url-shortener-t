package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Url struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Domain     string             `bson:"domain,omitempty" json:"domain"`
	Alias      string             `bson:"alias,omitempty" json:"alias"`
	LongUrl    string             `bson:"long_url,omitempty" json:"long_url"`
	TotalClick int                `bson:"total_click" json:"total_click"`
	CreatedAt  time.Time          `bson:"created_at,omitempty" json:"created_at"`
	UpdateAt   time.Time          `bson:"update_at,omitempty" json:"update_at"`
	ExpiredAt  time.Time          `bson:"expired_at,omitempty" json:"expired_at"`
}
