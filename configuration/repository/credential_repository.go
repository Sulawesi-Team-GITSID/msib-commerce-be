package repository

import (
	"backend-service/entity"
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// CredentialRepository connects entity.Credential with database.
type CredentialRepository struct {
	db *gorm.DB
}

// NewCredentialRepository creates an instance of RoleRepository.
func NewCredentialRepository(db *gorm.DB) *CredentialRepository {
	return &CredentialRepository{
		db: db,
	}
}

// Insert inserts Credential data to database.
func (repo *CredentialRepository) Insert(ctx context.Context, ent *entity.Credential) error {
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Credential{}).
		Create(ent).
		Error; err != nil {
		return errors.Wrap(err, "[CredentialRepository-Insert]")
	}
	return nil
}