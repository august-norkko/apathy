package controllers

import (
	_ "log"
	"net/http"
	"apathy/utils"
	"apathy/interfaces"
	"regexp"
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

	ok, _ := regexp.MatchString(`^[A-Za-z0-9-_=]+\.[A-Za-z0-9-_=]+\.?[A-Za-z0-9-_.+/=]*$`, token)
	if !ok {
		utils.Response(w, http.StatusBadRequest, "Validation failed")
		return
	}

	utils.ResponseToken(w, token)
	return
}
