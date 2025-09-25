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
	adminPeople.Get("/:id", handlers.AdminGetPeople)
	adminPeople.Post("/", handlers.AdminCreatePeople)
	adminPeople.Put("/:id", handlers.AdminUpdatePeople)
	adminPeople.Delete("/:id", handlers.AdminDeletePeople)
	adminPeople.Get("/list", handlers.AdminListPeople)

	// admin genres endpoint
	adminGenres := admin.Group("/genres", middlewares.AddHTMLResHeader)
	adminGenres.Get("/", handlers.AdminGenrePage)
	adminGenres.Get("/new", handlers.AdminGenreCreateModal)
	adminGenres.Get("/:id", handlers.AdminGetGenre)
	adminGenres.Post("/", handlers.AdminCreateGenre)
	adminGenres.Put("/:id", handlers.AdminUpdateGenre)
	adminGenres.Delete("/:id", handlers.AdminDeleteGenre)
	adminGenres.Get("/list", handlers.AdminListGenre)

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
