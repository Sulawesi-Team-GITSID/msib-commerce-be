package service

import (
	"backend-service/entity"
	"context"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

var (
	// ErrNilProfile occurs when a nil Profile is passed.
	ErrNilProfile = errors.New("Profile is nil")
)

// ProfileService responsible for any flow related to Profile.
// It also implements ProfileService.
type ProfileService struct {
	ProfileRepo ProfileRepository
}

// NewProfileService creates an instance of ProfileService.
func NewProfileService(ProfileRepo ProfileRepository) *ProfileService {
	return &ProfileService{
		ProfileRepo: ProfileRepo,
	}
}

type ProfileUseCase interface {
	Create(ctx context.Context, Profile *entity.Profile) error
	GetListProfile(ctx context.Context, limit, offset string) ([]*entity.Profile, error)
	GetDetailProfile(ctx context.Context, ID uuid.UUID) (*entity.Profile, error)
	UpdateProfile(ctx context.Context, Profile *entity.Profile) error
	DeleteProfile(ctx context.Context, ID uuid.UUID) error
}

type ProfileRepository interface {
	Insert(ctx context.Context, Profile *entity.Profile) error
	GetListProfile(ctx context.Context, limit, offset string) ([]*entity.Profile, error)
	GetDetailProfile(ctx context.Context, ID uuid.UUID) (*entity.Profile, error)
	UpdateProfile(ctx context.Context, Profile *entity.Profile) error
	DeleteProfile(ctx context.Context, ID uuid.UUID) error
}

func (svc ProfileService) Create(ctx context.Context, Profile *entity.Profile) error {
	// Checking nil Profile
	if Profile == nil {
		return ErrNilProfile
	}

	// Generate id if nil
	if Profile.Id == uuid.Nil {
		Profile.Id = uuid.New()
	}

	if err := svc.ProfileRepo.Insert(ctx, Profile); err != nil {
		return errors.Wrap(err, "[ProfileService-Create]")
	}
	return nil
}

func (svc ProfileService) GetListProfile(ctx context.Context, limit, offset string) ([]*entity.Profile, error) {
	Profile, err := svc.ProfileRepo.GetListProfile(ctx, limit, offset)
	if err != nil {
		return nil, errors.Wrap(err, "[ProfileService-Create]")
	}
	return Profile, nil
}

func (svc ProfileService) GetDetailProfile(ctx context.Context, ID uuid.UUID) (*entity.Profile, error) {
	Profile, err := svc.ProfileRepo.GetDetailProfile(ctx, ID)
	if err != nil {
		return nil, errors.Wrap(err, "[ProfileService-Create]")
	}
	return Profile, nil
}

func (svc ProfileService) DeleteProfile(ctx context.Context, ID uuid.UUID) error {
	err := svc.ProfileRepo.DeleteProfile(ctx, ID)
	if err != nil {
		return errors.Wrap(err, "[ProfileService-Create]")
	}
	return nil
}

func (svc ProfileService) UpdateProfile(ctx context.Context, Profile *entity.Profile) error {
	// Checking nil Profile
	if Profile == nil {
		return ErrNilProfile
	}

	// Generate id if nil
	if Profile.Id == uuid.Nil {
		Profile.Id = uuid.New()
	}

	if err := svc.ProfileRepo.UpdateProfile(ctx, Profile); err != nil {
		return errors.Wrap(err, "[ProfileService-Create]")
	}
	return nil
}
