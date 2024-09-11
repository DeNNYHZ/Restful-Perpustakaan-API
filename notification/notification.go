package notification

import "time"

// Notification represents a notification to a user
type Notification struct {
	ID        int       `json:"id" :"id"`
	UserID    int       `json:"user_id" :"user_id"`
	Name      string    `json:"name" :"name"`
	Message   string    `json:"message" :"message"`
	Timestamp time.Time `json:"timestamp" :"timestamp"`
	IsRead    bool      `json:"is_read" :"is_read"`
	Read      bool      `json:"read" :"read"`
	Category  string    `json:"category" :"category"`
	Receiver  string    `json:"receiver" :"receiver"`
	Sender    string    `json:"sender" : sender"`
}
