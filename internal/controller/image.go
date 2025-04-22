package controller

import (
	"io"
	"net/http"

	"github.com/redlex-spb/music-harvester/internal/model"
)

type ParseImage interface {
	Execute(model.ImageReq) error
}

func NewImageHandler(uc ParseImage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseMultipartForm(32 << 20)
		file, hdr, err := r.FormFile("file")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file.Close()
		data, _ := io.ReadAll(file)
		if err := uc.Execute(model.ImageReq{Filename: hdr.Filename, Bytes: data}); err != nil {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}
		w.WriteHeader(http.StatusAccepted)
	}
}
