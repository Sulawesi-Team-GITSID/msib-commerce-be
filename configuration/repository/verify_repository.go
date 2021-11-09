package repository

import (
	"backend-service/entity"
	"context"

	"github.com/google/uuid"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// VerificationRepository connects entity.Verification with database.
type VerificationRepository struct {
	db *gorm.DB
}

// NewVerificationRepository creates an instance of RoleRepository.
func NewVerificationRepository(db *gorm.DB) *VerificationRepository {
	return &VerificationRepository{
		db: db,
	}
}

// Insert inserts Verification data to database.
func (repo *VerificationRepository) Insert(ctx context.Context, ent *entity.Verification) error {
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Verification{}).
		Create(ent).
		Error; err != nil {
		return errors.Wrap(err, "[VerificationRepository-Insert]")
	}
	return nil
}

func (repo *VerificationRepository) GetListVerification(ctx context.Context, limit, offset string) ([]*entity.Verification, error) {
	var models []*entity.Verification
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Verification{}).
		Find(&models).
		Error; err != nil {
		return nil, errors.Wrap(err, "[VerificationRepository-FindAll]")
	}
	return models, nil
}

func (repo *VerificationRepository) GetDetailVerification(ctx context.Context, ID uuid.UUID) (*entity.Verification, error) {
	var models *entity.Verification
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Verification{}).
		Take(&models, ID).
		Error; err != nil {
		return nil, errors.Wrap(err, "[VerificationRepository-FindById]")
	}
	return models, nil
}

func (repo *VerificationRepository) DeleteVerification(ctx context.Context, ID uuid.UUID) error {
	if err := repo.db.
		WithContext(ctx).
		Delete(&entity.Verification{Id: ID}).Error; err != nil {
		return errors.Wrap(err, "[VerificationRepository-Delete]")
	}
	return nil
}

func (repo *VerificationRepository) UpdateVerification(ctx context.Context, ent *entity.Verification) error {
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Verification{Id: ent.Id}).
		Select("credential_id", "code", "expiresat").
		Updates(ent).Error; err != nil {
		return errors.Wrap(err, "[VerificationRepository-Update]")
	}
	return nil
}
