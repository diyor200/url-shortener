package shortener

import (
	"context"
	"crypto/md5"
	"github.com/diyor200/url-shortener/internal/domain"
	"math/big"
	"time"
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
		ShortenURL: shortURL(longURL),
		CreatedAt:  time.Now(),
		Long:       longURL,
	}

	id, err := uc.urlRepo.Create(ctx, data)
	if err != nil {
		return domain.URL{}, err
	}

	return domain.URL{
		ID:         id,
		Long:       longURL,
		ShortenURL: data.ShortenURL,
		CreatedAt:  data.CreatedAt,
	}, err
}

const base62 = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func shortURL(longURL string) string {
	hash := md5.Sum([]byte(longURL))
	num := new(big.Int).SetBytes(hash[:])

	var short string
	for num.Cmp(big.NewInt(0)) > 0 {
		mod := new(big.Int)
		num.DivMod(num, big.NewInt(62), mod)
		short = string(base62[mod.Int64()]) + short
	}

	return short[:7] // take first 7 chars
}

func (uc *UseCase) Get(ctx context.Context, short string) (domain.URL, error) {
	return uc.urlRepo.GetByShortURL(ctx, short)
}
