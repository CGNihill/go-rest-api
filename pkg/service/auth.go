package service

import (
	"crypto/sha1"
	"fmt"

	todo "github.com/CGNihill/go-rest-api"
	"github.com/CGNihill/go-rest-api/pkg/repository"
)

const salt string = "80234jljdhfblhjfuyg4trfqugf74g"

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

func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
