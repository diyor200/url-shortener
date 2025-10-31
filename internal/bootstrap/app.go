package bootstrap

import (
	"context"
	"github.com/diyor200/url-shortener/internal/driver/cache"
	"github.com/diyor200/url-shortener/internal/repository"
	"log"
	"net/http"
	"strconv"

	"github.com/diyor200/url-shortener/internal/config"
	"github.com/diyor200/url-shortener/internal/gateway/rest"
	_ "github.com/diyor200/url-shortener/internal/migrations"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func Run() {
	// parse config
	cfg := config.NewConfig()

	// connectDB
	dbConn, err := connectDB(cfg)

	redisCache, err := connectCache(cfg)
	if err != nil {
		log.Fatal("failed to connect cache", err)
	}

	// repo
	repo := repository.NewRepository(dbConn)

	// usecases
	usecases := NewUseCases(repo, redisCache)

	// start server
	handler := rest.NewHandler(usecases.ShortenUC)
	handler.RegisterRoutes()

	loggerMiddleware := handler.LoggingMiddleware(handler.Mux)

	log.Println("Starting server on port 8000")
	if err = http.ListenAndServe(cfg.HOST+":"+cfg.PORT, loggerMiddleware); err != nil {
		log.Fatal(err)
	}

	// graceful shutdown

	// close dbConn
	if err = dbConn.Disconnect(context.Background()); err != nil {
		log.Fatal("failed to disconnect from database", err)
		return
	}

	// close cache
	if err = redisCache.Close(); err != nil {
		log.Fatal("failed to close cache", err)
		return
	}

	log.Println("Server stopped")
}

// connect to database
func connectDB(cfg *config.Config) (*mongo.Client, error) {
	dbConn, err := mongo.Connect(options.Client().ApplyURI(cfg.Database.URI()))
	if err != nil {
		return nil, err
	}

	// ping dbConn
	if err = dbConn.Ping(context.Background(), nil); err != nil {
		return nil, err
	}
	log.Println("successfully connected to database")

	// migrate
	if err = migrateDB(cfg, dbConn); err != nil {
		return nil, err
	}

	return dbConn, nil
}

// connect cache
func connectCache(cfg *config.Config) (*cache.Cache, error) {
	cacheDB, _ := strconv.Atoi(cfg.Cache.DB)
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Cache.Addr,
		Password: cfg.Cache.Password,
		DB:       cacheDB,
	})

	// ping
	if err := client.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}

	log.Println("successfully connected to cache")
	return cache.NewCache(client), nil
}
