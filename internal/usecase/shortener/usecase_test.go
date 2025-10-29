package shortener

import (
	"testing"
	"unicode"
)

func TestShortURL(t *testing.T) {
	testCases := []struct {
		name string
		url  string
	}{
		{
			name: "first",
			url:  "https://google.com",
		},
		{
			name: "second",
			url:  "https://youtube.com",
		},
		{
			name: "third",
			url:  "https://example.com/very/long/url/with/query?x=1&y=2",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := shortURL(tt.url)
			if len(got) != 7 {
				t.Errorf("expected length 7, got %d", len(got))
			}
			if !isBase62(got) {
				t.Errorf("expected Base62 string, got %s", got)
			}
		})
	}

	t.Log("pass")
}

func isBase62(s string) bool {
	for _, c := range s {
		if !unicode.IsLetter(c) && (c < '0' || c > '9') {
			return false
		}
	}
	return true
}
