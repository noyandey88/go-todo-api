package middleware

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now().UTC()

		// Read the body (if any)
		var bodyBytes []byte
		if r.Body != nil {
			bodyBytes, _ = io.ReadAll(r.Body)
			// Restore the body so the next handler can read it
			r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		}

		// Call the next handler
		next.ServeHTTP(w, r)

		// Log method, path, duration, and body
		if len(bodyBytes) > 0 {
			log.Printf("%s %s (%s) Request Body: %s\n", r.Method, r.URL.Path, time.Since(start), string(bodyBytes))
		} else {
			log.Printf("%s %s (%s)\n", r.Method, r.URL.Path, time.Since(start))
		}
	})
}
