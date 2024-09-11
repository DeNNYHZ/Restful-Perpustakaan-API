package main

import (
	"github.com/gorilla/mux"

	book_handler "Restful-Perpustakaan-API/app/handlers/book_handler"
)

func SetupRoutes(r *mux.Router) {
	// Rute untuk Buku (book_handler.go)
	r.HandleFunc("/books", book_handler.GetAllBooks).Methods("GET")
	r.HandleFunc("/books/{id}", book_handler.GetBookByID).Methods("GET")
	r.HandleFunc("/books", book_handler.CreateBook).Methods("POST")
	r.HandleFunc("/books/{id}", book_handler.UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", book_handler.DeleteBook).Methods("DELETE")

	// Rute untuk Anggota (member_handler.go)
	r.HandleFunc("/members", member_handler.GetAllMembers).Methods("GET")
	r.HandleFunc("/members/{id}", member_handler.GetMemberByID).Methods("GET")
	r.HandleFunc("/members", member_handler.CreateMember).Methods("POST")
	r.HandleFunc("/members/{id}", member_handler.UpdateMember).Methods("PUT")
	r.HandleFunc("/members/{id}", member_handler.DeleteMember).Methods("DELETE")

	// Rute untuk Peminjaman (loan_handler.go)
	r.HandleFunc("/loans", loan_handler.GetAllLoans).Methods("GET")
	r.HandleFunc("/loans/{id}", loan_handler.GetLoanByID).Methods("GET")
	r.HandleFunc("/loans", loan_handler.CreateLoan).Methods("POST")
	r.HandleFunc("/loans/{id}", loan_handler.UpdateLoan).Methods("PUT")
	r.HandleFunc("/loans/{id}", loan_handler.DeleteLoan).Methods("DELETE")

	// Rute untuk Otentikasi (auth_handler.go)
	r.HandleFunc("/login", auth_handler.Login).Methods("POST")
	r.HandleFunc("/register", auth_handler.Register).Methods("POST")

	// Rute untuk Notifikasi (notification_handler.go)
	r.HandleFunc("/notifications", notification_handler.GetAllNotifications).Methods("GET")
	r.HandleFunc("/notifications/{id}", notification_handler.GetNotificationByID).Methods("GET")
	r.HandleFunc("/notifications", notification_handler.CreateNotification).Methods("POST")
	r.HandleFunc("/notifications/{id}", notification_handler.UpdateNotification).Methods("PUT")
	r.HandleFunc("/notifications/{id}", notification_handler.DeleteNotification).Methods("DELETE")
	r.HandleFunc("/notifications/mark-as-read/{id}", notification_handler.MarkNotificationAsRead).Methods("PUT")

	// Rute untuk Rekomendasi (recommendation_handler.go)
	r.HandleFunc("/recommendations", recommendation_handler.GetRecommendations).Methods("GET")
	r.HandleFunc("/recommendations/personalized/{memberId}", recommendation_handler.GetPersonalizedRecommendations).Methods("GET")

	// Rute untuk Ulasan (review_handler.go)
	r.HandleFunc("/reviews", review_handler.GetAllReviews).Methods("GET")
	r.HandleFunc("/reviews/{id}", review_handler.GetReviewByID).Methods("GET")
	r.HandleFunc("/reviews", review_handler.CreateReview).Methods("POST")
	r.HandleFunc("/reviews/{id}", review_handler.UpdateReview).Methods("PUT")
	r.HandleFunc("/reviews/{id}", review_handler.DeleteReview).Methods("DELETE")
	r.HandleFunc("/reviews/book/{bookId}", review_handler.GetReviewsForBook).Methods("GET")

	// Rute untuk Admin (admin_handler.go)
	r.HandleFunc("/admin/dashboard", admin_handler.GetDashboardData).Methods("GET")
	r.HandleFunc("/admin/reports", admin_handler.GenerateReports).Methods("GET")
	r.HandleFunc("/admin/users", admin_handler.ManageUsers).Methods("GET", "POST", "PUT", "DELETE")
	// ... tambahkan rute lain sesuai kebutuhan admin

	// Rute Umum Lainnya
	r.HandleFunc("/search", search_handler.Search).Methods("GET")                // Pencarian global
	r.HandleFunc("/contact", contact_handler.SendContactMessage).Methods("POST") // Mengirim pesan kontak
	r.HandleFunc("/", home_handler.GetHomePage).Methods("GET")                   // Halaman utama

	// tambahkan rute lain sesuai kebutuhan API Anda
}
