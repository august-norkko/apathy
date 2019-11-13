package controller

import (
	"io"
	"net/http"
)

func HealthcheckHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
    io.WriteString(w, `{"alive": true}`)
}