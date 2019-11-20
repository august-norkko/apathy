package controllers

import (
	_ "log"
	"net/http"
	"apathy/utils"
	"apathy/interfaces"
)

type UserController struct {
	interfaces.IUserService
}

func (controller *UserController) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	ok, err := controller.CreateUser(r)
	if err != nil {
		utils.Response(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	if !ok {
		utils.Response(w, http.StatusBadRequest, "Failed to create user")
		return
	}

	utils.Response(w, http.StatusOK, "Created user successfully")
	return
}

func (controller *UserController) LoginHandler(w http.ResponseWriter, r *http.Request) {
	token, err := controller.LoginUser(r)
	if err != nil {
		utils.Response(w, http.StatusBadRequest, "Incorrect credentials")
		return
	}

	if len(token) <= 0 {
		utils.Response(w, http.StatusBadRequest, "Unable to generate JWT token")
		return
	}

	utils.ResponseToken(w, token)
	return
}
