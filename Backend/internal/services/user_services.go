package services

import (
	"main/internal/models"
	"main/internal/repositories"
)

type IUserService interface {
	FindUser(email string) (*models.User, error)
}

type UserService struct {
	repo repositories.IUserRepository
}

func NewUserService(repo repositories.IUserRepository) IUserService {
	return &UserService{repo: repo}
}

func (s *UserService) FindUser(email string) (*models.User, error) {
	return s.repo.FindUser(email)
}
