package service

import (
	"backend-service/entity"
	"context"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

var (
	ErrNilSuperAdmin = errors.New("superAdmin is nil")
)

type SuperAdminUseCase interface {
	LoginAdmin(ctx context.Context, email string, password string) (*entity.SuperAdmin, error)
	Register(ctx context.Context, superAdmin *entity.SuperAdmin) error
	// GetProfile(ctx context.Context, ID uuid.UUID) (*entity.SuperAdmin, error)
}

type SuperAdminRepository interface {
	LoginAdmin(ctx context.Context, email string, password string) (*entity.SuperAdmin, error)
	Register(ctx context.Context, superAdmin *entity.SuperAdmin) error
	// GetProfile(ctx context.Context, ID uuid.UUID) (*entity.SuperAdmin, error)
}
type SuperAdminService struct {
	superAdminRepo SuperAdminRepository
}

// NewSuperAdminService creates an instance of SuperAdminService.
func NewSuperAdminService(superAdminRepo SuperAdminRepository) *SuperAdminService {
	return &SuperAdminService{
		superAdminRepo: superAdminRepo,
	}
}

func (svc SuperAdminService) Register(ctx context.Context, superAdmin *entity.SuperAdmin) error {
	// Checking nil barang
	if superAdmin == nil {
		return ErrNilSuperAdmin
	}

	// Generate id if nil
	if superAdmin.ID == uuid.Nil {
		superAdmin.ID = uuid.New()
	}

	if err := svc.superAdminRepo.Register(ctx, superAdmin); err != nil {
		return errors.Wrap(err, "[SuperAdminService-Register]")
	}
	return nil
}

func (svc SuperAdminService) LoginAdmin(ctx context.Context, email string, password string) (*entity.SuperAdmin, error) {

	userData, err := svc.superAdminRepo.LoginAdmin(ctx, email, password)

	if err != nil {
		return nil, err
	}

	return userData, nil
}
