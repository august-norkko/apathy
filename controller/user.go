package controller

import (
	_ "fmt"
	"log"
	"time"
	"encoding/json"
	"net/http"
	"apathy/utils"
	"golang.org/x/crypto/bcrypt"
	"github.com/dgrijalva/jwt-go"
)

var secret = []byte("secret") // temp

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
	Email 		string		`json:"email"`
	Password 	string		`json:"password"`
	CreatedAt 	time.Time	`json:"createdAt"`
	UpdatedAt 	time.Time	`json:"updatedAt"`
}

type Claim struct {
	Email	string	`json:"email"`
	jwt.StandardClaims
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	data := decodeJson(r)
	password := []byte(data.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
    if err != nil {
        log.Println(err)
	}

	data = User{
		Email:		data.Email,
		Password:	string(hashedPassword),
		CreatedAt:	time.Now(),
		UpdatedAt:	time.Now(),
	}

	expiration := time.Now().Add(10 * time.Minute)
	claim := &Claim{
		Email: data.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiration.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := token.SignedString(secret)
	if err != nil {
		log.Println(err)
	}

	log.Println(signedToken)
	msg := utils.Message(http.StatusOK, signedToken)
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
		msg := utils.Message(http.StatusForbidden, "Invalid email or password")
		utils.Response(w, msg)
		return
	}

	msg := utils.Message(http.StatusOK, "Successful login")
	utils.Response(w, msg)
	return

	// todo
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	msg := utils.Message(http.StatusOK, "Successful")
	utils.Response(w, msg)
	return
}
