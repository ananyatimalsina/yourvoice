package middleware

import "net/http"

func LogOriginalURL(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Store the original URL in the request header for later use
		r.Header.Set("X-Original-URL", r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
