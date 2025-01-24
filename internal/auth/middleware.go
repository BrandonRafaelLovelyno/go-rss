package auth

import (
	"fmt"
	"net/http"

	"github.com/BrandonRafaelLovelyno/go-rss/internal/database"
	"github.com/BrandonRafaelLovelyno/go-rss/internal/users"
	"github.com/BrandonRafaelLovelyno/go-rss/pkg/utils"
)

type authedHandler func(w http.ResponseWriter, r *http.Request, user database.User)

func Authenticate(handler authedHandler, query database.Queries) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		token, err := extractApiKey(r.Header)
		if err != nil {
			utils.RespondWithError(w, 401, fmt.Sprintf("invalid token: %v", err))
			return
		}

		user, err := users.GetByApiKey(query, r.Context(), token)
		if err != nil {
			utils.RespondWithError(w, 401, fmt.Sprintf("error getting user: %v", err))
		}

		handler(w, r, *user)
	}
}
