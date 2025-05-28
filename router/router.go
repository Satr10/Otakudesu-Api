package router

import (
	"github.com/Satr10/Otakudesu-Api/handlers"
	"github.com/gofiber/fiber/v2"
)

func InitRouter(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/", handlers.ApiIndex)
	api.Get("/home", handlers.HomeApi)
	api.Get("/ongoing/:page", handlers.OngoingApi)
	api.Get("/completed/:page", handlers.CompletedApi)
}
