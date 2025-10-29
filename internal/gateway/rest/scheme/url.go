package scheme

import "time"

type URLRequest struct {
	URL string `json:"url"`
}

type URL struct {
	ID        string `json:"id"`
	ShortURL  string `json:"short_url"`
	LongURL   string `json:"long_url"`
	CreatedAt string `json:"created_at"`
}

func ConvertTimeToString(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}
