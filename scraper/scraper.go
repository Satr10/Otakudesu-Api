package scraper

import (
	"fmt"
	"strings"

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

func ExtractAnime(h *colly.HTMLElement, isOngoing bool) (anime models.Anime) {
	anime.Title = h.ChildText(`h2.jdlflm`)
	anime.Episode = h.ChildText(`div.epz`)
	if isOngoing {
		schedule := h.ChildText(`div.epztipe`)
		anime.Schedule = &schedule
		anime.Status = "Ongoing"
		anime.Rating = nil
	} else {
		rating := h.ChildText(`div.epztipe`)
		anime.Rating = &rating
		anime.Status = "Completed"
		anime.Schedule = nil
	}
	anime.Date = h.ChildText(`div.newnime`)
	anime.Slug = extractSlug(h.ChildAttr(`a`, `href`))
	anime.Image = h.ChildAttr(`img`, `src`)
	anime.URL = h.ChildAttr(`a`, `href`)
	return anime
}

func (s *Scraper) HomePage() (animes []models.Anime, err error) {
	s.collector.OnHTML(`div.venz`, func(h *colly.HTMLElement) {
		h.ForEach(`li`, func(_ int, lih *colly.HTMLElement) {
			anime := ExtractAnime(lih, true)
			animes = append(animes, anime)
		})
	})
	err = s.collector.Visit(OtakudesuBaseURL)
	if err != nil {
		return nil, err
	}
	return animes, nil

}

func (s *Scraper) OngoingPage(page string) (animes []models.Anime, err error) {
	s.collector.OnHTML(`div.detpost`, func(h *colly.HTMLElement) {
		anime := ExtractAnime(h, true)
		animes = append(animes, anime)
	})
	err = s.collector.Visit(fmt.Sprintf("%v/ongoing-anime/page/%v", OtakudesuBaseURL, page))
	if err != nil {
		return nil, err
	}
	return animes, nil
}

func (s *Scraper) CompletedPage(page string) (animes []models.Anime, err error) {
	s.collector.OnHTML(`div.detpost`, func(h *colly.HTMLElement) {
		anime := ExtractAnime(h, false)
		animes = append(animes, anime)
	})
	err = s.collector.Visit(fmt.Sprintf("%v/complete-anime/page/%v", OtakudesuBaseURL, page))
	if err != nil {
		return nil, err
	}
	return animes, nil
}

func (s *Scraper) SearchPage(searchQuery string) (animes []models.Anime, err error) {
	s.collector.OnHTML(`ul.chivsrc`, func(h *colly.HTMLElement) {
		h.ForEach(`li`, func(i int, liH *colly.HTMLElement) {
			anime := models.Anime{}
			anime.Title = liH.ChildText(`h2`)
			anime.Status = strings.ReplaceAll(liH.ChildText(`div:nth-of-type(2)`), "Status : ", "")
			animeRating := strings.ReplaceAll(liH.ChildText(`div:nth-of-type(3)`), "Rating : ", "")
			anime.Rating = &animeRating
			anime.Slug = extractSlug(liH.ChildAttr(`a`, `href`))
			anime.Image = liH.ChildAttr(`img`, `src`)
			anime.URL = liH.ChildAttr(`a`, `href`)
			animes = append(animes, anime)
		})
	})
	err = s.collector.Visit(fmt.Sprintf("%v//?s=%v&post_type=anime", OtakudesuBaseURL, strings.ReplaceAll(searchQuery, " ", "+")))
	if err != nil {
		return nil, err
	}
	return animes, nil
}
