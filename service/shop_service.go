package service

import (
	"backend-service/entity"
	"context"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

var (
	// ErrNilShop occurs when a nil Shop is passed.
	ErrNilShop = errors.New("Shop is nil")
)

// ShopService responsible for any flow related to Shop.
// It also implements ShopService.
type ShopService struct {
	ShopRepo ShopRepository
}

// NewShopService creates an instance of ShopService.
func NewShopService(ShopRepo ShopRepository) *ShopService {
	return &ShopService{
		ShopRepo: ShopRepo,
	}
}

type ShopUseCase interface {
	Create(ctx context.Context, Shop *entity.Shop) error
	GetListShop(ctx context.Context, limit, offset string) ([]*entity.Shop, error)
	GetDetailShop(ctx context.Context, ID uuid.UUID) (*entity.Shop, error)
	SearchShop(ctx context.Context, search string) ([]*entity.Shop, error)
	UpdateShop(ctx context.Context, Shop *entity.Shop) error
	DeleteShop(ctx context.Context, ID uuid.UUID) error
}

type ShopRepository interface {
	Insert(ctx context.Context, Shop *entity.Shop) error
	GetListShop(ctx context.Context, limit, offset string) ([]*entity.Shop, error)
	GetDetailShop(ctx context.Context, ID uuid.UUID) (*entity.Shop, error)
	SearchShop(ctx context.Context, search string) ([]*entity.Shop, error)
	UpdateShop(ctx context.Context, Shop *entity.Shop) error
	DeleteShop(ctx context.Context, ID uuid.UUID) error
}

func (svc ShopService) Create(ctx context.Context, Shop *entity.Shop) error {
	// Checking nil Shop
	if Shop == nil {
		return ErrNilShop
	}

	// Generate id if nil
	if Shop.Id == uuid.Nil {
		Shop.Id = uuid.New()
	}

	if err := svc.ShopRepo.Insert(ctx, Shop); err != nil {
		return errors.Wrap(err, "[ShopService-Create]")
	}
	return nil
}

func (svc ShopService) GetListShop(ctx context.Context, limit, offset string) ([]*entity.Shop, error) {
	Shop, err := svc.ShopRepo.GetListShop(ctx, limit, offset)
	if err != nil {
		return nil, errors.Wrap(err, "[ShopService-GetListShop]")
	}
	return Shop, nil
}

func (svc ShopService) GetListGenre(ctx context.Context, limit, offset string) ([]*entity.Shop, error) {
	Shop, err := svc.ShopRepo.GetListShop(ctx, limit, offset)
	if err != nil {
		return nil, errors.Wrap(err, "[ShopService-GetListGenre]")
	}
	return Shop, nil
}

func (svc ShopService) GetDetailShop(ctx context.Context, ID uuid.UUID) (*entity.Shop, error) {
	Shop, err := svc.ShopRepo.GetDetailShop(ctx, ID)
	if err != nil {
		return nil, errors.Wrap(err, "[ShopService-GetDetailShop]")
	}
	return Shop, nil
}

func (svc ShopService) SearchShop(ctx context.Context, search string) ([]*entity.Shop, error) {
	Shop, err := svc.ShopRepo.SearchShop(ctx, search)
	if err != nil {
		return nil, errors.Wrap(err, "[ShopService-SearchShop]")
	}

	return Shop, nil
}

func (svc ShopService) DeleteShop(ctx context.Context, ID uuid.UUID) error {
	err := svc.ShopRepo.DeleteShop(ctx, ID)
	if err != nil {
		return errors.Wrap(err, "[ShopService-DeleteShop]")
	}
	return nil
}

func (svc ShopService) UpdateShop(ctx context.Context, Shop *entity.Shop) error {
	// Checking nil Shop
	if Shop == nil {
		return ErrNilShop
	}

	// Generate id if nil
	if Shop.Id == uuid.Nil {
		Shop.Id = uuid.New()
	}

	if err := svc.ShopRepo.UpdateShop(ctx, Shop); err != nil {
		return errors.Wrap(err, "[ShopService-UpdateShop]")
	}
	return nil
}
