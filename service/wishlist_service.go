package service

import (
	"backend-service/entity"
	"context"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

var (
	// ErrNilWishlist occurs when a nil Wishlist is passed.
	ErrNilWishlist = errors.New("Wishlist is nil")
)

// WishlistService responsible for any flow related to Wishlist.
// It also implements WishlistService.
type WishlistService struct {
	WishlistRepo WishlistRepository
}

// NewWishlistService creates an instance of WishlistService.
func NewWishlistService(WishlistRepo WishlistRepository) *WishlistService {
	return &WishlistService{
		WishlistRepo: WishlistRepo,
	}
}

type WishlistUseCase interface {
	Create(ctx context.Context, Wishlist *entity.Wishlist) error
	GetListWishlist(ctx context.Context, limit, offset string) ([]*entity.Wishlist, error)
	GetGame(ctx context.Context, ID uuid.UUID) ([]*entity.WishlistGame, error)
	GetDetailWishlist(ctx context.Context, ID uuid.UUID) (*entity.Wishlist, error)
	UpdateWishlist(ctx context.Context, Wishlist *entity.Wishlist) error
	DeleteWishlist(ctx context.Context, credential_id uuid.UUID, game string) error
}

type WishlistRepository interface {
	Insert(ctx context.Context, Wishlist *entity.Wishlist) error
	GetListWishlist(ctx context.Context, limit, offset string) ([]*entity.Wishlist, error)
	GetGame(ctx context.Context, ID uuid.UUID) ([]*entity.WishlistGame, error)
	GetDetailWishlist(ctx context.Context, ID uuid.UUID) (*entity.Wishlist, error)
	UpdateWishlist(ctx context.Context, Wishlist *entity.Wishlist) error
	DeleteWishlist(ctx context.Context, credential_id uuid.UUID, game string) error
}

func (svc WishlistService) Create(ctx context.Context, Wishlist *entity.Wishlist) error {
	// Checking nil Wishlist
	if Wishlist == nil {
		return ErrNilWishlist
	}

	if err := svc.WishlistRepo.Insert(ctx, Wishlist); err != nil {
		return errors.Wrap(err, "[WishlistService-Create]")
	}
	return nil
}

func (svc WishlistService) GetListWishlist(ctx context.Context, limit, offset string) ([]*entity.Wishlist, error) {
	Wishlist, err := svc.WishlistRepo.GetListWishlist(ctx, limit, offset)
	if err != nil {
		return nil, errors.Wrap(err, "[WishlistService-GetListWishlist]")
	}
	return Wishlist, nil
}

func (svc WishlistService) GetGame(ctx context.Context, ID uuid.UUID) ([]*entity.WishlistGame, error) {
	Wishlist, err := svc.WishlistRepo.GetGame(ctx, ID)
	if err != nil {
		return nil, errors.Wrap(err, "[WishlistService-GetGameTags]")
	}
	return Wishlist, nil
}
func (svc WishlistService) GetDetailWishlist(ctx context.Context, ID uuid.UUID) (*entity.Wishlist, error) {
	Wishlist, err := svc.WishlistRepo.GetDetailWishlist(ctx, ID)
	if err != nil {
		return nil, errors.Wrap(err, "[WishlistService-GetDetailWishlist]")
	}
	return Wishlist, nil
}

func (svc WishlistService) DeleteWishlist(ctx context.Context, credential_id uuid.UUID, game string) error {
	err := svc.WishlistRepo.DeleteWishlist(ctx, credential_id, game)
	if err != nil {
		return errors.Wrap(err, "[WishlistService-DeleteWishlist]")
	}
	return nil
}

func (svc WishlistService) UpdateWishlist(ctx context.Context, Wishlist *entity.Wishlist) error {
	// Checking nil Wishlist
	if Wishlist == nil {
		return ErrNilWishlist
	}

	if err := svc.WishlistRepo.UpdateWishlist(ctx, Wishlist); err != nil {
		return errors.Wrap(err, "[WishlistService-UpdateWishlist]")
	}
	return nil
}
