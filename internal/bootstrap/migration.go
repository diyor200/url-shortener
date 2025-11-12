package bootstrap

import (
	"context"
	"github.com/diyor200/url-shortener/internal/config"
	"github.com/rs/zerolog/log"
	migrate "github.com/xakep666/mongo-migrate"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func migrateDB(cfg *config.Config, dbConn *mongo.Client) error {
	db := dbConn.Database(cfg.Database.Name)

	// migrate
	migrate.SetDatabase(db)
	if err := migrate.Up(context.Background(), migrate.AllAvailable); err != nil {
		log.Fatal().Err(err).Msg("failed to migrate db")
		return err
	}

	log.Info().Msg("successfully applied migrations")
	return nil
}
