package models

type Anime struct {
	Title    string  `json:"title"`
	Episode  string  `json:"episode"`
	Status   string  `json:"status"`
	Schedule *string `json:"schedule"`
	Rating   *string `json:"rating"`
	Date     string  `json:"date"`
	Slug     string  `json:"slug"`
	Image    string  `json:"image"`
	URL      string  `json:"url"`
}

type AnimeDetail struct {
	Title         string    `json:"title"`
	JapaneseTitle string    `json:"japaneseTitle,omitempty"`
	Rating        string    `json:"rating,omitempty"`
	Producer      string    `json:"producer,omitempty"`
	Type          string    `json:"type,omitempty"` // anime quality type i.e BD or WEBDL
	Status        string    `json:"status,omitempty"`
	EpisodeTotal  string    `json:"episodeTotal,omitempty"`
	Duration      string    `json:"duration,omitempty"`
	ReleaseDate   string    `json:"releaseDate,omitempty"`
	Studio        string    `json:"studio,omitempty"`
	Genre         string    `json:"genre,omitempty"`
	Synopsis      string    `json:"synopsis"`
	Episodes      []Episode `json:"episodes,omitempty"`
}

type Episode struct {
	EpisodeTitle string `json:"episodeTitle"`
	Slug         string `json:"slug"`
	URL          string `json:"url"`
}
