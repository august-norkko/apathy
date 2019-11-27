package interfaces

import (
	"net/http"
)

type IAccountService interface {
	CreateAccount(r *http.Request) (bool, error)
	LoginAccount(r *http.Request) (string, error)
}