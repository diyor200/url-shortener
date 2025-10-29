package bootstrap

import (
	"github.com/diyor200/url-shortener/internal/repository"
	"github.com/diyor200/url-shortener/internal/usecase/shortener"
	"github.com/redis/go-redis/v9"
)

type UseCase struct {
	ShortenUC *shortener.UseCase
}

func NewUseCases(repo *repository.Repository, cache *redis.Client) *UseCase {
	shortenerUC := shortener.New(repo)

	return &UseCase{ShortenUC: shortenerUC}
}
