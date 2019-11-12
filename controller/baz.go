package controller

import (
	"net/http"
	"apathy/utils"
)

func BazHandler(w http.ResponseWriter, r *http.Request) {
	msg := utils.Message(200, "Successful")
	utils.Response(w, msg)
	return
}
