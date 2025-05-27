package main

import (
	"log"
	"net/http"
	"time"

	"github.com/migvas/go-tasks-api/config"
	"github.com/migvas/go-tasks-api/database"
	"github.com/migvas/go-tasks-api/internal/api"
	"github.com/migvas/go-tasks-api/internal/services"
)

func main() {
	cfg, err := config.LoadConfig()
	// Error handling for config loading
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	db := database.InitDB(cfg.DSN)

	// Initialize Services (Business Logic Layer)
	taskService := services.NewTaskService(db)
	// Initialize API Handlers (HTTP Layer)
	apiHandlers := api.NewAPIHandlers(taskService)

	// Create a new ServeMux (router)
	mux := http.NewServeMux()

	// Register routes to the mux
	api.SetupRoutes(mux, apiHandlers)

	server := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	log.Printf("Server starting on port %s...", cfg.Port)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server failed to start: %v", err)
	}
}
