package service

import (
	"backend-service/entity"
	"context"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

var (
	// ErrNilGame occurs when a nil Game is passed.
	ErrNilGame = errors.New("Game is nil")
)

// GameService responsible for any flow related to Game.
// It also implements GameService.
type GameService struct {
	GameRepo GameRepository
}

// NewGameService creates an instance of GameService.
func NewGameService(GameRepo GameRepository) *GameService {
	return &GameService{
		GameRepo: GameRepo,
	}
}

type GameUseCase interface {
	Create(ctx context.Context, Game *entity.Game) error
	GetListGame(ctx context.Context, limit, offset string) ([]*entity.Game, error)
	GetDetailGame(ctx context.Context, ID uuid.UUID) (*entity.Game, error)
	UpdateGame(ctx context.Context, Game *entity.Game) error
	DeleteGame(ctx context.Context, ID uuid.UUID) error
}

type GameRepository interface {
	Insert(ctx context.Context, Game *entity.Game) error
	GetListGame(ctx context.Context, limit, offset string) ([]*entity.Game, error)
	GetDetailGame(ctx context.Context, ID uuid.UUID) (*entity.Game, error)
	UpdateGame(ctx context.Context, Game *entity.Game) error
	DeleteGame(ctx context.Context, ID uuid.UUID) error
}

func (svc GameService) Create(ctx context.Context, Game *entity.Game) error {
	// Checking nil Game
	if Game == nil {
		return ErrNilGame
	}

	// Generate id if nil
	if Game.Id == uuid.Nil {
		Game.Id = uuid.New()
	}

	if err := svc.GameRepo.Insert(ctx, Game); err != nil {
		return errors.Wrap(err, "[GameService-Create]")
	}
	return nil
}

func (svc GameService) GetListGame(ctx context.Context, limit, offset string) ([]*entity.Game, error) {
	Game, err := svc.GameRepo.GetListGame(ctx, limit, offset)
	if err != nil {
		return nil, errors.Wrap(err, "[GameService-Create]")
	}
	return Game, nil
}

func (svc GameService) GetDetailGame(ctx context.Context, ID uuid.UUID) (*entity.Game, error) {
	Game, err := svc.GameRepo.GetDetailGame(ctx, ID)
	if err != nil {
		return nil, errors.Wrap(err, "[GameService-Create]")
	}
	return Game, nil
}

func (svc GameService) DeleteGame(ctx context.Context, ID uuid.UUID) error {
	err := svc.GameRepo.DeleteGame(ctx, ID)
	if err != nil {
		return errors.Wrap(err, "[GameService-Create]")
	}
	return nil
}

func (svc GameService) UpdateGame(ctx context.Context, Game *entity.Game) error {
	// Checking nil Game
	if Game == nil {
		return ErrNilGame
	}

	// Generate id if nil
	if Game.Id == uuid.Nil {
		Game.Id = uuid.New()
	}

	if err := svc.GameRepo.UpdateGame(ctx, Game); err != nil {
		return errors.Wrap(err, "[GameService-Create]")
	}
	return nil
}