package repository

import (
	"backend-service/entity"
	"context"

	"github.com/google/uuid"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// Tags_detailRepository connects entity.Tags_detail with database.
type Tags_detailRepository struct {
	db *gorm.DB
}

// NewTags_detailRepository creates an instance of RoleRepository.
func NewTags_detailRepository(db *gorm.DB) *Tags_detailRepository {
	return &Tags_detailRepository{
		db: db,
	}
}

// Insert inserts Tags_detail data to database.
func (repo *Tags_detailRepository) Insert(ctx context.Context, ent *entity.Tags_detail) error {
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Tags_detail{}).
		Create(ent).
		Error; err != nil {
		return errors.Wrap(err, "[Tags_detailRepository-Insert]")
	}
	return nil
}

func (repo *Tags_detailRepository) GetListTags_detail(ctx context.Context, limit, offset string) ([]*entity.Tags_detail, error) {
	var models []*entity.Tags_detail
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Tags_detail{}).
		Find(&models).
		Error; err != nil {
		return nil, errors.Wrap(err, "[Tags_detailRepository-FindAll]")
	}
	return models, nil
}

func (repo *Tags_detailRepository) GetGameTags(ctx context.Context, ID uuid.UUID) ([]*entity.GameTags, error) {
	var models []*entity.GameTags
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Tags_detail{}).
		Select("game_id, tags.name as tags").
		Joins("inner join tags on tags.id = tags_id").Where("game_id", ID).
		Order("tags.name asc").
		Find(&models).
		Error; err != nil {
		return nil, errors.Wrap(err, "[Tags_detailRepository-FindAll]")
	}
	return models, nil
}

func (repo *Tags_detailRepository) GetDetailTags_detail(ctx context.Context, ID uuid.UUID) (*entity.Tags_detail, error) {
	var models *entity.Tags_detail
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Tags_detail{}).
		Find(&models, ID).
		Error; err != nil {
		return nil, errors.Wrap(err, "[Tags_detailRepository-FindById]")
	}
	return models, nil
}

func (repo *Tags_detailRepository) DeleteTags_detail(ctx context.Context, ID uuid.UUID) error {
	if err := repo.db.
		WithContext(ctx).
		Delete(&entity.Tags_detail{Game_id: ID}).Where("game_id", ID).
		Error; err != nil {
		return errors.Wrap(err, "[Tags_detailRepository-Delete]")
	}
	return nil
}

func (repo *Tags_detailRepository) UpdateTags_detail(ctx context.Context, ent *entity.Tags_detail) error {
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Tags_detail{Game_id: ent.Game_id}).
		Select("tags_id").
		Updates(ent).Error; err != nil {
		return errors.Wrap(err, "[Tags_detailRepository-Update]")
	}
	return nil
}
