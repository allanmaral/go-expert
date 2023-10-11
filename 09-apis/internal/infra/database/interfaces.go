package database

import "github.com/allanmaral/go-expert/09-apis/internal/entity"

type UserRepository interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}
