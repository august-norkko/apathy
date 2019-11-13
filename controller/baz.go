package controller

import (
	"net/http"
	"apathy/utils"
)

func BazHandler(w http.ResponseWriter, r *http.Request) {
	msg := utils.Message(http.StatusOK, "Successful")
	utils.Response(w, msg)
	return
}
