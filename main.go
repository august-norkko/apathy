package main

import (
	"log"
	"net/http"
	"apathy/database"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
		log.Fatal("Error loading .env file")
	}

	database.Initialize()
	http.ListenAndServe(":3000", MuxRouter().InitializeRouter())
}
