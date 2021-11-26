package repository

import (
	"backend-service/entity"
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

//ReviewRepository connects entity.Review with database.
type ReviewRepository struct {
	db *gorm.DB
}

// NewReviewRepository creates an instance of RoleRepository.
func NewReviewRepository(db *gorm.DB) *ReviewRepository {
	return &ReviewRepository{
		db: db,
	}
}

// Insert insertsReview data to database.
func (repo *ReviewRepository) Insert(ctx context.Context, ent *entity.Review) error {
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Review{}).
		Create(ent).
		Error; err != nil {
		return errors.Wrap(err, "[ReviewRepository-Insert]")
	}
	return nil
}

func (repo *ReviewRepository) GetListReview(ctx context.Context, limit, offset string) ([]*entity.Review, error) {
	var models []*entity.Review
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Review{}).
		Find(&models).
		Error; err != nil {
		return nil, errors.Wrap(err, "[ReviewRepository-FindAll]")
	}
	return models, nil
}
