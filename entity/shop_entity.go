package entity

import "github.com/google/uuid"

const (
	ShopTableName = "shop"
)

// ShopModel is a model for entity.Shop
type Shop struct {
	Id            uuid.UUID   `gorm:"type:uuid;primary_key;unique" json:"id"`
	Credential_id uuid.UUID   `gorm:"type:uuid;primary_key" json:"credential_id"`
	Image_url     string      `gorm:"type:varchar;null" json:"image_url"`
	Name          string      `gorm:"type:varchar;not_null;unique" json:"name"`
	Phone         string      `gorm:"type:varchar;not_null;unique" json:"phone"`
	Location      string      `gorm:"type:varchar;not_null;unique" json:"location"`
	Credential    *Credential `gorm:"foreignKey:Credential_id"`
}

func NewShop(id, credential_id uuid.UUID, image_url, name, phone, location string) *Shop {
	return &Shop{
		Id:            id,
		Credential_id: credential_id,
		Image_url:     image_url,
		Name:          name,
		Phone:         phone,
		Location:      location,
	}
}

// TableName specifies table name for ShopModel.
func (model *Shop) TableName() string {
	return ShopTableName
}
