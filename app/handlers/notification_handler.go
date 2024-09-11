package handlers

import (
	"Restful-Perpustakaan-API/database"
	"Restful-Perpustakaan-API/notification"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// GetAllNotifications retrieves all notifications from the database.
func GetAllNotifications(w http.ResponseWriter, r *http.Request) {
	notifications, err := database.GetAllNotifications()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notifications)
}

// GetNotificationByID retrieves a notification by ID from the database.
func GetNotificationByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid notification ID", http.StatusBadRequest)
		return
	}

	notification, err := database.GetNotificationByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notification)
}

// CreateNotification creates a new notification from the request body and saves it to the database.
func CreateNotification(w http.ResponseWriter, r *http.Request) {
	var newNotification notification.Notification
	err := json.NewDecoder(r.Body).Decode(&newNotification)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = database.CreateNotification(&newNotification)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newNotification)
}

// UpdateNotification updates a notification by ID from the request body.
func UpdateNotification(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid notification ID", http.StatusBadRequest)
		return
	}

	var updatedNotification notification.Notification
	err = json.NewDecoder(r.Body).Decode(&updatedNotification)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	updatedNotification.ID = id

	err = database.UpdateNotification(&updatedNotification)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedNotification)
}

// DeleteNotification deletes a notification by ID from the database.
func DeleteNotification(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid notification ID", http.StatusBadRequest)
		return
	}

	err = database.DeleteNotification(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// MarkNotificationAsRead marks a notification as read.
func MarkNotificationAsRead(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid notification ID", http.StatusBadRequest)
		return
	}

	err = database.MarkNotificationAsRead(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// MarkAllNotificationsAsRead marks all notifications as read.
func MarkAllNotificationsAsRead(w http.ResponseWriter, r *http.Request) {
	err := database.MarkAllNotificationsAsRead()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// GetUnreadNotificationsCount returns the count of unread notifications.
func GetUnreadNotificationsCount(w http.ResponseWriter, r *http.Request) {
	count, err := database.GetUnreadNotificationsCount()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(count)
}

// GetAllUnreadNotificationsCount returns the count of all unread notifications.
func GetAllUnreadNotificationsCount(w http.ResponseWriter, r *http.Request) {
	count, err := database.GetAllUnreadNotificationsCount()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(count)
}

// GetAllNotificationsCount returns the count of all notifications.
func GetAllNotificationsCount(w http.ResponseWriter, r *http.Request) {
	count, err := database.GetAllNotificationsCount()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(count)
}

// GetNotificationByName retrieves a notification by name from the database.
func GetNotificationByName(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]
	notifications, err := database.GetNotificationByName(name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notifications)
}

// GetNotificationByCategory retrieves a notification by category from the database.
func GetNotificationByCategory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	category := params["category"]
	notifications, err := database.GetNotificationByCategory(category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notifications)
}

// GetNotificationByReceiver retrieves a notification by receiver from the database.
func GetNotificationByReceiver(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	receiver := params["receiver"]
	notifications, err := database.GetNotificationByReceiver(receiver)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notifications)
}

// GetNotificationBySender retrieves a notification by sender from the database.
func GetNotificationBySender(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	sender := params["sender"]
	notifications, err := database.GetNotificationBySender(sender)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notifications)
}
