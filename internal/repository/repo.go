package repository

import (
	"github.com/diyor200/url-shortener/internal/domain"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Repository struct {
	db *mongo.Client
}

func NewRepository(db *mongo.Client) *Repository {
	return &Repository{db: db}
}

func (r Repository) Create(data domain.URL) (domain.URL, error) {
	//TODO implement me
	panic("implement me")
}
