package services

import (
	"Restful-Perpustakaan-API/app/models"
	"Restful-Perpustakaan-API/app/repositories"
)

// RecommendationService provides methods for getting book recommendations
type RecommendationService struct {
	bookRepository repositories.BookRepository
}

// NewRecommendationService creates a new RecommendationService instance
func NewRecommendationService(bookRepository repositories.BookRepository) *RecommendationService {
	return &RecommendationService{bookRepository: bookRepository}
}

// GetRecommendations mengambil rekomendasi buku umum
func (rs *RecommendationService) GetRecommendations() ([]models.Book, error) {
	return rs.bookRepository.GetRecommendations()
}

// GetPersonalizedRecommendations mengambil rekomendasi buku yang dipersonalisasi berdasarkan ID anggota
func (rs *RecommendationService) GetPersonalizedRecommendations(memberID int) ([]models.Book, error) {
	return rs.bookRepository.GetPersonalizedRecommendations(memberID)
}
