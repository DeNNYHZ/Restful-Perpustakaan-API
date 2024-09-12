package repositories

import (
	"Restful-Perpustakaan-API/app/models"
	"database/sql"
	"errors"
	_ "github.com/lib/pq" // Import driver PostgreSQL
)

// BookRepository provides methods for interacting with the book data in the database
type BookRepository interface {
	GetAllBooks() ([]models.Book, error)
	GetBookByID(id int) (*models.Book, error)
	CreateBook(b *models.Book) error
	UpdateBook(b *models.Book) error
	DeleteBook(id int) error
	GetRecommendations() ([]models.Book, error)
	GetPersonalizedRecommendations(id int) ([]models.Book, error)
	GetTotalBooks() (interface{}, interface{})
}

// NewBookRepository creates a new BookRepository instance
func NewBookRepository(db *sql.DB) *bookRepository {
	return &bookRepository{db: db}
}

type bookRepository struct {
	db *sql.DB
}

// GetAllBooks retrieves all books from the database
func (br *bookRepository) GetAllBooks() ([]models.Book, error) {
	const query = `
        SELECT id, title, author, publisher, published_year, isbn, genre, description, cover_image 
        FROM books
    `

	rows, err := br.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var b models.Book
		if err := rows.Scan(&b.ID, &b.Title, &b.Author, &b.Publisher, &b.PublishedYear, &b.ISBN, &b.Genre, &b.Description, &b.CoverImage); err != nil {
			return nil, err
		}
		books = append(books, b)
	}

	return books, nil
}

// GetBookByID retrieves a book by ID from the database
func (br *bookRepository) GetBookByID(id int) (*models.Book, error) {
	const query = `
        SELECT id, title, author, publisher, published_year, isbn, genre, description, cover_image 
        FROM books 
        WHERE id = $1
    `

	var b models.Book
	err := br.db.QueryRow(query, id).Scan(&b.ID, &b.Title, &b.Author, &b.Publisher, &b.PublishedYear, &b.ISBN, &b.Genre, &b.Description, &b.CoverImage)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("book not found")
		}
		return nil, err
	}

	return &b, nil
}

// CreateBook creates a new book in the database
func (br *bookRepository) CreateBook(b *models.Book) error {
	const query = `
        INSERT INTO books (title, author, publisher, published_year, isbn, genre, description, cover_image) 
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
        RETURNING id 
    `

	err := br.db.QueryRow(query, b.Title, b.Author, b.Publisher, b.PublishedYear, b.ISBN, b.Genre, b.Description, b.CoverImage).Scan(&b.ID)
	if err != nil {
		return err
	}

	return nil
}

// UpdateBook updates a book in the database
func (br *bookRepository) UpdateBook(b *models.Book) error {
	const query = `
        UPDATE books 
        SET title = $1, author = $2, publisher = $3, published_year = $4, isbn = $5, genre = $6, description = $7, cover_image = $8
        WHERE id = $9
    `

	_, err := br.db.Exec(query, b.Title, b.Author, b.Publisher, b.PublishedYear, b.ISBN, b.Genre, b.Description, b.CoverImage, b.ID)
	return err
}

// DeleteBook deletes a book from the database
func (br *bookRepository) DeleteBook(id int) error {
	const query = "DELETE FROM books WHERE id = $1"

	_, err := br.db.Exec(query, id)
	return err
}

// GetRecommendations retrieves book recommendations
func (br *bookRepository) GetRecommendations() ([]models.Book, error) {
	return nil, nil
}

// GetPersonalizedRecommendations retrieves personalized book recommendations
func (br *bookRepository) GetPersonalizedRecommendations(id int) ([]models.Book, error) {
	return nil, nil
}

func (br *bookRepository) GetTotalBooks() (interface{}, interface{}) {
	return nil, nil
}
