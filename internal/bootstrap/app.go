package bootstrap

import (
	"fmt"
	"github.com/diyor200/url-shortener/internal/gateway/rest"
	"net/http"
)

func App() {
	// parse config
	// connect db
	// connect cache
	// start server
	handler := rest.NewHandler()
	handler.RegisterRoutes()

	loggerMiddleware := handler.LoggingMiddleware(handler.Mux)

	fmt.Println("Starting server on port 8080")
	if err := http.ListenAndServe("localhost:8080", loggerMiddleware); err != nil {
		panic(err)
	}

	// graceful shutdown
	// close db
	// close cache
}
