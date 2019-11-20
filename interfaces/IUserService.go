package interfaces

import (
	"net/http"
)

type IUserService interface {
	CreateUser(r *http.Request) (int, string, error)
	LoginUser(r *http.Request) (string, error)
}