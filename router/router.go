package router

import (
	"github.com/Satr10/Otakudesu-Api/handlers"
	"github.com/gofiber/fiber/v2"
)

func InitRouter(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/", handlers.ApiIndex)
	api.Get("/home", handlers.HomeApi)
	api.Get("/ongoing/:page?", handlers.OngoingApi)
	api.Get("/completed/:page?", handlers.CompletedApi)
	api.Get("/search/:query", handlers.SearchApi)
	api.Get("/anime/:slug", handlers.AnimeDetailApi)
	api.Get("/episode/:slug", handlers.EpisodeApi)
	api.Get("/genre-list", handlers.GenreLIstApi)
	api.Get("/genre/:slug/:page?", handlers.GenreApi)
}
