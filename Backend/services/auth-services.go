package services

import (
	"main/models"
	"main/repositories"
	"main/utils"
)

type IAuthService interface {
	Signup(email string, password string) error
	Login(email string, password string) (*string, error)
}

type AuthService struct {
	repo repositories.IAuthRepository
}

func NewAuthService(repo repositories.IAuthRepository) IAuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) Signup(email string, password string) error {
	hashPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}
	user := &models.User{
		Email:    email,
		Password: string(hashPassword),
	}
	return s.repo.CreateUser(user)
}

func (s *AuthService) Login(email string, password string) (*string, error) {
	foundUser, err := s.repo.FindUser(email)
	if err != nil {
		return nil, err
	}

	err = utils.CheckPassword(foundUser.Password, password)
	if err != nil {
		return nil, err
	}

	token, err := utils.GenerateToken(foundUser.ID, foundUser.Email)
	if err != nil {
		return nil, err
	}
	return token, err
}
