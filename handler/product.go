package handler

func ProductHandler(
	c *fiber.Ctx,
) error {
	return c.SendString("Product Handler")
}