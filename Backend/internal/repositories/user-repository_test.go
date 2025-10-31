package repositories_test

import (
	"main/internal/models"
	"main/internal/repositories"
	"main/internal/testutils"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestMain(m *testing.M) {
	testutils.RunTests(m)
}

func SetupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to test database")
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		t.Fatalf("Failed to migrate schema")
	}
	return db
}

func TestFindUser_Success(t *testing.T) {
	db := SetupTestDB(t)
	repo := repositories.NewUserRepository(db)

	//sample
	db.Create(&models.User{Email: "find@example.com", Password: "password123"})
	user, err := repo.FindUser("find@example.com")

	assert.NoError(t, err)
	assert.Equal(t, "find@example.com", user.Email)
}

func TestFindUser_NotFound(t *testing.T) {
	db := SetupTestDB(t)
	repo := repositories.NewUserRepository(db)

	user, err := repo.FindUser("NotFound@example.com")

	assert.Nil(t, user)
	assert.Error(t, err)
	assert.Equal(t, gorm.ErrRecordNotFound, err)
}
