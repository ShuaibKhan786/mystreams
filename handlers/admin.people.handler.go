package handlers

import (
	"fmt"

	"github.com/ShuaibKhan786/mystreams/models"
	"github.com/ShuaibKhan786/mystreams/utils"
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
	personID := getIDfromParms(c)
	if personID == nil {
		return c.SendString("invalid id")
	}

	person := models.ReadPersonByID(c.Context(), personID)
	if person == nil {
		return c.SendString("invalid id")
	}

	fmt.Fprintf(c, "Hi I am %s and I am %s", *person.Name, *person.Gender)
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
	// partial := c.QueryBool("partial")
	page := c.QueryInt("page")
	size := c.QueryInt("size")
	filter := c.Query("filter")
	sort := c.Query("sort")

	var err error
	var filterQuery *utils.FilterQuery
	if len(filter) > 0 {
		filterQuery, err = utils.ParseFilterQuery(filter)
		if err != nil {
			// return some message
		}
	}

	var sortQueries []*utils.SortQuery
	if len(sort) > 0 {
		sortQueries, err = utils.ParseSortQuery(sort)
		if err != nil {
		}
	}

	if size == 0 {
		size = utils.DEFAULT_PAGINATION_SIZE
	}

	people := models.ReadPeople(c.Context(), page, size, filterQuery, sortQueries)
	if people != nil {
	}

	for _, person := range people {
		fmt.Println(*person.ID, *person.Name)
	}

	return nil
}
