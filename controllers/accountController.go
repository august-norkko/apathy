package controllers

import (
	"net/http"
	"fmt"
	"regexp"
	"apathy/response"
	"apathy/interfaces"
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

func (controller *AccountController) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	ok, err := controller.UpdateAccount(r)
	if !ok {
		response.Send(w, http.StatusBadRequest, fmt.Sprint(err))
		return
	}

	response.Send(w, http.StatusOK, "Successfully updated account")
	return
}

func (controller *AccountController) DashboardHandler(w http.ResponseWriter, r *http.Request) {
	data, err := controller.FetchAccount(r)
	if err != nil {
		response.Send(w, http.StatusBadRequest, "Unable to fetch account")
		return
	}

	response.SendConstructedObject(w, http.StatusOK, map[string]interface{} {
		"username": data.Username,
		"email": data.Email,
		"about": data.About,
		"location": data.Location,
	})
	return
}

func (controller *AccountController) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	ok, _ := controller.DeleteAccount(r)
	if !ok {
		response.Send(w, http.StatusBadRequest, "Unable to delete account")
		return
	}

	response.Send(w, http.StatusOK, "Successfully deleted account")
	return
}