package handler

func UserHandler(
	c *fiber.Ctx,
) error {
	return c.SendString("User Handler")
}