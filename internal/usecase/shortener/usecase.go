package shortener

import "github.com/diyor200/url-shortener/internal/domain"

type urlRepo interface {
	Create(data domain.URL) (domain.URL, error)
}

type UseCase struct {
	urlRepo urlRepo
}

func New(urlRepo urlRepo) *UseCase {
	return &UseCase{urlRepo: urlRepo}
}

func (u *UseCase) Shorten(longURL string) (domain.URL, error) {
	// logic here
	return u.urlRepo.Create(domain.URL{})
}
