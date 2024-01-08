package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/allanmaral/go-expert/09-apis/internal/dto"
	"github.com/allanmaral/go-expert/09-apis/internal/entity"
	"github.com/allanmaral/go-expert/09-apis/internal/infra/database"
	"github.com/go-chi/jwtauth"
)

type UsersHandler struct {
	UserRepository database.UserRepository
	dummyUser      *entity.User
	JWT            *jwtauth.JWTAuth
	JWTExpiresIn   int
}

func NewUsersHandler(db database.UserRepository, jwt *jwtauth.JWTAuth, jwtexp int) *UsersHandler {
	u, _ := entity.NewUser("dummy", "d@d.com", "password to avoid early return")

	return &UsersHandler{
		UserRepository: db,
		dummyUser:      u,
		JWT:            jwt,
		JWTExpiresIn:   jwtexp,
	}
}

func (h *UsersHandler) Login(w http.ResponseWriter, r *http.Request) {
	var user dto.LoginInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u, err := h.UserRepository.FindByEmail(user.Email)
	if err != nil {
		u = h.dummyUser
	}

	if !u.ValidatePassword(user.Password) || u.ID == h.dummyUser.ID {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	_, token, err := h.JWT.Encode(map[string]interface{}{
		"sub": u.ID.String(),
		"iat": time.Now().Add(time.Second * time.Duration(h.JWTExpiresIn)).Unix(),
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	o := dto.AuthOutput{AccessToken: token}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(o)
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
