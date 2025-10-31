package controllers_test

import (
	"errors"
	"main/internal/controllers"
	"main/internal/middleware"
	"main/internal/models"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
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

func TestMain(m *testing.M) {
	//ダミーロガーの設定
	middleware.Log, _ = zap.NewDevelopment()

	code := m.Run()
	os.Exit(code)
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

func TestDetUserInfo_Unauthorized(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockService := new(MockUserService)
	controller := controllers.NewUserController(mockService)

	//return mock user
	// ↓emailをセットしない（これがUnauthorizedの原因）
	// mockService.On("FindUser", "").Return(nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	controller.GetUserInfo(c)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Contains(t, w.Body.String(), "Unauthorized")
}

func TestDetUserInfo_ServiceError(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockService := new(MockUserService)
	controller := controllers.NewUserController(mockService)

	mockService.On("FindUser", "error@test.com").Return(nil, errors.New("DB Error"))

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("email", "error@test.com")

	controller.GetUserInfo(c)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), "DB Error")
}
