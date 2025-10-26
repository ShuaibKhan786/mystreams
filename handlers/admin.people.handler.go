package handlers

import (
	"fmt"

	"github.com/ShuaibKhan786/mystreams/models"
	"github.com/ShuaibKhan786/mystreams/utils"
	"github.com/ShuaibKhan786/mystreams/views/components/toast"
	"github.com/ShuaibKhan786/mystreams/views/layouts"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func AdminPeoplePage(c *fiber.Ctx) error {
	partial := c.QueryBool("partial")

	paginationQuery := utils.NewPaginationQuery()
	paginationQuery.Page = utils.DEFAULT_PAGINATION_PAGE
	paginationQuery.Size = utils.DEFAULT_PAGINATION_SIZE

	paginatedResults := models.ReadPeople(c.Context(), &paginationQuery)
	totalPage := utils.CalculatePaginationPages(
		utils.SanitizeNilPointer(paginatedResults.TotalCount),
		utils.DEFAULT_PAGINATION_SIZE,
	)
	tablePagination := layouts.AdminPeoplePaginationTableLayout(
		"/admin/people/list", "",
		utils.DEFAULT_PAGINATION_PAGE,
		totalPage,
		paginatedResults.People,
	)
	peopleLayout := layouts.AdminPeopleLayout(tablePagination)

	if partial {
		return peopleLayout.Render(c.Context(), c)
	}

	return layouts.AdminLayout(peopleLayout).Render(c.Context(), c)
}

func AdminPeopleCreateModal(c *fiber.Ctx) error {
	return layouts.AdminPeopleCreateLayout().Render(c.Context(), c)
}

func AdminPeopleEditModal(c *fiber.Ctx) error {
	personID := getIDfromParms(c)
	if personID == nil {
		return toast.ToastNotification(toast.FAILURE, "Update Person", "Invalid credentials").
			Render(c.Context(), c)
	}

	personResult := models.ReadPersonByID(c.Context(), personID)
	if personResult == nil {
		return toast.ToastNotification(toast.FAILURE, "Update Person", "Invalid credentials").
			Render(c.Context(), c)
	}

	return layouts.AdminPeopleEditLayout(personResult).Render(c.Context(), c)
}

func AdminPeopleRemoveModal(c *fiber.Ctx) error {
	personID := getIDfromParms(c)
	if personID == nil {
		return toast.ToastNotification(toast.FAILURE, "Delete Person", "Invalid credentials").
			Render(c.Context(), c)
	}

	return layouts.AdminPeopleRemoveLayout(*personID).Render(c.Context(), c)
}

func AdminGetPeople(c *fiber.Ctx) error {
	personID := getIDfromParms(c)
	if personID == nil {
		return toast.ToastNotification(toast.FAILURE, "Get Person", "Invalid credentials").
			Render(c.Context(), c)
	}

	person := models.ReadPersonByID(c.Context(), personID)
	if person == nil {
		return toast.ToastNotification(toast.FAILURE, "Get Person", "Invalid credentials").
			Render(c.Context(), c)
	}

	fmt.Fprintf(c, "Hi I am %s and I am %s", *person.Name, *person.Gender)
	return nil
}

func AdminCreatePeople(c *fiber.Ctx) error {
	var person models.Person

	err := c.BodyParser(&person)
	if err != nil {
		log.Error(err)
		return toast.ToastNotification(toast.FAILURE, "Create Person", "Invalid credentials").
			Render(c.Context(), c)
	}

	if !utils.Validate(&person) {
		return toast.ToastNotification(toast.FAILURE, "Create Person", "Incomplete credentials").
			Render(c.Context(), c)
	}

	err = person.Create(c.Context())
	if err != nil {
		log.Error("Failed to run an INSERT Query")
		return toast.ToastNotification(toast.FAILURE, "Create Person", "Person already exists").
			Render(c.Context(), c)
	}

	return toast.ToastNotification(toast.SUCCESS, "Create Person", "Successfully created").
		Render(c.Context(), c)
}

func AdminUpdatePeople(c *fiber.Ctx) error {
	var person models.Person
	err := c.BodyParser(&person)
	if err != nil {
		return toast.ToastNotification(toast.FAILURE, "Update Person", "Invalid credentials").
			Render(c.Context(), c)
	}

	err = person.Update(c.Context())
	if err != nil {
		log.Error("Failed to run an UPDATE query")
		return toast.ToastNotification(toast.FAILURE, "Update Person", "Invalid credentials").
			Render(c.Context(), c)
	}

	return toast.ToastNotification(toast.SUCCESS, "Update Person", "Successfully updated").
		Render(c.Context(), c)
}

func AdminDeletePeople(c *fiber.Ctx) error {
	personID := getIDfromParms(c)
	if personID == nil {
		return toast.ToastNotification(toast.FAILURE, "Delete Person", "Invalid credentials").
			Render(c.Context(), c)
	}

	person := models.Person{ID: personID}
	err := person.Delete(c.Context())
	if err != nil {
		log.Error("Failed to run an DELETE query ")
		return toast.ToastNotification(toast.SUCCESS, "Delete Person", "Invalid credentials").
			Render(c.Context(), c)
	}

	return toast.ToastNotification(toast.SUCCESS, "Delete Person", "Successfully deleted").
		Render(c.Context(), c)
}

// /admin/people/list?page=1&size=10&filter=gender;male&sort=asc,name;desc (dont used this)
// /admin/people/list?page=1&size=10&filter[gender]=male&sort[name]=desc
func AdminListPeople(c *fiber.Ctx) error {
	partial := c.QueryBool("partial")
	mode := c.Query("mode")

	paginationQuery := utils.NewPaginationQuery()

	if err := paginationQuery.Parse(c); err != nil {
		log.Error(err)
		return err
	}

	if paginationQuery.Size == 0 {
		paginationQuery.Size = utils.DEFAULT_PAGINATION_SIZE
	}

	paginatedResults := models.ReadPeople(c.Context(), &paginationQuery)
	totalPage := utils.CalculatePaginationPages(
		utils.SanitizeNilPointer(paginatedResults.TotalCount),
		paginationQuery.Size,
	)

	if mode == "select" && partial {
		return layouts.MovieSelectPeopleListLayout(paginatedResults.People).
			Render(c.Context(), c)
	}

	tablePagination := layouts.AdminPeoplePaginationTableLayout(
		"/admin/people/list",
		fmt.Sprintf("%s%s", paginationQuery.FilterQuery(), paginationQuery.SortQuery()),
		paginationQuery.Page,
		totalPage,
		paginatedResults.People,
	)

	if partial {
		return tablePagination.Render(c.Context(), c)
	}

	peopleLayout := layouts.AdminPeopleLayout(tablePagination)
	return layouts.AdminLayout(peopleLayout).Render(c.Context(), c)
}
