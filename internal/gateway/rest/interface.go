package rest

import (
	"context"

	"github.com/diyor200/url-shortener/internal/domain"
)

type shortenUC interface {
	Shorten(ctx context.Context, longURL string) (domain.URL, error)
}
