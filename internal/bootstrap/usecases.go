package bootstrap

import (
	"github.com/diyor200/url-shortener/internal/driver/cache"
	"github.com/diyor200/url-shortener/internal/repository"
	"github.com/diyor200/url-shortener/internal/usecase/shortener"
)

type UseCase struct {
	ShortenUC *shortener.UseCase
}

func NewUseCases(repo *repository.Repository, cache *cache.Cache) *UseCase {
	shortenerUC := shortener.New(repo)

	return &UseCase{ShortenUC: shortenerUC}
}
