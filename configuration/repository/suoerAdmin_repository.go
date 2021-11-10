package repository

import (
	"backend-service/entity"
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type SuperAdminRepository struct {
	db *gorm.DB
}

func NewSuperAdminRepository(db *gorm.DB) *SuperAdminRepository {
	return &SuperAdminRepository{
		db: db,
	}
}

func (repo *SuperAdminRepository) Register(ctx context.Context, ent *entity.SuperAdmin) error {
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.SuperAdmin{}).
		Create(ent); err.Error != nil {
		return errors.Wrap(err.Error, "[SuperAdminRepository-Register]")
	}
	return nil
}

func (repo *SuperAdminRepository) LoginAdmin(ctx context.Context, email string, password string) (*entity.SuperAdmin, error) {
	var models *entity.SuperAdmin
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.SuperAdmin{}).
		Where(`email`, email).
		Where(`password`, password).
		First(&models).
		Error; err != nil {
		return nil, errors.Wrap(err, "[SuperAdminRepository-Login]")
	}
	return models, nil
}
