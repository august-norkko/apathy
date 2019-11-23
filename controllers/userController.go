package controllers

import (
	_ "log"
	"net/http"
	"apathy/response"
	"apathy/interfaces"
	"regexp"
)

type UserController struct {
	interfaces.IUserService
}

func (controller *UserController) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	ok, err := controller.CreateUser(r)
	if err != nil {
		response.Send(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	if !ok {
		response.Send(w, http.StatusBadRequest, "Failed to create user")
		return
	}

	response.Send(w, http.StatusOK, "Created user successfully")
	return
}

func (controller *UserController) LoginHandler(w http.ResponseWriter, r *http.Request) {
	token, err := controller.LoginUser(r)
	if err != nil {
		response.Send(w, http.StatusBadRequest, "Incorrect credentials")
		return
	}

	ok, _ := regexp.MatchString(`^[A-Za-z0-9-_=]+\.[A-Za-z0-9-_=]+\.?[A-Za-z0-9-_.+/=]*$`, token)
	if !ok {
		response.Send(w, http.StatusBadRequest, "Validation failed")
		return
	}

	response.SendToken(w, token)
	return
}
