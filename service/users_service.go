package service

import (
	"backend-service/entity"
	"context"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

var (
	ErrNilUsers = errors.New("users is nil")
)

type UsersUseCase interface {
	Login(ctx context.Context, email string, password string) (*entity.Users, error)
	Register(ctx context.Context, users *entity.Users) error
	// GetProfile(ctx context.Context, ID uuid.UUID) (*entity.Users, error)
}

type UsersRepository interface {
	Login(ctx context.Context, email string, password string) (*entity.Users, error)
	Register(ctx context.Context, users *entity.Users) error
	// GetProfile(ctx context.Context, ID uuid.UUID) (*entity.Users, error)
}
type UsersService struct {
	usersRepo UsersRepository
}

// NewUsersService creates an instance of UsersService.
func NewUsersService(usersRepo UsersRepository) *UsersService {
	return &UsersService{
		usersRepo: usersRepo,
	}
}

func (svc UsersService) Register(ctx context.Context, users *entity.Users) error {
	// Checking nil barang
	if users == nil {
		return ErrNilUsers
	}

	// Generate id if nil
	if users.ID == uuid.Nil {
		users.ID = uuid.New()
	}

	if err := svc.usersRepo.Register(ctx, users); err != nil {
		return errors.Wrap(err, "[UsersService-Register]")
	}
	return nil
}

func (svc UsersService) Login(ctx context.Context, email string, password string) (*entity.Users, error) {

	userData, err := svc.usersRepo.Login(ctx, email, password)

	if err != nil {
		return nil, err
	}

	return userData, nil
}
