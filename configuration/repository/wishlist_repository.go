package repository

import (
	"backend-service/entity"
	"context"

	"github.com/google/uuid"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// WishlistRepository connects entity.Wishlist with database.
type WishlistRepository struct {
	db *gorm.DB
}

// NewWishlistRepository creates an instance of RoleRepository.
func NewWishlistRepository(db *gorm.DB) *WishlistRepository {
	return &WishlistRepository{
		db: db,
	}
}

// Insert inserts Wishlist data to database.
func (repo *WishlistRepository) Insert(ctx context.Context, ent *entity.Wishlist) error {
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Wishlist{}).
		Create(ent).
		Error; err != nil {
		return errors.Wrap(err, "[WishlistRepository-Insert]")
	}
	return nil
}

func (repo *WishlistRepository) GetListWishlist(ctx context.Context, limit, offset string) ([]*entity.Wishlist, error) {
	var models []*entity.Wishlist
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Wishlist{}).
		Find(&models).
		Error; err != nil {
		return nil, errors.Wrap(err, "[WishlistRepository-FindAll]")
	}
	return models, nil
}

func (repo *WishlistRepository) GetGame(ctx context.Context, ID uuid.UUID) ([]*entity.WishlistGame, error) {
	var models []*entity.WishlistGame
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Wishlist{}).
		Select("credential_id, game_id, game.nama_game, game.image_url, game.harga").
		Joins("inner join game on game.id = game_id").Where("credential_id", ID).
		Order("game.nama_game asc").
		Find(&models).
		Error; err != nil {
		return nil, errors.Wrap(err, "[WishlistRepository-FindAll]")
	}
	return models, nil
}

func (repo *WishlistRepository) GetDetailWishlist(ctx context.Context, ID uuid.UUID) (*entity.Wishlist, error) {
	var models *entity.Wishlist
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Wishlist{}).
		Find(&models, ID).
		Error; err != nil {
		return nil, errors.Wrap(err, "[WishlistRepository-FindById]")
	}
	return models, nil
}

func (repo *WishlistRepository) DeleteWishlist(ctx context.Context, credential_id uuid.UUID, game string) error {
	if err := repo.db.
		WithContext(ctx).Where("credential_id = '" + credential_id.String() + "' AND game_id = '" + game + "'").
		Delete(&entity.Wishlist{}).
		Error; err != nil {
		return errors.Wrap(err, "[WishlistRepository-Delete]")
	}
	return nil
}

func (repo *WishlistRepository) UpdateWishlist(ctx context.Context, ent *entity.Wishlist) error {
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Wishlist{Game_id: ent.Game_id}).
		Select("game_id").
		Updates(ent).Error; err != nil {
		return errors.Wrap(err, "[WishlistRepository-Update]")
	}
	return nil
}
