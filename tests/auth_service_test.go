package tests

import (
	"fmt"
	"testing"

	"github.com/caoyong2619/elotus/internal/database"
	"github.com/caoyong2619/elotus/internal/services"
)

func TestAuthServiceRegister(t *testing.T) {
	err := authService.Register(testUsername, testPassword)

	if err != nil {
		t.Fatal(err)
	}

	var user database.User
	var count int64
	count, err = database.Engine.Where("username = ? and password = ?", "test", "123456").Count(&user)
	if err != nil {
		t.Fatal(err)
	}

	if count != 1 {
		t.Fatal(fmt.Errorf("user not registered"))
	}
}

func TestAuthServiceLogin(t *testing.T) {
	token, err := authService.Login(testUsername, testPassword)

	if err != nil {
		t.Fatal(err)
	}

	if token == `` {
		t.Fatal(fmt.Errorf("token is empty"))
	}

	// check the token
	parsed, err := authService.ParseToken(token)
	if err != nil {
		t.Fatal(err)
	}

	claims := parsed.Claims.(*services.ElotusClaims)
	if claims.Username != testUsername {
		t.Fatal(fmt.Errorf("username not match"))
	}
}
