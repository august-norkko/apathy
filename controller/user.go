package controller

import (
	"net/http"
	"apathy/utils"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	msg := utils.Message(200, "Successful")
	utils.Response(w, msg)
	return
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	msg := utils.Message(200, "Successful")
	utils.Response(w, msg)
	return
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	msg := utils.Message(200, "Successful")
	utils.Response(w, msg)
	return
}