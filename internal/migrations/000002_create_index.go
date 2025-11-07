package migrations

import (
	"context"

	"github.com/rs/zerolog/log"
	migrate "github.com/xakep666/mongo-migrate"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func createIndex() {
	if err := migrate.Register(
		func(ctx context.Context, db *mongo.Database) error {
			_, err := db.Collection("url_mapping").Indexes().CreateOne(ctx, mongo.IndexModel{
				Keys:    bson.D{{Key: "short_url", Value: 1}, {Key: "long_url", Value: 1}},
				Options: options.Index().SetUnique(true),
			})
			if err != nil {
				log.Fatal().Err(err)
				return err
			}

			return nil
		},
		func(ctx context.Context, db *mongo.Database) error {
			err := db.Collection("url_mapping").Indexes().DropOne(ctx, "short_url")
			if err != nil {
				log.Fatal().Err(err)
				return err
			}

			return nil
		},
	); err != nil {
		log.Fatal().Str("failed to register migrations", err.Error())
		return
	}

	log.Info().Str("msg", "setting index to short_url and long_url is successful!")
}
