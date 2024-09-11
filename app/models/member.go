package models

import "time"

// Member represents a library member.
type Member struct {
	ID               int       `json:"id"`
	Name             string    `json:"name"`
	Email            string    `json:"email"`
	Password         string    `json:"password"`
	PhoneNumber      string    `json:"phone_number"`
	Address          string    `json:"address"`
	RegistrationDate time.Time `json:"registration_date"` // e.g., "2022-01-01T00:00:00Z"
	MembershipType   string    `json:"membership_type"`   // e.g., "Regular", "Student", "Premium"
	Username         string    `json:"username"`          // e.g., "johndoe"
	Gender           string    `json:"gender"`            // e.g., "Male", "Female", "Other" or any

	// ... tambahkan field lain sesuai kebutuhan
}

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"` // Ensure passwords are securely handled
}

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Simulate storage with maps for quick lookups.
var (
	members = make(map[int]Member)
)

// Helper function to generate new IDs.
func generateID(m interface{}) int {
	switch m := m.(type) {
	case map[int]Member:
		if len(m) == 0 {
			return 1
		}
		var maxID int
		for id := range m {
			if id > maxID {
				maxID = id
			}
		}
		return maxID + 1
	default:
		return 1
	}
}
