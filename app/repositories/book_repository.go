package repositories

import (
	"Restful-Perpustakaan-API/app/models"
	"database/sql"
	"errors"
	_ "github.com/lib/pq" // Import driver PostgreSQL
)

// BookRepository provides methods for interacting with the book data in the database
type BookRepository struct {
	db *sql.DB
}

// NewBookRepository creates a new BookRepository instance
func NewBookRepository(db *sql.DB) *BookRepository {
	return &BookRepository{db: db}
}

// GetAllBooks mengambil semua buku dari database
func (br *BookRepository) GetAllBooks() ([]models.Book, error) {
	query := `
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
		err := rows.Scan(&b.ID, &b.Title, &b.Author, &b.Publisher, &b.PublishedYear, &b.ISBN, &b.Genre, &b.Description, &b.CoverImage)
		if err != nil {
			return nil, err
		}
		books = append(books, b)
	}

	return books, nil
}

// GetBookByID mengambil buku berdasarkan ID dari database
func (br *BookRepository) GetBookByID(id int) (*models.Book, error) {
	query := `
        SELECT id, title, author, publisher, published_year, isbn, genre, description, cover_image 
        FROM books 
        WHERE id = $1
    ` // Menggunakan placeholder $1 untuk PostgreSQL

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

// CreateBook membuat buku baru di database
func (br *BookRepository) CreateBook(b *models.Book) error {
	query := `
        INSERT INTO books (title, author, publisher, published_year, isbn, genre, description, cover_image) 
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
        RETURNING id 
    ` // Menggunakan placeholder $1, $2, dll. dan RETURNING id untuk PostgreSQL

	err := br.db.QueryRow(query, b.Title, b.Author, b.Publisher, b.PublishedYear, b.ISBN, b.Genre, b.Description, b.CoverImage).Scan(&b.ID)
	if err != nil {
		return err
	}

	return nil
}

// UpdateBook memperbarui buku di database
func (br *BookRepository) UpdateBook(b *models.Book) error {
	query := `
        UPDATE books 
        SET title = $1, author = $2, publisher = $3, published_year = $4, isbn = $5, genre = $6, description = $7, cover_image = $8
        WHERE id = $9
    ` // Menggunakan placeholder

	_, err := br.db.Exec(query, b.Title, b.Author, b.Publisher, b.PublishedYear, b.ISBN, b.Genre, b.Description, b.CoverImage, b.ID)
	return err
}

// DeleteBook menghapus buku dari database
func (br *BookRepository) DeleteBook(id int) error {
	query := "DELETE FROM books WHERE id = $1" // Menggunakan placeholder
	_, err := br.db.Exec(query, id)
	return err
}

// GetRecommendations mengambil rekomendasi buku
func (br *BookRepository) GetRecommendations() ([]models.Book, error) {
	return nil, nil
}

// PersonalizedRecommendations
func (br *BookRepository) GetPersonalizedRecommendations(id int) ([]models.Book, error) {
	return nil, nil
}

// ... (fungsi lain yang mungkin Anda butuhkan, seperti GetBooksByTitle, GetBooksByAuthor, dll.)
