package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/redlex-spb/music-harvester/internal/controller"
	"github.com/redlex-spb/music-harvester/internal/gateway"
	"github.com/redlex-spb/music-harvester/internal/usecase"
)

func main() {
	prod := gateway.NewKafkaProducer([]string{"kafka:9092"})
	h := controller.Handlers{
		Sources: controller.NewSourcesHandler(usecase.NewParseSources(prod)),
		Songs:   controller.NewSongsHandler(usecase.NewParseSongs(prod)),
		Image:   controller.NewImageHandler(usecase.NewParseImage(prod)),
	}
	r := controller.NewRouter(h)
	log.Println("harvester :8080")
	http.ListenAndServe(":8080", r)
}
