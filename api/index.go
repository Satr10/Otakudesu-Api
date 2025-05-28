// THIS IS FOR VERCEL
package handler

import (
	"net/http"

	"github.com/Satr10/Otakudesu-Api/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

// Handler is the main entry point of the application. Think of it like the main() method
func Handler(w http.ResponseWriter, r *http.Request) {
	// This is needed to set the proper request path in `*fiber.Ctx`
	r.RequestURI = r.URL.String()

	handler().ServeHTTP(w, r)
}

// building the fiber application
func handler() http.HandlerFunc {
	app := fiber.New()

	router.InitRouter(app)

	return adaptor.FiberApp(app)
}
