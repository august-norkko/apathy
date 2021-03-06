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
	accountController := container.InjectAccountController()
	r.HandleFunc("/new", accountController.RegisterHandler).Methods("POST")
	r.HandleFunc("/login", accountController.LoginHandler).Methods("POST")

	r.HandleFunc("/account", accountController.DashboardHandler).Methods("GET")
	r.HandleFunc("/account", accountController.UpdateHandler).Methods("PUT")
	r.HandleFunc("/account", accountController.DeleteHandler).Methods("DELETE")
	return r
}
