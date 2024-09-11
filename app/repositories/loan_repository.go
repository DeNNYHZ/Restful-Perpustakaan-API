package repositories

import (
	"Restful-Perpustakaan-API/app/models"
	"database/sql"
	"errors"
	"time"

	_ "github.com/lib/pq" // Import driver PostgreSQL
)

// LoanRepository provides methods for interacting with loan data in the database.
type LoanRepository struct {
	db *sql.DB
}

// NewLoanRepository creates a new LoanRepository instance
func NewLoanRepository(db *sql.DB) *LoanRepository {
	return &LoanRepository{db: db}
}

// GetAllLoans mengambil semua peminjaman dari database
func (lr *LoanRepository) GetAllLoans() ([]models.Loan, error) {
	query := `
		SELECT id, member_id, book_id, borrow_date, due_date, return_date
		FROM loans
	`

	rows, err := lr.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var loans []models.Loan
	for rows.Next() {
		var l models.Loan
		err := rows.Scan(&l.ID, &l.MemberID, &l.BookID, &l.BorrowDate, &l.DueDate, &l.ReturnDate)
		if err != nil {
			return nil, err
		}
		loans = append(loans, l)
	}

	return loans, nil
}

// GetLoanByID mengambil peminjaman berdasarkan ID dari database
func (lr *LoanRepository) GetLoanByID(id int) (*models.Loan, error) {
	query := `
		SELECT id, member_id, book_id, borrow_date, due_date, return_date
		FROM loans
		WHERE id = $1
	`

	var l models.Loan
	err := lr.db.QueryRow(query, id).Scan(&l.ID, &l.MemberID, &l.BookID, &l.BorrowDate, &l.DueDate, &l.ReturnDate)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("loan not found")
		}
		return nil, err
	}

	return &l, nil
}

// CreateLoan membuat peminjaman baru di database
func (lr *LoanRepository) CreateLoan(l *models.Loan) error {
	query := `
		INSERT INTO loans (member_id, book_id, borrow_date, due_date) 
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`

	err := lr.db.QueryRow(query, l.MemberID, l.BookID, time.Now(), l.DueDate).Scan(&l.ID)
	if err != nil {
		return err
	}

	return nil
}

// UpdateLoan memperbarui peminjaman di database
func (lr *LoanRepository) UpdateLoan(l *models.Loan) error {
	query := `
		UPDATE loans
		SET member_id = $1, book_id = $2, borrow_date = $3, due_date = $4, return_date = $5
		WHERE id = $6
	`

	_, err := lr.db.Exec(query, l.MemberID, l.BookID, l.BorrowDate, l.DueDate, l.ReturnDate, l.ID)
	return err
}

// DeleteLoan menghapus peminjaman dari database
func (lr *LoanRepository) DeleteLoan(id int) error {
	query := "DELETE FROM loans WHERE id = $1"
	_, err := lr.db.Exec(query, id)
	return err
}

func (lr *LoanRepository) GetLoansByMemberID(id int) ([]models.Loan, error) {
	// Implementasi metode GetLoansByMemberID

	return nil, nil

}

func (lr *LoanRepository) GetLoansByBookID(id int) ([]models.Loan, error) {
	// Implementasi metode GetLoansByBookID

	return nil, nil

}

func (lr *LoanRepository) GetTotalLoans() (interface{}, interface{}) {
	// Implementasi metode GetTotalLoans

	return nil, nil
}

func (lr *LoanRepository) GetOverdueLoans() (interface{}, interface{}) {
	// Implementasi metode GetOverdueLoans

	return nil, nil

}

// ... (fungsi lain yang mungkin Anda butuhkan, seperti GetLoansByMemberID, GetLoansByBookID, dll.)
