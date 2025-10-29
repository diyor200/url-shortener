package repository

import (
	"github.com/diyor200/url-shortener/internal/domain"
	"go.mongodb.org/mongo-driver/v2/bson"
	"time"
)

type url struct {
	ID           bson.ObjectID `bson:"_id,omitempty"`
	LongURL      string        `bson:"long_url"`
	ShortenedURL string        `bson:"shortened_url"`
	CreatedAt    time.Time     `bson:"created_at"`
}

func (u *url) toModel() domain.URL {
	return domain.URL{
		ID:         u.ID.String(),
		Long:       u.LongURL,
		ShortenURL: u.ShortenedURL,
		CreatedAt:  u.CreatedAt,
	}
}
