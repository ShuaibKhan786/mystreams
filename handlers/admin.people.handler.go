package handlers

import (
	"github.com/ShuaibKhan786/mystreams/models"
	"github.com/ShuaibKhan786/mystreams/views/layouts"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func AdminPeoplePage(c *fiber.Ctx) error {
	partial := c.QueryBool("partial")

	peopleLayout := layouts.AdminPeopleLayout()

	if partial {
		return peopleLayout.Render(c.Context(), c)
	}

	return layouts.AdminLayout(peopleLayout).Render(c.Context(), c)
}

func AdminPeopleCreateModal(c *fiber.Ctx) error {
	return layouts.AdminPeopleCreateLayout().Render(c.Context(), c)
}

func AdminGetPeople(c *fiber.Ctx) error {
	return nil
}

func AdminCreatePeople(c *fiber.Ctx) error {
	var person models.Person

	err := c.BodyParser(&person)
	if err != nil {
		log.Error(err)
		return nil
	}

	// requiredFields := [...]string{"Name", "Gender"}
	return person.Create(c.Context())
}

func AdminUpdatePeople(c *fiber.Ctx) error {
	return nil
}

func AdminDeletePeople(c *fiber.Ctx) error {
	return nil
}

func AdminListPeople(c *fiber.Ctx) error {
	return nil
}
