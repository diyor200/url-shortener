package rest

import "github.com/diyor200/url-shortener/internal/domain"

type shortenUC interface {
	Shorten(longURL string) (domain.URL, error)
}
