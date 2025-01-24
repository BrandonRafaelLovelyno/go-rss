package users

import (
	"errors"
	"net/http"

	"github.com/BrandonRafaelLovelyno/go-rss/pkg/utils"
)

type CreateUserParams struct {
	Name string `json:"name"`
}

func getCreateUserParams(r *http.Request) (*CreateUserParams, error) {
	parameter := CreateUserParams{}
	if err := utils.ReadParams(r, &parameter); err != nil {
		return &CreateUserParams{}, errors.New("invalid request body")
	}

	return &parameter, nil
}
