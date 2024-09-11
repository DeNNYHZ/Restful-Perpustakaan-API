package repositories

import (
	"Restful-Perpustakaan-API/app/models"
	"database/sql"
	"errors"
	"time"

	_ "github.com/lib/pq" // Import driver PostgreSQL
)

// NotificationRepository provides methods for interacting with notification data in the database
type NotificationRepository struct {
	db *sql.DB
}

// NewNotificationRepository creates a new NotificationRepository instance
func NewNotificationRepository(db *sql.DB) *NotificationRepository {
	return &NotificationRepository{db: db}
}

// GetAllNotifications mengambil semua notifikasi dari database
func (nr *NotificationRepository) GetAllNotifications() ([]models.Notification, error) {
	query := "SELECT id, user_id, message, timestamp, is_read FROM notifications"

	rows, err := nr.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notifications []models.Notification
	for rows.Next() {
		var n models.Notification
		err := rows.Scan(&n.ID, &n.UserID, &n.Message, &n.Timestamp, &n.IsRead)
		if err != nil {
			return nil, err
		}
		notifications = append(notifications, n)
	}

	return notifications, nil
}

// GetNotificationByID mengambil notifikasi berdasarkan ID dari database
func (nr *NotificationRepository) GetNotificationByID(id int) (*models.Notification, error) {
	query := "SELECT id, user_id, message, timestamp, is_read FROM notifications WHERE id = $1"

	var n models.Notification
	err := nr.db.QueryRow(query, id).Scan(&n.ID, &n.UserID, &n.Message, &n.Timestamp, &n.IsRead)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("notification not found")
		}
		return nil, err
	}

	return &n, nil
}

// CreateNotification membuat notifikasi baru di database
func (nr *NotificationRepository) CreateNotification(n *models.Notification) error {
	query := `
        INSERT INTO notifications (user_id, message, timestamp, is_read) 
        VALUES ($1, $2, $3, $4)
        RETURNING id
    `

	err := nr.db.QueryRow(query, n.UserID, n.Message, time.Now(), n.IsRead).Scan(&n.ID)
	if err != nil {
		return err
	}

	return nil
}

// UpdateNotification memperbarui notifikasi di database
func (nr *NotificationRepository) UpdateNotification(n *models.Notification) error {
	query := `
        UPDATE notifications 
        SET user_id = $1, message = $2, timestamp = $3, is_read = $4
        WHERE id = $5
    `

	_, err := nr.db.Exec(query, n.UserID, n.Message, n.Timestamp, n.IsRead, n.ID)
	return err
}

// DeleteNotification menghapus notifikasi dari database
func (nr *NotificationRepository) DeleteNotification(id int) error {
	query := "DELETE FROM notifications WHERE id = $1"
	_, err := nr.db.Exec(query, id)
	return err
}

// MarkNotificationAsRead menandai notifikasi sebagai telah dibaca
func (nr *NotificationRepository) MarkNotificationAsRead(id int) error {
	query := "UPDATE notifications SET is_read = true WHERE id = $1"
	_, err := nr.db.Exec(query, id)
	return err
}

// ... (fungsi lain yang mungkin Anda butuhkan, seperti GetNotificationsByUserID, dll.)
