package repository

import (
	"sync"

	"web-server/internal/domain/constants"
	"web-server/internal/domain/entity"
)

type InMemoryUserRepository struct {
	users map[string]*entity.User
	mutex sync.RWMutex
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: make(map[string]*entity.User),
	}
}

func (r *InMemoryUserRepository) Create(user *entity.User) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.users[user.ID]; exists {
		return constants.ErrUserAlreadyExists
	}

	r.users[user.ID] = user
	return nil
}

func (r *InMemoryUserRepository) GetByID(id string) (*entity.User, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	user, exists := r.users[id]
	if !exists {
		return nil, constants.ErrUserNotFound
	}

	return user, nil
}

func (r *InMemoryUserRepository) Update(user *entity.User) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.users[user.ID]; !exists {
		return constants.ErrUserNotFound
	}

	r.users[user.ID] = user
	return nil
}

func (r *InMemoryUserRepository) Delete(id string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.users[id]; !exists {
		return constants.ErrUserNotFound
	}

	delete(r.users, id)
	return nil
}

func (r *InMemoryUserRepository) List() ([]*entity.User, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	users := make([]*entity.User, 0, len(r.users))
	for _, user := range r.users {
		users = append(users, user)
	}

	return users, nil
}

func (r *InMemoryUserRepository) GetByEmail(email string) (*entity.User, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	for _, user := range r.users {
		if user.Email == email {
			return user, nil
		}
	}

	return nil, constants.ErrUserNotFound
}
