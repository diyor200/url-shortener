package domain

import "time"

const RateLimit = 20

type URL struct {
	ID        string
	Long      string
	Short     string
	Counter   int
	CreatedAt time.Time
}
