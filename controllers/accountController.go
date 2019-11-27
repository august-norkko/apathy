package controllers

import (
	"net/http"
	"fmt"
	"regexp"
	"apathy/response"
	"apathy/interfaces"
	"apathy/security"
)

type AccountController struct {
	interfaces.IAccountService
}

func (controller *AccountController) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	ok, err := controller.CreateAccount(r)
	if err != nil {
		response.Send(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	if !ok {
		response.Send(w, http.StatusBadRequest, "Failed to create account")
		return
	}

	response.Send(w, http.StatusOK, "Created account successfully")
	return
}

func (controller *AccountController) LoginHandler(w http.ResponseWriter, r *http.Request) {
	token, err := controller.LoginAccount(r)
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

func (controller *AccountController) DashboardHandler(w http.ResponseWriter, r *http.Request) {
	_, claims, _ := security.ParseToken(r.Header.Get("Authorization"))
	id := claims.Id
	response.Send(w, http.StatusOK, fmt.Sprint(id))
	return
}
