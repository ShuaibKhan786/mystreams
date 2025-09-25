package handlers

import (
	"github.com/ShuaibKhan786/mystreams/views/layouts"
	"github.com/gofiber/fiber/v2"
)

func GetProfilePage(c *fiber.Ctx) error {
	partial := c.QueryBool("partial")

	profileLayout := layouts.ProfileLayout()
	if partial {
		return profileLayout.Render(c.Context(), c)
	}

	return layouts.Layout(profileLayout).Render(c.Context(), c)
}
