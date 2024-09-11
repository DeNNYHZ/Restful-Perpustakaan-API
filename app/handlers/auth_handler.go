package handlers

import (
	"Restful-Perpustakaan-API/app/models"
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"Restful-Perpustakaan-API/database"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// Login menangani permintaan login anggota.
func Login(w http.ResponseWriter, r *http.Request) {
	var credentials models.Credentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validasi input
	if credentials.Email == "" || credentials.Password == "" {
		http.Error(w, "Email and password are required", http.StatusBadRequest)
		return
	}

	// Lakukan validasi kredensial (misalnya, cek di database)
	member, err := database.GetMemberByEmail(credentials.Email)
	if err != nil {
		http.Error(w, "Member not found", http.StatusNotFound)
		return
	}

	// Bandingkan password dengan hash yang tersimpan di database
	err = bcrypt.CompareHashAndPassword([]byte(member.Password), []byte(credentials.Password))
	if err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	// Buat token JWT
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
		Subject:   strconv.Itoa(member.ID),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("your_secret_key"))
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	// Kirim token JWT dalam response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"token": tokenString,
	})
}

// Register handles new user registration requests.
func Register(w http.ResponseWriter, r *http.Request) {
	var newMember models.Member // Use member.Member directly
	err := json.NewDecoder(r.Body).Decode(&newMember)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validasi data anggota baru
	if newMember.Name == "" || newMember.Email == "" || newMember.Password == "" {
		http.Error(w, "Name, email, and password are required", http.StatusBadRequest)
		return
	}

	// Validasi format email (opsional)
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(newMember.Email) {
		http.Error(w, "Invalid email format", http.StatusBadRequest)
		return
	}

	// Cek apakah email sudah ada
	existingMember, _ := database.GetMemberByEmail(newMember.Email)
	if existingMember != nil {
		http.Error(w, "Email already exists", http.StatusConflict)
		return
	}

	// Hash password sebelum disimpan
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newMember.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}
	newMember.Password = string(hashedPassword)

	// Simpan anggota baru ke database
	err = database.CreateMember(&newMember) // Pass a pointer to member.Member
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newMember)
}
