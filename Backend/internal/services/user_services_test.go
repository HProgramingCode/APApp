package services_test

import (
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

	mockUser := &models.User{Model: gorm.Model{ID: 9}, Email: "test@example.com"}
	mockRepo.On("FindUser", "test@example.com").Return(mockUser, nil)

	user, err := service.FindUser("test@example.com")

	assert.NoError(t, err)
	assert.Equal(t, "test@example.com", user.Email)
	mockRepo.AssertExpectations(t)
}
