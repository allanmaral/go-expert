package database

import (
	"testing"

	"github.com/allanmaral/go-expert/09-apis/internal/entity"
	generators "github.com/allanmaral/go-expert/09-apis/test"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func makeSUT(t *testing.T) (UserRepository, *gorm.DB) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.User{})
	userRepository := NewGormUserRepository(db)

	return userRepository, db
}

func makeUser() *entity.User {
	name := generators.RandonName()
	email := generators.RandonEmail()
	password := generators.RandonString(16)
	user, _ := entity.NewUser(name, email, password)

	return user
}

func TestCreateUser(t *testing.T) {
	user := makeUser()
	sut, db := makeSUT(t)

	err := sut.Create(user)

	assert.Nil(t, err)
	var userFound entity.User
	err = db.First(&userFound, "id = ?", user.ID).Error
	assert.Nil(t, err)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
	assert.Equal(t, user.Password, userFound.Password)
}

func TestFindByEmail(t *testing.T) {
	email := "some_specific@emai.com"
	user := makeUser()
	user.Email = email
	sut, db := makeSUT(t)
	db.Create(&user)

	foundUser, err := sut.FindByEmail(email)

	assert.Nil(t, err)
	assert.Equal(t, user.ID, foundUser.ID)
	assert.Equal(t, user.Name, foundUser.Name)
	assert.Equal(t, user.Email, foundUser.Email)
	assert.Equal(t, user.Password, foundUser.Password)
}
