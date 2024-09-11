package services

import (
	"Restful-Perpustakaan-API/app/models"
	"Restful-Perpustakaan-API/app/repositories"
	"time"
)

// LoanService provides methods for managing loans
type LoanService struct {
	loanRepository repositories.LoanRepository
}

// NewLoanService creates a new LoanService instance
func NewLoanService(loanRepository repositories.LoanRepository) *LoanService {
	return &LoanService{loanRepository: loanRepository}
}

// GetAllLoans mengambil semua peminjaman
func (ls *LoanService) GetAllLoans() ([]models.Loan, error) {
	return ls.loanRepository.GetAllLoans()
}

// GetLoanByID mengambil peminjaman berdasarkan ID
func (ls *LoanService) GetLoanByID(id int) (*models.Loan, error) {
	return ls.loanRepository.GetLoanByID(id)
}

// CreateLoan membuat peminjaman baru
func (ls *LoanService) CreateLoan(l *models.Loan) error {
	// Anda dapat menambahkan logika validasi atau bisnis lainnya di sini sebelum menyimpan peminjaman ke database
	// Contoh: Memastikan buku tersedia, memeriksa batas peminjaman anggota, dll.

	// Set tanggal peminjaman saat ini
	l.BorrowDate = time.Now()

	return ls.loanRepository.CreateLoan(l)
}

// UpdateLoan memperbarui peminjaman
func (ls *LoanService) UpdateLoan(l *models.Loan) error {
	// Anda dapat menambahkan logika validasi atau bisnis lainnya di sini sebelum memperbarui peminjaman di database
	// Contoh: Memeriksa apakah peminjaman sudah dikembalikan, menghitung denda jika terlambat, dll.

	// Jika buku dikembalikan, set tanggal pengembalian
	if l.Returned {
		now := time.Now()
		l.ReturnDate = &now // Take the address of the time.Time value
	}

	return ls.loanRepository.UpdateLoan(l)
}

// DeleteLoan menghapus peminjaman
func (ls *LoanService) DeleteLoan(id int) error {
	return ls.loanRepository.DeleteLoan(id)
}

// ... (fungsi lain yang mungkin Anda butuhkan, seperti GetLoansByMemberID, GetOverdueLoans, dll.)
