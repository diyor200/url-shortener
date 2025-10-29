package domain

import "time"

type URL struct {
	ID         string
	Long       string
	ShortenURL string
	CreatedAt  time.Time
}
