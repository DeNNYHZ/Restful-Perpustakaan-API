package handlers

import (
	"Restful-Perpustakaan-API/app/converters"
	"Restful-Perpustakaan-API/app/models"
	"Restful-Perpustakaan-API/database" // Update with the actual path
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"time"
)

// DashboardData represents the data for the admin dashboard.
type DashboardData struct {
	TotalMembers int `json:"total_members"`
	TotalBooks   int `json:"total_books"`
	TotalLoans   int `json:"total_loans"`
	OverdueLoans int `json:"overdue_loans"`
}

// GetDashboardData handles the GET request for the admin dashboard data.
func GetDashboardData(w http.ResponseWriter, r *http.Request) {
	totalMembers, _ := database.GetTotalMembers()
	totalBooks, _ := database.GetTotalBooks()
	totalLoans, _ := database.GetTotalLoans()
	overdueLoans, _ := database.GetOverdueLoans()

	dashboardData := DashboardData{
		TotalMembers: totalMembers,
		TotalBooks:   totalBooks,
		TotalLoans:   totalLoans,
		OverdueLoans: overdueLoans,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dashboardData)
}

// LoanReport represents the loan report data.
type LoanReport struct {
	TotalLoans         int          `json:"total_loans"`
	OverdueLoans       int          `json:"overdue_loans"`
	TotalLoanAmount    float64      `json:"total_loan_amount"`
	TotalOverdueAmount float64      `json:"total_overdue_amount"`
	Loans              []LoanDetail `json:"loans"`
}

// LoanDetail provides details about individual loans within the report.
type LoanDetail struct {
	LoanID       int        `json:"loan_id"`
	BookTitle    string     `json:"book_title"`
	BorrowerName string     `json:"borrower_name"`
	LoanDate     time.Time  `json:"loan_date"`
	DueDate      time.Time  `json:"due_date"`
	ReturnDate   *time.Time `json:"return_date,omitempty"`
	Amount       float64    `json:"amount"`
	IsOverdue    bool       `json:"is_overdue"`
}

// GenerateReports handles the GET request for the loan reports.
func GenerateReports(w http.ResponseWriter, r *http.Request) {
	startDateStr := r.URL.Query().Get("start_date")
	endDateStr := r.URL.Query().Get("end_date")

	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		http.Error(w, "Invalid start date format", http.StatusBadRequest)
		return
	}

	endDate, err := time.Parse("2006-01-02", endDateStr)
	if err != nil {
		http.Error(w, "Invalid end date format", http.StatusBadRequest)
		return
	}

	loanReport, err := database.GetLoanReport(startDate, endDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(loanReport)
}

// ManageUsers handles the CRUD operations for users.
func ManageUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		users, err := database.GetAllUsers()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)

	case http.MethodPost:
		var newUser models.User
		err := json.NewDecoder(r.Body).Decode(&newUser)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// POST request for user creation is handled by the Register function.
		http.Error(w, "Use /register endpoint for user registration", http.StatusMethodNotAllowed)

	case http.MethodPut:
		params := mux.Vars(r)
		memberID, err := strconv.Atoi(params["id"])
		if err != nil {
			http.Error(w, "Invalid member ID", http.StatusBadRequest)
			return
		}

		var updatedUser models.User
		err = json.NewDecoder(r.Body).Decode(&updatedUser)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Convert user.User to member.Member
		updatedMember := converters.ConvertUserToMember(&updatedUser)
		updatedMember.ID = memberID

		err = database.UpdateMember(updatedMember)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(updatedUser)

	case http.MethodDelete:
		params := mux.Vars(r)
		userID, err := strconv.Atoi(params["id"])
		if err != nil {
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
			return
		}

		err = database.DeleteUser(userID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
