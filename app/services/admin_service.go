package services

import (
	"Restful-Perpustakaan-API/app/models"
	"Restful-Perpustakaan-API/app/repositories"
	"errors"
	"time"
)

// AdminService provides methods for admin operations
type AdminService struct {
	userRepository   repositories.UserRepository
	bookRepository   repositories.BookRepository
	memberRepository repositories.MemberRepository
	loanRepository   repositories.LoanRepository
}

// NewAdminService creates a new AdminService instance
func NewAdminService(userRepository repositories.UserRepository, bookRepository repositories.BookRepository, memberRepository repositories.MemberRepository, loanRepository database.LoanRepository) *AdminService {
	return &AdminService{
		userRepository:   userRepository,
		bookRepository:   bookRepository,
		memberRepository: memberRepository,
		loanRepository:   loanRepository,
	}
}

// GetDashboardData mengambil data untuk dashboard admin
func (as *AdminService) GetDashboardData() (map[string]interface{}, error) {
	totalMembers, err := as.memberRepository.GetTotalMembers()
	if err != nil {
		return nil, err
	}

	totalBooks, err := as.bookRepository.GetTotalBooks()
	if err != nil {
		return nil, err
	}

	totalLoans, err := as.loanRepository.GetTotalLoans()
	if err != nil {
		return nil, err
	}

	overdueLoans, err := as.loanRepository.GetOverdueLoans()
	if err != nil {
		return nil, err
	}

	// Anda bisa menambahkan metrik lain sesuai kebutuhan

	dashboardData := map[string]interface{}{
		"total_members": totalMembers,
		"total_books":   totalBooks,
		"total_loans":   totalLoans,
		"overdue_loans": overdueLoans,
		// ...
	}

	return dashboardData, nil
}

// GenerateReports menghasilkan laporan-laporan yang dibutuhkan admin
func (as *AdminService) GenerateReports(reportType string, startDate, endDate time.Time) (interface{}, error) {
	switch reportType {
	case "loan_report":
		return as.loanRepository.GetLoanReport(startDate, endDate)
	case "member_report":
		// ... logika untuk menghasilkan laporan anggota
		// Anda bisa menggunakan as.memberRepository untuk mengambil data anggota yang diperlukan
	case "book_report":
		// ... logika untuk menghasilkan laporan buku
		// Anda bisa menggunakan as.bookRepository untuk mengambil data buku yang diperlukan
	// ... tambahkan jenis laporan lain sesuai kebutuhan
	default:
		return nil, errors.New("invalid report type")
	}
}

// AddBook menambahkan buku baru ke perpustakaan
func (as *AdminService) AddBook(newBook *models.Book) error {
	// Lakukan validasi data buku baru (misalnya, cek apakah ISBN sudah ada)
	// ...

	return as.bookRepository.CreateBook(newBook)
}

// UpdateBook memperbarui informasi buku yang ada
func (as *AdminService) UpdateBook(updatedBook *models.Book) error {
	// Lakukan validasi data buku yang diperbarui
	// ...

	return as.bookRepository.UpdateBook(updatedBook)
}

// DeleteBook menghapus buku dari perpustakaan
func (as *AdminService) DeleteBook(bookID int) error {
	// Anda mungkin ingin menambahkan logika untuk menangani peminjaman yang terkait dengan buku ini sebelum menghapusnya
	// ...

	return as.bookRepository.DeleteBook(bookID)
}

// GetMemberLoans mengambil riwayat peminjaman seorang anggota
func (as *AdminService) GetMemberLoans(memberID int) ([]models.Loan, error) {
	return as.loanRepository.GetLoansByMemberID(memberID)
}

// GetBookLoans mengambil daftar peminjaman untuk sebuah buku
func (as *AdminService) GetBookLoans(bookID int) ([]models.Loan, error) {
	return as.loanRepository.GetLoansByBookID(bookID)
}

// IssueFine mengenakan denda kepada anggota
func (as *AdminService) IssueFine(memberID int, amount float64, reason string) error {
	// 1. Ambil data anggota dari database
	member, err := as.memberRepository.GetMemberByID(memberID)
	if err != nil {
		return err
	}

	// 2. Update field `FineAmount` pada struct `Member`
	member.FineAmount += amount

	// 3. Simpan perubahan ke database
	err = as.memberRepository.UpdateMember(member)
	if err != nil {
		return err
	}

	// 4. Buat catatan transaksi denda di database (opsional)
	// ...

	return nil
}

// ManageUsers menangani operasi CRUD untuk pengguna (admin dan pustakawan)
func (as *AdminService) ManageUsers(method string, user *models.User) error {
	switch method {
	case "GET":
		return as.userRepository.GetAllUsers()
	case "POST":
		// Lakukan validasi data pengguna baru (misalnya, cek apakah username sudah ada)
		// ...

		// Hash password sebelum disimpan (jika diperlukan)
		// ...

		return as.userRepository.CreateUser(user)
	case "PUT":
		// Lakukan validasi data pengguna yang diperbarui
		// ...

		return as.userRepository.UpdateUser(user)
	case "DELETE":
		return as.userRepository.DeleteUser(user.ID)
	default:
		return errors.New("invalid method")
	}
}
