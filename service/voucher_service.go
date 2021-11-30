package service

import (
	"backend-service/entity"
	"context"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

var (
	// ErrNilVoucher occurs when a nil Voucher is passed.
	ErrNilVoucher = errors.New("Voucher is nil")
)

// VoucherService responsible for any flow related to Voucher.
// It also implements VoucherService.
type VoucherService struct {
	VoucherRepo VoucherRepository
}

// NewVoucherService creates an instance of VoucherService.
func NewVoucherService(VoucherRepo VoucherRepository) *VoucherService {
	return &VoucherService{
		VoucherRepo: VoucherRepo,
	}
}

type VoucherUseCase interface {
	Create(ctx context.Context, Voucher *entity.Voucher) error
	GetListVoucher(ctx context.Context, limit, offset string) ([]*entity.ListVoucher, error)
	GetListVoucherShop(ctx context.Context, ID uuid.UUID) ([]*entity.VoucherShop, error)
	SortVoucher(ctx context.Context, order, sort string) ([]*entity.ListVoucher, error)
	SortVoucherByShop(ctx context.Context, order, sort string, ID uuid.UUID) ([]*entity.VoucherShop, error)
	GetDetailVoucher(ctx context.Context, ID uuid.UUID) (*entity.Voucher, error)
	SearchVoucher(ctx context.Context, search string) ([]*entity.ListVoucher, error)
	UpdateVoucher(ctx context.Context, Voucher *entity.Voucher) error
	DeleteVoucher(ctx context.Context, ID uuid.UUID) error
}

type VoucherRepository interface {
	Insert(ctx context.Context, Voucher *entity.Voucher) error
	GetListVoucher(ctx context.Context, limit, offset string) ([]*entity.ListVoucher, error)
	GetListVoucherShop(ctx context.Context, ID uuid.UUID) ([]*entity.VoucherShop, error)
	SortVoucher(ctx context.Context, order, sort string) ([]*entity.ListVoucher, error)
	SortVoucherByShop(ctx context.Context, order, sort string, ID uuid.UUID) ([]*entity.VoucherShop, error)
	GetDetailVoucher(ctx context.Context, ID uuid.UUID) (*entity.Voucher, error)
	SearchVoucher(ctx context.Context, search string) ([]*entity.ListVoucher, error)
	UpdateVoucher(ctx context.Context, Voucher *entity.Voucher) error
	DeleteVoucher(ctx context.Context, ID uuid.UUID) error
}

func (svc VoucherService) Create(ctx context.Context, Voucher *entity.Voucher) error {
	// Checking nil Voucher
	if Voucher == nil {
		return ErrNilVoucher
	}

	// Generate id if nil
	if Voucher.Id == uuid.Nil {
		Voucher.Id = uuid.New()
	}

	if err := svc.VoucherRepo.Insert(ctx, Voucher); err != nil {
		return errors.Wrap(err, "[VoucherService-Create]")
	}
	return nil
}

func (svc VoucherService) GetListVoucher(ctx context.Context, limit, offset string) ([]*entity.ListVoucher, error) {
	Voucher, err := svc.VoucherRepo.GetListVoucher(ctx, limit, offset)
	if err != nil {
		return nil, errors.Wrap(err, "[VoucherService-GetListVoucher]")
	}
	return Voucher, nil
}

func (svc VoucherService) GetListVoucherShop(ctx context.Context, ID uuid.UUID) ([]*entity.VoucherShop, error) {
	Voucher, err := svc.VoucherRepo.GetListVoucherShop(ctx, ID)
	if err != nil {
		return nil, errors.Wrap(err, "[VoucherService-GetListVoucherShop]")
	}
	return Voucher, nil
}

func (svc VoucherService) SortVoucher(ctx context.Context, order, sort string) ([]*entity.ListVoucher, error) {
	Voucher, err := svc.VoucherRepo.SortVoucher(ctx, order, sort)
	if err != nil {
		return nil, errors.Wrap(err, "[VoucherService-SortByAsc]")
	}
	return Voucher, nil
}

func (svc VoucherService) SortVoucherByShop(ctx context.Context, order, sort string, ID uuid.UUID) ([]*entity.VoucherShop, error) {
	Voucher, err := svc.VoucherRepo.SortVoucherByShop(ctx, order, sort, ID)
	if err != nil {
		return nil, errors.Wrap(err, "[VoucherService-SortVoucherByShop]")
	}
	return Voucher, nil
}

func (svc VoucherService) GetDetailVoucher(ctx context.Context, ID uuid.UUID) (*entity.Voucher, error) {
	Voucher, err := svc.VoucherRepo.GetDetailVoucher(ctx, ID)
	if err != nil {
		return nil, errors.Wrap(err, "[VoucherService-GetDetailVoucher]")
	}
	return Voucher, nil
}

func (svc VoucherService) DeleteVoucher(ctx context.Context, ID uuid.UUID) error {
	err := svc.VoucherRepo.DeleteVoucher(ctx, ID)
	if err != nil {
		return errors.Wrap(err, "[VoucherService-DeleteVoucher]")
	}
	return nil
}

func (svc VoucherService) SearchVoucher(ctx context.Context, search string) ([]*entity.ListVoucher, error) {
	Voucher, err := svc.VoucherRepo.SearchVoucher(ctx, search)
	if err != nil {
		return nil, errors.Wrap(err, "[VoucherService-SearchVoucher]")
	}

	return Voucher, nil
}

func (svc VoucherService) UpdateVoucher(ctx context.Context, Voucher *entity.Voucher) error {
	// Checking nil Voucher
	if Voucher == nil {
		return ErrNilVoucher
	}

	// Generate id if nil
	if Voucher.Id == uuid.Nil {
		Voucher.Id = uuid.New()
	}

	if err := svc.VoucherRepo.UpdateVoucher(ctx, Voucher); err != nil {
		return errors.Wrap(err, "[VoucherService-UpdateVoucher]")
	}
	return nil
}
