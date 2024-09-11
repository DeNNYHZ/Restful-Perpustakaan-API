package user

// User represents a user entity.
type User struct {
	ID       int    `json:"id" :"ID"`
	Username string `json:"username" :"username"`
	Email    string `json:"email" :"email"`
	Password string `json:"password" :"password"` // Ensure passwords are securely handled
	Name     string `:"name"`
}

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Simulate storage with maps for quick lookups.
var (
	users = make(map[int]User)
)
