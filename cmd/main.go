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
	"Restful-Perpustakaan-API/app/handlers"
	"Restful-Perpustakaan-API/app/middleware"
	"Restful-Perpustakaan-API/app/repositories"
	"Restful-Perpustakaan-API/app/services"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq" // PostgreSQL driver
)

// main is the entry point of the application.
func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("postgres", cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repos := initializeRepositories(db)
	services := initializeServices(repos, cfg)
	handlers := initializeHandlers(services)

	router := mux.NewRouter()
	registerRoutes(router, handlers)

	log.Fatal(startServer(router, cfg))
}

// initializeRepositories initializes the repositories with the given database connection.
func initializeRepositories(db *sql.DB) map[string]repositories.Repository {
	repos := make(map[string]repositories.Repository)

	repos["book"] = repositories.NewBookRepository(db)
	repos["member"] = repositories.NewMemberRepository(db)
	repos["loan"] = repositories.NewLoanRepository(db)
	repos["notification"] = repositories.NewNotificationRepository(db)
	repos["review"] = repositories.NewReviewRepository(db)

	return repos
}

// initializeServices initializes the services with the given repositories and configuration.
func initializeServices(repos map[string]repositories.Repository, cfg config.Config) map[string]services.Service {
	services := make(map[string]services.Service)

	services["book"] = services.NewBookService(repos["book"].(repositories.BookRepository))
	services["member"] = services.NewMemberService(repos["member"])
	services["loan"] = services.NewLoanService(repos["loan"])
	services["notification"] = services.NewNotificationService(repos["notification"])
	services["review"] = services.NewReviewService(repos["review"])
	services["auth"] = services.NewAuthService(repos["member"], []byte(cfg.JWTSecretKey))
	services["admin"] = services.NewAdminService(repos["member"], repos["book"], repos["member"], repos["loan"])

	return services
}

// initializeHandlers initializes the handlers with the given services.
func initializeHandlers(services map[string]services.Service) map[string]handlers.Handler {
	handlers := make(map[string]handlers.Handler)

	handlers["book"] = handlers.NewBookHandler(services["book"])
	handlers["member"] = handlers.NewMemberHandler(services["member"])
	handlers["loan"] = handlers.NewLoanHandler(services["loan"])
	handlers["notification"] = handlers.NewNotificationHandler(services["notification"])
	handlers["review"] = handlers.NewReviewHandler(services["review"])
	handlers["auth"] = handlers.NewAuthHandler(services["auth"])
	handlers["admin"] = handlers.NewAdminHandler(services["admin"])

	return handlers
}

// registerRoutes registers the routes with the given router and handlers.
func registerRoutes(router *mux.Router, handlers map[string]handlers.Handler) {
	// Book routes
	router.HandleFunc("/books", handlers["book"].GetAllBooks).Methods("GET")
	router.HandleFunc("/books/{id}", handlers["book"].GetBookByID).Methods("GET")
	router.HandleFunc("/books/search", handlers["book"].SearchBooks).Methods("GET")
	router.HandleFunc("/books/{id}/reviews", handlers["review"].GetReviewsForBook).Methods("GET")

	// Member routes
	router.HandleFunc("/members", handlers["member"].GetAllMembers).Methods("GET")
	router.HandleFunc("/members/{id}", handlers["member"].GetMemberByID).Methods("GET")
	router.HandleFunc("/members/search", handlers["member"].SearchMembers).Methods("GET")
	router.HandleFunc("/members/{id}/loans", handlers["loan"].GetLoansByMemberID).Methods("GET")

	// Loan routes
	router.HandleFunc("/loans", handlers["loan"].GetAllLoans).Methods("GET")
	router.HandleFunc("/loans/{id}", handlers["loan"].GetLoanByID).Methods("GET")
	router.HandleFunc("/loans/overdue", handlers["loan"].GetOverdueLoans).Methods("GET")
	router.HandleFunc("/loans/member/{memberId}", handlers["loan"].GetLoansByMemberID).Methods("GET")

	// Notification routes
	router.HandleFunc("/notifications", handlers["notification"].GetAllNotifications).Methods("GET")
	router.HandleFunc("/notifications/member/{memberId}", handlers["notification"].GetNotificationsByMemberID).Methods("GET")

	// Review routes
	router.HandleFunc("/reviews", handlers["review"].GetAllReviews).Methods("GET")
	router.HandleFunc("/reviews/book/{bookId}/average-rating", handlers["review"].GetAverageRatingForBook).Methods("GET")

	// Auth routes
	router.HandleFunc("/login", handlers["auth"].Login).Methods("POST")
	router.HandleFunc("/register", handlers["auth"].Register).Methods("POST")

	// Admin routes (publicly accessible)
	router.HandleFunc("/admin/dashboard", handlers["admin"].GetDashboardData).Methods("GET")
	router.HandleFunc("/admin/books", handlers["admin"].ManageBooks).Methods("GET", "POST", "PUT", "DELETE")
	router.HandleFunc("/admin/members", handlers["admin"].ManageMembers).Methods("GET", "POST", "PUT", "DELETE")

	// Create a new router for authenticated routes
	authenticatedRouter := router.PathPrefix("/authenticated").Subrouter()
	authenticatedRouter.Use(middleware.AuthMiddleware)

	// Register routes that require authentication
	authenticatedRouter.HandleFunc("/admin/dashboard", handlers["admin"].GetDashboardData).Methods("GET")
	authenticatedRouter.HandleFunc("/admin/books", handlers["admin"].ManageBooks).Methods("GET", "POST", "PUT", "DELETE")
}

// startServer starts the server with the given router and configuration.
func startServer(router *mux.Router, cfg config.Config) error {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.ServerPort),
		Handler: router,
	}

	go func() {
		log.Printf("Server is running on port %d", cfg.ServerPort)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}

	log.Println("Server exited")
	return nil
}
