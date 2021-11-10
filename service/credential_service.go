package service

import (
	"backend-service/entity"
	"context"
	"strconv"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

var (
	// ErrNilCredential occurs when a nil Credential is passed.
	ErrNilCredential = errors.New("Credential is nil")
)

// CredentialService responsible for any flow related to Credential.
// It also implements CredentialService.
type CredentialService struct {
	CredentialRepo CredentialRepository
}

// NewCredentialService creates an instance of CredentialService.
func NewCredentialService(CredentialRepo CredentialRepository) *CredentialService {
	return &CredentialService{
		CredentialRepo: CredentialRepo,
	}
}

type CredentialUseCase interface {
	Create(ctx context.Context, Credential *entity.Credential) error
	Login(ctx context.Context, email string, password string) (*entity.Credential, error)
	GetListCredential(ctx context.Context, limit, offset string) ([]*entity.Credential, error)
}

type CredentialRepository interface {
	Insert(ctx context.Context, Credential *entity.Credential) error
	Login(ctx context.Context, email string, password string) (*entity.Credential, error)
	GetListCredential(ctx context.Context, limit, offset string) ([]*entity.Credential, error)
}

func (svc CredentialService) Create(ctx context.Context, Credential *entity.Credential) error {
	// Checking nil Credential
	if Credential == nil {
		return ErrNilCredential
	}

	// Generate id if nil
	if Credential.Id == uuid.Nil {
		Credential.Id = uuid.New()
	}

	if val := strconv.FormatBool(Credential.Seller); val == "" {
		Credential.Seller = false
	}

	if val := strconv.FormatBool(Credential.Verified); val == "" {
		Credential.Verified = false
	}

	if err := svc.CredentialRepo.Insert(ctx, Credential); err != nil {
		return errors.Wrap(err, "[CredentialService-Create]")
	}
	return nil
}

func (svc CredentialService) GetListCredential(ctx context.Context, limit, offset string) ([]*entity.Credential, error) {
	Profile, err := svc.CredentialRepo.GetListCredential(ctx, limit, offset)
	if err != nil {
		return nil, errors.Wrap(err, "[CredentialService-GetList]")
	}
	return Profile, nil
}

func (svc CredentialService) Login(ctx context.Context, email string, password string) (*entity.Credential, error) {

	CredentialData, err := svc.CredentialRepo.Login(ctx, email, password)

	if err != nil {
		return nil, err
	}

	return CredentialData, nil
}
