package services

import (
	"fmt"
	"time"

	"github.com/caoyong2619/elotus/internal/database"
	"github.com/golang-jwt/jwt/v5"
	"xorm.io/xorm"
)

type ElotusClaims struct {
	jwt.RegisteredClaims
	ID       int64  `json:"id"`
	Username string `json:"username"`
}

func NewAuthService(e *xorm.Engine, secret []byte) *AuthService {
	return &AuthService{
		engine: e,
		secret: secret,
	}
}

type AuthService struct {
	engine *xorm.Engine
	secret []byte
}

func (s *AuthService) Register(username, password string) error {
	// make a simple example using plain text password
	_, err := s.engine.InsertOne(&database.User{
		Username: username,
		Password: password,
	})

	return err
}

// login with username and password
// generate a new jwt token every time
func (s *AuthService) Login(username, password string) (string, error) {
	var user database.User

	has, err := s.engine.Where("username = ?", username).Get(&user)
	if err != nil {
		return ``, err
	}

	if !has {
		return ``, fmt.Errorf("user not found")
	}

	if user.Password != password {
		return ``, fmt.Errorf("password not match")
	}

	exp := time.Now().Add(10 * time.Minute)
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, &ElotusClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(10 * time.Minute)),
		},
		ID:       user.Id,
		Username: user.Username,
	})

	signedToken, err := t.SignedString(s.secret)
	if err != nil {
		return ``, err
	}

	_, err = s.engine.InsertOne(&database.AuthToken{
		Token:     signedToken,
		UserId:    user.Id,
		ExpiredAt: exp.Unix(),
	})
	if err != nil {
		return ``, err
	}

	return signedToken, nil
}
