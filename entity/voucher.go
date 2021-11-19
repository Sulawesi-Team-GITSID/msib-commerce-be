package entity

import "github.com/google/uuid"

const (
	VoucherTableName = "voucher"
)

// VoucherModel is a model for entity.Voucher
type Voucher struct {
	Id          uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	Game_id     uuid.UUID `gorm:"type:uuid;not_null" json:"game_id"`
	Shop_id     uuid.UUID `gorm:"type:uuid;not_null" json:"shop_id"`
	VoucherName string    `gorm:"type:varchar;not_null" json:"voucher_name"`
	Harga       int       `gorm:"type:int;not_null" json:"harga"`
	Game        *Game     `gorm:"foreignKey:Game_id"`
	Shop        *Shop     `gorm:"foreignKey:Shop_id"`
}

type VoucherShop struct {
	Id          uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	Shop_id     uuid.UUID `gorm:"type:uuid;not_null" json:"shop_id"`
	VoucherName string    `gorm:"type:varchar;not_null" json:"voucher_name"`
	Harga       int       `gorm:"type:int;not_null" json:"harga"`
	Shop        string    `gorm:"type:varchar;not_null" json:"shop"`
}

func NewVoucher(id, game_id, shop_id uuid.UUID, voucher_name string, harga int) *Voucher {
	return &Voucher{
		Id:          id,
		Game_id:     game_id,
		Shop_id:     shop_id,
		VoucherName: voucher_name,
		Harga:       harga,
	}
}

// TableName specifies table name for VoucherModel.
func (model *Voucher) TableName() string {
	return VoucherTableName
}
