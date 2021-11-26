package service

import (
	"backend-service/entity"
	"context"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

var (
	// ErrNilReview occurs when a nil Review is passed.
	ErrNilReview = errors.New("Review is nil")
)

// ReviewService responsible for any flow related to Review.
// It also implements ReviewService.
type ReviewService struct {
	ReviewRepo ReviewRepository
}

// NewReviewService creates an instance of ReviewService.
func NewReviewService(ReviewRepo ReviewRepository) *ReviewService {
	return &ReviewService{
		ReviewRepo: ReviewRepo,
	}
}

type ReviewUseCase interface {
	Create(ctx context.Context, Review *entity.Review) error
	GetListReview(ctx context.Context, limit, offset string) ([]*entity.Review, error)
	// GetListGenre(ctx context.Context, limit, offset string) ([]*entity.Review, error)
	// GetListTrendReview(ctx context.Context, limit, offset string) ([]*entity.Review, error)
	// GetDetailReview(ctx context.Context, ID uuid.UUID) (*entity.Review, error)
	// UpdateReview(ctx context.Context, Review *entity.Review) error
	// DeleteReview(ctx context.Context, ID uuid.UUID) error
}

type ReviewRepository interface {
	Insert(ctx context.Context, Review *entity.Review) error
	GetListReview(ctx context.Context, limit, offset string) ([]*entity.Review, error)
	// GetListGenre(ctx context.Context, limit, offset string) ([]*entity.Review, error)
	// GetListTrendReview(ctx context.Context, limit, offset string) ([]*entity.Review, error)
	// GetDetailReview(ctx context.Context, ID uuid.UUID) (*entity.Review, error)
	// UpdateReview(ctx context.Context, Review *entity.Review) error
	// DeleteReview(ctx context.Context, ID uuid.UUID) error
}

func (svc ReviewService) Create(ctx context.Context, Review *entity.Review) error {
	// Checking nil Review
	if Review == nil {
		return ErrNilReview
	}

	// Generate id if nil
	if Review.Id == uuid.Nil {
		Review.Id = uuid.New()
	}

	if Review.Rating < 0 {
		Review.Rating = 1
	}
	if Review.Rating > 5 {
		Review.Rating = 5
	}

	if err := svc.ReviewRepo.Insert(ctx, Review); err != nil {
		return errors.Wrap(err, "[ReviewService-Create]")
	}
	return nil
}

func (svc ReviewService) GetListReview(ctx context.Context, limit, offset string) ([]*entity.Review, error) {
	Review, err := svc.ReviewRepo.GetListReview(ctx, limit, offset)
	if err != nil {
		return nil, errors.Wrap(err, "[ReviewService-GetListReview]")
	}
	return Review, nil
}

// func (svc ReviewService) GetDetailReview(ctx context.Context, ID uuid.UUID) (*entity.Review, error) {
// 	Review, err := svc.ReviewRepo.GetDetailReview(ctx, ID)
// 	if err != nil {
// 		return nil, errors.Wrap(err, "[ReviewService-GetDetailReview]")
// 	}
// 	return Review, nil
// }

// func (svc ReviewService) DeleteReview(ctx context.Context, ID uuid.UUID) error {
// 	err := svc.ReviewRepo.DeleteReview(ctx, ID)
// 	if err != nil {
// 		return errors.Wrap(err, "[ReviewService-DeleteReview]")
// 	}
// 	return nil
// }

// func (svc ReviewService) UpdateReview(ctx context.Context, Review *entity.Review) error {
// 	// Checking nil Review
// 	if Review == nil {
// 		return ErrNilReview
// 	}

// 	// Generate id if nil
// 	if Review.Id == uuid.Nil {
// 		Review.Id = uuid.New()
// 	}

// 	if err := svc.ReviewRepo.UpdateReview(ctx, Review); err != nil {
// 		return errors.Wrap(err, "[ReviewService-UpdateReview]")
// 	}
// 	return nil
// }
