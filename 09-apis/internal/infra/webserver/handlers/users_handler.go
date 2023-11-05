package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/allanmaral/go-expert/09-apis/internal/dto"
	"github.com/allanmaral/go-expert/09-apis/internal/entity"
	"github.com/allanmaral/go-expert/09-apis/internal/infra/database"
)

type UsersHandler struct {
	UserRepository database.UserRepository
}

func NewUsersHandler(db database.UserRepository) *UsersHandler {
	return &UsersHandler{
		UserRepository: db,
	}
}

func (h *UsersHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user dto.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u, err := entity.NewUser(user.Name, user.Email, user.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.UserRepository.Create(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
