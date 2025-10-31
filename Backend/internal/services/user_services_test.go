package services_test

import (
	"errors"
	"main/internal/models"
	"main/internal/services"
	"main/internal/testutils"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type MockUserRepository struct {
	mock.Mock
}

func TestMain(m *testing.M) {
	testutils.RunTests(m)
}

func (m *MockUserRepository) CreateUser(user *models.User) error {
	panic("unimplemented")
}

func (m *MockUserRepository) FindUser(email string) (*models.User, error) {
	args := m.Called(email)
	if user, ok := args.Get(0).(*models.User); ok {
		return user, args.Error(1)
	}
	return nil, args.Error(1)
}

func TestFindUser_Success(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := services.NewUserService(mockRepo)

	ReturnMockUser := &models.User{Model: gorm.Model{ID: 9}, Email: "test@example.com"}
	mockRepo.On("FindUser", "test@example.com").Return(ReturnMockUser, nil)

	user, err := service.FindUser("test@example.com")

	assert.NoError(t, err)
	assert.Equal(t, "test@example.com", user.Email)
	mockRepo.AssertExpectations(t)
}

func TestFindUser_Error(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := services.NewUserService(mockRepo)

	mockRepo.On("FindUser", "Error@example.com").Return(nil, errors.New("not found"))

	user, err := service.FindUser("Error@example.com")

	assert.Nil(t, user)
	assert.EqualError(t, err, "not found")
	mockRepo.AssertExpectations(t)
}
