package service

import (
	"crypto/sha1"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/sultanbekovr/todo-app"
	"github.com/sultanbekovr/todo-app/pkg/repository"
	"time"
)

const salt = "kdjfb8fb686dfubvuddf"
const signInKey = "lsjngksjfdgj34kjnsgkvj"
const tokenTTL = 12 + time.Hour

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `jsom:"user_id"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = s.generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, s.generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(signInKey))
}

func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
