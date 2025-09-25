package middlewares

import "github.com/gofiber/fiber/v2"

func AddHTMLResHeader(c *fiber.Ctx) error {
	c.Set("Content-Type", "text/html")
	return c.Next()
}
