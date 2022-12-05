package main

import (
	"ll/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	user := models.User{
		Name:     "Caio",
		Email:    "caio@gmail.com",
		Password: "12345678",
	}

	assert.Equal(t, user.Name, "Caio", "OK response is expected")
}
