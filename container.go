package main

import (
	"apathy/services"
	"apathy/repositories"
	"apathy/controllers"
)

type IContainer interface {
	InjectAccountController() controllers.AccountController
}

type container struct{}

func Container() IContainer {
	return &container{}
}

func (container *container) InjectAccountController() controllers.AccountController {
	accountRepository := &repositories.AccountRepository{}
	accountService := &services.AccountService{accountRepository}
	accountController := controllers.AccountController{accountService}
	return accountController
}