package repository

import (
	"context"
	"errors"

	"github.com/rs/zerolog/log"

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

func (r *Repository) Create(ctx context.Context, data domain.URL) (domain.URL, error) {
	collection := r.db.Database("url_shortener").Collection("url_mapping")
	_, err := collection.
		InsertOne(ctx, url{
			LongURL:   data.Long,
			ShortURL:  data.Short,
			CreatedAt: data.CreatedAt,
		})
	if err != nil {
		if !mongo.IsDuplicateKeyError(err) {
			log.Info().Str("failed to insert url data: err:", err.Error())
			return domain.URL{}, err
		}
	}

	var record url
	err = collection.FindOne(ctx, bson.M{"short_url": data.Short}).Decode(&record)
	if err != nil {
		log.Error().Err(err).Msg("failed to get url data")
		return domain.URL{}, err
	}

	return record.toModel(), nil
}

func (r *Repository) Get(ctx context.Context, data domain.URL) (domain.URL, error) {
	collection := r.db.Database("url_shortener").Collection("url_mapping")
	filter := bson.D{}

	if data.ID != "" {
		filter = append(filter, bson.E{Key: "_id", Value: data.ID})
	}

	if data.Long != "" {
		filter = append(filter, bson.E{Key: "long_url", Value: data.Long})
	}

	if data.Short != "" {
		filter = append(filter, bson.E{Key: "short_url", Value: data.Short})
	}

	cur, err := collection.Find(ctx, filter)
	if err != nil {
		log.Error().Str("failed to find url data: err:", err.Error())
		if errors.Is(err, mongo.ErrNoDocuments) {
			return domain.URL{}, errs.ErrNotFound
		}
		return domain.URL{}, err
	}

	var res []url
	if err = cur.All(ctx, &res); err != nil {
		log.Error().Str("failed to find url data: err:", err.Error())
		return domain.URL{}, err
	}

	if len(res) == 0 {
		return domain.URL{}, errs.ErrNotFound
	}

	// increment
	_, err = collection.UpdateByID(ctx, res[0].ID, bson.D{{"$inc", bson.D{{"counter", 1}}}})
	if err != nil {
		log.Error().Str("failed to update url data: err:", err.Error())
		return domain.URL{}, err
	}

	return res[0].toModel(), nil
}

func (r *Repository) IncrementCounter(ctx context.Context, shortURL string) error {
	collection := r.db.Database("url_shortener").
		Collection("url_mapping")

	_, err := collection.UpdateOne(ctx, bson.M{"short_url": shortURL},
		bson.D{{"$inc", bson.D{{"counter", 1}}}})
	if err != nil {
		log.Error().Str("failed to update url data: err:", err.Error())
		return err
	}

	return nil
}
