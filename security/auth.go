package security

import (
	"log"
	"fmt"
	"net/http"
	"regexp"
	"apathy/utils"
)

// All request goes through this middleware
func Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		allowed := []string{"/new", "/login", "/"}
		for _, current := range allowed {
			if current == r.URL.Path {
				next.ServeHTTP(w, r)
				return
			}
		}
		
		// Check for empty header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			msg := utils.Message(http.StatusForbidden, "Missing Authorization Header")
			utils.Respond(w, msg)
			return
		}

		// Check for malformed header
		match, _ := regexp.MatchString(`^Bearer [A-Za-z0-9-_=]+\.[A-Za-z0-9-_=]+\.?[A-Za-z0-9-_.+/=]*$`, authHeader)
		if match == false {
			msg := utils.Message(http.StatusForbidden, "Malformed Authorization Header")
			utils.Respond(w, msg)
			return
		}

		token, err := ParseToken(authHeader)
		if err != nil {
			log.Println(err)
			msg := utils.Message(http.StatusForbidden, fmt.Sprint(err))
			utils.Respond(w, msg)
			return
		}

		if !token.Valid {
			log.Println(err)
			msg := utils.Message(http.StatusForbidden, fmt.Sprint(err))
			utils.Respond(w, msg)
			return
		}

		next.ServeHTTP(w, r)
		return
	})
}
