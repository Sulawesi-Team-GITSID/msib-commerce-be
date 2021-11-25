package repository

import (
	"backend-service/entity"
	"context"

	"github.com/google/uuid"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// GenreRepository connects entity.Genre with database.
type GenreRepository struct {
	db *gorm.DB
}

// NewGenreRepository creates an instance of RoleRepository.
func NewGenreRepository(db *gorm.DB) *GenreRepository {
	return &GenreRepository{
		db: db,
	}
}

// Insert inserts Genre data to database.
func (repo *GenreRepository) Insert(ctx context.Context, ent *entity.Genre) error {
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Genre{}).
		Create(ent).
		Error; err != nil {
		return errors.Wrap(err, "[GenreRepository-Insert]")
	}
	return nil
}

func (repo *GenreRepository) GetListGenre(ctx context.Context, limit, offset string) ([]*entity.Genre, error) {
	var models []*entity.Genre
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Genre{}).Where("deleted", false).
		Find(&models).
		Error; err != nil {
		return nil, errors.Wrap(err, "[GenreRepository-FindAll]")
	}
	return models, nil
}

func (repo *GenreRepository) GetDetailGenre(ctx context.Context, ID uuid.UUID) (*entity.Genre, error) {
	var models *entity.Genre
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Genre{}).
		Take(&models, ID).
		Error; err != nil {
		return nil, errors.Wrap(err, "[GenreRepository-FindById]")
	}
	return models, nil
}

func (repo *GenreRepository) DeleteGenre(ctx context.Context, ID uuid.UUID) error {
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Genre{}).Where("id = ?", ID).
		Update("deleted", true).Error; err != nil {
		return errors.Wrap(err, "[GenreRepository-Delete]")
	}
	return nil
}

func (repo *GenreRepository) UpdateGenre(ctx context.Context, ent *entity.Genre) error {
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Genre{Id: ent.Id}).
		Select("name").
		Updates(ent).Error; err != nil {
		return errors.Wrap(err, "[GenreRepository-Update]")
	}
	return nil
}
