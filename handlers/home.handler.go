package handlers

import (
	"github.com/ShuaibKhan786/mystreams/views/layouts"
	"github.com/gofiber/fiber/v2"
)

func GetHomePage(c *fiber.Ctx) error {
	partial := c.QueryBool("partial")

	homeLayout := layouts.HomeLayout()
	if partial {
		return homeLayout.Render(c.Context(), c)
	}

	return layouts.Layout(homeLayout).Render(c.Context(), c)
}
