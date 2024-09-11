package handlers

import (
	"Restful-Perpustakaan-API/database"
	"Restful-Perpustakaan-API/member"
	"encoding/json"
	"net/http"
	"strconv"

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
func (h *MemberHandler) GetAllMembers(w http.ResponseWriter, r *http.Request) {
	members, err := h.DB.GetAllMembers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(members)
}

// GetMemberByID handles GET requests for a member by ID
func (h *MemberHandler) GetMemberByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid member ID", http.StatusBadRequest)
		return
	}

	member, err := h.DB.GetMemberByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(member)
}

// CreateMember handles POST requests to create a new member
func (h *MemberHandler) CreateMember(w http.ResponseWriter, r *http.Request) {
	var newMember member.Member
	err := json.NewDecoder(r.Body).Decode(&newMember)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.DB.CreateMember(&newMember)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newMember)
}

// UpdateMember handles PUT requests to update an existing member
func (h *MemberHandler) UpdateMember(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid member ID", http.StatusBadRequest)
		return
	}

	var updatedMember member.Member
	err = json.NewDecoder(r.Body).Decode(&updatedMember)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	updatedMember.ID = id

	err = h.DB.UpdateMember(&updatedMember)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedMember)
}

// DeleteMember handles DELETE requests to remove a member
func (h *MemberHandler) DeleteMember(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid member ID", http.StatusBadRequest)
		return
	}

	err = h.DB.DeleteMember(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// GetMemberByName handles GET requests to find members by name
func (h *MemberHandler) GetMemberByName(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]
	members, err := h.DB.GetMemberByName(name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(members)
}

// GetMemberByAddress handles GET requests to find members by address
func (h *MemberHandler) GetMemberByAddress(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	address := params["address"]
	members, err := h.DB.GetMemberByAddress(address)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(members)
}

// GetMemberByPhoneNumber handles GET requests to find members by phone number
func (h *MemberHandler) GetMemberByPhoneNumber(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	phoneNumber := params["phoneNumber"]
	members, err := h.DB.GetMemberByPhoneNumber(phoneNumber)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(members)
}

// GetMemberByEmail handles GET requests to find members by email
func (h *MemberHandler) GetMemberByEmail(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	email := params["email"]
	members, err := h.DB.GetMemberByEmail(email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(members)
}

// GetMemberByGender handles GET requests to find members by gender
func (h *MemberHandler) GetMemberByGender(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	gender := params["gender"]
	members, err := h.DB.GetMemberByGender(gender)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(members)
}

// GetMembersByMembershipType handles GET requests to find members by membership type
func (h *MemberHandler) GetMembersByMembershipType(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	membershipType := params["membershipType"]
	members, err := h.DB.GetMembersByMembershipType(membershipType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(members)
}

// GetMembersByRegistrationDate handles GET requests to find members by registration date
func (h *MemberHandler) GetMembersByRegistrationDate(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	registrationDate := params["registrationDate"]
	members, err := h.DB.GetMembersByRegistrationDate(registrationDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(members)
}
