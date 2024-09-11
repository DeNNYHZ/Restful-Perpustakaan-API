package services

import (
	"Restful-Perpustakaan-API/app/models"
	"Restful-Perpustakaan-API/app/repositories"
)

// ReviewService provides methods for managing reviews
type ReviewService struct {
	reviewRepository repositories.ReviewRepository
}

// NewReviewService creates a new ReviewService instance
func NewReviewService(reviewRepository repositories.ReviewRepository) *ReviewService {
	return &ReviewService{reviewRepository: reviewRepository}
}

// GetAllReviews mengambil semua ulasan
func (rs *ReviewService) GetAllReviews() ([]models.Review, error) {
	return rs.reviewRepository.GetAllReviews()
}

// GetReviewByID mengambil ulasan berdasarkan ID
func (rs *ReviewService) GetReviewByID(id int) (*models.Review, error) {
	return rs.reviewRepository.GetReviewByID(id)
}

// CreateReview membuat ulasan baru
func (rs *ReviewService) CreateReview(r *models.Review) error {
	// Anda dapat menambahkan logika validasi atau bisnis lainnya di sini sebelum menyimpan ulasan ke database
	// Contoh: Memastikan pengguna sudah meminjam buku tersebut, dll.
	return rs.reviewRepository.CreateReview(r)
}

// UpdateReview memperbarui ulasan
func (rs *ReviewService) UpdateReview(r *models.Review) error {
	// Anda dapat menambahkan logika validasi atau bisnis lainnya di sini sebelum memperbarui ulasan di database
	// Contoh: Memastikan pengguna yang memperbarui adalah pemilik ulasan, dll
	return rs.reviewRepository.UpdateReview(r)
}

// DeleteReview menghapus ulasan
func (rs *ReviewService) DeleteReview(id int) error {
	return rs.reviewRepository.DeleteReview(id)
}

// GetReviewsForBook mengambil semua ulasan untuk buku tertentu berdasarkan ID buku
func (rs *ReviewService) GetReviewsForBook(bookID int) ([]models.Review, error) {
	return rs.reviewRepository.GetReviewsForBook(bookID)
}
