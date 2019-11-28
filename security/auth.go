package security

import (
	"fmt"
	"net/http"
	"regexp"
	"context"
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

		token, claims, err := ParseToken(authHeader)
		if err != nil {
			response.Send(w, http.StatusBadRequest, fmt.Sprint(err))
			return
		}

		if !token.Valid {
			response.Send(w, http.StatusBadRequest, fmt.Sprint(err))
			return
		}

		c := context.WithValue(r.Context(), "id", claims.Id)
		r = r.WithContext(c)

		next.ServeHTTP(w, r)
		return
	})
}
