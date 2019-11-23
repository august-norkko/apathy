package utils

import (
	"net/http"
	"encoding/json"
	_ "fmt"
)

func Response(w http.ResponseWriter, status int, m string) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]interface{} { "message": m })
}

func ResponseToken(w http.ResponseWriter, token string) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{} { "token": token })
}