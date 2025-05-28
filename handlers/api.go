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

func OngoingApi(c *fiber.Ctx) error {
	page := c.Params("page")
	if page == "" {
		page = "1"
	}
	NewScraper := scraper.NewScraper()
	animes, err := NewScraper.OngoingPage(page)
	if err != nil {
		return c.JSON(ApiResponse{
			Status:  "failed",
			Message: "failed to fetch animes",
			Data:    nil,
		})
	}
	return c.JSON(ApiResponse{
		Status:  "success",
		Message: "",
		Data:    animes,
	})
}

func CompletedApi(c *fiber.Ctx) error {
	page := c.Params("page")
	if page == "" {
		page = "1"
	}
	NewScraper := scraper.NewScraper()
	animes, err := NewScraper.CompletedPage(page)
	if err != nil {
		return c.JSON(ApiResponse{
			Status:  "failed",
			Message: "failed to fetch animes",
			Data:    nil,
		})
	}
	return c.JSON(ApiResponse{
		Status:  "success",
		Message: "",
		Data:    animes,
	})
}

func SearchApi(c *fiber.Ctx) error {
	searchQuery := c.Params("query")

	newScraper := scraper.NewScraper()
	result, err := newScraper.SearchPage(searchQuery)
	if err != nil {
		return c.JSON(ApiResponse{
			Status:  "failed",
			Message: "failed to fetch anime",
			Data:    nil,
		})
	}

	return c.JSON(ApiResponse{
		Status:  "success",
		Message: "",
		Data:    result,
	})
}

func AnimeDetailApi(c *fiber.Ctx) error {
	slug := c.Params("slug")

	newScraper := scraper.NewScraper()
	result, err := newScraper.AnimePage(slug)
	if err != nil {
		return c.JSON(ApiResponse{
			Status:  "failed",
			Message: "failed to fetch anime",
			Data:    nil,
		})
	}

	return c.JSON(ApiResponse{
		Status:  "success",
		Message: "",
		Data:    result,
	})

}

func EpisodeApi(c *fiber.Ctx) error {
	slug := c.Params("slug")

	newScraper := scraper.NewScraper()
	result, err := newScraper.EpisodeDetailPage(slug)
	if err != nil {
		return c.JSON(ApiResponse{
			Status:  "failed",
			Message: "failed to fetch anime",
			Data:    nil,
		})
	}
	return c.JSON(ApiResponse{
		Status:  "success",
		Message: "",
		Data:    result,
	})

}

func GenreLIstApi(c *fiber.Ctx) error {
	result, err := scraper.NewScraper().GenresPage()
	if err != nil {
		return c.JSON(ApiResponse{
			Status:  "failed",
			Message: "failed to fetch anime",
			Data:    nil,
		})
	}
	return c.JSON(ApiResponse{
		Status:  "success",
		Message: "",
		Data:    result,
	})
}

func GenreApi(c *fiber.Ctx) error {
	slug := c.Params("slug")
	page := c.Params("page")
	if page == "" {
		page = "1"
	}
	result, err := scraper.NewScraper().GenrePage(slug, page)
	if err != nil {
		return c.JSON(ApiResponse{
			Status:  "failed",
			Message: "failed to fetch anime",
			Data:    nil,
		})
	}
	return c.JSON(ApiResponse{
		Status:  "success",
		Message: "",
		Data:    result,
	})
}
