package scraper

import (
	"github.com/Satr10/Otakudesu-Api/models"
	"github.com/gocolly/colly"
)

const OtakudesuBaseURL string = "https://otakudesu.cloud"

type Scraper struct {
	collector *colly.Collector
}

func NewScraper() *Scraper {
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36"),
		colly.MaxDepth(1),
	)

	return &Scraper{collector: c}
}

func ExtractAnime(_ int, h *colly.HTMLElement, isOngoing bool) (anime models.Anime) {
	anime.Title = h.ChildText(`h2.jdlflm`)
	anime.Episode = h.ChildText(`div.epz`)
	if isOngoing {
		schedule := h.ChildText(`div.epztipe`)
		anime.Schedule = &schedule
		anime.Rating = nil
	} else {
		rating := h.ChildText(`div.epztipe`)
		anime.Rating = &rating
		anime.Schedule = nil
	}
	anime.Date = h.ChildText(`div.newnime`)
	anime.Slug = h.ChildAttr(`a`, `href`)
	anime.Image = h.ChildAttr(`img`, `src`)
	anime.URL = h.ChildAttr(`a`, `href`)
	return anime
}

func (s *Scraper) HomePage() (animes []models.Anime, err error) {
	s.collector.OnHTML(`div.venz`, func(h *colly.HTMLElement) {
		h.ForEach(`li`, func(i int, lih *colly.HTMLElement) {
			anime := ExtractAnime(i, lih, true)
			animes = append(animes, anime)
		})
	})
	err = s.collector.Visit(OtakudesuBaseURL)
	if err != nil {
		return nil, err
	}
	return animes, nil

}
