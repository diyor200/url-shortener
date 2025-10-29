package bootstrap

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/diyor200/url-shortener/internal/config"
	"github.com/diyor200/url-shortener/internal/gateway/rest"
	_ "github.com/diyor200/url-shortener/internal/migrations"
	"github.com/redis/go-redis/v9"
	migrate "github.com/xakep666/mongo-migrate"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func App() {
	// parse config
	cfg := config.NewConfig()

	// connect dbConn
	dbConn, err := mongo.Connect(options.Client().ApplyURI(cfg.Database.URI()))
	if err != nil {
		log.Fatal("failed to connect to database", err)
		return
	}

	// ping dbConn
	if err = dbConn.Ping(context.Background(), nil); err != nil {
		log.Fatal("failed to ping to database", err)
		return
	}
	log.Println("successfully connected to database")

	db := dbConn.Database(cfg.Database.Name)

	// migrate
	migrate.SetDatabase(db)
	if err = migrate.Up(context.Background(), migrate.AllAvailable); err != nil {
		log.Fatal("failed to apply migrations", err)
		return
	}
	log.Println("successfully applied migrations")

	// connect cache
	casheDB, _ := strconv.Atoi(cfg.Cache.DB)
	cache := redis.NewClient(&redis.Options{
		Addr:     cfg.Cache.Addr,
		Password: cfg.Cache.Password,
		DB:       casheDB,
	})

	// ping
	if err = cache.Ping(context.Background()).Err(); err != nil {
		log.Fatal("failed to ping to cache", err)
		return
	}
	log.Println("successfully connected to cache")

	// start server
	handler := rest.NewHandler()
	handler.RegisterRoutes()

	loggerMiddleware := handler.LoggingMiddleware(handler.Mux)

	log.Println("Starting server on port 8000")
	if err := http.ListenAndServe("localhost:8000", loggerMiddleware); err != nil {
		panic(err)
	}

	// graceful shutdown
	// close dbConn
	if err = dbConn.Disconnect(context.Background()); err != nil {
		log.Fatal("failed to disconnect from database", err)
		return
	}

	// close cache
	if err = cache.Close(); err != nil {
		log.Fatal("failed to close cache", err)
		return
	}

	log.Println("Server stopped")
}
