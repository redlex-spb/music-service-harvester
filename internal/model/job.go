package model

import "time"

type JobKind string

const (
	JobSources JobKind = "sources"
	JobSongs   JobKind = "songs"
	JobImage   JobKind = "image"
)

type Job struct {
	ID        string    `json:"id"`
	Kind      JobKind   `json:"kind"`
	Payload   any       `json:"payload"`
	CreatedAt time.Time `json:"createdAt"`
	UserID    string    `json:"userId,omitempty"`
}

type SourcePayload struct {
	URLs []string `json:"urls"`
	Lang string   `json:"lang,omitempty"`
}

type SongPayload struct {
	Songs []SongID `json:"songs"`
	Lang  string   `json:"lang,omitempty"`
}

type ImagePayload struct {
	Filename   string `json:"filename"`
	ObjectPath string `json:"objectPath"`
}
