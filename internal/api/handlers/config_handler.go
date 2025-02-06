package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sonjinho/otel-deploy-backend/internal/services"
)

type ConfigHandler struct {
	Service *services.ConfigService
}

func NewConfigHandler(service *services.ConfigService) *ConfigHandler {
	return &ConfigHandler{
		Service: service,
	}
}

func (h *ConfigHandler) GetConfig(c *fiber.Ctx) error {
	key := c.Params("key")
	value, err := h.Service.GetConfig(key)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Config not found"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"key": key, "value": value})
}
