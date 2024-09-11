package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"Restful-Perpustakaan-API/app/config"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq" // PostgreSQL driver
)

func main() {
	// 1. Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load configuration:", err)
	}

	// 2. Connect to database
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName))
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// 3. Initialize repositories, services, and handlers
	bookRepository := repositories.NewBookRepository(db)
	bookService := services.NewBookService(bookRepository)
	bookHandler := handlers.NewBookHandler(bookService)

	// ... initialize other repositories, services, and handlers

	// 4. Initialize router
	router := mux.NewRouter()

	// 5. Register routes
	router.HandleFunc("/books", bookHandler.GetAllBooks).Methods("GET")
	router.HandleFunc("/books/{id}", bookHandler.GetBookByID).Methods("GET")
	// ... register other routes

	// Apply middleware (e.g., logging, authentication)
	router.Use(middleware.LoggingMiddleware)
	// ... apply other middleware

	// 6. Start server
	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", cfg.ServerAddress, cfg.ServerPort),
		Handler: router,
	}

	go func() {
		log.Printf("Server is running on %s:%d", cfg.ServerAddress, cfg.ServerPort)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Failed to start server:", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
	log.Println("Server exited")
}
