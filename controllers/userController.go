package controller

import (
	"log"
	"net/http"
	"apathy/utils"
	"apathy/interfaces"
)

type UserController struct {
	interfaces.IUserService
}

func (controller *UserController) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	status, msg, err := controller.CreateUser(r)
	if err != nil {
		log.Println(err)
	}

	utils.Respond(w, utils.Message(status, msg))
	return
}

func (controller *UserController) LoginHandler(w http.ResponseWriter, r *http.Request) {
	token, err := controller.LoginUser(r)
	if len(token) == 0 || err != nil {
		log.Println(err)
		utils.Respond(w, utils.Message(http.StatusBadRequest, "Incorrect email or password"))
		return
	}

	utils.Respond(w, utils.GiveToken(token))
	return
}

func (controller *UserController) UserHandler(w http.ResponseWriter, r *http.Request) {
	utils.Respond(w, utils.Message(http.StatusOK, "Authenticated"))
}