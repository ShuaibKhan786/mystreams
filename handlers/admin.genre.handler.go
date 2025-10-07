package handlers

import (
	"github.com/ShuaibKhan786/mystreams/models"
	"github.com/ShuaibKhan786/mystreams/views/layouts"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func AdminGenrePage(c *fiber.Ctx) error {
	partial := c.QueryBool("partial")

	genreLayout := layouts.AdminGenreLayout()

	if partial {
		return genreLayout.Render(c.Context(), c)
	}

	return layouts.AdminLayout(genreLayout).Render(c.Context(), c)
}

func AdminGenreCreateModal(c *fiber.Ctx) error {
	return layouts.AdminGenreCreateLayout().Render(c.Context(), c)
}

func AdminGetGenre(c *fiber.Ctx) error {
	return nil
}

func AdminCreateGenre(c *fiber.Ctx) error {
	var genre models.Genre

	err := c.BodyParser(&genre)
	if err != nil {
		log.Errorf("Failed to parse a body in genre create :  ", err)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	genre.Create(c.Context())

	return nil
}

func AdminUpdateGenre(c *fiber.Ctx) error {
	return nil
}

func AdminDeleteGenre(c *fiber.Ctx) error {
	return nil
}

func AdminListGenre(c *fiber.Ctx) error {
	return nil
}
