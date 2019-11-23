package main

import (
	"log"
	"os"
	"io"
	"net/http"
	"github.com/joho/godotenv"
	"apathy/database"
)

func main() {
	file, err := os.OpenFile("logs.log", os.O_CREATE | os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	mw := io.MultiWriter(os.Stdout, file)
	log.SetOutput(mw)

	err = godotenv.Load()
	if err != nil {
		log.Println(err)
		log.Fatal("Error loading .env file")
	}

	database.Initialize()
	http.ListenAndServe(":3000", MuxRouter().InitializeRouter())
}
