package repository

import (
	"time"

	"github.com/diyor200/url-shortener/internal/domain"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type url struct {
	ID        bson.ObjectID `bson:"_id,omitempty"`
	LongURL   string        `bson:"long_url"`
	ShortURL  string        `bson:"short_url"`
	Counter   int           `bson:"counter"`
	CreatedAt time.Time     `bson:"created_at"`
}

func (u *url) toModel() domain.URL {
	return domain.URL{
		ID:        u.ID.String(),
		Counter:   u.Counter,
		Long:      u.LongURL,
		Short:     u.ShortURL,
		CreatedAt: u.CreatedAt,
	}
}
