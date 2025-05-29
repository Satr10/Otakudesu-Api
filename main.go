package main

import (
	"log"
	"time"

	"github.com/Satr10/Otakudesu-Api/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
)

func main() {
	app := fiber.New()
	app.Use(cache.New(cache.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.Query("noCache") == "true"
		},
		Expiration:   1 * time.Minute,
		CacheControl: true,
	}))
	router.InitRouter(app)

	log.Fatal(app.Listen(":5001"))
}
