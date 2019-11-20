package main

import (
	"github.com/gorilla/mux"
	"apathy/security"
)

type IMuxRouter interface {
	InitializeRouter() *mux.Router
}

type router struct{}

func MuxRouter() IMuxRouter {
	return &router{}
}

func (router *router) InitializeRouter() *mux.Router {
	r := mux.NewRouter()
	r.Use(security.Authentication)

	container := Container()
	userController := container.InjectUserController()
	r.HandleFunc("/new", userController.RegisterHandler)
	r.HandleFunc("/login", userController.LoginHandler)
	
	return r
}