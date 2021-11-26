package entity

import "github.com/google/uuid"

const (
	GiftCardTableName = "giftCard"
)

// GiftCardModel is a model for entity.GiftCard
type GiftCard struct {
	Id    uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	Nama  string    `gorm:"type:varchar;not_null" json:"nama"`
	Harga int       `gorm:"type:int;not_null" json:"harga"`
}

func NewGiftCard(id uuid.UUID, nama string, harga int) *GiftCard {
	return &GiftCard{
		Id:    id,
		Nama:  nama,
		Harga: harga,
	}
}

// TableName specifies table name for GiftCardModel.
func (model *GiftCard) TableName() string {
	return GiftCardTableName
}
