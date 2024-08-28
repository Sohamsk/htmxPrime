package middleware

import (
	"log"
	"net/http"
)

// type middleware struct{}

func RequestLoggerMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL)
		next.ServeHTTP(w, r)
	}
}
