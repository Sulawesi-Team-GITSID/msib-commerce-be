package entity

import "github.com/google/uuid"

const (
	VoucherTableName = "voucher"
)

// VoucherModel is a model for entity.Voucher
type Voucher struct {
	Id          uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	Game_id     uuid.UUID `gorm:"type:uuid;not_null" json:"game_id"`
	VoucherName string    `gorm:"type:varchar;not_null" json:"nama_voucher"`
	Harga       int       `gorm:"type:int;not_null" json:"harga"`
	Game        *Game     `gorm:"foreignKey:Game_id"`
}

func NewVoucher(id, game_id uuid.UUID, nama_voucher string, harga int) *Voucher {
	return &Voucher{
		Id:          id,
		Game_id:     game_id,
		VoucherName: nama_voucher,
		Harga:       harga,
	}
}

// TableName specifies table name for VoucherModel.
func (model *Voucher) TableName() string {
	return VoucherTableName
}
