package main

import (
	"log"
	"github.com/Kunal726/marketmosaic-product-service-go/cmd/marketmosaic-product-service/app"
	"github.com/Kunal726/marketmosaic-product-service-go/pkg/api/router"
	"os"

	commonApp "github.com/Kunal726/market-mosaic-common-lib-go/pkg/app"
	"github.com/Kunal726/market-mosaic-common-lib-go/pkg/auth"
	authMiddleware "github.com/Kunal726/market-mosaic-common-lib-go/pkg/middleware/auth"
	"github.com/Kunal726/market-mosaic-common-lib-go/pkg/server"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	// Initialize application
	application, err := commonApp.NewApplication()
	if err != nil {
		return err
	}
	defer application.Cleanup()

	// Initialize dependencies
	repositories := app.NewRepositories(application.DB, application.RedisManager)
	services := app.NewServices(repositories)
	handlers := app.NewHandlers(services)

	// Initialize auth client and middleware
	authClient := auth.NewClient(os.Getenv("AUTH_SERVICE_URL"), application.Logger)
	authMiddleware := authMiddleware.NewMiddleware(authClient, application.Logger)

	// Initialize server
	server := server.NewServer(application.Logger, os.Getenv("PORT"))

	// Setup routes
	router.RegisterRoutes(server.Engine(), handlers.Product, authMiddleware)

	// Start server
	return server.Start()
}
