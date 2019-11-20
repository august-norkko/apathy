package interfaces

import (
	"net/http"
)

type IUserService interface {
	CreateUser(r *http.Request) (bool, error)
	LoginUser(r *http.Request) (string, error)
}