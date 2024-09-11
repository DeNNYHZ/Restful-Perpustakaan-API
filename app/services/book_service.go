package services

import (
	"Restful-Perpustakaan-API/app/models"
	"Restful-Perpustakaan-API/app/repositories"
)

// BookService provides methods for managing books
type BookService struct {
	bookRepository repositories.BookRepository
}

// NewBookService creates a new BookService instance
func NewBookService(bookRepository repositories.BookRepository) *BookService {
	return &BookService{bookRepository: bookRepository}
}

// GetAllBooks mengambil semua buku
func (bs *BookService) GetAllBooks() ([]models.Book, error) {
	return bs.bookRepository.GetAllBooks()
}

// GetBookByID mengambil buku berdasarkan ID
func (bs *BookService) GetBookByID(id int) (*models.Book, error) {
	return bs.bookRepository.GetBookByID(id)
}

// CreateBook membuat buku baru
func (bs *BookService) CreateBook(b *models.Book) error {
	// Anda dapat menambahkan logika validasi atau bisnis lainnya di sini sebelum menyimpan buku ke database
	return bs.bookRepository.CreateBook(b)
}

// UpdateBook memperbarui buku
func (bs *BookService) UpdateBook(b *models.Book) error {
	// Anda dapat menambahkan logika validasi atau bisnis lainnya di sini sebelum memperbarui buku di database
	return bs.bookRepository.UpdateBook(b)
}

// DeleteBook menghapus buku
func (bs *BookService) DeleteBook(id int) error {
	return bs.bookRepository.DeleteBook(id)
}

// ... (fungsi lain yang mungkin Anda butuhkan, seperti SearchBooks, dll.)
