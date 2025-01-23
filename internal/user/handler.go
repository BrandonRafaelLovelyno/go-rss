package user

import (
	"fmt"
	"github.com/BrandonRafaelLovelyno/go-rss/internal/database"
	"github.com/BrandonRafaelLovelyno/go-rss/pkg/utils"
	"github.com/google/uuid"
	"net/http"
	"time"
)

type UserHandler struct {
	Query   *database.Queries
	service *UserService
}

func NewUserHandler(query *database.Queries) *UserHandler {
	return &UserHandler{
		Query:   query,
		service: NewUserService(query),
	}
}

func (h *UserHandler) HandleCreateUser(w http.ResponseWriter, r *http.Request) (hr *utils.HandlerReturn) {
	parameter, err := getCreateUserParams(r)
	if err != nil {
		return &utils.HandlerReturn{
			Code:  400,
			Error: fmt.Errorf("error reading parameters: %w", err),
		}
	}

	user, err := h.service.Create(r.Context(), database.CreateUserParams{
		Name:      parameter.Name,
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return &utils.HandlerReturn{
			Code:  500,
			Error: fmt.Errorf("error creating user: %w", err),
		}
	}

	return &utils.HandlerReturn{
		Code:    201,
		Payload: &user,
	}
}

func (h *UserHandler) HandleGetUserByApiKey(w http.ResponseWriter, r *http.Request) (hr *utils.HandlerReturn) {
	parameter, err := getCreateUserParams(r)
	if err != nil {
		return &utils.HandlerReturn{
			Code:  400,
			Error: fmt.Errorf("error reading parameters: %w", err),
		}
	}

	user, err := h.service.Create(r.Context(), database.CreateUserParams{
		Name:      parameter.Name,
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return &utils.HandlerReturn{
			Code:  500,
			Error: fmt.Errorf("error creating user: %w", err),
		}
	}

	return &utils.HandlerReturn{
		Code:    201,
		Payload: &user,
	}
}
