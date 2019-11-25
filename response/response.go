package response

import (
	"net/http"
	"encoding/json"
)

func Send(w http.ResponseWriter, status int, m string) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]interface{} { "message": m })
}

func SendToken(w http.ResponseWriter, token string) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{} { "token": token })
}