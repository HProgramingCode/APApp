package services

import (
	"main/models"
	"main/repositories"
)

type IAuthService interface {
	Signup(email string, password string) error
}

type AuthService struct {
	repo repositories.IAuthRepository
}

func NewAuthService(repo repositories.IAuthRepository) IAuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) Signup(email string, password string) error {
	user := &models.User{
		Email:    email,
		Password: password,
	}
	return s.repo.CreateUser(user)
}
