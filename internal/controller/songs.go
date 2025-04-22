package controller

import (
	"encoding/json"
	"github.com/redlex-spb/music-harvester/internal/usecase"
	"github.com/redlex-spb/music-harvester/pkg"
	"net/http"
)

func NewSongsHandler(uc *usecase.ParseSongs) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req pkg.SongReq
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		if err := uc.Execute(req); err != nil {
			http.Error(w, err.Error(), 422)
			return
		}
		w.WriteHeader(202)
	}
}
