package utils

import (
	"net/http"
	"encoding/json"
	"apathy/entity"
)

func Decode(r *http.Request) (*entity.User, error) {
	var data entity.User
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
