package repository

import (
	"backend-service/entity"
	"context"

	"github.com/google/uuid"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// GameRepository connects entity.Game with database.
type GameRepository struct {
	db *gorm.DB
}

// NewGameRepository creates an instance of RoleRepository.
func NewGameRepository(db *gorm.DB) *GameRepository {
	return &GameRepository{
		db: db,
	}
}

// Insert inserts Game data to database.
func (repo *GameRepository) Insert(ctx context.Context, ent *entity.Game) error {
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Game{}).
		Create(ent).
		Error; err != nil {
		return errors.Wrap(err, "[GameRepository-Insert]")
	}
	return nil
}

func (repo *GameRepository) GetListGame(ctx context.Context, limit, offset string) ([]*entity.ListGame, error) {
	var models []*entity.ListGame
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Game{}).
		Select("game.id", "game.shop_id", "game.nama_game", "game.harga", "genre.name as genre").
		Joins("inner join genre on game.genre_id = genre.id").Where("game.deleted", false).
		Order("game.nama_game desc").
		Find(&models).
		Error; err != nil {
		return nil, errors.Wrap(err, "[GameRepository-FindAll]")
	}
	return models, nil
}

func (repo *GameRepository) GetListGameShop(ctx context.Context, ID uuid.UUID) ([]*entity.GameShop, error) {
	var models []*entity.GameShop
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Game{}).
		Select("game.id", "game.shop_id", "game.nama_game", "game.harga", "shop.name as shop").
		Joins("inner join shop on game.shop_id = shop.id").Where("game.shop_id = '" + ID.String() + "' AND game.deleted = false").
		Order("game.nama_game desc").
		Find(&models).
		Error; err != nil {
		return nil, errors.Wrap(err, "[GameRepository-FindAll]")
	}
	return models, nil
}

func (repo *GameRepository) SortByAsc(ctx context.Context, ID uuid.UUID) ([]*entity.GameShop, error) {
	var models []*entity.GameShop
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Game{}).
		Select("game.id", "game.shop_id", "game.nama_game", "game.harga", "shop.name as shop").
		Joins("inner join shop on game.shop_id = shop.id").Where("game.shop_id = '" + ID.String() + "' AND game.deleted = false").
		Order("game.harga asc").
		Find(&models).
		Error; err != nil {
		return nil, errors.Wrap(err, "[GameRepository-FindAll]")
	}
	return models, nil
}

func (repo *GameRepository) SortByDesc(ctx context.Context, ID uuid.UUID) ([]*entity.GameShop, error) {
	var models []*entity.GameShop
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Game{}).
		Select("game.id", "game.shop_id", "game.nama_game", "game.harga", "shop.name as shop").
		Joins("inner join shop on game.shop_id = shop.id").Where("game.shop_id = '" + ID.String() + "' AND game.deleted = false").
		Order("game.harga desc").
		Find(&models).
		Error; err != nil {
		return nil, errors.Wrap(err, "[GameRepository-FindAll]")
	}
	return models, nil
}

func (repo *GameRepository) GetListGenre(ctx context.Context, limit, offset string) ([]*entity.Genre, error) {
	var models []*entity.Genre
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Game{}).
		Distinct("genre").Order("genre asc").
		Find(&models).
		Error; err != nil {
		return nil, errors.Wrap(err, "[GameRepository-FindGenre]")
	}
	return models, nil
}

func (repo *GameRepository) GetListTrendGame(ctx context.Context, limit, offset string) ([]*entity.TrendGame, error) {
	var models []*entity.TrendGame
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Game{}).
		Select("game.id", "nama_game", "harga", "avg(rating) as rating").Group("game.id").
		Joins("inner join review on review.game_id = game.id").Order("rating desc").
		Find(&models).
		Error; err != nil {
		return nil, errors.Wrap(err, "[GameRepository-FindTrendGame]")
	}
	return models, nil
}

func (repo *GameRepository) GetDetailGame(ctx context.Context, ID uuid.UUID) (*entity.Game, error) {
	var models *entity.Game
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Game{}).
		Take(&models, ID).
		Error; err != nil {
		return nil, errors.Wrap(err, "[GameRepository-FindById]")
	}
	return models, nil
}

func (repo *GameRepository) SearchGame(ctx context.Context, search string) ([]*entity.ListGame, error) {
	var models []*entity.ListGame
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Game{}).
		Select("game.id", "game.shop_id", "game.nama_game", "game.harga", "genre.name as genre").
		Joins("inner join genre on game.genre_id = genre.id").
		Where("lower(nama_game) LIKE lower('%" + search + "%') AND game.deleted = false").Order("game.nama_game desc").
		Find(&models).
		Error; err != nil {
		return nil, errors.Wrap(err, "[GameRepository-FindAll]")
	}
	return models, nil
}

func (repo *GameRepository) DeleteGame(ctx context.Context, ID uuid.UUID) error {
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Game{}).Where("id = ?", ID).
		Update("deleted", true).Error; err != nil {
		return errors.Wrap(err, "[GameRepository-Delete]")
	}
	return nil
}

func (repo *GameRepository) UpdateGame(ctx context.Context, ent *entity.Game) error {
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Game{Id: ent.Id}).
		Select("nama_game", "harga", "genre").
		Updates(ent).Error; err != nil {
		return errors.Wrap(err, "[GameRepository-Update]")
	}
	return nil
}
