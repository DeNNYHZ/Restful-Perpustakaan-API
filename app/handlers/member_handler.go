package handlers

import (
	"Restful-Perpustakaan-API/app/models"
	"Restful-Perpustakaan-API/database"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// MemberHandler handles HTTP requests related to members
type MemberHandler struct {
	DB *database.Database
}

// NewMemberHandler creates a new instance of MemberHandler
func NewMemberHandler(db *database.Database) *MemberHandler {
	return &MemberHandler{DB: db}
}

// GetAllMembers handles GET requests for all members
// Returns a JSON response with a list of members
func (h *MemberHandler) GetAllMembers(w http.ResponseWriter, r *http.Request) ([]models.Member, error) {
	members, err := h.DB.GetAllMembers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil, err
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(members)
	return members, nil
}

// GetMemberByID handles GET requests for a member by ID
// Returns a JSON response with a single member
func (h *MemberHandler) GetMemberByID(w http.ResponseWriter, r *http.Request) (*models.Member, error) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid member ID", http.StatusBadRequest)
		return nil, err
	}

	member, err := h.DB.GetMemberByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return nil, err
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(member)
	return member, nil
}

// CreateMember handles POST requests to create a new member
// Returns a JSON response with the created member
func (h *MemberHandler) CreateMember(w http.ResponseWriter, r *http.Request) (*models.Member, error) {
	var newMember models.Member
	err := json.NewDecoder(r.Body).Decode(&newMember)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return nil, err
	}
	err = h.DB.CreateMember(&newMember)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil, err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newMember)
	return &newMember, nil
}

// UpdateMember handles PUT requests to update an existing member
// Returns a JSON response with the updated member
func (h *MemberHandler) UpdateMember(w http.ResponseWriter, r *http.Request) (*models.Member, error) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid member ID", http.StatusBadRequest)
		return nil, err
	}

	var updatedMember models.Member
	err = json.NewDecoder(r.Body).Decode(&updatedMember)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return nil, err
	}
	updatedMember.ID = id

	err = h.DB.UpdateMember(&updatedMember)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil, err
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedMember)
	return &updatedMember, nil
}

// DeleteMember handles DELETE requests to remove a member
// Returns a JSON response with no content
func (h *MemberHandler) DeleteMember(w http.ResponseWriter, r *http.Request) error {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid member ID", http.StatusBadRequest)
		return err
	}

	err = h.DB.DeleteMember(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}

	w.WriteHeader(http.StatusNoContent)
	return nil
}

// GetMemberByName handles GET requests to find members by name
// Returns a JSON response with a list of members
func (h *MemberHandler) GetMemberByName(w http.ResponseWriter, r *http.Request) ([]models.Member, error) {
	params := mux.Vars(r)
	name := params["name"]
	members, err := h.DB.GetMemberByName(name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil, err
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(members)
	return members, nil
}

// GetMemberByAddress handles GET requests to find members by address
// Returns a JSON response with a list of members
func (h *MemberHandler) GetMemberByAddress(w http.ResponseWriter, r *http.Request) ([]models.Member, error) {
	params := mux.Vars(r)
	address := params["address"]
	members, err := h.DB.GetMemberByAddress(address)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil, err
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(members)
	return members, nil
}

// GetMemberByPhoneNumber handles GET requests to find members by phone number
// Returns a JSON response with a list of members
func (h *MemberHandler) GetMemberByPhoneNumber(w http.ResponseWriter, r *http.Request) ([]models.Member, error) {
	params := mux.Vars(r)
	phoneNumber := params["phoneNumber"]
	members, err := h.DB.GetMemberByPhoneNumber(phoneNumber)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil, err
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(members)
	return members, nil
}

// GetMemberByEmail handles GET requests to find members by email
// Returns a JSON response with a list of members
func (h *MemberHandler) GetMemberByEmail(w http.ResponseWriter, r *http.Request) ([]models.Member, error) {
	params := mux.Vars(r)
	email := params["email"]
	members, err := h.DB.GetMemberByEmail(email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil, err
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(members)
	return members, nil
}

// GetMemberByGender handles GET requests to find members by gender
// Returns a JSON response with a list of members
func (h *MemberHandler) GetMemberByGender(w http.ResponseWriter, r *http.Request) ([]models.Member, error) {
	params := mux.Vars(r)
	gender := params["gender"]
	members, err := h.DB.GetMemberByGender(gender)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil, err
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(members)
	return members, nil
}

// GetMembersByMembershipType handles GET requests to find members by membership type
// Returns a JSON response with a list of members
func (h *MemberHandler) GetMembersByMembershipType(w http.ResponseWriter, r *http.Request) ([]models.Member, error) {
	params := mux.Vars(r)
	membershipType := params["membershipType"]
	members, err := h.DB.GetMembersByMembershipType(membershipType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil, err
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(members)
	return members, nil
}

// GetMembersByRegistrationDate handles GET requests to find members by registration date
// Returns a JSON response with a list of members
func (h *MemberHandler) GetMembersByRegistrationDate(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	registrationDateStr := params["registrationDate"]
	layout := "2006-01-02" // assuming the date format is YYYY-MM-DD
	registrationDate, err := time.Parse(layout, registrationDateStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	members, err := h.DB.GetMembersByRegistrationDate(registrationDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(members)
}

// GetMembersByLastLoginDate handles GET requests to find members by last login date
// Returns a JSON response with a list of members
func (h *MemberHandler) GetMembersByLastLoginDate(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	lastLoginDateStr := params["lastLoginDate"]
	layout := "2006-01-02" // assuming the date format is YYYY-MM-DD
	lastLoginDate, err := time.Parse(layout, lastLoginDateStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	members, err := h.DB.GetMembersByLastLoginDate(lastLoginDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(members)
}
