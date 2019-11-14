package main

import (
	"log"
	"os"
	"io"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"apathy/security"
	"apathy/controller"
	"apathy/database"
)

func main() {
	setup()
	router := mux.NewRouter()

	router.HandleFunc("/health", controller.HealthcheckHandler)
	router.HandleFunc("/user", controller.UserHandler).Methods("POST")
	router.HandleFunc("/user/new", controller.RegisterHandler).Methods("POST")
	router.HandleFunc("/user/login", controller.LoginHandler).Methods("POST")
	
	router.Use(security.Authentication)
	http.ListenAndServe(":3000", router)
}


func setup() {
	file, err := os.OpenFile("logs.log", os.O_CREATE | os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	mw := io.MultiWriter(os.Stdout, file)
	log.SetOutput(mw)
	log.Print("Logging has begun.")

	err = godotenv.Load()
	if err != nil {
		log.Println(err)
		log.Fatal("Error loading environment file")
	}
	log.Println("Loaded environment file")

	database.Initialize()
	return
}