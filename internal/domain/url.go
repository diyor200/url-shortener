package domain

import "time"

type URL struct {
	ID        string
	Long      string
	Short     string
	CreatedAt time.Time
}
