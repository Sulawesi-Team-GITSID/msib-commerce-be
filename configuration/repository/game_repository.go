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

func (repo *GameRepository) GetListGame(ctx context.Context, limit, offset string) ([]*entity.Game, error) {
	var models []*entity.Game
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Game{}).
		Find(&models).
		Error; err != nil {
		return nil, errors.Wrap(err, "[GameRepository-FindAll]")
	}
	return models, nil
}

func (repo *GameRepository) GetListGenre(ctx context.Context, limit, offset string) ([]*entity.Game, error) {
	var models []*entity.Game
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

func (repo *GameRepository) GetListTrendGame(ctx context.Context, limit, offset string) ([]*entity.Game, error) {
	var models []*entity.Game
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Game{}).
		Select("game.id", "nama_game", "harga", "review.rating").
		Joins("inner join review on review.game_id = game.id").
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

func (repo *GameRepository) DeleteGame(ctx context.Context, ID uuid.UUID) error {
	if err := repo.db.
		WithContext(ctx).
		Delete(&entity.Game{Id: ID}).Error; err != nil {
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
