package controller

import (
	"encoding/json"
	"net/http"

	"github.com/redlex-spb/music-harvester/internal/model"
	"github.com/redlex-spb/music-harvester/internal/usecase"
)

func NewSongsHandler(uc *usecase.ParseSongs) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req model.SongReq
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := uc.Execute(req); err != nil {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}
		w.WriteHeader(http.StatusAccepted)
	}
}
