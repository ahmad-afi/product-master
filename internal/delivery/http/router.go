package http

import (
	"net/http"
	"product-master/internal/delivery/http/handler"
	"product-master/internal/infrastructure/container"

	"github.com/gofiber/fiber/v2"
)

func SetupRouter(f *fiber.App, cont container.Container) {
	h := handler.SetupHandler(cont)

	f.Get("", healthCheck)
	v1api := f.Group("/api/v1")

	v1api.Get("/category", h.ProductHandler.ListCategory)

	productGroup := v1api.Group("/product")
	{
		productGroup.Get("", h.ProductHandler.ListProduct)
		productGroup.Post("", h.ProductHandler.CreateProduct)
	}
}

func healthCheck(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"data": "Server is up and running",
	})
}
