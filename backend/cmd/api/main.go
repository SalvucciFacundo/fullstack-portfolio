package main

import (
	"log"
	"portfolio-backend/internal/adapters/handlers"
	"portfolio-backend/internal/config"

	"github.com/labstack/echo/v4"
)

func main() {
	// 1. Setup DI
	injector := config.SetupDI()
	defer injector.Shutdown()

	// 2. Setup Server
	e := echo.New()
	e.HTTPErrorHandler = handlers.CustomHTTPErrorHandler

	// 3. Setup Routes
	handlers.SetupRouter(e, injector)

	// 4. Start Server
	port := config.GetEnv("PORT", "8080")
	log.Printf("Starting server on :%s...\n", port)
	if err := e.Start(":" + port); err != nil {
		log.Fatal("Shutting down the server", err)
	}
}
