package models

type Anime struct {
	Title    string  `json:"title"`
	Episode  string  `json:"episode"`
	Schedule *string `json:"schedule"`
	Rating   *string `json:"rating"`
	Date     string  `json:"date"`
	Slug     string  `json:"slug"`
	Image    string  `json:"image"`
	URL      string  `json:"url"`
}
