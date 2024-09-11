package database

import (
	"Restful-Perpustakaan-API/app/book"
	"Restful-Perpustakaan-API/app/loan"
	"Restful-Perpustakaan-API/app/user"
	"Restful-Perpustakaan-API/member"
	"Restful-Perpustakaan-API/notification"
	"Restful-Perpustakaan-API/review"
	"database/sql"
	"errors"
	"sync"
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

type GetAllMembers func() ([]Member, error)

// Database provides methods for interacting with the library database
type Database struct {
	members       []Member
	books         []Book
	loans         []Loan
	loanHistory   []LoanHistory
	notifications []Notification
	reviews       []review.Review
}

// Notification represents a notification to a user
type Notification struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
	Read      bool      `json:"read"`
	IsRead    bool      `json:"is_read"`
	MemberID  int       `json:"member_id"`
}

// NewDatabase returns a new instance of the database
func NewDatabase() *Database {
	return &Database{
		members:       make([]Member, 0),
		books:         make([]Book, 0),
		loans:         make([]Loan, 0),
		loanHistory:   make([]LoanHistory, 0),
		notifications: make([]Notification, 0),
	}
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

// GetAllNotifications returns all notifications
func (db *Database) GetAllNotifications() ([]Notification, error) {
	return db.notifications, nil
}

// CreateNotification creates a new notification
func (db *Database) CreateNotification(n *Notification) error {
	n.ID = len(db.notifications) + 1
	db.notifications = append(db.notifications, *n)
	return nil
}

// GetAllNotifications returns all notifications
func (db *Database) GetNotificationsByMemberID(memberID int) ([]Notification, error) {
	var result []Notification
	for _, n := range db.notifications {
		if n.MemberID == memberID {
			result = append(result, n)
		}
	}
	if len(result) == 0 {
		return nil, errors.New("notification not found")
	}
	return result, nil
}

// GetNotificationByID returns a notification by ID
func (db *Database) GetNotificationByID(id int) (*Notification, error) {
	for _, n := range db.notifications {
		if n.ID == id {
			return &n, nil
		}
	}
	return nil, errors.New("notification not found")
}

// UpdateNotification updates a notification
func (db *Database) UpdateNotification(updatedNotification *Notification) error {
	for i, n := range db.notifications {
		if n.ID == updatedNotification.ID {
			db.notifications[i] = *updatedNotification
			return nil
		}
	}
	return errors.New("notification not found")
}

// DeleteNotification deletes a notification
func (db *Database) DeleteNotification(id int) error {
	for i, n := range db.notifications {
		if n.ID == id {
			db.notifications = append(db.notifications[:i], db.notifications[i+1:]...)
			return nil
		}
	}
	return errors.New("notification not found")
}

func (db *Database) GetMembersByLastLoginDate(date time.Time) (interface{}, interface{}) {
	return db.members, nil

}

var (
	notifications = map[int]notification.Notification{}
	nextID        = 1
	mu            sync.Mutex
	reviews       = map[int]review.Review{}
	nextReviewID  = 1
	nextUserID    = 1
)

func GetAllNotifications() ([]notification.Notification, error) {
	mu.Lock()
	defer mu.Unlock()
	var result []notification.Notification
	for _, notif := range notifications {
		result = append(result, notif)
	}
	return result, nil
}

func GetNotificationByID(id int) (*notification.Notification, error) {
	mu.Lock()
	defer mu.Unlock()
	notif, exists := notifications[id]
	if !exists {
		return nil, errors.New("notification not found")
	}
	return &notif, nil
}

func CreateNotification(notif *notification.Notification) error {
	mu.Lock()
	defer mu.Unlock()
	notif.ID = nextID
	nextID++
	notifications[notif.ID] = *notif
	return nil
}

func UpdateNotification(notif *notification.Notification) error {
	mu.Lock()
	defer mu.Unlock()
	_, exists := notifications[notif.ID]
	if !exists {
		return errors.New("notification not found")
	}
	notifications[notif.ID] = *notif
	return nil
}

func DeleteNotification(id int) error {
	mu.Lock()
	defer mu.Unlock()
	_, exists := notifications[id]
	if !exists {
		return errors.New("notification not found")
	}
	delete(notifications, id)
	return nil
}

func MarkNotificationAsRead(id int) error {
	mu.Lock()
	defer mu.Unlock()
	notif, exists := notifications[id]
	if !exists {
		return errors.New("notification not found")
	}
	notif.Read = true
	notifications[id] = notif
	return nil
}

func MarkAllNotificationsAsRead() error {
	mu.Lock()
	defer mu.Unlock()
	for id, notif := range notifications {
		notif.Read = true
		notifications[id] = notif
	}
	return nil
}

func GetUnreadNotificationsCount() (int, error) {
	mu.Lock()
	defer mu.Unlock()
	count := 0
	for _, notif := range notifications {
		if !notif.Read {
			count++
		}
	}
	return count, nil
}

func GetAllUnreadNotificationsCount() (int, error) {
	return GetUnreadNotificationsCount() // For simplicity, assume all are unread
}

func GetAllNotificationsCount() (int, error) {
	mu.Lock()
	defer mu.Unlock()
	return len(notifications), nil
}

// GetNotificationByName returns a notification by name
func GetNotificationByName(name string) (*notification.Notification, error) {
	mu.Lock()
	defer mu.Unlock()
	for _, notif := range notifications {
		if notif.Name == name {
			return &notif, nil
		}
	}
	return nil, errors.New("notification not found")
}

// GetNotificationByCategory returns a notification by category
func GetNotificationByCategory(category string) (*notification.Notification, error) {
	mu.Lock()
	defer mu.Unlock()
	for _, notif := range notifications {
		if notif.Category == category {
			return &notif, nil
		}
	}
	return nil, errors.New("notification not found")
}

// GetNotificationByReceiver returns a notification by receiver
func GetNotificationByReceiver(receiver string) (*notification.Notification, error) {
	mu.Lock()
	defer mu.Unlock()
	for _, notif := range notifications {
		if notif.Receiver == receiver {
			return &notif, nil
		}
	}
	return nil, errors.New("notification not found")
}

// GetNotificationBySender returns a notification by sender
func GetNotificationBySender(sender string) (*notification.Notification, error) {
	mu.Lock()
	defer mu.Unlock()
	for _, notif := range notifications {
		if notif.Sender == sender {
			return &notif, nil
		}
	}
	return nil, errors.New("notification not found")
}

// GetRecommendations mengambil rekomendasi buku umum dari database.
func GetRecommendations(db *sql.DB) ([]book.Book, error) {
	// Query untuk mengambil buku-buku dengan rating tertinggi, misalnya 10 buku teratas
	query := "SELECT * FROM books ORDER BY rating DESC LIMIT 10"

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var recommendations []book.Book
	for rows.Next() {
		var b book.Book
		// Pastikan kolom-kolom dalam query sesuai dengan field-field dalam struct book.Book
		err := rows.Scan(&b.ID, &b.Title, &b.Author, &b.Genre, &b.Year, &b.Rating)
		if err != nil {
			return nil, err
		}
		recommendations = append(recommendations, b)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return recommendations, nil
}

// GetPersonalizedRecommendations mengambil rekomendasi buku yang dipersonalisasi berdasarkan ID anggota
func GetPersonalizedRecommendations(db *sql.DB, memberID int) ([]book.Book, error) {
	// Query untuk mengambil buku-buku yang pernah dipinjam oleh anggota
	queryBorrowedBooks := "SELECT book_id FROM borrowed_books WHERE member_id = ?"

	rows, err := db.Query(queryBorrowedBooks, memberID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	genreMap := make(map[string]int)
	for rows.Next() {
		var bookID int
		err := rows.Scan(&bookID)
		if err != nil {
			return nil, err
		}

		// Query untuk mengambil genre buku berdasarkan bookID
		queryBookGenre := "SELECT genre FROM books WHERE id = ?"
		genreRow := db.QueryRow(queryBookGenre, bookID)
		var genre string
		err = genreRow.Scan(&genre)
		if err != nil {
			return nil, err
		}
		genreMap[genre]++
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	var recommendations []book.Book
	// Loop melalui genre-genre yang pernah dipinjam oleh anggota
	for genre, count := range genreMap {
		if count > 1 { // Hanya pertimbangkan genre yang dipinjam lebih dari sekali
			// Query untuk mengambil buku-buku dengan genre yang sama dan rating tertinggi
			queryRecommendedBooks := "SELECT * FROM books WHERE genre = ? ORDER BY rating DESC LIMIT 3"

			genreRows, err := db.Query(queryRecommendedBooks, genre)
			if err != nil {
				return nil, err
			}
			defer genreRows.Close()

			for genreRows.Next() {
				var b book.Book
				err := genreRows.Scan(&b.ID, &b.Title, &b.Author, &b.Genre, &b.Year, &b.Rating)
				if err != nil {
					return nil, err
				}
				recommendations = append(recommendations, b)
			}

			if err := genreRows.Err(); err != nil {
				return nil, err
			}
		}
	}

	// Jika rekomendasi kurang dari 10, tambahkan rekomendasi umum
	if len(recommendations) < 10 {
		additionalRecs, err := GetRecommendations(db)
		if err != nil {
			return nil, err
		}
		recommendations = append(recommendations, additionalRecs[:10-len(recommendations)]...)
	}

	return recommendations, nil
}

var (
	books      = map[int]book.Book{}
	nextBookID = 1
	members    = map[int]member.Member{}
	loans      = map[int]loan.Loan{}
)

// GetAllBooks retrieves all books from the database.
func GetAllBooks() ([]book.Book, error) {
	mu.Lock()
	defer mu.Unlock()
	var result []book.Book
	for _, b := range books {
		result = append(result, b)
	}
	return result, nil
}

// GetBookByID retrieves a book by its ID.
func GetBookByID(id int) (*book.Book, error) {
	mu.Lock()
	defer mu.Unlock()
	b, exists := books[id]
	if !exists {
		return nil, errors.New("book not found")
	}
	return &b, nil
}

// CreateBook adds a new book to the database.
func CreateBook(b *book.Book) error {
	mu.Lock()
	defer mu.Unlock()
	b.ID = nextBookID
	nextBookID++
	books[b.ID] = *b
	return nil
}

// UpdateBook updates an existing book in the database.
func UpdateBook(b *book.Book) error {
	mu.Lock()
	defer mu.Unlock()
	_, exists := books[b.ID]
	if !exists {
		return errors.New("book not found")
	}
	books[b.ID] = *b
	return nil
}

// DeleteBook removes a book from the database by its ID.
func DeleteBook(id int) error {
	mu.Lock()
	defer mu.Unlock()
	_, exists := books[id]
	if !exists {
		return errors.New("book not found")
	}
	delete(books, id)
	return nil
}

// GetAllReviews retrieves all reviews from the database.
func GetAllReviews() ([]review.Review, error) {
	mu.Lock()
	defer mu.Unlock()
	var result []review.Review
	for _, r := range reviews {
		result = append(result, r)
	}
	return result, nil
}

// GetReviewsByID retrieves a review by its ID.
func GetReviewByID(id int) (*review.Review, error) {
	mu.Lock()
	defer mu.Unlock()
	r, exists := reviews[id]
	if !exists {
		return nil, errors.New("review not found")
	}
	return &r, nil
}

// CreateReview adds a new review to the database.
func CreateReview(r *review.Review) error {
	mu.Lock()
	defer mu.Unlock()
	r.ID = nextReviewID
	nextReviewID++
	reviews[r.ID] = *r
	return nil
}

// UpdateReview updates an existing review in the database.
func UpdateReview(r *review.Review) error {
	mu.Lock()
	defer mu.Unlock()
	_, exists := reviews[r.ID]
	if !exists {
		return errors.New("review not found")
	}
	reviews[r.ID] = *r
	return nil
}

// DeleteReview removes a review from the database by its ID.
func DeleteReview(id int) error {
	mu.Lock()
	defer mu.Unlock()
	_, exists := reviews[id]
	if !exists {
		return errors.New("review not found")
	}
	delete(reviews, id)
	return nil
}

// GetReviewsByBookID retrieves all reviews for a given book ID.
func GetReviewsByBookID(bookID int) ([]review.Review, error) {
	mu.Lock()
	defer mu.Unlock()
	var result []review.Review
	for _, r := range reviews {
		if r.BookID == bookID {
			result = append(result, r)
		}
	}
	return result, nil
}

// GetTotalMembers returns the total number of members.
func GetTotalMembers() (int, error) {
	mu.Lock()
	defer mu.Unlock()
	return len(members), nil
}

// GetTotalBooks returns the total number of books.
func GetTotalBooks() (int, error) {
	mu.Lock()
	defer mu.Unlock()
	return len(books), nil
}

// GetTotalLoans returns the total number of loans.
func GetTotalLoans() (int, error) {
	mu.Lock()
	defer mu.Unlock()
	return len(loans), nil
}

// GetOverdueLoans returns the number of overdue loans.
func GetOverdueLoans() (int, error) {
	mu.Lock()
	defer mu.Unlock()
	count := 0
	for _, loan := range loans {
		if isOverdue(loan) {
			count++
		}
	}
	return count, nil
}

// Helper function to check if a loan is overdue
func isOverdue(l loan.Loan) bool {
	return time.Now().After(l.DueDate) && !l.Returned
}

// DeleteUser removes a user from the database by its ID.
func DeleteUser(id int) error {
	mu.Lock()
	defer mu.Unlock()
	_, exists := members[id]
	if !exists {
		return errors.New("user not found")
	}
	delete(members, id)
	return nil
}

// UpdateUser updates an existing user in the database.
func UpdateUser(u *user.User) error {
	mu.Lock()
	defer mu.Unlock()
	_, exists := members[u.ID]
	if !exists {
		return errors.New("user not found")
	}

	// Convert user.User to member.Member
	var m member.Member
	m.ID = u.ID
	m.Name = u.Name
	m.Email = u.Email
	// ... assign other fields as needed

	members[u.ID] = m // Now assign the converted member
	return nil
}

// GetReviewsForBook retrieves all reviews for a given book ID.
func GetReviewsForBook(bookID int) ([]review.Review, error) {
	mu.Lock()
	defer mu.Unlock()
	var result []review.Review
	for _, r := range reviews {
		if r.BookID == bookID {
			result = append(result, r)
		}
	}
	return result, nil
}

// CreateUser adds a new user to the database.
func CreateUser(m *member.Member) error {
	mu.Lock()
	defer mu.Unlock()

	// Check if the user already exists based on some unique field (e.g., Username)
	for _, existingUser := range members {
		if existingUser.Username == m.Username {
			return errors.New("user already exists")
		}
	}
	m.ID = nextUserID
	nextUserID++
	members[m.ID] = *m
	return nil
}

// CreateUser adds a new user to the database.
func CreateMember(m *member.Member) error {
	mu.Lock()
	defer mu.Unlock()
	nextMemberID := len(members) + 1
	m.ID = nextMemberID
	nextMemberID++
	members[m.ID] = *m
	return nil
}

// GetLoanReport
func GetLoanReport(startDate, endDate time.Time) ([]loan.Loan, error) {
	mu.Lock()
	defer mu.Unlock()
	var result []loan.Loan
	for _, l := range loans {
		if l.DueDate.After(startDate) && l.DueDate.Before(endDate) {
			result = append(result, l)
		}
	}
	return result, nil
}

// GetAllUsers retrieves all users from the database.
func GetAllUsers() ([]user.User, error) {
	mu.Lock()
	defer mu.Unlock()
	var users []user.User
	for _, m := range members {
		users = append(users, user.User{
			ID:       m.ID,
			Name:     m.Name,
			Email:    m.Email,
			Password: m.Password,
		})
	}
	return users, nil
}

// UpdateMember updates an existing member in the database.
func UpdateMember(m *member.Member) error {
	mu.Lock()
	defer mu.Unlock()
	_, exists := members[m.ID]
	if !exists {
		return errors.New("member not found")
	}
	members[m.ID] = *m
	return nil
}

// GetMemberByEmail retrieves a member by email.
func GetMemberByEmail(email string) (*member.Member, error) {
	mu.Lock() // Lock the mutex to prevent concurrent access
	defer mu.Unlock()
	for _, m := range members {
		if m.Email == email {
			return &m, nil
		}
	}
	return nil, errors.New("member not found")
}
