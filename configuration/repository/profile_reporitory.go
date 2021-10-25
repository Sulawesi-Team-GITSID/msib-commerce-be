package repository

import (
	"backend-service/entity"
	"context"

	"github.com/google/uuid"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// ProfileRepository connects entity.Profile with database.
type ProfileRepository struct {
	db *gorm.DB
}

// NewProfileRepository creates an instance of RoleRepository.
func NewProfileRepository(db *gorm.DB) *ProfileRepository {
	return &ProfileRepository{
		db: db,
	}
}

// Insert inserts Profile data to database.
func (repo *ProfileRepository) Insert(ctx context.Context, ent *entity.Profile) error {
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Profile{}).
		Create(ent).
		Error; err != nil {
		return errors.Wrap(err, "[ProfileRepository-Insert]")
	}
	return nil
}

func (repo *ProfileRepository) GetListProfile(ctx context.Context, limit, offset string) ([]*entity.Profile, error) {
	var models []*entity.Profile
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Profile{}).
		Find(&models).
		Error; err != nil {
		return nil, errors.Wrap(err, "[ProfileRepository-FindAll]")
	}
	return models, nil
}

func (repo *ProfileRepository) GetDetailProfile(ctx context.Context, ID uuid.UUID) (*entity.Profile, error) {
	var models *entity.Profile
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Profile{}).
		Take(&models, ID).
		Error; err != nil {
		return nil, errors.Wrap(err, "[ProfileRepository-FindById]")
	}
	return models, nil
}

func (repo *ProfileRepository) DeleteProfile(ctx context.Context, ID uuid.UUID) error {
	if err := repo.db.
		WithContext(ctx).
		Delete(&entity.Profile{Id: ID}).Error; err != nil {
		return errors.Wrap(err, "[ProfileRepository-Delete]")
	}
	return nil
}

func (repo *ProfileRepository) UpdateProfile(ctx context.Context, ent *entity.Profile) error {
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Profile{Id: ent.Id}).
		Select("first_name", "last_name", "phone", "gender", "birthday").
		Updates(ent).Error; err != nil {
		return errors.Wrap(err, "[ProfileRepository-Update]")
	}
	return nil
}
