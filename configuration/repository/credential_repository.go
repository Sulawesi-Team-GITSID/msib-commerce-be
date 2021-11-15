package repository

import (
	"backend-service/entity"
	"context"

	"github.com/google/uuid"
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

func (repo *CredentialRepository) GetListCredential(ctx context.Context, limit, offset string) ([]*entity.Credential, error) {
	var models []*entity.Credential
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Credential{}).
		Find(&models).
		Error; err != nil {
		return nil, errors.Wrap(err, "[CredentialRepository-FindAll]")
	}
	return models, nil
}

func (repo *CredentialRepository) Login(ctx context.Context, email string, password string) (*entity.Credential, error) {
	var models *entity.Credential
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Credential{}).
		Where(`email`, email).
		Where(`password`, password).
		First(&models).
		Error; err != nil {
		return nil, errors.Wrap(err, "[UsersRepository-Login]")
	}
	return models, nil
}

func (repo *CredentialRepository) UpdateCredential(ctx context.Context, ent *entity.Credential) error {
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Credential{}).Where("id = ?", ent.Id).
		Update("verified", "true").Error; err != nil {
		return errors.Wrap(err, "[GameRepository-Update]")
	}
	return nil
}

func (repo *CredentialRepository) GetDetailCredential(ctx context.Context, ID uuid.UUID) (*entity.Credential, error) {
	var models *entity.Credential
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Credential{}).
		Take(&models, ID).
		Error; err != nil {
		return nil, errors.Wrap(err, "[CredentialRepository-FindById]")
	}
	return models, nil
}

func (repo *CredentialRepository) ForgotPassword(ctx context.Context, ent *entity.Credential) error {
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Credential{}).
		Where("id = ?", ent.Id).
		Update("password", ent.Password).Error; err != nil {
		return errors.Wrap(err, "[CredentialRepository-ForgotPassword]")
	}
	return nil
}
