package repository

import (
	"context"
	"web-server/internal/domain/entity"
	"web-server/prisma/db"
)

type PrismaUserRepository struct {
	client *db.PrismaClient
	ctx    context.Context
}

func NewPrismaUserRepository(client *db.PrismaClient) *PrismaUserRepository {
	return &PrismaUserRepository{
		client: client,
		ctx:    context.Background(),
	}
}

func (r *PrismaUserRepository) Create(user *entity.User) error {
	_, err := r.client.User.CreateOne(
		db.User.Email.Set(user.Email),
		db.User.Username.Set(user.Username),
		db.User.Password.Set(user.Password),
		db.User.Role.Set(user.Role),
		db.User.ID.Set(user.ID),
	).Exec(r.ctx)

	return err
}

func (r *PrismaUserRepository) GetByID(id string) (*entity.User, error) {
	user, err := r.client.User.FindUnique(
		db.User.ID.Equals(id),
	).Exec(r.ctx)

	if err != nil {
		return nil, err
	}

	return &entity.User{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
		Role:     user.Role,
	}, nil
}

func (r *PrismaUserRepository) GetByEmail(email string) (*entity.User, error) {
	user, err := r.client.User.FindUnique(
		db.User.Email.Equals(email),
	).Exec(r.ctx)

	if err != nil {
		return nil, err
	}

	return &entity.User{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
		Role:     user.Role,
	}, nil
}

func (r *PrismaUserRepository) Update(user *entity.User) error {
	_, err := r.client.User.FindUnique(
		db.User.ID.Equals(user.ID),
	).Update(
		db.User.Username.Set(user.Username),
		db.User.Email.Set(user.Email),
		db.User.Password.Set(user.Password),
		db.User.Role.Set(user.Role),
	).Exec(r.ctx)

	return err
}

func (r *PrismaUserRepository) Delete(id string) error {
	_, err := r.client.User.FindUnique(
		db.User.ID.Equals(id),
	).Delete().Exec(r.ctx)

	return err
}

func (r *PrismaUserRepository) List() ([]*entity.User, error) {
	prismaUsers, err := r.client.User.FindMany().Exec(r.ctx)
	if err != nil {
		return nil, err
	}

	users := make([]*entity.User, len(prismaUsers))
	for i, u := range prismaUsers {
		users[i] = &entity.User{
			ID:       u.ID,
			Username: u.Username,
			Email:    u.Email,
			Password: u.Password,
			Role:     u.Role,
		}
	}

	return users, nil
}
