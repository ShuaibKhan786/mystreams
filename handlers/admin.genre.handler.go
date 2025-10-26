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

func AdminGenrePage(c *fiber.Ctx) error {
	partial := c.QueryBool("partial")

	paginationQuery := utils.NewPaginationQuery()
	paginationQuery.Page = utils.DEFAULT_PAGINATION_PAGE
	paginationQuery.Size = utils.DEFAULT_PAGINATION_SIZE

	paginatedResults := models.ReadGenres(c.Context(), &paginationQuery)
	totalPage := utils.CalculatePaginationPages(
		utils.SanitizeNilPointer(paginatedResults.TotalCount),
		paginationQuery.Size,
	)
	tablePagination := layouts.AdminGenrePaginationTableLayout(
		"/admin/genres/list", "",
		utils.DEFAULT_PAGINATION_PAGE,
		totalPage,
		paginatedResults.Genres,
	)

	genreLayout := layouts.AdminGenreLayout(tablePagination)

	if partial {
		return genreLayout.Render(c.Context(), c)
	}

	return layouts.AdminLayout(genreLayout).Render(c.Context(), c)
}

func AdminGenreCreateModal(c *fiber.Ctx) error {
	return layouts.AdminGenreCreateLayout().Render(c.Context(), c)
}

func AdminGenreEditModal(c *fiber.Ctx) error {
	genreID := getIDfromParms(c)
	if genreID == nil {
		return toast.ToastNotification(toast.FAILURE, "Update Genre", "Invalid credentials").
			Render(c.Context(), c)
	}

	genreResult := models.ReadGenreByID(c.Context(), genreID)
	if genreResult == nil {
		return toast.ToastNotification(toast.FAILURE, "Update Genre", "Invalid credentials").
			Render(c.Context(), c)
	}

	return layouts.AdminGenreEditLayout(genreResult).Render(c.Context(), c)
}

func AdminGenreRemoveModal(c *fiber.Ctx) error {
	genreID := getIDfromParms(c)
	if genreID == nil {
		return toast.ToastNotification(toast.FAILURE, "Delete Genre", "Invalid credentials").
			Render(c.Context(), c)
	}

	return layouts.AdminGenreRemoveLayout(*genreID).Render(c.Context(), c)
}

func AdminGetGenre(c *fiber.Ctx) error {
	return nil
}

func AdminCreateGenre(c *fiber.Ctx) error {
	var genre models.Genre

	err := c.BodyParser(&genre)
	if err != nil {
		toast.ToastNotification(toast.FAILURE, "Create Genre", "Invalid credentials").
			Render(c.Context(), c)
	}

	if !utils.Validate(&genre) {
		return toast.ToastNotification(toast.FAILURE, "Create Genre", "Incomplete credentials").
			Render(c.Context(), c)
	}

	err = genre.Create(c.Context())
	if err != nil {
		return toast.ToastNotification(toast.FAILURE, "Create Genre", "Genre already exists").
			Render(c.Context(), c)
	}

	return toast.ToastNotification(toast.SUCCESS, "Create Genre", "Successfully created a genre").
		Render(c.Context(), c)
}

func AdminUpdateGenre(c *fiber.Ctx) error {
	var genre models.Genre
	err := c.BodyParser(&genre)
	if err != nil {
		return toast.ToastNotification(toast.FAILURE, "Update Genre", "Invalid credentials").
			Render(c.Context(), c)
	}

	err = genre.Update(c.Context())
	if err != nil {
		log.Error("Failed to run an UPDATE query")
		return toast.ToastNotification(toast.FAILURE, "Update Genre", "Invalid credentials").
			Render(c.Context(), c)
	}

	return toast.ToastNotification(toast.SUCCESS, "Update Genre", "Successfully updated").
		Render(c.Context(), c)
}

func AdminDeleteGenre(c *fiber.Ctx) error {
	genreID := getIDfromParms(c)
	if genreID == nil {
		return toast.ToastNotification(toast.FAILURE, "Delete Genre", "Invalid credentials").
			Render(c.Context(), c)
	}

	genre := models.Genre{ID: genreID}
	err := genre.Delete(c.Context())
	if err != nil {
		log.Error("Failed to run an DELETE query ")
		return toast.ToastNotification(toast.SUCCESS, "Delete Genre", "Invalid credentials").
			Render(c.Context(), c)
	}

	return toast.ToastNotification(toast.SUCCESS, "Delete Genre", "Successfully deleted").
		Render(c.Context(), c)
}

// /admin/genres/list?page=1&size=10&filter=gender;male&sort=asc,name;desc,created_at
func AdminListGenre(c *fiber.Ctx) error {
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

	paginatedResults := models.ReadGenres(c.Context(), &paginationQuery)
	totalPage := utils.CalculatePaginationPages(
		utils.SanitizeNilPointer(paginatedResults.TotalCount),
		paginationQuery.Size,
	)

	if mode == "select" && partial {
		return layouts.MovieSelectGenreListLayout(paginatedResults.Genres).
			Render(c.Context(), c)
	}

	tablePagination := layouts.AdminGenrePaginationTableLayout(
		"/admin/genres/list",
		fmt.Sprintf("%s%s", paginationQuery.FilterQuery(), paginationQuery.SortQuery()),
		paginationQuery.Page,
		totalPage,
		paginatedResults.Genres,
	)

	if partial {
		return tablePagination.Render(c.Context(), c)
	}

	peopleLayout := layouts.AdminGenreLayout(tablePagination)
	return layouts.AdminLayout(peopleLayout).Render(c.Context(), c)
}
