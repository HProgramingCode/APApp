package controllers_test

import (
	"bytes"
	"errors"
	"main/internal/controllers"
	"main/internal/testutils"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockAuthService struct {
	mock.Mock
}

func TestMain(m *testing.M) {
	testutils.RunTests(m)
}

func (m *MockAuthService) Signup(email string, password string) error {
	args := m.Called(email, password)
	return args.Error(0)
}

func (m *MockAuthService) Login(email string, password string) (*string, error) {
	args := m.Called(email, password)
	if token, ok := args.Get(0).(*string); ok {
		return token, args.Error(1)
	}
	return nil, args.Error(1)
}

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	return gin.Default()
}

func TestSignup_Success(t *testing.T) {
	mockService := new(MockAuthService)
	mockService.On("Signup", "test@example.com", "password123").Return(nil)

	controller := controllers.NewAuthController(mockService)
	router := setupRouter()
	router.POST("/auth/signup", controller.Signup)

	body := []byte(`{"email": "test@example.com", "password" : "password123"}`)
	req, _ := http.NewRequest(http.MethodPost, "/auth/signup", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusCreated, rec.Code)
	mockService.AssertExpectations(t)
}

func TestSignup_EmptyEmail(t *testing.T) {
	mockService := new(MockAuthService)
	mockService.On("Signup", "", "password123").Return(nil)

	controller := controllers.NewAuthController(mockService)
	router := setupRouter()
	router.POST("/auth/signup", controller.Signup)

	body := []byte(`{"email": "", "password" : "password123"}`)
	req, _ := http.NewRequest(http.MethodPost, "/auth/signup", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	mockService.AssertNotCalled(t, "Signup")
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestSignup_EmptyPassword(t *testing.T) {
	mockService := new(MockAuthService)
	mockService.On("Signup", "EmptyPass@example.com", "").Return(nil)

	controller := controllers.NewAuthController(mockService)
	router := setupRouter()
	router.POST("/auth/signup", controller.Signup)

	body := []byte(`{"email": "EmptyPass@example.com", "password" : ""}`)
	req, _ := http.NewRequest(http.MethodPost, "/auth/signup", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	mockService.AssertNotCalled(t, "Signup")
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestSignup_ShortPassword(t *testing.T) {
	mockService := new(MockAuthService)
	mockService.On("Signup", "shortPass@example.com", "123").Return(nil)

	controller := controllers.NewAuthController(mockService)
	router := setupRouter()
	router.POST("/auth/signup", controller.Signup)

	body := []byte(`{"email": "shortPass@example.com", "password" : "123"}`)
	req, _ := http.NewRequest(http.MethodPost, "/auth/signup", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	mockService.AssertNotCalled(t, "Signup")
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestLogin_Success(t *testing.T) {
	mockService := new(MockAuthService)
	controller := controllers.NewAuthController(mockService)
	token := "dummy_token"

	mockService.On("Login", "LoginTest@example.com", "password123").Return(&token, nil)

	router := setupRouter()
	router.POST("/auth/login", controller.Login)

	body := []byte(`{"email" : "LoginTest@example.com", "password" : "password123"}`)
	req, _ := http.NewRequest(http.MethodPost, "/auth/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), "dummy_token")

	//モックに設定した期待値が満たされたかのチェック
	mockService.AssertExpectations(t)
}

func TestLogin_BindError(t *testing.T) {
	mockService := new(MockAuthService)
	controller := controllers.NewAuthController(mockService)

	router := setupRouter()
	router.POST("/auth/login", controller.Login)

	req, _ := http.NewRequest(http.MethodPost, "/auth/login", bytes.NewBuffer([]byte(`{invalied}`)))
	req.Header.Set("Content-Type", "application/Json")

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestLogin_ServiceError(t *testing.T) {
	mockService := new(MockAuthService)
	controller := controllers.NewAuthController(mockService)

	router := setupRouter()
	router.POST("/auth/login", controller.Login)

	mockService.On("Login", "serviceError@example.com", "password123").Return(nil, errors.New("login failed"))

	body := []byte(`{"email": "serviceError@example.com", "password" : "password123"}`)

	req, _ := http.NewRequest(http.MethodPost, "/auth/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/Json")

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}
func TestLogin_EmptyEmail(t *testing.T) {
	mockService := new(MockAuthService)
	controller := controllers.NewAuthController(mockService)
	router := setupRouter()
	router.POST("/auth/login", controller.Login)

	body := []byte(`{"email": "", "password" : "password123"}`)
	req, _ := http.NewRequest(http.MethodPost, "/auth/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestLogin_ShortPasswordError(t *testing.T) {
	mockService := new(MockAuthService)
	controller := controllers.NewAuthController(mockService)

	router := setupRouter()
	router.POST("/auth/login", controller.Login)

	body := []byte(`{"email": "serviceError@example.com", "password" : "123"}`)

	req, _ := http.NewRequest(http.MethodPost, "/auth/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/Json")

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}
