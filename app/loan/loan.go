package loan

import (
	"Restful-Perpustakaan-API/database"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"time"
)

// LoanService provides operations related to loans.
type LoanService struct {
	DB *database.Database
}

// NewLoanService creates a new instance of LoanService.
func NewLoanService(db *database.Database) *LoanService {
	return &LoanService{DB: db}
}

// Loan represents a book loan.
type Loan struct {
	ID         int        `json:"id"`
	MemberID   int        `json:"member_id"`
	BookID     int        `json:"book_id"`
	BorrowDate time.Time  `json:"borrow_date"`
	DueDate    time.Time  `json:"due_date"`
	ReturnDate *time.Time `json:"return_date,omitempty"` // Can be null if not returned
	Returned   bool       `json:"returned"`
}

// LoanHistory represents a loan history.
type LoanHistory struct {
	ID         int        `json:"id"`
	MemberID   int        `json:"member_id"`
	BookID     int        `json:"book_id"`
	BorrowDate time.Time  `json:"borrow_date"`
	DueDate    time.Time  `json:"due_date"`
	ReturnDate *time.Time `json:"return_date,omitempty"` // Can be null if not returned
}

// Member represents a library member.
type Member struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Book represents a book in the library.
type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

// Simulate storage with maps for quick lookups.
var (
	loans       = make(map[int]Loan)
	loanHistory = make(map[int]LoanHistory)
	members     = make(map[int]Member)
	books       = make(map[int]Book)
)

// Helper function to generate new IDs.
func generateID(m interface{}) int {
	switch m := m.(type) {
	case map[int]Member:
		if len(m) == 0 {
			return 1
		}
		var maxID int
		for id := range m {
			if id > maxID {
				maxID = id
			}
		}
		return maxID + 1
	case map[int]Book:
		if len(m) == 0 {
			return 1
		}
		var maxID int
		for id := range m {
			if id > maxID {
				maxID = id
			}
		}
		return maxID + 1
	default:
		return 1
	}
}

// GetAllMembers handles GET requests for all members.
func GetAllMembers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	membersList := make([]Member, 0, len(members))
	for _, member := range members {
		membersList = append(membersList, member)
	}
	json.NewEncoder(w).Encode(membersList)
}

// GetMemberByID handles GET requests for a member by ID.
func GetMemberByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	member, exists := members[id]
	if !exists {
		http.NotFound(w, r)
		return
	}
	json.NewEncoder(w).Encode(member)
}

// CreateMember handles POST requests to create a new member.
func CreateMember(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newMember Member
	if err := json.NewDecoder(r.Body).Decode(&newMember); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	newMember.ID = generateID(members)
	members[newMember.ID] = newMember
	json.NewEncoder(w).Encode(newMember)
}

// CreateBook handles POST requests to create a new book.
func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newBook Book
	if err := json.NewDecoder(r.Body).Decode(&newBook); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	newBook.ID = generateID(books)
	books[newBook.ID] = newBook
	json.NewEncoder(w).Encode(newBook)
}

// UpdateMember handles PUT requests to update an existing member.
func UpdateMember(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	var updatedMember Member
	if err := json.NewDecoder(r.Body).Decode(&updatedMember); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	updatedMember.ID = id
	if _, exists := members[id]; !exists {
		http.NotFound(w, r)
		return
	}
	members[id] = updatedMember
	json.NewEncoder(w).Encode(updatedMember)
}

// DeleteMember handles DELETE requests to remove a member.
func DeleteMember(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	if _, exists := members[id]; !exists {
		http.NotFound(w, r)
		return
	}
	delete(members, id)
	w.WriteHeader(http.StatusNoContent)
}

// GetAllBooks handles GET requests for all books.
func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	booksList := make([]Book, 0, len(books))
	for _, book := range books {
		booksList = append(booksList, book)
	}
	json.NewEncoder(w).Encode(booksList)
}

// GetBookByID handles GET requests for a book by ID.
func GetBookByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	book, exists := books[id]
	if !exists {
		http.NotFound(w, r)
		return
	}
	json.NewEncoder(w).Encode(book)
}

// UpdateBook handles PUT requests to update an existing book.
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	var updatedBook Book
	if err := json.NewDecoder(r.Body).Decode(&updatedBook); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	updatedBook.ID = id
	if _, exists := books[id]; !exists {
		http.NotFound(w, r)
		return
	}
	books[id] = updatedBook
	json.NewEncoder(w).Encode(updatedBook)
}

// DeleteBook handles DELETE requests to remove a book.
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	if _, exists := books[id]; !exists {
		http.NotFound(w, r)
		return
	}
	delete(books, id)
	w.WriteHeader(http.StatusNoContent)
}
