package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/example/cipher-arsenal/internal/config"
	storepkg "github.com/example/cipher-arsenal/internal/store"
	api "github.com/example/cipher-arsenal/pkg/api"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	// Initialize DB
	db, err := storepkg.New(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("failed to initialize db: %v", err)
	}
	defer storepkg.Close(db)

	// Gin
	router := gin.Default()

	// Register handlers
	h := api.NewHandler(db, cfg)
	h.RegisterRoutes(router.Group("/v1"))

	// Server with graceful shutdown
	srv := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	log.Printf("server running on :%s", cfg.Port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("server exiting")
}
