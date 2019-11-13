package utils

import (
	"io"
	"net/http"
	"encoding/json"
)

func Response(w http.ResponseWriter, data map[string] interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func Message(status int, message string) (map[string]interface{}) {
	return map[string]interface{} {"status": status, "message": message }
}

func HealthcheckHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
    io.WriteString(w, `{"alive": true}`)
}