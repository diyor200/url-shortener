package shortener

import (
	"context"
	"errors"
	"github.com/diyor200/url-shortener/internal/errs"
	"time"

	"github.com/diyor200/url-shortener/internal/domain"
	"github.com/diyor200/url-shortener/internal/helpers"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

type urlRepo interface {
	Create(ctx context.Context, data domain.URL) (domain.URL, error)
	Get(ctx context.Context, data domain.URL) (domain.URL, error)
	IncrementCounter(ctx context.Context, shortURL string) error
}

type cache interface {
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	Get(ctx context.Context, key string, val interface{}) error
}

type UseCase struct {
	urlRepo urlRepo
	cache   cache
}

func New(urlRepo urlRepo, c cache) *UseCase {
	return &UseCase{urlRepo: urlRepo, cache: c}
}

func (uc *UseCase) Shorten(ctx context.Context, longURL string) (domain.URL, error) {
	// get from cache
	var v cacheURL
	err := uc.cache.Get(ctx, longURL, &v)
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			log.Error().Err(err).Msg("failed to get from cache")
			return domain.URL{}, err
		}
	}

	if v.Short != "" {
		// increment
		err = uc.urlRepo.IncrementCounter(ctx, v.Short)
		if err != nil {
			log.Error().Err(err)
		}
		return domain.URL{Short: v.Short}, nil
	}

	// get from db if exists
	res, err := uc.urlRepo.Get(ctx, domain.URL{Long: longURL})
	if err != nil {
		if !errors.Is(err, errs.ErrNotFound) {
			log.Error().Err(err)
			return domain.URL{}, err
		}
	}

	if res.ID == "" {
		v.Short = helpers.ShortURL(longURL)
		data := domain.URL{
			Short:     v.Short,
			CreatedAt: time.Now(),
			Long:      longURL,
		}

		res, err = uc.urlRepo.Create(ctx, data)
		if err != nil {
			return domain.URL{}, err
		}
	}

	if res.Counter >= domain.RateLimit {
		// save to cache
		err = uc.cache.Set(ctx, res.Long, &cacheURL{Short: res.Short}, time.Hour*12)
		if err != nil {
			log.Error().Err(err).Msg("failed to set data")
		}
	}

	return res, err
}
