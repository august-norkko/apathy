package controller

import (
	"net/http"
	"apathy/utils"
)

func FooHandler(w http.ResponseWriter, r *http.Request) {
	msg := utils.Message(200, "Successful")
	utils.Response(w, msg)
	return
}
