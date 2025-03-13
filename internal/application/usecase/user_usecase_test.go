package usecase

import (
	"testing"

	"web-server/internal/domain/entity"
	"web-server/internal/infrastructure/repository"

	"github.com/stretchr/testify/assert"
)

func TestUserUseCase(t *testing.T) {
	repo := repository.NewInMemoryUserRepository()
	useCase := NewUserUseCase(repo)

	t.Run("Create User", func(t *testing.T) {
		user := &entity.User{
			ID:       "1",
			Email:    "test@example.com",
			Password: "password123",
		}

		err := useCase.CreateUser(user)
		assert.NoError(t, err)

		// Verify user was created
		created, err := useCase.GetUser("1")
		assert.NoError(t, err)
		assert.Equal(t, user.Email, created.Email)
	})

	t.Run("Get User", func(t *testing.T) {
		// Test getting existing user
		user, err := useCase.GetUser("1")
		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, "test@example.com", user.Email)

		// Test getting non-existent user
		user, err = useCase.GetUser("999")
		assert.Error(t, err)
		assert.Nil(t, user)
	})

	t.Run("Update User", func(t *testing.T) {
		user := &entity.User{
			ID:       "1",
			Email:    "updated@example.com",
			Password: "newpassword",
		}

		err := useCase.UpdateUser(user)
		assert.NoError(t, err)

		// Verify update
		updated, err := useCase.GetUser("1")
		assert.NoError(t, err)
		assert.Equal(t, "updated@example.com", updated.Email)
	})

	t.Run("Delete User", func(t *testing.T) {
		err := useCase.DeleteUser("1")
		assert.NoError(t, err)

		// Verify deletion
		user, err := useCase.GetUser("1")
		assert.Error(t, err)
		assert.Nil(t, user)
	})

	t.Run("List Users", func(t *testing.T) {
		// Create some test users
		users := []*entity.User{
			{ID: "2", Email: "user2@example.com", Password: "pass2"},
			{ID: "3", Email: "user3@example.com", Password: "pass3"},
		}
		for _, u := range users {
			err := useCase.CreateUser(u)
			assert.NoError(t, err)
		}

		// Test listing
		list, err := useCase.ListUsers()
		assert.NoError(t, err)
		assert.Len(t, list, 2)
	})

	t.Run("Get User By Email", func(t *testing.T) {
		// Test getting existing user
		user, err := useCase.GetUserByEmail("user2@example.com")
		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, "2", user.ID)

		// Test getting non-existent user
		user, err = useCase.GetUserByEmail("nonexistent@example.com")
		assert.Error(t, err)
		assert.Nil(t, user)
	})
}
