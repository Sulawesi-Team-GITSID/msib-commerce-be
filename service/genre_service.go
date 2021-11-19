package service

import (
	"backend-service/entity"
	"context"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

var (
	// ErrNilGenre occurs when a nil Genre is passed.
	ErrNilGenre = errors.New("Genre is nil")
)

// GenreService responsible for any flow related to Genre.
// It also implements GenreService.
type GenreService struct {
	GenreRepo GenreRepository
}

// NewGenreService creates an instance of GenreService.
func NewGenreService(GenreRepo GenreRepository) *GenreService {
	return &GenreService{
		GenreRepo: GenreRepo,
	}
}

type GenreUseCase interface {
	Create(ctx context.Context, Genre *entity.Genre) error
	GetListGenre(ctx context.Context, limit, offset string) ([]*entity.Genre, error)
	GetDetailGenre(ctx context.Context, ID uuid.UUID) (*entity.Genre, error)
	UpdateGenre(ctx context.Context, Genre *entity.Genre) error
	DeleteGenre(ctx context.Context, ID uuid.UUID) error
}

type GenreRepository interface {
	Insert(ctx context.Context, Genre *entity.Genre) error
	GetListGenre(ctx context.Context, limit, offset string) ([]*entity.Genre, error)
	GetDetailGenre(ctx context.Context, ID uuid.UUID) (*entity.Genre, error)
	UpdateGenre(ctx context.Context, Genre *entity.Genre) error
	DeleteGenre(ctx context.Context, ID uuid.UUID) error
}

func (svc GenreService) Create(ctx context.Context, Genre *entity.Genre) error {
	// Checking nil Genre
	if Genre == nil {
		return ErrNilGenre
	}

	// Generate id if nil
	if Genre.Id == uuid.Nil {
		Genre.Id = uuid.New()
	}

	if err := svc.GenreRepo.Insert(ctx, Genre); err != nil {
		return errors.Wrap(err, "[GenreService-Create]")
	}
	return nil
}

func (svc GenreService) GetListGenre(ctx context.Context, limit, offset string) ([]*entity.Genre, error) {
	Genre, err := svc.GenreRepo.GetListGenre(ctx, limit, offset)
	if err != nil {
		return nil, errors.Wrap(err, "[GenreService-GetListGenre]")
	}
	return Genre, nil
}

func (svc GenreService) GetDetailGenre(ctx context.Context, ID uuid.UUID) (*entity.Genre, error) {
	Genre, err := svc.GenreRepo.GetDetailGenre(ctx, ID)
	if err != nil {
		return nil, errors.Wrap(err, "[GenreService-GetDetailGenre]")
	}
	return Genre, nil
}

func (svc GenreService) DeleteGenre(ctx context.Context, ID uuid.UUID) error {
	err := svc.GenreRepo.DeleteGenre(ctx, ID)
	if err != nil {
		return errors.Wrap(err, "[GenreService-DeleteGenre]")
	}
	return nil
}

func (svc GenreService) UpdateGenre(ctx context.Context, Genre *entity.Genre) error {
	// Checking nil Genre
	if Genre == nil {
		return ErrNilGenre
	}

	// Generate id if nil
	if Genre.Id == uuid.Nil {
		Genre.Id = uuid.New()
	}

	if err := svc.GenreRepo.UpdateGenre(ctx, Genre); err != nil {
		return errors.Wrap(err, "[GenreService-UpdateGenre]")
	}
	return nil
}
