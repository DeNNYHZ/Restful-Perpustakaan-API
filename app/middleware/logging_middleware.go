package middleware

import (
	"log"
	"net/http"
	"time"
)

// LoggingMiddleware adalah middleware untuk mencatat log request dan response HTTP.
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Log informasi request
		log.Printf("Request: %s %s from %s", r.Method, r.URL.Path, r.RemoteAddr)

		// Panggil handler selanjutnya
		next.ServeHTTP(w, r)

		// Log informasi response
		duration := time.Since(start)
		log.Printf("Response: %d %s in %v", w.(*responseWriter).statusCode, http.StatusText(w.(*responseWriter).statusCode), duration)
	})
}

// responseWriter adalah struct untuk menyimpan status code response.
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

// WriteHeader meng-override fungsi WriteHeader untuk menyimpan status code response.
func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}
