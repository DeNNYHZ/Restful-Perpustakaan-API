package main

import (
	"Restful-Perpustakaan-API/app/handlers"
	"Restful-Perpustakaan-API/app/middleware"
	"Restful-Perpustakaan-API/app/repositories"
	"Restful-Perpustakaan-API/app/services"
	"Restful-Perpustakaan-API/database"
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
	db, err := sql.Open("postgres", "postgres://user:password@localhost/library?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 3. Initialize repositories, services, and handlers
	bookRepository := repositories.NewBookRepository(db)
	reviewRepo := database.NewReviewRepository(db)
	bookService := services.NewBookService(bookRepository)
	bookHandler := handlers.NewBookHandler(bookService)

	// ... initialize other repositories, services, and handlers

	// 4. Initialize router
	router := mux.NewRouter()

	// 5. Register routes
	router.HandleFunc("/books", bookHandler.GetAllBooks).Methods("GET")
	router.HandleFunc("/books/{id}", bookHandler.GetBookByID).Methods("GET")
	router.HandleFunc("/notifications", handlers.GetAllNotifications).Methods("GET")
	router.HandleFunc("/notifications/{id:[0-9]+}", handlers.GetNotificationByID).Methods("GET")
	router.HandleFunc("/notifications", handlers.CreateNotification).Methods("POST")
	router.HandleFunc("/notifications/{id:[0-9]+}", handlers.UpdateNotification).Methods("PUT")
	router.HandleFunc("/notifications/{id:[0-9]+}", handlers.DeleteNotification).Methods("DELETE")
	router.HandleFunc("/notifications/{id:[0-9]+}/read", handlers.MarkNotificationAsRead).Methods("POST")
	router.HandleFunc("/notifications/read", handlers.MarkAllNotificationsAsRead).Methods("POST")
	router.HandleFunc("/notifications/unread/count", handlers.GetUnreadNotificationsCount).Methods("GET")
	router.HandleFunc("/notifications/all/unread/count", handlers.GetAllUnreadNotificationsCount).Methods("GET")
	router.HandleFunc("/notifications/count", handlers.GetAllNotificationsCount).Methods("GET")
	router.HandleFunc("/notifications/name/{name}", handlers.GetNotificationByName).Methods("GET")
	router.HandleFunc("/notifications/category/{category}", handlers.GetNotificationByCategory).Methods("GET")
	router.HandleFunc("/notifications/receiver/{receiver}", handlers.GetNotificationByReceiver).Methods("GET")
	router.HandleFunc("/notifications/sender/{sender}", handlers.GetNotificationBySender).Methods("GET")
	router.HandleFunc("/notifications/status/{status}", handlers.GetNotificationByStatus).Methods("GET")
	router.HandleFunc("/notifications/receiver/{receiver}/status/{status}", handlers.GetNotificationByReceiverAndStatus).Methods("GET")
	router.HandleFunc("/notifications/receiver/{receiver}/category/{category}", handlers.GetNotificationByReceiverAndCategory).Methods("GET")
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
