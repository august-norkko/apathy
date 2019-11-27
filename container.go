package main

import (
	"apathy/services"
	"apathy/repositories"
	"apathy/controllers"
)

type IContainer interface {
	InjectUserController() controllers.UserController
}

type container struct{}

func Container() IContainer {
	return &container{}
}

func (c *container) InjectUserController() controllers.UserController {
	userRepository := &repositories.UserRepository{}
	userService := &services.UserService{userRepository}
	userController := controllers.UserController{userService}
	return userController
}