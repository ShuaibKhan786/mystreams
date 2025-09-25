package handlers

import (
	"github.com/ShuaibKhan786/mystreams/views/layouts"
	"github.com/gofiber/fiber/v2"
)

func GetShowPage(c *fiber.Ctx) error {
	partial := c.QueryBool("partial")

	showLayout := layouts.ShowLayout()

	if partial {
		return showLayout.Render(c.Context(), c)
	}

	return layouts.Layout(showLayout).Render(c.Context(), c)
}
