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

// Login godoc
// @Summary		Login and generate a JWT
// @Description	Login and generate a JWT
// @Tags		users
// @Accept		json
// @Produce		json
// @Param		request		body		dto.LoginInput	true	"user credentials"
// @Success		200			{object}	dto.AuthOutput
// @Failure		401			{object}	dto.ErrorOutput
// @Failure		500			{object}	dto.ErrorOutput
// @Router		/auth/login [post]
func (h *UsersHandler) Login(w http.ResponseWriter, r *http.Request) {
	var user dto.LoginInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(dto.ErrorOutput{Message: "Invalid JSON input"})
		return
	}

	u, err := h.UserRepository.FindByEmail(user.Email)
	if err != nil {
		u = h.dummyUser
	}

	if !u.ValidatePassword(user.Password) || u.ID == h.dummyUser.ID {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(dto.ErrorOutput{Message: "Invalid email or password"})
		return
	}

	_, token, err := h.JWT.Encode(map[string]interface{}{
		"sub": u.ID.String(),
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Second * time.Duration(h.JWTExpiresIn)).Unix(),
	})
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(dto.ErrorOutput{Message: "Failed to sign JWT"})
		return
	}

	o := dto.AuthOutput{AccessToken: token}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(o)
}

// Create user godoc
// @Summary		Create user
// @Description Create user
// @Tags		users
// @Accept		json
// @Produce		json
// @Param		request	body		dto.CreateUserInput		true	"user request"
// @Success		201
// @Failure		500		{object}	dto.ErrorOutput
// @Router		/users [post]
func (h *UsersHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user dto.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(dto.ErrorOutput{Message: err.Error()})
		return
	}

	u, err := entity.NewUser(user.Name, user.Email, user.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(dto.ErrorOutput{Message: err.Error()})
		return
	}

	err = h.UserRepository.Create(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(dto.ErrorOutput{Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusCreated)
}
