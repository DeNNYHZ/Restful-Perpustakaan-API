package utils

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

// AppError represents a custom application error with an HTTP status code and a user-friendly message
type AppError struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

// Error implements the error interface
func (e *AppError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.StatusCode, e.Message)
}

// NewAppError creates a new AppError
func NewAppError(statusCode int, message string) *AppError {
	return &AppError{StatusCode: statusCode, Message: message}
}

// HandleError handles an error by sending an appropriate HTTP response
func HandleError(w http.ResponseWriter, err error) {
	var appErr *AppError
	if errors.As(err, &appErr) {
		// If it's a custom AppError, use its status code and message
		http.Error(w, appErr.Message, appErr.StatusCode)
	} else {
		// For other errors, log the error and send a generic internal server error response
		log.Println(err) // You might want to use a more robust logging mechanism
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// HandleError2 handles an error by sending an appropriate HTTP response
func HandleError2(w http.ResponseWriter, err error) {
	var appErr *AppError
	if errors.As(err, &appErr) {
		// If it's a custom AppError, use its status code and message
		http.Error(w, appErr.Message, appErr.StatusCode)
	} else {
		// For other errors, log the error and send a generic internal server error response
		log.Println(err) // You might want to use a more robust logging mechanism
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// HandleError3 handles an error by sending an appropriate HTTP response
func HandleError3(w http.ResponseWriter, err error) {
	var appErr *AppError
	if errors.As(err, &appErr) {
		// If it's a custom AppError, use its status code and message
		http.Error(w, appErr.Message, appErr.StatusCode)
	} else {
		// For other errors, log the error and send a generic internal server error response
		log.Println(err) // You might want to use a more robust logging mechanism
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
