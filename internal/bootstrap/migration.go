package bootstrap

import (
	"context"
	"github.com/diyor200/url-shortener/internal/config"
	migrate "github.com/xakep666/mongo-migrate"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"log"
)

func migrateDB(cfg *config.Config, dbConn *mongo.Client) error {
	db := dbConn.Database(cfg.Database.Name)

	// migrate
	migrate.SetDatabase(db)
	if err := migrate.Up(context.Background(), migrate.AllAvailable); err != nil {
		return err
	}

	log.Println("successfully applied migrations")
	return nil
}
