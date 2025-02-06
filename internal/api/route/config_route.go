package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sonjinho/otel-deploy-backend/internal/api/handlers"
)

func SetupConfigRoutes(router fiber.Router, handler *handlers.ConfigHandler) {
	configGroup := router.Group("/config")

	configGroup.Get("/:key", handler.GetConfig)
	configGroup.Post("/", handler.SaveConfig)
}
