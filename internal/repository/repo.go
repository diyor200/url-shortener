package repository

import (
	"context"
	"log"

	"github.com/diyor200/url-shortener/internal/domain"
	"github.com/diyor200/url-shortener/internal/errs"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Repository struct {
	db *mongo.Client
}

func NewRepository(db *mongo.Client) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(ctx context.Context, data domain.URL) (string, error) {
	res, err := r.db.
		Database("url_shortener").
		Collection("url_mapping").
		InsertOne(ctx, url{
			LongURL:   data.Long,
			ShortURL:  data.Short,
			CreatedAt: data.CreatedAt,
		})
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return "", errs.ErrDuplicateData
		}
		
		log.Println("failed to insert url data: err:", err)
		return "", err
	}

	id := res.InsertedID.(bson.ObjectID)
	return id.String(), nil
}

func (r *Repository) GetByShortURL(ctx context.Context, shortURL string) (domain.URL, error) {
	cur, err := r.db.
		Database("url_shortener").
		Collection("url_mapping").
		Find(ctx, bson.M{"short_url": shortURL})
	if err != nil {
		log.Println("failed to find url data: err:", err)
		return domain.URL{}, err
	}

	var res []url
	if err = cur.All(ctx, &res); err != nil {
		log.Println("failed to find url data: err:", err)
		return domain.URL{}, err
	}

	if len(res) == 0 {
		return domain.URL{}, errs.ErrNotFound
	}

	return res[0].toModel(), nil
}
