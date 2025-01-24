package users

import (
	"fmt"
	"github.com/BrandonRafaelLovelyno/go-rss/internal/database"
	"github.com/BrandonRafaelLovelyno/go-rss/pkg/utils"
	"net/http"
)

type UserHandler struct {
	Service *UserService
}

func NewUserHandler(query *database.Queries) *UserHandler {
	return &UserHandler{
		Service: newUserService(query),
	}
}

func (h *UserHandler) HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	parameter, err := getCreateUserParams(r)
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("error reading request: %s", err))
		return
	}

	user, err := h.Service.create(r.Context(), parameter.Name)
	if err != nil {
		utils.RespondWithError(w, 500, fmt.Sprintf("error creating user: %s", err))
		return
	}

	utils.RespondWithJson(w, 201, user)
}

func (h *UserHandler) HandleGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	utils.RespondWithJson(w, 200, user)
}
