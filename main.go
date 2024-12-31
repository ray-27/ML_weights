package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	handler "weights/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error occured while loading the .env file")
	}

	//init router
	gin.SetMode(os.Getenv("GIN_MODE"))
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if PORT env is not set
	}

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	handler.Routes(router)

	// Create an HTTP server with the Gin router
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: router,
	}

	// Run the server in a goroutine to allow graceful shutdown
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %s\n", err)
		}
	}()
	fmt.Println("Server is running on port :8080")

	// Wait for interrupt signal to shut down gracefully
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt) // Listen for SIGINT (Ctrl+C)
	<-quit
	fmt.Println("\nShutting down server...")

	// Create a context with a timeout for the shutdown process
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Attempt a graceful shutdown
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %s\n", err)
	}

	fmt.Println("Server exited gracefully.")
}
