package services

import (
	"Restful-Perpustakaan-API/app/models"
	"Restful-Perpustakaan-API/app/repositories"
	"errors"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// AuthService provides methods for authentication and authorization
type AuthService struct {
	memberRepository repositories.MemberRepository
	secretKey        []byte // Kunci rahasia untuk JWT
}

// NewAuthService creates a new AuthService instance
func NewAuthService(memberRepository repositories.MemberRepository, secretKey []byte) *AuthService {
	return &AuthService{
		memberRepository: memberRepository,
		secretKey:        secretKey,
	}
}

// Login melakukan otentikasi anggota dan mengembalikan token JWT jika berhasil
func (as *AuthService) Login(credentials models.Credentials) (string, error) {
	// Validasi input
	if credentials.Email == "" || credentials.Password == "" {
		return "", errors.New("email and password are required")
	}

	// Ambil anggota dari database berdasarkan email
	member, err := as.memberRepository.GetMemberByEmail(credentials.Email)
	if err != nil {
		return "", err
	}

	// Bandingkan password dengan hash yang tersimpan di database
	err = bcrypt.CompareHashAndPassword([]byte(member.Password), []byte(credentials.Password))
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	// Buat token JWT
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
		Subject:   strconv.Itoa(member.ID),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(as.secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Register mendaftarkan anggota baru dan mengembalikan data anggota jika berhasil
func (as *AuthService) Register(newMember *models.Member) error {
	// Validasi data anggota baru
	if newMember.Name == "" || newMember.Email == "" || newMember.Password == "" {
		return errors.New("name, email, and password are required")
	}

	// Validasi format email (opsional)
	// ...

	// Cek apakah email sudah ada
	existingMember, _ := as.memberRepository.GetMemberByEmail(newMember.Email)
	if existingMember != nil {
		return errors.New("email already exists")
	}

	// Hash password sebelum disimpan
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newMember.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	newMember.Password = string(hashedPassword)

	// Simpan anggota baru ke database
	err = as.memberRepository.CreateMember(newMember)
	if err != nil {
		return err
	}

	return nil
}
