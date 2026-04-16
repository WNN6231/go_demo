package main

import (
	"context"
	"exchangeapp/config"
	"exchangeapp/router"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	config.InitConfig()

	r := router.SetupRouter()
	port := config.AppConfig.App.Port

	if port == "" {
		port = ":8080"
	}

	srv := &http.Server{
		Addr:    port,
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown failed: %s\n", err)
	}
	log.Println("Server exited properly")
}
