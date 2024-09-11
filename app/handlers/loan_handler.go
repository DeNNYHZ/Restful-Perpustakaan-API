package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"Restful-Perpustakaan-API/app/loan"
	"github.com/gorilla/mux"
)

// LoanHandlers holds the handlers for loan-related endpoints
type LoanHandlers struct {
	loanService loan.LoanService
}

// NewLoanHandlers returns a new instance of LoanHandlers
func NewLoanHandlers(loanService loan.LoanService) *LoanHandlers {
	return &LoanHandlers{loanService: loanService}
}

// GetAllLoans handles GET requests to retrieve all loans
func (lh *LoanHandlers) GetAllLoans(w http.ResponseWriter, r *http.Request) {
	loans, err := lh.loanService.GetAllLoans()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, loans)
}

// GetLoanByID handles GET requests to retrieve a loan by ID
func (lh *LoanHandlers) GetLoanByID(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromParams(r)
	if err != nil {
		http.Error(w, "Invalid loan ID", http.StatusBadRequest)
		return
	}
	loan, err := lh.loanService.GetLoanByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, loan)
}

// CreateLoan handles POST requests to create a new loan
func (lh *LoanHandlers) CreateLoan(w http.ResponseWriter, r *http.Request) {
	var newLoan loan.Loan
	err := json.NewDecoder(r.Body).Decode(&newLoan)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = lh.loanService.CreateLoan(&newLoan)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	writeJSON(w, newLoan)
}

// UpdateLoan handles PUT requests to update a loan
func (lh *LoanHandlers) UpdateLoan(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromParams(r)
	if err != nil {
		http.Error(w, "Invalid loan ID", http.StatusBadRequest)
		return
	}
	var updatedLoan loan.Loan
	err = json.NewDecoder(r.Body).Decode(&updatedLoan)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	updatedLoan.ID = id
	err = lh.loanService.UpdateLoan(&updatedLoan)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, updatedLoan)
}

// DeleteLoan handles DELETE requests to delete a loan
func (lh *LoanHandlers) DeleteLoan(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromParams(r)
	if err != nil {
		http.Error(w, "Invalid loan ID", http.StatusBadRequest)
		return
	}
	err = lh.loanService.DeleteLoan(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// GetLoansByMemberID handles GET requests to retrieve loans by member ID
func (lh *LoanHandlers) GetLoansByMemberID(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromParams(r)
	if err != nil {
		http.Error(w, "Invalid member ID", http.StatusBadRequest)
		return
	}
	loans, err := lh.loanService.GetLoansByMemberID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, loans)
}

// GetLoansByBookID handles GET requests to retrieve loans by book ID
func (lh *LoanHandlers) GetLoansByBookID(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromParams(r)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}
	loans, err := lh.loanService.GetLoansByBookID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, loans)
}

// Helper functions
func getIDFromParams(r *http.Request) (int, error) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	return id, err
}

func writeJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
