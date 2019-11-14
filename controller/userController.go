package controller

import (
	"log"
	"net/http"
	"apathy/utils"
	"apathy/services"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	status, msg, err := services.UserService().CreateUser(r)
	if err != nil {
		log.Println(err)
	}

	utils.Respond(w, utils.Message(status, msg))
	return
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	status, msg, err := services.UserService().LoginUser(r)
	if err != nil {
		log.Println(err)
	}

	utils.Respond(w, utils.Message(status, msg))
	return
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	msg := utils.Message(http.StatusOK, "Successful")
	utils.Respond(w, msg)
	return
}
