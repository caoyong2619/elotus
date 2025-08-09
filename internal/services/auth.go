package services

import (
	"github.com/caoyong2619/elotus/internal/database"
	"xorm.io/xorm"
)

func NewAuthService(e *xorm.Engine) *AuthService {
	return &AuthService{
		engine: e,
	}
}

type AuthService struct {
	engine *xorm.Engine
}

func (s AuthService) Register(username, password string) error {
	// make a simple example using plain text password
	_, err := s.engine.InsertOne(&database.User{
		Username: username,
		Password: password,
	})

	return err
}
