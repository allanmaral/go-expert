package entity

import (
	"testing"

	generators "github.com/allanmaral/go-expert/09-apis/test"
	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	name := generators.RandonName()
	email := generators.RandonEmail()
	password := generators.RandonString(12)

	user, err := NewUser(name, email, password)

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, name, user.Name)
	assert.Equal(t, email, user.Email)
}

func TestUserValidatePasswordReturnTrueOnSamePassword(t *testing.T) {
	tests := []string{
		"sample",
		"@5uPe#5Tr0ngP^a$s",
		"1234567",
		"a",
	}

	for _, input := range tests {
		sut, _ := NewUser(generators.RandonName(), generators.RandonEmail(), input)

		valid := sut.ValidatePassword(input)

		assert.True(t, valid)
	}
}

func TestUserValidatePasswordReturnFalseOnDifferentPassword(t *testing.T) {
	tests := []struct {
		password   string
		wrongInput string
	}{
		{"sample", "ample"},
		{"sample", "ssample"},
		{"sample", "Sample"},
		{"@5uPe#5Tr0ngP^a$s", ""},
		{"@5uPe#5Tr0ngP^a$s", "123"},
		{"@5uPe#5Tr0ngP^a$s", "@5uPe#5Tr0ngP^a$x"},
		{"@5uPe#5Tr0ngP^a$s", "@5uPe#5ngP^a$s"},
	}

	for _, test := range tests {
		sut, _ := NewUser(generators.RandonName(), generators.RandonEmail(), test.password)

		valid := sut.ValidatePassword(test.wrongInput)

		assert.False(t, valid)
	}
}
