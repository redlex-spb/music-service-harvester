package model
type SourceReq struct{ URLs []string `json:"urls"`; Lang string `json:"lang,omitempty"`; UserID string `json:"userId,omitempty"`}
type SongID struct{Artist,string;Title,string}
type SongReq struct{ Songs []SongID `json:"songs"`; Lang string `json:"lang,omitempty"`; UserID string `json:"userId,omitempty"`}
type ImageReq struct{ Filename string; Bytes []byte; UserID string}
