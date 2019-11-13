package auth

import (
	"log"
	"net/http"
	"apathy/utils"
)

// requests goes through this middleware
func Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Print("Incoming request")

		allowed := []string{"/foo", "/baz", "/user/new", "/user/login"}
		for _, current := range allowed {
			if current == r.URL.Path {
				next.ServeHTTP(w, r)
				return
			}
		}

		log.Print(r.URL.Path)
		//log.Print("Request being authenticated: ", r)

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			msg := utils.Message(401, "Unauthorized, missing JWT token")
			utils.Response(w, msg)
			return
		}

		next.ServeHTTP(w, r)
	})
}