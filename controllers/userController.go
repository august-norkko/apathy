package controllers

import (
	"net/http"
	"fmt"
	"regexp"
	"apathy/response"
	"apathy/interfaces"
	"apathy/security"
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

func (controller *UserController) DashboardHandler(w http.ResponseWriter, r *http.Request) {
	_, claims, _ := security.ParseToken(r.Header.Get("Authorization"))
	id := claims.Id
	response.Send(w, http.StatusOK, fmt.Sprint(id))
	return
}
