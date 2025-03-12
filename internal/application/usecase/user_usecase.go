package usecase

import (
	"web-server/internal/domain/entity"
	"web-server/internal/domain/repository"
)

type UserUseCase struct {
	userRepo repository.UserRepository
}

func NewUserUseCase(repo repository.UserRepository) *UserUseCase {
	return &UserUseCase{
		userRepo: repo,
	}
}

func (uc *UserUseCase) CreateUser(user *entity.User) error {
	return uc.userRepo.Create(user)
}

func (uc *UserUseCase) GetUser(id string) (*entity.User, error) {
	return uc.userRepo.GetByID(id)
}

func (uc *UserUseCase) UpdateUser(user *entity.User) error {
	return uc.userRepo.Update(user)
}

func (uc *UserUseCase) DeleteUser(id string) error {
return uc.userRepo.Delete(id)
}

func (uc *UserUseCase) ListUsers() ([]*entity.User, error) {
return uc.userRepo.List()
}

func (uc *UserUseCase) GetUserByEmail(email string) (*entity.User, error) {
return uc.userRepo.GetByEmail(email)
}