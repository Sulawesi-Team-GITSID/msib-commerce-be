package repository

import (
	"backend-service/entity"
	"context"

	"github.com/google/uuid"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// TagsRepository connects entity.Tags with database.
type TagsRepository struct {
	db *gorm.DB
}

// NewTagsRepository creates an instance of RoleRepository.
func NewTagsRepository(db *gorm.DB) *TagsRepository {
	return &TagsRepository{
		db: db,
	}
}

// Insert inserts Tags data to database.
func (repo *TagsRepository) Insert(ctx context.Context, ent *entity.Tags) error {
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Tags{}).
		Create(ent).
		Error; err != nil {
		return errors.Wrap(err, "[TagsRepository-Insert]")
	}
	return nil
}

func (repo *TagsRepository) GetListTags(ctx context.Context, limit, offset string) ([]*entity.Tags, error) {
	var models []*entity.Tags
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Tags{}).
		Find(&models).
		Error; err != nil {
		return nil, errors.Wrap(err, "[TagsRepository-FindAll]")
	}
	return models, nil
}

func (repo *TagsRepository) GetDetailTags(ctx context.Context, ID uuid.UUID) (*entity.Tags, error) {
	var models *entity.Tags
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Tags{}).
		Take(&models, ID).
		Error; err != nil {
		return nil, errors.Wrap(err, "[TagsRepository-FindById]")
	}
	return models, nil
}

func (repo *TagsRepository) DeleteTags(ctx context.Context, ID uuid.UUID) error {
	if err := repo.db.
		WithContext(ctx).
		Delete(&entity.Tags{Id: ID}).Error; err != nil {
		return errors.Wrap(err, "[TagsRepository-Delete]")
	}
	return nil
}

func (repo *TagsRepository) UpdateTags(ctx context.Context, ent *entity.Tags) error {
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Tags{Id: ent.Id}).
		Select("name").
		Updates(ent).Error; err != nil {
		return errors.Wrap(err, "[TagsRepository-Update]")
	}
	return nil
}
