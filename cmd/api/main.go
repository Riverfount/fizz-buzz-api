package main

import (
	"context"
	"log"
	"os"

	"github.com/Riverfount/fizz-buzz-api/internal/server"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // fallback
	}

	app := server.New(port)
	app.RegisterRoutes()

	serverErr := make(chan error, 1)
	shutdownErr := make(chan error, 1)

	go func() {
		serverErr <- app.Start()
	}()

	go func() {
		shutdownErr <- app.GracefulShutdown(context.Background())
	}()

	select {
	case err := <-serverErr:
		if err != nil {
			log.Fatalf("server error: %v", err)
		}
		log.Println("server stopped")
	case err := <-shutdownErr:
		if err != nil {
			log.Fatalf("shutdown error: %v", err)
		}
		log.Println("graceful shutdown completed")
	}
}
