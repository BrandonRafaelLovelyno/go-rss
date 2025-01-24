package users

import (
	"fmt"
	"github.com/BrandonRafaelLovelyno/go-rss/internal/database"
	"github.com/BrandonRafaelLovelyno/go-rss/pkg/utils"
	"github.com/google/uuid"
	"net/http"
	"time"
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
	}

	user, err := h.Service.create(r.Context(), database.CreateUserParams{
		Name:      parameter.Name,
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		utils.RespondWithError(w, 500, fmt.Sprintf("error creating user: %s", err))
	}

	utils.RespondWithJson(w, 201, user)
}

func (h *UserHandler) HandleGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	utils.RespondWithJson(w, 200, user)
}
