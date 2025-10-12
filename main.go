package main

import (
	"github.com/ShuaibKhan786/mystreams/handlers"
	"github.com/ShuaibKhan786/mystreams/middlewares"
	"github.com/ShuaibKhan786/mystreams/services/database"
	"github.com/gofiber/fiber/v2"

	_ "github.com/lib/pq"
)

func main() {
	db, err := database.ConnectPSQL()
	if err != nil {
		return
	}
	defer db.Close()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/home", fiber.StatusSeeOther)
	})

	app.Static("/static", "./static/")

	// admin endpoint
	admin := app.Group("/admin")
	admin.Get("/", func(c *fiber.Ctx) error {
		// later route to dashboard
		return c.Redirect("/admin/movies", fiber.StatusSeeOther)
	})

	// admin movie endpoint
	adminMovies := admin.Group("/movies", middlewares.AddHTMLResHeader)
	adminMovies.Get("/", handlers.AdminMoviePage)
	adminMovies.Get("/new", handlers.AdminMovieCreateLayout)
	adminMovies.Get("/:id", handlers.AdminGetMovie)
	adminMovies.Post("/", handlers.AdminCreateMovie)
	adminMovies.Put("/:id", handlers.AdminUpdateMovie)
	adminMovies.Delete("/:id", handlers.AdminDeleteMovie)
	adminMovies.Get("/list", handlers.AdminListMovies)

	// admin person/people endpoint
	adminPeople := admin.Group("/people", middlewares.AddHTMLResHeader)
	adminPeople.Get("/", handlers.AdminPeoplePage)
	adminPeople.Get("/new", handlers.AdminPeopleCreateModal)
	adminPeople.Get("/:id/edit", handlers.AdminPeopleEditModal)
	adminPeople.Get("/:id/remove", handlers.AdminPeopleRemoveModal)
	adminPeople.Get("/list", handlers.AdminListPeople)
	adminPeople.Post("/", handlers.AdminCreatePeople)
	adminPeople.Get("/:id", handlers.AdminGetPeople)
	adminPeople.Put("/:id", handlers.AdminUpdatePeople)
	adminPeople.Delete("/:id", handlers.AdminDeletePeople)

	// admin genres endpoint
	adminGenres := admin.Group("/genres", middlewares.AddHTMLResHeader)
	adminGenres.Get("/", handlers.AdminGenrePage)
	adminGenres.Get("/new", handlers.AdminGenreCreateModal)
	adminGenres.Get("/list", handlers.AdminListGenre)
	adminGenres.Post("/", handlers.AdminCreateGenre)
	adminGenres.Get("/:id", handlers.AdminGetGenre)
	adminGenres.Put("/:id", handlers.AdminUpdateGenre)
	adminGenres.Delete("/:id", handlers.AdminDeleteGenre)

	// home endpoint
	home := app.Group("/home", middlewares.AddHTMLResHeader)
	home.Get("/", handlers.GetHomePage)

	// movie endpoint
	movies := app.Group("/movies", middlewares.AddHTMLResHeader)
	movies.Get("/", handlers.GetMoviePage)

	// shows endpoint
	shows := app.Group("/shows", middlewares.AddHTMLResHeader)
	shows.Get("/", handlers.GetShowPage)

	// profile endpoint
	profile := app.Group("/profile", middlewares.AddHTMLResHeader)
	profile.Get("/", handlers.GetProfilePage)

	app.Listen(":4000")
}
