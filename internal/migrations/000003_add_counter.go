package migrations

import (
	"context"
	"github.com/rs/zerolog/log"
	migrate "github.com/xakep666/mongo-migrate"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func addCounter() {
	if err := migrate.Register(
		func(ctx context.Context, db *mongo.Database) error {
			coll := db.Collection("url_mapping").
		},
		func(ctx context.Context, db *mongo.Database) error {

		}); err != nil {
		log.Fatal(err)
	}
}
