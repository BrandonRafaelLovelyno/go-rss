package user

import (
	"fmt"
	"github.com/BrandonRafaelLovelyno/go-rss/internal/database"
	"net/http"
)

type UserHandler struct {
	Query   *database.Queries
	service *UserService
}

func (h *UserHandler) HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}
