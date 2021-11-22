package service

import (
	"backend-service/entity"
	"context"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

var (
	// ErrNilTags occurs when a nil Tags is passed.
	ErrNilTags = errors.New("Tags is nil")
)

// TagsService responsible for any flow related to Tags.
// It also implements TagsService.
type TagsService struct {
	TagsRepo TagsRepository
}

// NewTagsService creates an instance of TagsService.
func NewTagsService(TagsRepo TagsRepository) *TagsService {
	return &TagsService{
		TagsRepo: TagsRepo,
	}
}

type TagsUseCase interface {
	Create(ctx context.Context, Tags *entity.Tags) error
	GetListTags(ctx context.Context, limit, offset string) ([]*entity.Tags, error)
	GetDetailTags(ctx context.Context, ID uuid.UUID) (*entity.Tags, error)
	UpdateTags(ctx context.Context, Tags *entity.Tags) error
	DeleteTags(ctx context.Context, ID uuid.UUID) error
}

type TagsRepository interface {
	Insert(ctx context.Context, Tags *entity.Tags) error
	GetListTags(ctx context.Context, limit, offset string) ([]*entity.Tags, error)
	GetDetailTags(ctx context.Context, ID uuid.UUID) (*entity.Tags, error)
	UpdateTags(ctx context.Context, Tags *entity.Tags) error
	DeleteTags(ctx context.Context, ID uuid.UUID) error
}

func (svc TagsService) Create(ctx context.Context, Tags *entity.Tags) error {
	// Checking nil Tags
	if Tags == nil {
		return ErrNilTags
	}

	// Generate id if nil
	if Tags.Id == uuid.Nil {
		Tags.Id = uuid.New()
	}

	if err := svc.TagsRepo.Insert(ctx, Tags); err != nil {
		return errors.Wrap(err, "[TagsService-Create]")
	}
	return nil
}

func (svc TagsService) GetListTags(ctx context.Context, limit, offset string) ([]*entity.Tags, error) {
	Tags, err := svc.TagsRepo.GetListTags(ctx, limit, offset)
	if err != nil {
		return nil, errors.Wrap(err, "[TagsService-GetListTags]")
	}
	return Tags, nil
}

func (svc TagsService) GetDetailTags(ctx context.Context, ID uuid.UUID) (*entity.Tags, error) {
	Tags, err := svc.TagsRepo.GetDetailTags(ctx, ID)
	if err != nil {
		return nil, errors.Wrap(err, "[TagsService-GetDetailTags]")
	}
	return Tags, nil
}

func (svc TagsService) DeleteTags(ctx context.Context, ID uuid.UUID) error {
	err := svc.TagsRepo.DeleteTags(ctx, ID)
	if err != nil {
		return errors.Wrap(err, "[TagsService-DeleteTags]")
	}
	return nil
}

func (svc TagsService) UpdateTags(ctx context.Context, Tags *entity.Tags) error {
	// Checking nil Tags
	if Tags == nil {
		return ErrNilTags
	}

	// Generate id if nil
	if Tags.Id == uuid.Nil {
		Tags.Id = uuid.New()
	}

	if err := svc.TagsRepo.UpdateTags(ctx, Tags); err != nil {
		return errors.Wrap(err, "[TagsService-UpdateTags]")
	}
	return nil
}
