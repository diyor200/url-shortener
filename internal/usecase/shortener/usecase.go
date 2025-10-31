package shortener

import (
	"context"
	"errors"
	"time"

	"github.com/diyor200/url-shortener/internal/domain"
	"github.com/diyor200/url-shortener/internal/errs"
	"github.com/diyor200/url-shortener/internal/helpers"
)

type urlRepo interface {
	Create(ctx context.Context, data domain.URL) (string, error)
	GetByShortURL(ctx context.Context, shortURL string) (domain.URL, error)
}

type UseCase struct {
	urlRepo urlRepo
}

func New(urlRepo urlRepo) *UseCase {
	return &UseCase{urlRepo: urlRepo}
}

func (uc *UseCase) Shorten(ctx context.Context, longURL string) (domain.URL, error) {
	// logic here
	data := domain.URL{
		Short:     helpers.ShortURL(longURL),
		CreatedAt: time.Now(),
		Long:      longURL,
	}

	id, err := uc.urlRepo.Create(ctx, data)
	if err != nil {
		// if duplicate get from db and return it
		if errors.Is(err, errs.ErrDuplicateData) {
			return uc.Get(ctx, data.Short)
		}

		return domain.URL{}, err
	}

	return domain.URL{
		ID:        id,
		Long:      longURL,
		Short:     data.Short,
		CreatedAt: data.CreatedAt,
	}, err
}

func (uc *UseCase) Get(ctx context.Context, short string) (domain.URL, error) {
	return uc.urlRepo.GetByShortURL(ctx, short)
}
