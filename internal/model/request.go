package model

// ---- REST input DTOs ----

type SourceReq struct {
	URLs   []string `json:"urls"`
	Lang   string   `json:"lang,omitempty"`
	UserID string   `json:"userId,omitempty"`
}

type SongReq struct {
	Songs  []SongID `json:"songs"`
	Lang   string   `json:"lang,omitempty"`
	UserID string   `json:"userId,omitempty"`
}

type ImageReq struct {
	Filename string
	Bytes    []byte
	UserID   string
}

// helper
type SongID struct {
	Artist string `json:"artist"`
	Title  string `json:"title"`
}
