package repository

import (
	"backend-service/entity"
	"context"

	"github.com/google/uuid"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// ShopRepository connects entity.Shop with database.
type ShopRepository struct {
	db *gorm.DB
}

// NewShopRepository creates an instance of RoleRepository.
func NewShopRepository(db *gorm.DB) *ShopRepository {
	return &ShopRepository{
		db: db,
	}
}

// Insert inserts Shop data to database.
func (repo *ShopRepository) Insert(ctx context.Context, ent *entity.Shop) error {
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Shop{}).
		Create(ent).
		Error; err != nil {
		return errors.Wrap(err, "[ShopRepository-Insert]")
	}
	return nil
}

func (repo *ShopRepository) GetListShop(ctx context.Context, limit, offset string) ([]*entity.Shop, error) {
	var models []*entity.Shop
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Shop{}).
		Find(&models).
		Error; err != nil {
		return nil, errors.Wrap(err, "[ShopRepository-FindAll]")
	}
	return models, nil
}

func (repo *ShopRepository) GetDetailShop(ctx context.Context, ID uuid.UUID) (*entity.Shop, error) {
	var models *entity.Shop
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Shop{}).
		Take(&models, ID).
		Error; err != nil {
		return nil, errors.Wrap(err, "[ShopRepository-FindById]")
	}
	return models, nil
}

func (repo *ShopRepository) DeleteShop(ctx context.Context, ID uuid.UUID) error {
	if err := repo.db.
		WithContext(ctx).
		Delete(&entity.Shop{Id: ID}).Error; err != nil {
		return errors.Wrap(err, "[ShopRepository-Delete]")
	}
	return nil
}

func (repo *ShopRepository) UpdateShop(ctx context.Context, ent *entity.Shop) error {
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Shop{Id: ent.Id}).
		Select("name", "location").
		Updates(ent).Error; err != nil {
		return errors.Wrap(err, "[ShopRepository-Update]")
	}
	return nil
}
