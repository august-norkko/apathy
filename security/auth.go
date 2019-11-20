package security

import (
	"log"
	"fmt"
	"net/http"
	"regexp"
	"apathy/utils"
)

func Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		allowed := []string{"/new", "/login", "/"}
		for _, current := range allowed {
			if current == r.URL.Path {
				next.ServeHTTP(w, r)
				return
			}
		}
		
		authHeader := r.Header.Get("Authorization")
		if len(authHeader) <= 0 {
			msg := utils.Message(http.StatusForbidden, "Missing Authorization Header")
			utils.Respond(w, msg)
			return
		}

		ok, _ := regexp.MatchString(`^Bearer [A-Za-z0-9-_=]+\.[A-Za-z0-9-_=]+\.?[A-Za-z0-9-_.+/=]*$`, authHeader)
		if !ok {
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
