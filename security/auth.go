package security

import (
	"fmt"
	"net/http"
	"regexp"
	"apathy/response"
)

func Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		allowed := []string{
			"/new",
			"/login",
			"/",
		}

		for _, current := range allowed {
			if current == r.URL.Path {
				next.ServeHTTP(w, r)
				return
			}
		}
		
		authHeader := r.Header.Get("Authorization")
		if len(authHeader) <= 0 {
			response.Send(w, http.StatusBadRequest, "Missing Authorization header")
			return
		}

		ok, _ := regexp.MatchString(`^Bearer [A-Za-z0-9-_=]+\.[A-Za-z0-9-_=]+\.?[A-Za-z0-9-_.+/=]*$`, authHeader)
		if !ok {
			response.Send(w, http.StatusBadRequest, "Malformed Authorization header")
			return
		}

		token, _, err := ParseToken(authHeader)
		if err != nil {
			response.Send(w, http.StatusBadRequest, fmt.Sprint(err))
			return
		}

		if !token.Valid {
			response.Send(w, http.StatusBadRequest, fmt.Sprint(err))
			return
		}

		next.ServeHTTP(w, r)
		return
	})
}
