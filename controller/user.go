package controller

import (
	_ "fmt"
	"log"
	"time"
	"os"
	"encoding/json"
	"net/http"
	"golang.org/x/crypto/bcrypt"
	"github.com/dgrijalva/jwt-go"
	"apathy/utils"
	"apathy/database"
)

var secret = []byte(os.Getenv("JWT_SECRET")) // temp

func decodeJson(r *http.Request) database.User {
	var data database.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)
	if err != nil {
		log.Fatal(err)
	}
	return data
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

		// no validation
	database.Mysql().Create(&database.User{ Email: data.Email, Password: string(hashedPassword) })

	msg := utils.Message(http.StatusOK, "User created successfully")
	utils.Response(w, msg)
	return
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	data := decodeJson(r)
	password := []byte(data.Password)

	user := &database.User{}
	db := database.Mysql()
	err := db.Table("users").Where("email = ?", data.Email).First(user).Error
	if err != nil {
		log.Println(err)
		msg := utils.Message(http.StatusForbidden, "Email not found")
		utils.Response(w, msg)
		return
	}

    err = bcrypt.CompareHashAndPassword([]byte(user.Password), password)
    if err != nil {
		log.Println(err) // no match
		msg := utils.Message(http.StatusForbidden, "Invalid email or password")
		utils.Response(w, msg)
		return
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

	msg := utils.Message(http.StatusOK, signedToken)
	utils.Response(w, msg)
	return
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	msg := utils.Message(http.StatusOK, "Successful")
	utils.Response(w, msg)
	return
}
