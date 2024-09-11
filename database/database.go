package database

import (
	"errors"
	"time"
)

// Member represents a library member
type Member struct {
	ID               int       `json:"id"`
	Name             string    `json:"name"`
	Gender           string    `json:"gender"`
	Email            string    `json:"email"`
	PhoneNumber      string    `json:"phone_number"`
	Address          string    `json:"address"`
	RegistrationDate time.Time `json:"registration_date"` // e.g., "2022-01-01T00:00:00Z"
	MembershipType   string    `json:"membership_type"`   // e.g., "Regular", "Student", "Premium"
}

// Book represents a library book
type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

// Loan represents a library loan
type Loan struct {
	ID         int        `json:"id"`
	MemberID   int        `json:"member_id"`
	BookID     int        `json:"book_id"`
	BorrowDate time.Time  `json:"borrow_date"`
	DueDate    time.Time  `json:"due_date"`
	ReturnDate *time.Time `json:"return_date,omitempty"` // Can be null if not returned
}

// LoanHistory represents a library loan history
type LoanHistory struct {
	ID         int        `json:"id"`
	MemberID   int        `json:"member_id"`
	BookID     int        `json:"book_id"`
	BorrowDate time.Time  `json:"borrow_date"`
	DueDate    time.Time  `json:"due_date"`
	ReturnDate *time.Time `json:"return_date,omitempty"` // Can be null if not returned
}

// Database provides methods for interacting with the library database
type Database struct {
	members     []Member
	books       []Book
	loans       []Loan
	loanHistory []LoanHistory
}

// NewDatabase returns a new instance of the database
func NewDatabase() *Database {
	return &Database{
		members:     make([]Member, 0),
		books:       make([]Book, 0),
		loans:       make([]Loan, 0),
		loanHistory: make([]LoanHistory, 0),
	}
}

// GetAllMembers returns all members from the database
func (db *Database) GetAllMembers() ([]Member, error) {
	return db.members, nil
}

// GetMemberByID returns a member by ID from the database
func (db *Database) GetMemberByID(id int) (*Member, error) {
	for _, m := range db.members {
		if m.ID == id {
			return &m, nil
		}
	}
	return nil, errors.New("member not found")
}

// CreateMember creates a new member in the database
func (db *Database) CreateMember(m *Member) error {
	m.ID = len(db.members) + 1
	db.members = append(db.members, *m)
	return nil
}

// UpdateMember updates a member in the database
func (db *Database) UpdateMember(m *Member) error {
	for i, existingMember := range db.members {
		if existingMember.ID == m.ID {
			db.members[i] = *m
			return nil
		}
	}
	return errors.New("member not found")
}

// DeleteMember deletes a member from the database
func (db *Database) DeleteMember(id int) error {
	for i, m := range db.members {
		if m.ID == id {
			db.members = append(db.members[:i], db.members[i+1:]...)
			return nil
		}
	}
	return errors.New("member not found")
}

// GetAllBooks returns all books from the database
func (db *Database) GetAllBooks() ([]Book, error) {
	return db.books, nil
}

// GetBookByID returns a book by ID from the database
func (db *Database) GetBookByID(id int) (*Book, error) {
	for _, b := range db.books {
		if b.ID == id {
			return &b, nil
		}
	}
	return nil, errors.New("book not found")
}

// CreateBook creates a new book in the database
func (db *Database) CreateBook(b *Book) error {
	b.ID = len(db.books) + 1
	db.books = append(db.books, *b)
	return nil
}

// UpdateBook updates a book in the database
func (db *Database) UpdateBook(b *Book) error {
	for i, existingBook := range db.books {
		if existingBook.ID == b.ID {
			db.books[i] = *b
			return nil
		}
	}
	return errors.New("book not found")
}

// DeleteBook deletes a book from the database
func (db *Database) DeleteBook(id int) error {
	for i, b := range db.books {
		if b.ID == id {
			db.books = append(db.books[:i], db.books[i+1:]...)
			return nil
		}
	}
	return errors.New("book not found")
}

// GetAllLoans returns all loans from the database
func (db *Database) GetAllLoans() ([]Loan, error) {
	return db.loans, nil
}

// GetLoanByID returns a loan by ID from the database
func (db *Database) GetLoanByID(id int) (*Loan, error) {
	for _, l := range db.loans {
		if l.ID == id {
			return &l, nil
		}
	}
	return nil, errors.New("loan not found")
}

// CreateLoan creates a new loan in the database
func (db *Database) CreateLoan(l *Loan) error {
	l.ID = len(db.loans) + 1
	db.loans = append(db.loans, *l)
	return nil
}

// UpdateLoan updates a loan in the database
func (db *Database) UpdateLoan(l *Loan) error {
	for i, existingLoan := range db.loans {
		if existingLoan.ID == l.ID {
			db.loans[i] = *l
			return nil
		}
	}
	return errors.New("loan not found")
}

// DeleteLoan deletes a loan from the database
func (db *Database) DeleteLoan(id int) error {
	for i, l := range db.loans {
		if l.ID == id {
			db.loans = append(db.loans[:i], db.loans[i+1:]...)
			return nil
		}
	}
	return errors.New("loan not found")
}

// GetAllLoanHistories returns all loan histories from the database
func (db *Database) GetAllLoanHistories() ([]LoanHistory, error) {
	return db.loanHistory, nil
}

// GetLoanHistoryByID returns a loan history by ID from the database
func (db *Database) GetLoanHistoryByID(id int) (*LoanHistory, error) {
	for _, lh := range db.loanHistory {
		if lh.ID == id {
			return &lh, nil
		}
	}
	return nil, errors.New("loan history not found")
}

// CreateLoanHistory creates a new loan history in the database
func (db *Database) CreateLoanHistory(lh *LoanHistory) error {
	lh.ID = len(db.loanHistory) + 1
	db.loanHistory = append(db.loanHistory, *lh)
	return nil
}

// UpdateLoanHistory updates a loan history in the database
func (db *Database) UpdateLoanHistory(lh *LoanHistory) error {
	for i, existingLoanHistory := range db.loanHistory {
		if existingLoanHistory.ID == lh.ID {
			db.loanHistory[i] = *lh
			return nil
		}
	}
	return errors.New("loan history not found")
}

// DeleteLoanHistory deletes a loan history from the database
func (db *Database) DeleteLoanHistory(id int) error {
	for i, lh := range db.loanHistory {
		if lh.ID == id {
			db.loanHistory = append(db.loanHistory[:i], db.loanHistory[i+1:]...)
			return nil
		}
	}
	return errors.New("loan history not found")
}

// Example method to find members by name
func (db *Database) GetMembersByName(name string) ([]Member, error) {
	var result []Member
	for _, m := range db.members {
		if m.Name == name {
			result = append(result, m)
		}
	}
	if len(result) == 0 {
		return nil, errors.New("member not found")
	}
	return result, nil
}

// Example method to find books by title
func (db *Database) GetBooksByTitle(title string) ([]Book, error) {
	var result []Book
	for _, b := range db.books {
		if b.Title == title {
			result = append(result, b)
		}
	}
	if len(result) == 0 {
		return nil, errors.New("book not found")
	}
	return result, nil
}

// GetMemberByEmail returns a member by email
func (db *Database) GetMemberByEmail(email string) (*Member, error) {
	for _, m := range db.members {
		if m.Email == email {
			return &m, nil
		}
	}
	return nil, errors.New("member not found")
}

// GetMemberByAddress returns a member by address
func (db *Database) GetMemberByAddress(address string) (*Member, error) {
	for _, m := range db.members {
		if m.Address == address {
			return &m, nil
		}
	}
	return nil, errors.New("member not found")
}

// GetMemberByPhoneNumber returns a member by phone number
func (db *Database) GetMemberByPhoneNumber(phoneNumber string) (*Member, error) {
	for _, m := range db.members {
		if m.PhoneNumber == phoneNumber {
			return &m, nil
		}
	}
	return nil, errors.New("member not found")
}

// GetMemberByGender returns a member by gender
func (db *Database) GetMemberByGender(gender string) (*Member, error) {
	for _, m := range db.members {
		if m.Gender == gender {
			return &m, nil
		}
	}
	return nil, errors.New("member not found")
}

// GetMemberByMembershipType returns a member by membership type
func (db *Database) GetMembersByMembershipType(membershipType string) (*Member, error) {
	for _, m := range db.members {
		if m.MembershipType == membershipType {
			return &m, nil
		}
	}
	return nil, errors.New("member not found")
}

// GetMembersByRegistrationDate returns a member by registration date
func (db *Database) GetMembersByRegistrationDate(registrationDate time.Time) (*Member, error) {
	for _, m := range db.members {
		if m.RegistrationDate == registrationDate {
			return &m, nil
		}
	}
	return nil, errors.New("member not found")
}
