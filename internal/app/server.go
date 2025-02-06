package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sonjinho/otel-deploy-backend/internal/api/handlers"
	"github.com/sonjinho/otel-deploy-backend/internal/api/route"
)

func SetupServer(app *Application) *fiber.App {
	server := fiber.New()

	configHandler := handlers.NewConfigHandler(app.ConfigService)

	route.SetupRoutes(server, configHandler)
	return server
}
