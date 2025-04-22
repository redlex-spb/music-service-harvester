package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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
	srv := &http.Server{
		Addr:    ":8080",
		Handler: controller.NewRouter(h),
	}

	go func() {
		log.Println("harvester :8080")
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("server: %v", err)
		}
	}()

	// graceful shutdown
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
	log.Println("shutting down...")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	_ = srv.Shutdown(ctx)
	_ = prod.Close()
}
