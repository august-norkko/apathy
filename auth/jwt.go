package auth

import (
	"log"
	"net/http"
	_ "apathy/utils"
)

// requests goes through this middleware
var Authentication = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Print("Request received", r)
		next.ServeHTTP(w, r)
	})
}