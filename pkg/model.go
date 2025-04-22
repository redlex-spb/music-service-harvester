package pkg

type SourceReq struct {
	URLs []string `json:"urls"`
}
type SongReq struct {
	Songs []string `json:"songs"`
}
type ImageReq struct {
	Filename string
	Bytes    []byte
}
