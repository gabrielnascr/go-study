package middleware

import (
	"fmt"
	"net/http"
)

func BasicMiddleware(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Basic Middleware")
		h.ServeHTTP(w, r)
	}
}
