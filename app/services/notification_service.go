package services

import (
	"Restful-Perpustakaan-API/app/models"
	"Restful-Perpustakaan-API/app/repositories"
)

// NotificationService provides methods for managing notifications
type NotificationService struct {
	notificationRepository repositories.NotificationRepository
}

// NewNotificationService creates a new NotificationService instance
func NewNotificationService(notificationRepository repositories.NotificationRepository) *NotificationService {
	return &NotificationService{notificationRepository: notificationRepository}
}

// GetAllNotifications mengambil semua notifikasi
func (ns *NotificationService) GetAllNotifications() ([]models.Notification, error) {
	return ns.notificationRepository.GetAllNotifications()
}

// GetNotificationByID mengambil notifikasi berdasarkan ID
func (ns *NotificationService) GetNotificationByID(id int) (*models.Notification, error) {
	return ns.notificationRepository.GetNotificationByID(id)
}

// CreateNotification membuat notifikasi baru
func (ns *NotificationService) CreateNotification(n *models.Notification) error {
	// Anda dapat menambahkan logika validasi atau bisnis lainnya di sini sebelum menyimpan notifikasi ke database
	return ns.notificationRepository.CreateNotification(n)
}

// UpdateNotification memperbarui notifikasi
func (ns *NotificationService) UpdateNotification(n *models.Notification) error {
	// Anda dapat menambahkan logika validasi atau bisnis lainnya di sini sebelum memperbarui notifikasi di database
	return ns.notificationRepository.UpdateNotification(n)
}

// DeleteNotification menghapus notifikasi
func (ns *NotificationService) DeleteNotification(id int) error {
	return ns.notificationRepository.DeleteNotification(id)
}

// MarkNotificationAsRead menandai notifikasi sebagai telah dibaca
func (ns *NotificationService) MarkNotificationAsRead(id int) error {
	return ns.notificationRepository.MarkNotificationAsRead(id)
}

// ... (fungsi lain yang mungkin Anda butuhkan, seperti GetNotificationsByUserID, dll.)
