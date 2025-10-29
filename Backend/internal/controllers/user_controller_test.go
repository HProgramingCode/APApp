package controllers_test

import (
	"main/internal/controllers"
	"main/internal/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

// MockUserService
type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) FindUser(email string) (*models.User, error) {
	args := m.Called(email)
	if user, ok := args.Get(0).(*models.User); ok {
		return user, args.Error(1)
	}
	return nil, args.Error(1)
}

// TestGetUserInfo

func TestGetUserInfo_success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockService := new(MockUserService)
	controller := controllers.NewUserController(mockService)

	//return mock user
	mockService.On("FindUser", "test@example.com").Return(&models.User{Model: gorm.Model{ID: 1}, Email: "test@example.com"}, nil)

	// Gin Context
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("email", "test@example.com")

	controller.GetUserInfo(c)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"email":"test@example.com"`)
}
