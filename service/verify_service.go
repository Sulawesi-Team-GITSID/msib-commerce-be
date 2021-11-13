package service

import (
	"backend-service/entity"
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

var (
	// ErrNilVerification occurs when a nil Verification is passed.
	ErrNilVerification = errors.New("Verification is nil")

	//If user verify after expiration time
	ErrNilExpired = errors.New("Sorry, link is expired")
)

// VerificationService responsible for any flow related to Verification.
// It also implements VerificationService.
type VerificationService struct {
	VerificationRepo VerificationRepository
}

// NewVerificationService creates an instance of VerificationService.
func NewVerificationService(VerificationRepo VerificationRepository) *VerificationService {
	return &VerificationService{
		VerificationRepo: VerificationRepo,
	}
}

type VerificationUseCase interface {
	Create(ctx context.Context, Verification *entity.Verification) error
	Verify(ctx context.Context, credential_id string) (*entity.Getcode, error)
	GetListVerification(ctx context.Context, limit, offset string) ([]*entity.Verification, error)
	GetDetailVerification(ctx context.Context, ID uuid.UUID) (*entity.Verification, error)
	UpdateVerification(ctx context.Context, Verification *entity.Verification) error
	DeleteVerification(ctx context.Context, ID uuid.UUID) error
}

type VerificationRepository interface {
	Insert(ctx context.Context, Verification *entity.Verification) error
	Verify(ctx context.Context, credential_id string) (*entity.Getcode, error)
	GetListVerification(ctx context.Context, limit, offset string) ([]*entity.Verification, error)
	GetDetailVerification(ctx context.Context, ID uuid.UUID) (*entity.Verification, error)
	UpdateVerification(ctx context.Context, Verification *entity.Verification) error
	DeleteVerification(ctx context.Context, ID uuid.UUID) error
}

func (svc VerificationService) Create(ctx context.Context, Verification *entity.Verification) error {
	// Checking nil Verification
	if Verification == nil {
		return ErrNilVerification
	}

	// Generate id if nil
	if Verification.Id == uuid.Nil {
		Verification.Id = uuid.New()
	}

	if err := svc.VerificationRepo.Insert(ctx, Verification); err != nil {
		return errors.Wrap(err, "[VerificationService-Create]")
	}
	return nil
}

func (svc VerificationService) GetListVerification(ctx context.Context, limit, offset string) ([]*entity.Verification, error) {
	Verification, err := svc.VerificationRepo.GetListVerification(ctx, limit, offset)
	if err != nil {
		return nil, errors.Wrap(err, "[VerificationService-GetListVerification]")
	}
	return Verification, nil
}

func (svc VerificationService) GetDetailVerification(ctx context.Context, ID uuid.UUID) (*entity.Verification, error) {
	Verification, err := svc.VerificationRepo.GetDetailVerification(ctx, ID)
	if err != nil {
		return nil, errors.Wrap(err, "[VerificationService-GetDetailVerification]")
	}
	return Verification, nil
}

func (svc VerificationService) DeleteVerification(ctx context.Context, ID uuid.UUID) error {
	err := svc.VerificationRepo.DeleteVerification(ctx, ID)
	if err != nil {
		return errors.Wrap(err, "[VerificationService-DeleteVerification]")
	}
	return nil
}

func (svc VerificationService) UpdateVerification(ctx context.Context, Verification *entity.Verification) error {
	// Checking nil Verification
	if Verification == nil {
		return ErrNilVerification
	}

	// Generate id if nil
	if Verification.Id == uuid.Nil {
		Verification.Id = uuid.New()
	}

	if err := svc.VerificationRepo.UpdateVerification(ctx, Verification); err != nil {
		return errors.Wrap(err, "[VerificationService-UpdateVerification]")
	}
	return nil
}

func (svc VerificationService) Verify(ctx context.Context, credential_id string) (*entity.Getcode, error) {

	VerificationData, err := svc.VerificationRepo.Verify(ctx, credential_id)

	if VerificationData.Expiresat.Before(time.Now()) {
		return nil, ErrNilExpired
	}

	if err != nil {
		return nil, err
	}

	return VerificationData, nil
}
