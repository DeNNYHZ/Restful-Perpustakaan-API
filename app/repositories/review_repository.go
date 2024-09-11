package repositories

import (
	"Restful-Perpustakaan-API/app/models"
	"database/sql"
	"errors"
	"time"

	_ "github.com/lib/pq" // Import driver PostgreSQL
)

// ReviewRepository provides methods for interacting with review data in the database
type ReviewRepository struct {
	db *sql.DB
}

// NewReviewRepository creates a new ReviewRepository instance
func NewReviewRepository(db *sql.DB) *ReviewRepository {
	return &ReviewRepository{db: db}
}

// GetAllReviews mengambil semua ulasan dari database
func (rr *ReviewRepository) GetAllReviews() ([]models.Review, error) {
	query := "SELECT id, user_id, book_id, rating, comment, timestamp FROM reviews"

	rows, err := rr.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reviews []models.Review
	for rows.Next() {
		var r models.Review
		err := rows.Scan(&r.ID, &r.UserID, &r.BookID, &r.Rating, &r.Comment, &r.Timestamp)
		if err != nil {
			return nil, err
		}
		reviews = append(reviews, r)
	}

	return reviews, nil
}

// GetReviewByID mengambil ulasan berdasarkan ID dari database
func (rr *ReviewRepository) GetReviewByID(id int) (*models.Review, error) {
	query := "SELECT id, user_id, book_id, rating, comment, timestamp FROM reviews WHERE id = $1"

	var r models.Review
	err := rr.db.QueryRow(query, id).Scan(&r.ID, &r.UserID, &r.BookID, &r.Rating, &r.Comment, &r.Timestamp)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("review not found")
		}
		return nil, err
	}

	return &r, nil
}

// CreateReview membuat ulasan baru di database
func (rr *ReviewRepository) CreateReview(r *models.Review) error {
	query := `
		INSERT INTO reviews (user_id, book_id, rating, comment, timestamp)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`

	err := rr.db.QueryRow(query, r.UserID, r.BookID, r.Rating, r.Comment, time.Now()).Scan(&r.ID)
	if err != nil {
		return err
	}

	return nil
}

// UpdateReview memperbarui ulasan di database
func (rr *ReviewRepository) UpdateReview(r *models.Review) error {
	query := `
		UPDATE reviews
		SET user_id = $1, book_id = $2, rating = $3, comment = $4, timestamp = $5
		WHERE id = $6
	`

	_, err := rr.db.Exec(query, r.UserID, r.BookID, r.Rating, r.Comment, r.Timestamp, r.ID)
	return err
}

// DeleteReview menghapus ulasan dari database
func (rr *ReviewRepository) DeleteReview(id int) error {
	query := "DELETE FROM reviews WHERE id = $1"
	_, err := rr.db.Exec(query, id)
	return err
}

// GetReviewsForBook mengambil semua ulasan untuk buku tertentu berdasarkan ID buku
func (rr *ReviewRepository) GetReviewsForBook(bookID int) ([]models.Review, error) {
	query := "SELECT id, user_id, book_id, rating, comment, timestamp FROM reviews WHERE book_id = $1"

	rows, err := rr.db.Query(query, bookID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reviews []models.Review
	for rows.Next() {
		var r models.Review
		err := rows.Scan(&r.ID, &r.UserID, &r.BookID, &r.Rating, &r.Comment, &r.Timestamp)
		if err != nil {
			return nil, err
		}
		reviews = append(reviews, r)
	}

	return reviews, nil
}
