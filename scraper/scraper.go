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

// NewScraper creates a new Scraper instance.
//
// The instance is configured with a default user agent and a maximum depth
// of 1. This means that the scraper will only traverse one level deep when
// scraping the website.
func NewScraper() *Scraper {
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36"),
		colly.MaxDepth(1),
	)

	return &Scraper{collector: c}
}

// ExtractAnime takes a colly.HTMLElement and a boolean indicating whether the
// anime is ongoing or not, and returns a populated models.Anime instance.
//
// The boolean argument is used to determine whether to extract the schedule
// or rating from the HTML element. If the anime is ongoing, the schedule
// will be extracted; if it's completed, the rating will be extracted.
//
// The method will return a models.Anime instance with the following fields
// populated:
//
//   - Title: The title of the anime, extracted from the element with class
//     "jdlflm".
//   - Episode: The episode number of the anime, extracted from the element with
//     class "epz".
//   - Schedule (if ongoing): The schedule of the anime, extracted from the
//     element with class "epztipe".
//   - Rating (if completed): The rating of the anime, extracted from the element
//     with class "epztipe".
//   - Date: The release date of the anime, extracted from the element with class
//     "newnime".
//   - Slug: The slug of the anime, extracted from the href attribute of the
//     element with class "jdlflm".
//   - Image: The image URL of the anime, extracted from the src attribute of the
//     element with class "lazyload".
//   - URL: The URL of the anime, extracted from the href attribute of the element
//     with class "jdlflm".
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

// HomePage scrapes the home page of Otakudesu and returns a list of models.Anime
// instances. The list contains the anime that are currently being shown on the
// home page.
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

// OngoingPage scrapes the ongoing page of Otakudesu and returns a list of
// models.Anime instances. The list contains the anime that are currently
// being shown on the ongoing page for the given page number.
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

// CompletedPage scrapes the completed page of Otakudesu and returns a list of
// models.Anime instances. The list contains the anime that are currently
// being shown on the completed page for the given page number.
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

// SearchPage scrapes the search page of Otakudesu and returns a list of models.Anime
// instances. The list contains the anime that are currently being shown on the
// search page for the given search query.
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
	err = s.collector.Visit(fmt.Sprintf("%v//?s=%v&post_type=anime", OtakudesuBaseURL, strings.ReplaceAll(searchQuery, " ", "+"))) // replace space with plus
	if err != nil {
		return nil, err
	}
	return animes, nil
}

// AnimePage scrapes the anime page of Otakudesu and returns a models.AnimeDetail instance.
//
// The method takes a slug as a parameter, which is the URL path of the anime
// page. For example, if the URL of the anime page is https://otakudesu.cloud/anime/naruto,
// then the slug is "naruto".
//
// The returned models.AnimeDetail instance contains the following fields:
//
//   - Title: The title of the anime.
//   - JapaneseTitle: The Japanese title of the anime.
//   - Rating: The rating of the anime.
//   - Producer: The producer of the anime.
//   - Type: The type of the anime (e.g. TV, OVA, etc.).
//   - Status: The status of the anime (e.g. Ongoing, Completed, etc.).
//   - EpisodeTotal: The total number of episodes of the anime.
//   - Duration: The duration of each episode of the anime.
//   - ReleaseDate: The release date of the anime.
//   - Studio: The studio that produced the anime.
//   - Genre: The genre of the anime.
//   - Synopsis: The synopsis of the anime.
//   - Episodes: A slice of models.Episode instances, which contain the episode title,
//     slug, and URL.
func (s *Scraper) AnimePage(slug string) (anime models.AnimeDetail, err error) {
	s.collector.OnHTML(`div.infozingle`, func(h *colly.HTMLElement) {
		anime.Title = strings.ReplaceAll(h.ChildText(`p:nth-of-type(1)`), "Judul: ", "")
		anime.JapaneseTitle = strings.ReplaceAll(h.ChildText(`p:nth-of-type(2)`), "Japanese: ", "")
		anime.Rating = strings.ReplaceAll(h.ChildText(`p:nth-of-type(3)`), "Skor: ", "")
		anime.Producer = strings.ReplaceAll(h.ChildText(`p:nth-of-type(4)`), "Produser: ", "")
		anime.Type = strings.ReplaceAll(h.ChildText(`p:nth-of-type(5)`), "Tipe: ", "")
		anime.Status = strings.ReplaceAll(h.ChildText(`p:nth-of-type(6)`), "Status: ", "")
		anime.EpisodeTotal = strings.ReplaceAll(h.ChildText(`p:nth-of-type(7)`), "Total Episode: ", "")
		anime.Duration = strings.ReplaceAll(h.ChildText(`p:nth-of-type(8)`), "Durasi: ", "")
		anime.ReleaseDate = strings.ReplaceAll(h.ChildText(`p:nth-of-type(9)`), "Tanggal Rilis: ", "")
		anime.Studio = strings.ReplaceAll(h.ChildText(`p:nth-of-type(10)`), "Studio: ", "")
		anime.Genre = strings.ReplaceAll(h.ChildText(`p:nth-of-type(11)`), "Genre: ", "")
		// anime.Genre = h.ChildText(`p:nth-of-type(12)`)
	})

	s.collector.OnHTML(`div.sinopc`, func(h *colly.HTMLElement) {
		anime.Synopsis = h.Text
	})

	s.collector.OnHTML(`div.episodelist`, func(h *colly.HTMLElement) {
		h.ForEach(`li`, func(i int, lih *colly.HTMLElement) {
			episode := models.Episode{}
			episode.EpisodeTitle = lih.ChildText(`a`)
			episode.Slug = extractSlug(lih.ChildAttr(`a`, `href`))
			episode.URL = lih.ChildAttr(`a`, `href`)
			anime.Episodes = append(anime.Episodes, episode)
		})
	})

	err = s.collector.Visit(fmt.Sprintf("%v/anime/%v", OtakudesuBaseURL, slug))
	if err != nil {
		return models.AnimeDetail{}, err
	}
	return anime, nil
}

// EpisodeDetailPage scrapes the episode page of Otakudesu and returns a models.EpisodePage
// instance.
//
// The method takes a slug as a parameter, which is the URL path of the episode
// page. For example, if the URL of the episode page is https://otakudesu.cloud/episode/naruto-episode-1,
// then the slug is "naruto-episode-1".
//
// The returned models.EpisodePage instance contains the following fields:
//
//   - Downloads: A slice of models.EpisodeDownloads instances, which contain the
//     quality, size, and download URLs of the episode.
//   - StreamingURL: The streaming URL of the episode, if available.
func (s *Scraper) EpisodeDetailPage(slug string) (episode models.EpisodePage, err error) {
	// find streaming URL
	s.collector.OnHTML(`div.responsive-embed-stream iframe`, func(h *colly.HTMLElement) {
		episodeDL := models.EpisodeDownloads{}
		episodeDL.Quality = "Streaming"
		episodeDL.Size = "N/A" // Size is not available for streaming
		stream := models.Download{}
		stream.Provider = "Streaming URL"
		stream.DownloadURL = h.Attr(`src`)
		episodeDL.Downloads = append(episodeDL.Downloads, stream)
		episode.Downloads = append(episode.Downloads, episodeDL)
	})

	s.collector.OnHTML(`div.download ul li`, func(h *colly.HTMLElement) {
		episodeDL := models.EpisodeDownloads{}
		episodeDL.Quality = h.ChildText(`strong`)
		episodeDL.Size = h.ChildText(`i`)
		h.ForEach(`a`, func(_ int, h *colly.HTMLElement) {
			download := models.Download{}
			download.DownloadURL = h.Attr(`href`)
			download.Provider = h.Text
			episodeDL.Downloads = append(episodeDL.Downloads, download)
		})
		episode.Downloads = append(episode.Downloads, episodeDL)
	})

	err = s.collector.Visit(fmt.Sprintf("%v/episode/%v", OtakudesuBaseURL, slug))
	if err != nil {
		return models.EpisodePage{}, err
	}

	return episode, nil

}

func (s *Scraper) GenresPage() (genres models.Genres, err error) {
	s.collector.OnHTML(`ul.genres li`, func(h *colly.HTMLElement) {
		h.ForEach(`a`, func(i int, h *colly.HTMLElement) {
			genre := models.Genre{}
			genre.Slug = extractSlug(h.Attr(`href`))
			genre.Title = h.Text
			genre.URL = fmt.Sprintf("%v%v", OtakudesuBaseURL, h.Attr(`href`))
			genres.Genres = append(genres.Genres, genre)
		})
	})
	err = s.collector.Visit(fmt.Sprintf("%v/genre-list", OtakudesuBaseURL))
	if err != nil {
		return models.Genres{}, err
	}

	return genres, nil

}

func (s *Scraper) GenrePage(slug string, page string) (animes []models.Anime, err error) {
	s.collector.OnHTML(`div.col-anime`, func(h *colly.HTMLElement) {
		anime := models.Anime{}
		anime.Title = h.ChildText(`div.col-anime-title`)
		anime.Episode = h.ChildText(`col-anime-eps`)
		// anime.Status = ""
		animeRating := h.ChildText(`div.col-anime-rating`)
		anime.Rating = &animeRating
		anime.Slug = extractSlug(h.ChildAttr(`div.col-anime-title a`, `href`))
		anime.Image = h.ChildAttr(`div.col-anime-cover img`, `src`)
		anime.Episode = strings.ReplaceAll(h.ChildText(`div.col-anime-eps`), " Eps", "")
		anime.URL = h.ChildAttr(`div.col-anime-title a`, `href`)
		animes = append(animes, anime)
	})

	err = s.collector.Visit(fmt.Sprintf("%v/genres/%v/page/%v", OtakudesuBaseURL, slug, page))
	if err != nil {
		return nil, err
	}
	return animes, nil
}
