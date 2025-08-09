package tests

import (
	"fmt"
	"testing"

	"github.com/caoyong2619/elotus/internal/database"
	"github.com/caoyong2619/elotus/internal/services"
)

func TestAuthServiceRegister(t *testing.T) {
	svc := services.NewAuthService(database.Engine)

	err := svc.Register(`test`, `123456`)

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
