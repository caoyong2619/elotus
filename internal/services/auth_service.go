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
	claims := &ElotusClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
		},
		ID:       user.Id,
		Username: user.Username,
	}
	t := s.GenerateToken(claims)

	signedToken, err := t.SignedString(s.secret)
	if err != nil {
		return ``, err
	}

	// record the token, but it seems to have no effect in this example :)
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

func (s *AuthService) GenerateToken(claims jwt.Claims) *jwt.Token {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
}

func (s *AuthService) ParseToken(token string) (*jwt.Token, error) {
	claims := &ElotusClaims{}
	return jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return s.secret, nil
	})
}
