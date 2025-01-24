package feeds

import (
	"errors"
	"net/http"

	"github.com/BrandonRafaelLovelyno/go-rss/pkg/utils"
)

type CreateFeedParam struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func getCreateFeedParams(r *http.Request) (*CreateFeedParam, error) {
	parameter := CreateFeedParam{}
	if err := utils.ReadParams(r, &parameter); err != nil {
		return &CreateFeedParam{}, errors.New("invalid request body")
	}

	return &parameter, nil
}
