package handlers

import (
	"Restful-Perpustakaan-API/app/models"
	_ "Restful-Perpustakaan-API/app/utils"
	"Restful-Perpustakaan-API/database"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Book represents a book entity
type Book struct {
	ID            int     `json:"id"`
	Title         string  `json:"title"`
	Year          int     `json:"year"`
	Rating        float64 `json:"rating"`
	Author        string  `json:"author"`
	Publisher     string  `json:"publisher"`
	PublishedYear int     `json:"published_year"`
	ISBN          string  `json:"isbn"`
	Genre         string  `json:"genre"`
	Description   string  `json:"description"`
	CoverImage    string  `json:"cover_image"` // URL or path to cover image
}

// books is a simulated in-memory book store
var books []Book

// AddBook adds a new book to the store
func AddBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	book.ID = len(books) + 1
	books = append(books, book)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

// GetAllBooks retrieves all books from the database
func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books, err := database.GetAllBooks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// GetBook retrieves a book by its ID
func GetBook(id int) models.BookModel {
	book, err := database.GetBookByID(id)
	if err != nil {
		return models.BookModel{}
	}
	return models.BookModel{
		ID:            book.ID,
		Title:         book.Title,
		Author:        book.Author,
		Year:          book.Year,
		Rating:        book.Rating,
		Publisher:     book.Publisher,
		PublishedYear: book.PublishedYear,
		ISBN:          book.ISBN,
		Genre:         book.Genre,
		Description:   book.Description,
		CoverImage:    book.CoverImage, // URL atau path ke gambar sampul
		// Copy other fields as needed
	}
}

// GetBookByID retrieves a book by its ID
func GetBookByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}
	book, err := database.GetBookByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

// CreateBook adds a new book to the database
func CreateBook(w http.ResponseWriter, r *http.Request) {
	var newBook models.Book
	err := json.NewDecoder(r.Body).Decode(&newBook)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = database.CreateBook(&newBook)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newBook)
}

// UpdateBook updates an existing book in the database
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}
	var updatedBook models.Book
	err = json.NewDecoder(r.Body).Decode(&updatedBook)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	updatedBook.ID = id
	err = database.UpdateBook(&updatedBook)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedBook)
}

// DeleteBook removes a book from the database by its ID
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}
	err = database.DeleteBook(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
