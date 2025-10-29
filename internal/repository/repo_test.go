package repository

import (
	"context"
	"github.com/diyor200/url-shortener/internal/config"
	"github.com/diyor200/url-shortener/internal/domain"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"log"
	"testing"
	"time"
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
	_, err := client.Create(context.Background(), domain.URL{
		Long:       "https://test.uz/very/long/url",
		ShortenURL: "/knaksjfhfasa",
		CreatedAt:  time.Now(),
	})
	if err != nil {
		t.Error(err)
	}

	t.Log("successfully created url")
}
