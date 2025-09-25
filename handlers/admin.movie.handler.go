package handlers

import (
	"encoding/json"
	"fmt"

	"github.com/ShuaibKhan786/mystreams/models"
	"github.com/ShuaibKhan786/mystreams/views/layouts"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func AdminMoviePage(c *fiber.Ctx) error {
	partial := c.QueryBool("partial")

	adminLayout := layouts.AdminMovieLayout()
	if partial {
		return adminLayout.Render(c.Context(), c)
	}

	return layouts.AdminLayout(adminLayout).Render(c.Context(), c)
}

func AdminMovieCreateLayout(c *fiber.Ctx) error {
	partial := c.QueryBool("partial")

	adminCreateLayout := layouts.AdminMovieCreateLayout()
	if partial {
		return adminCreateLayout.Render(c.Context(), c)
	}

	return layouts.AdminLayout(adminCreateLayout).Render(c.Context(), c)
}

func AdminGetMovie(c *fiber.Ctx) error {
	return nil
}

func AdminCreateMovie(c *fiber.Ctx) error {
	var movie models.Movie

	err := c.BodyParser(&movie)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	roles := c.FormValue("json_data")
	err = json.Unmarshal([]byte(roles), &movie)
	if err != nil {
		log.Error(err)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	fmt.Println(roles)

	for _, m := range movie.Roles {
		fmt.Println(*m.Name, " ", *m.RoleType, " ", *m.Character)
	}

	return nil
}

func AdminUpdateMovie(c *fiber.Ctx) error {
	return nil
}

func AdminDeleteMovie(c *fiber.Ctx) error {
	return nil
}

func AdminListMovies(c *fiber.Ctx) error {
	return nil
}
