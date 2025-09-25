package handlers

import (
	"github.com/ShuaibKhan786/mystreams/views/layouts"
	"github.com/gofiber/fiber/v2"
)

func GetMoviePage(c *fiber.Ctx) error {
	partial := c.QueryBool("partial")

	movieLayout := layouts.MovieLayout()
	if partial {
		return movieLayout.Render(c.Context(), c)
	}
	return layouts.Layout(movieLayout).Render(c.Context(), c)
}
