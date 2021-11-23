package service

import (
	"backend-service/entity"
	"context"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

var (
	// ErrNilTags_detail occurs when a nil Tags_detail is passed.
	ErrNilTags_detail = errors.New("Tags_detail is nil")
)

// Tags_detailService responsible for any flow related to Tags_detail.
// It also implements Tags_detailService.
type Tags_detailService struct {
	Tags_detailRepo Tags_detailRepository
}

// NewTags_detailService creates an instance of Tags_detailService.
func NewTags_detailService(Tags_detailRepo Tags_detailRepository) *Tags_detailService {
	return &Tags_detailService{
		Tags_detailRepo: Tags_detailRepo,
	}
}

type Tags_detailUseCase interface {
	Create(ctx context.Context, Tags_detail *entity.Tags_detail) error
	GetListTags_detail(ctx context.Context, limit, offset string) ([]*entity.Tags_detail, error)
	GetGameTags(ctx context.Context, ID uuid.UUID) ([]*entity.GameTags, error)
	GetDetailTags_detail(ctx context.Context, ID uuid.UUID) (*entity.Tags_detail, error)
	UpdateTags_detail(ctx context.Context, Tags_detail *entity.Tags_detail) error
	DeleteTags_detail(ctx context.Context, ID uuid.UUID) error
}

type Tags_detailRepository interface {
	Insert(ctx context.Context, Tags_detail *entity.Tags_detail) error
	GetListTags_detail(ctx context.Context, limit, offset string) ([]*entity.Tags_detail, error)
	GetGameTags(ctx context.Context, ID uuid.UUID) ([]*entity.GameTags, error)
	GetDetailTags_detail(ctx context.Context, ID uuid.UUID) (*entity.Tags_detail, error)
	UpdateTags_detail(ctx context.Context, Tags_detail *entity.Tags_detail) error
	DeleteTags_detail(ctx context.Context, ID uuid.UUID) error
}

func (svc Tags_detailService) Create(ctx context.Context, Tags_detail *entity.Tags_detail) error {
	// Checking nil Tags_detail
	if Tags_detail == nil {
		return ErrNilTags_detail
	}

	if err := svc.Tags_detailRepo.Insert(ctx, Tags_detail); err != nil {
		return errors.Wrap(err, "[Tags_detailService-Create]")
	}
	return nil
}

func (svc Tags_detailService) GetListTags_detail(ctx context.Context, limit, offset string) ([]*entity.Tags_detail, error) {
	Tags_detail, err := svc.Tags_detailRepo.GetListTags_detail(ctx, limit, offset)
	if err != nil {
		return nil, errors.Wrap(err, "[Tags_detailService-GetListTags_detail]")
	}
	return Tags_detail, nil
}

func (svc Tags_detailService) GetGameTags(ctx context.Context, ID uuid.UUID) ([]*entity.GameTags, error) {
	Tags_detail, err := svc.Tags_detailRepo.GetGameTags(ctx, ID)
	if err != nil {
		return nil, errors.Wrap(err, "[Tags_detailService-GetGameTags]")
	}
	return Tags_detail, nil
}
func (svc Tags_detailService) GetDetailTags_detail(ctx context.Context, ID uuid.UUID) (*entity.Tags_detail, error) {
	Tags_detail, err := svc.Tags_detailRepo.GetDetailTags_detail(ctx, ID)
	if err != nil {
		return nil, errors.Wrap(err, "[Tags_detailService-GetDetailTags_detail]")
	}
	return Tags_detail, nil
}

func (svc Tags_detailService) DeleteTags_detail(ctx context.Context, ID uuid.UUID) error {
	err := svc.Tags_detailRepo.DeleteTags_detail(ctx, ID)
	if err != nil {
		return errors.Wrap(err, "[Tags_detailService-DeleteTags_detail]")
	}
	return nil
}

func (svc Tags_detailService) UpdateTags_detail(ctx context.Context, Tags_detail *entity.Tags_detail) error {
	// Checking nil Tags_detail
	if Tags_detail == nil {
		return ErrNilTags_detail
	}

	if err := svc.Tags_detailRepo.UpdateTags_detail(ctx, Tags_detail); err != nil {
		return errors.Wrap(err, "[Tags_detailService-UpdateTags_detail]")
	}
	return nil
}
