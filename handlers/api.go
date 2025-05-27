package handlers

import (
	"github.com/Satr10/Otakudesu-Api/scraper"
	"github.com/gofiber/fiber/v2"
)

func ApiIndex(c *fiber.Ctx) error {
	return c.JSON(ApiResponse{
		Status:  "success",
		Message: "Welcome to Otakudesu API, (THIS IS AN UNOFFICIAL API)",
		Data:    nil,
	})
}

// /home to show what animes is in otakdesu home page
func HomeApi(c *fiber.Ctx) error {
	NewScraper := scraper.NewScraper()
	animes, err := NewScraper.HomePage()
	if err != nil {
		return c.JSON(ApiResponse{
			Status:  "failed",
			Message: err.Error(),
			Data:    nil,
		})
	}
	return c.JSON(ApiResponse{
		Status:  "success",
		Message: "",
		Data:    animes,
	})
}
