package main

import (
	"log"

	"github.com/Satr10/Otakudesu-Api/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	router.InitRouter(app)

	log.Fatal(app.Listen(":5001"))
}
