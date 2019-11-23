package response

import (
	"net/http"
	"encoding/json"
	"apathy/models"
)

func Decode(r *http.Request) (*models.User, error) { // move from response package
	var data models.User
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
