package middleware

import (
	"net/http"
)

func SetCORSHeaders(w http.ResponseWriter) {
	// CORS header because frontend is on a different port
	// no more "CORS MISSING ALLOW ORIGIN" errors
	w.Header().Set("Content-Type", "application/json") // expected request format is JSON
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}
