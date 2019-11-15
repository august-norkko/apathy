package controller

import (
	"log"
	"net/http"
	"apathy/utils"
	"apathy/services"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	service := services.UserService()
	status, msg, err := service.CreateUser(r)
	if err != nil {
		log.Println(err)
	}

	utils.Respond(w, utils.Message(status, msg))
	return
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	service := services.UserService()
	token, err := service.LoginUser(r)
	if len(token) == 0 || err != nil {
		log.Println(err)
		utils.Respond(w, utils.Message(http.StatusBadRequest, "Incorrect email or password"))
		return
	}

	utils.Respond(w, utils.GiveToken(token))
	return
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	header := r.Header.Get("Authorization")

	service := services.UserService()
	user, _ := service.User(header) // returns user entity

	msg := utils.Message(http.StatusOK, user.Email)
	utils.Respond(w, msg)
	return
}
