package controller

import (
	_ "fmt"
	"log"
	"time"
	"encoding/json"
	"net/http"
	"apathy/utils"
	"golang.org/x/crypto/bcrypt"
)

func decodeJson(r *http.Request) User {
	decoder := json.NewDecoder(r.Body)
	var data User
	err := decoder.Decode(&data)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

type User struct {
	Email string
	Password string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	data := decodeJson(r)
	password := []byte(data.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
    if err != nil {
        panic(err)
	}
	data.Password = string(hashedPassword)
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()
	
	// validation, save in db
	log.Println(data)

	msg := utils.Message(200, "Successful")
	utils.Response(w, msg)
	return
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	data := decodeJson(r)
	password := []byte(data.Password)

	// actually fetch from db
	hashedPassword := []byte("$2a$10$Bb40xvsagDfOr6XS5V2vyOaf6.qBGToycDNisWOCiFJRUSnt9vKb.")
	
    err := bcrypt.CompareHashAndPassword(hashedPassword, password)
    if err != nil {
		log.Println(err) // no match
		msg := utils.Message(301, "Invalid email or password")
		utils.Response(w, msg)
		return
	}

	msg := utils.Message(200, "Successful login")
	utils.Response(w, msg)
	return

	// todo
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	msg := utils.Message(200, "Successful")
	utils.Response(w, msg)
	return
}
