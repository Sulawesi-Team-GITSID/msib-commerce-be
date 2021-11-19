package repository

import (
	"backend-service/entity"
	"context"

	"github.com/google/uuid"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// VoucherRepository connects entity.Voucher with database.
type VoucherRepository struct {
	db *gorm.DB
}

// NewVoucherRepository creates an instance of RoleRepository.
func NewVoucherRepository(db *gorm.DB) *VoucherRepository {
	return &VoucherRepository{
		db: db,
	}
}

// Insert inserts Voucher data to database.
func (repo *VoucherRepository) Insert(ctx context.Context, ent *entity.Voucher) error {
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Voucher{}).
		Create(ent).
		Error; err != nil {
		return errors.Wrap(err, "[VoucherRepository-Insert]")
	}
	return nil
}

func (repo *VoucherRepository) GetListVoucher(ctx context.Context, limit, offset string) ([]*entity.Voucher, error) {
	var models []*entity.Voucher
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Voucher{}).
		Find(&models).
		Error; err != nil {
		return nil, errors.Wrap(err, "[VoucherRepository-FindAll]")
	}
	return models, nil
}

func (repo *VoucherRepository) GetListVoucherShop(ctx context.Context, limit, offset string) ([]*entity.VoucherShop, error) {
	var models []*entity.VoucherShop
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Voucher{}).
		Select("voucher.id", "shop_id", "voucher_name", "harga", "shop.name as shop").
		Joins("inner join shop on voucher.shop_id = shop.id").Order("voucher_name desc").
		Find(&models).
		Error; err != nil {
		return nil, errors.Wrap(err, "[VoucherRepository-FindAll]")
	}
	return models, nil
}

func (repo *VoucherRepository) GetDetailVoucher(ctx context.Context, ID uuid.UUID) (*entity.Voucher, error) {
	var models *entity.Voucher
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Voucher{}).
		Take(&models, ID).
		Error; err != nil {
		return nil, errors.Wrap(err, "[VoucherRepository-FindById]")
	}
	return models, nil
}

func (repo *VoucherRepository) DeleteVoucher(ctx context.Context, ID uuid.UUID) error {
	if err := repo.db.
		WithContext(ctx).
		Delete(&entity.Voucher{Id: ID}).Error; err != nil {
		return errors.Wrap(err, "[VoucherRepository-Delete]")
	}
	return nil
}

func (repo *VoucherRepository) UpdateVoucher(ctx context.Context, ent *entity.Voucher) error {
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Voucher{Id: ent.Id}).
		Select("voucher_name", "harga").
		Updates(ent).Error; err != nil {
		return errors.Wrap(err, "[VoucherRepository-Update]")
	}
	return nil
}
