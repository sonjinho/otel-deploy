package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sonjinho/otel-deploy-backend/internal/api/handlers"
)

func SetupRoutes(app *fiber.App, configHandler *handlers.ConfigHandler) {
	api := app.Group("/api/v1")

	SetupConfigRoutes(api, configHandler)

}
