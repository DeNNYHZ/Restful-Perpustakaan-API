package utils

import (
	"os"

	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

// InitializeLogger menginisialisasi logger dengan pengaturan yang sesuai.
func InitializeLogger() {
	logger = logrus.New()

	// Set output ke stdout
	logger.SetOutput(os.Stdout)

	// Set format log menjadi JSON
	logger.SetFormatter(&logrus.JSONFormatter{})

	// Set level log (opsional)
	// logger.SetLevel(logrus.DebugLevel) // Untuk logging yang lebih detail
}

// GetLogger mengembalikan instance logger yang telah diinisialisasi.
func GetLogger() *logrus.Logger {
	if logger == nil {
		InitializeLogger()
	}
	return logger
}
