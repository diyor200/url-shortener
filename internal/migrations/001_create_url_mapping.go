package migrations

import (
	"context"
	"log"

	migrate "github.com/xakep666/mongo-migrate"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func init() {
	if err := migrate.Register(
		func(ctx context.Context, db *mongo.Database) error {
			log.Println("dbname = ", db.Name())
			return db.CreateCollection(ctx, "url_mapping", nil)
		},
		func(ctx context.Context, db *mongo.Database) error {
			return db.Collection("url_mapping").Drop(ctx)
		}); err != nil {
		log.Fatal("failed to register migrations", err)
	}

	log.Println("✅ Migration registered: create url_mapping collection")
}
