package entity

import "github.com/google/uuid"

const (
	VoucherTableName = "voucher"
)

// VoucherModel is a model for entity.Voucher
type Voucher struct {
	Id           uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	Game_id      uuid.UUID `gorm:"type:uuid;not_null" json:"game_id"`
	Shop_id      uuid.UUID `gorm:"type:uuid;not_null" json:"shop_id"`
	Voucher_name string    `gorm:"type:varchar;not_null" json:"voucher_name"`
	Harga        int       `gorm:"type:int;not_null" json:"harga"`
	Deleted      bool      `gorm:"type:bool;default:false;not_null" json:"deleted"`
	Game         *Game     `gorm:"foreignKey:Game_id"`
	Shop         *Shop     `gorm:"foreignKey:Shop_id"`
}

type ListVoucher struct {
	Id           uuid.UUID `json:"id"`
	Game_id      uuid.UUID `json:"game_id"`
	Game_name    string    `json:"game_name"`
	Shop_id      uuid.UUID `json:"shop_id"`
	Shop_name    string    `json:"shop_name"`
	Voucher_name string    `json:"voucher_name"`
	Harga        int       `json:"harga"`
}

func NewVoucher(id, game_id, shop_id uuid.UUID, voucher_name string, harga int, deleted bool) *Voucher {
	return &Voucher{
		Id:           id,
		Game_id:      game_id,
		Shop_id:      shop_id,
		Voucher_name: voucher_name,
		Harga:        harga,
		Deleted:      deleted,
	}
}

// TableName specifies table name for VoucherModel.
func (model *Voucher) TableName() string {
	return VoucherTableName
}
