package repository

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/diyor200/url-shortener/internal/config"
	"github.com/diyor200/url-shortener/internal/domain"
	"github.com/diyor200/url-shortener/internal/errs"
	"github.com/diyor200/url-shortener/internal/helpers"
	"github.com/go-playground/assert/v2"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// connect db
func connectDB() *mongo.Client {
	cfg := config.NewConfig()
	client, err := db(cfg)
	if err != nil {
		panic(err)
	}

	return client
}

func db(cfg *config.Config) (*mongo.Client, error) {
	dbConn, err := mongo.Connect(options.Client().ApplyURI(cfg.Database.URI()))
	if err != nil {
		return nil, err
	}

	// ping dbConn
	if err = dbConn.Ping(context.Background(), nil); err != nil {
		return nil, err
	}
	log.Println("successfully connected to database")
	return dbConn, nil
}

var client = NewRepository(connectDB())

func TestRepository_Create(t *testing.T) {
	testCases := []struct {
		name    string
		data    domain.URL
		wantErr error
	}{
		{
			name: "success",
			data: domain.URL{
				Long:      "https://hmb.atlassian.net/jira/software/projects/HPL/boards/424",
				Short:     helpers.ShortURL("https://hmb.atlassian.net/jira/software/projects/HPL/boards/424"),
				CreatedAt: time.Now(),
			},
			wantErr: nil,
		},
		{
			name: "duplicate error",
			data: domain.URL{
				Long:      "https://hmb.atlassian.net/jira/software/projects/HPL/boards/424",
				Short:     helpers.ShortURL("https://hmb.atlassian.net/jira/software/projects/HPL/boards/424"),
				CreatedAt: time.Now(),
			},
			wantErr: errs.ErrDuplicateData,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := client.Create(context.Background(), tc.data)
			assert.Equal(t, tc.wantErr, err)
		})
	}

	t.Log("pass!")
}
