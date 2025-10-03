package services

import (
	"main/models"
	"main/repositories"

	"golang.org/x/crypto/bcrypt"
)

type IAuthService interface {
	Signup(email string, password string) error
	Login(email string, password string) error
}

type AuthService struct {
	repo repositories.IAuthRepository
}

func NewAuthService(repo repositories.IAuthRepository) IAuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) Signup(email string, password string) error {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user := &models.User{
		Email:    email,
		Password: string(hashPassword),
	}
	return s.repo.CreateUser(user)
}

func (s *AuthService) Login(email string, password string) error {
	foundUser, err := s.repo.FindUser(email)
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(password))
	if err != nil {
		return err
	}
	return nil
}
