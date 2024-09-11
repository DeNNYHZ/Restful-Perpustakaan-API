package review

import "time"

// Review represents a review for a book
type Review struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	BookID    int       `json:"book_id"`
	Rating    int       `json:"rating"`
	Comment   string    `json:"comment"`
	Timestamp time.Time `json:"timestamp"`
}
