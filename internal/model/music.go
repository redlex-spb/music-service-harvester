package model

type TrackMeta struct {
	Artist      string   `json:"artist"`
	Title       string   `json:"title"`
	Album       string   `json:"album,omitempty"`
	Year        int      `json:"year,omitempty"`
	DurationSec int      `json:"durationSec,omitempty"`
	CoverPath   string   `json:"coverPath,omitempty"`
	Genres      []string `json:"genres,omitempty"`
}
