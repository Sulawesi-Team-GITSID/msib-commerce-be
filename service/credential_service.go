package service

import (
	"backend-service/entity"
	"context"

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
}

type CredentialRepository interface {
	Insert(ctx context.Context, Credential *entity.Credential) error
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

	if err := svc.CredentialRepo.Insert(ctx, Credential); err != nil {
		return errors.Wrap(err, "[CredentialService-Create]")
	}
	return nil
}
