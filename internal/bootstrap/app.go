package bootstrap

import (
	"fmt"
	"net/http"

	"github.com/diyor200/url-shortener/internal/config"
	"github.com/diyor200/url-shortener/internal/gateway/rest"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func App() {
	// parse config
	cfg := config.NewConfig()

	// connect db
	client, err := mongo.Connect(options.Client().ApplyURI())
	// connect cache
	// start server
	handler := rest.NewHandler()
	handler.RegisterRoutes()

	loggerMiddleware := handler.LoggingMiddleware(handler.Mux)

	fmt.Println("Starting server on port 8000")
	if err := http.ListenAndServe("localhost:8000", loggerMiddleware); err != nil {
		panic(err)
	}

	// graceful shutdown
	// close db
	// close cache
}
