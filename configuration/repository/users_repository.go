package repository

import (
	"backend-service/entity"
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type UsersRepository struct {
	db *gorm.DB
}

func NewUsersRepository(db *gorm.DB) *UsersRepository {
	return &UsersRepository{
		db: db,
	}
}

func (repo *UsersRepository) Register(ctx context.Context, ent *entity.Users) error {
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Users{}).
		Create(ent); err.Error != nil {
		return errors.Wrap(err.Error, "[UsersRepository-Register]")
	}
	return nil
}

func (repo *UsersRepository) Login(ctx context.Context, email string, password string) (*entity.Users, error) {
	var models *entity.Users
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Users{}).
		Where(`email`, email).
		Where(`password`, password).
		First(&models).
		Error; err != nil {
		return nil, errors.Wrap(err, "[UsersRepository-Login]")
	}
	return models, nil
}
