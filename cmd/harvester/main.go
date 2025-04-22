package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/redlex-spb/music-harvester/internal/controller"
	"github.com/redlex-spb/music-harvester/internal/gateway"
	"github.com/redlex-spb/music-harvester/internal/usecase"
	"log"
	"net/http"
)

func main() {
	p := gateway.NewKafkaProducer([]string{"kafka:9092"})
	r := chi.NewRouter()
	r.Post("/v1/harvest/sources", controller.NewSourcesHandler(usecase.NewParseSources(p)))
	r.Post("/v1/harvest/songs", controller.NewSongsHandler(usecase.NewParseSongs(p)))
	r.Post("/v1/harvest/image", controller.NewImageHandler(usecase.NewParseImage(p)))
	log.Println("harvester :8080")
	http.ListenAndServe(":8080", r)
}
