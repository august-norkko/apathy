package utils

import (
	"net/http"
	"encoding/json"
	_ "fmt"
)

func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func Message(status int, message string) (map[string]interface{}) {
	return map[string]interface{} {"status": status, "message": message }
}

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