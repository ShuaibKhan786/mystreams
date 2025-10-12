package handlers

import (
	"encoding/json"
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

	people := models.ReadPeople(c.Context(), &paginationQuery)

	tablePagination := layouts.AdminPeoplePaginationTableLayout(
		"", "",
		utils.DEFAULT_PAGINATION_PAGE,
		1,
		people,
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
		return toast.ToastNotification(toast.FAILURE, "Create Person", "Invalid payload").
			Render(c.Context(), c)
	}

	// requiredFields := [...]string{"Name", "Gender"}

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
		return toast.ToastNotification(toast.FAILURE, "Update Person", "Invalid payload").
			Render(c.Context(), c)
	}

	err = person.Update(c.Context())
	if err != nil {
		log.Error("Failed to run an UPDATE query")
		return toast.ToastNotification(toast.FAILURE, "Update Person", "Invalid payload").
			Render(c.Context(), c)
	}

	return toast.ToastNotification(toast.SUCCESS, "Update Person", "Successfully updated ").
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

// /admin/prople?page=1&size=10&filter=gender;male&sort=asc,name;desc,created_at
func AdminListPeople(c *fiber.Ctx) error {
	paginationQuery := utils.NewPaginationQuery()

	if err := paginationQuery.Parse(c); err != nil {
		log.Error(err)
		return err
	}

	if paginationQuery.Size == 0 {
		paginationQuery.Size = utils.DEFAULT_PAGINATION_SIZE
	}

	people := models.ReadPeople(c.Context(), &paginationQuery)
	if people != nil {
	}

	for _, person := range people {
		fmt.Println(*person.ID, *person.Name)
	}

	jsonPeople, err := json.Marshal(&people)
	if err != nil {
		log.Error(err)
	}

	return c.Send(jsonPeople)
}
