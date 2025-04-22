package controller

import (
	"github.com/redlex-spb/music-harvester/internal/usecase"
	"github.com/redlex-spb/music-harvester/pkg"
	"io"
	"net/http"
)

func NewImageHandler(uc *usecase.ParseImage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseMultipartForm(32 << 20)
		file, hdr, _ := r.FormFile("file")
		defer file.Close()
		b, _ := io.ReadAll(file)
		req := pkg.ImageReq{Filename: hdr.Filename, Bytes: b}
		if err := uc.Execute(req); err != nil {
			http.Error(w, err.Error(), 422)
			return
		}
		w.WriteHeader(202)
	}
}
