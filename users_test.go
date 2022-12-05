package main

import (
	"ll/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Users struct {
	us models.UserService
}

func TestCreateUser(t *testing.T) {
	user := models.User{
		Name:     "Caio",
		Email:    "caio@gmail.com",
		Password: "12345678",
	}

	assert.Equal(t, user.Name, "Caio", "OK response is expected")
}

func (u *Users) TestCreateUserWithSmallPassword(t *testing.T) {
	user := models.User{
		Name:     "Caio",
		Email:    "caio@gmail.com",
		Password: "123456",
	}
	err := u.us.Create(&user)

	assert.Equal(t, err, "models: password must be at least 8 characters long", "OK response is expected")
}

func (u *Users) TestCreateUserPasswordRequired(t *testing.T) {
	user := models.User{
		Name:     "Caio",
		Email:    "caio@gmail.com",
		Password: "",
	}
	err := u.us.Create(&user)

	assert.Equal(t, err, "models: password is required", "OK response is expected")
}

func (u *Users) TestCreateUserEmailFormat(t *testing.T) {
	user := models.User{
		Name:     "Caio",
		Email:    "caiogmailcom",
		Password: "12345678",
	}
	err := u.us.Create(&user)

	assert.Equal(t, err, "models: email address is not valid", "OK response is expected")
}

func (u *Users) TestCreateUserEmailNotAvailable(t *testing.T) {
	user := models.User{
		Name:     "Alex",
		Email:    "caio@gmail.com",
		Password: "12345678",
	}
	_ = u.us.Create(&user)
	err := u.us.Create(&user)

	assert.Equal(t, err, "models: resource not found", "OK response is expected")
}

func (u *Users) TestUpdateUserSmallPassword(t *testing.T) {
	user := models.User{
		Name:     "Caio",
		Email:    "caio@gmail.com",
		Password: "123456",
	}
	err := u.us.Update(&user)

	assert.Equal(t, err, "models: password must be at least 8 characters long", "OK response is expected")
}

func (u *Users) TestUpdateUserRequiredEmail(t *testing.T) {
	user := models.User{
		Name:     "Caio",
		Email:    "",
		Password: "123456",
	}
	err := u.us.Update(&user)

	assert.Equal(t, err, "models: email address is required", "OK response is expected")
}

func (u *Users) TestUpdateUserEmailNotAvailable(t *testing.T) {
	user := models.User{
		Name:     "Alex",
		Email:    "caio@gmail.com",
		Password: "12345678",
	}
	_ = u.us.Update(&user)
	err := u.us.Update(&user)

	assert.Equal(t, err, "models: resource not found", "OK response is expected")
}

func (u *Users) TestUpdateUserEmailFormat(t *testing.T) {
	user := models.User{
		Name:     "Caio",
		Email:    "caiogmailcom",
		Password: "123456",
	}
	err := u.us.Update(&user)

	assert.Equal(t, err, "models: email address is not valid", "OK response is expected")
}
