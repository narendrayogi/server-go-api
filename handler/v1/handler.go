package v1

import (
	"github.com/gofiber/fiber/v3"
)

func RouteHandler(v1Router fiber.Router) {
	v1Router.Get("/product", ListProductHandler)      // /api/v1/product
	v1Router.Get("/user", ListUserHandler)      // /api/v1/user
}