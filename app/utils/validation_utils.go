package utils

import (
	"regexp"
	"unicode"
)

// ValidateEmail memvalidasi format alamat email menggunakan regular expression
func ValidateEmail(email string) bool {
	// Regular expression untuk validasi email yang umum digunakan
	// Anda dapat menyesuaikannya jika perlu
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

// ValidatePassword memvalidasi kekuatan password
// Anda dapat menyesuaikan aturan validasi sesuai kebutuhan keamanan Anda
func ValidatePassword(password string) bool {
	// Contoh aturan validasi:
	// - Minimal 8 karakter
	// - Mengandung setidaknya satu huruf besar, satu huruf kecil, dan satu angka

	if len(password) < 8 {
		return false
	}

	hasUpper := false
	hasLower := false
	hasNumber := false

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		}
	}

	return hasUpper && hasLower && hasNumber
}

// ValidateStringLength memvalidasi panjang string
// minLength adalah panjang minimum yang diizinkan
// maxLength adalah panjang maksimum yang diizinkan
func ValidateStringLength(str string, minLength, maxLength int) bool {
	return len(str) >= minLength && len(str) <= maxLength
}
