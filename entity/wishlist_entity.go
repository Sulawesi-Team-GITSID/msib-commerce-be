package entity

import "github.com/google/uuid"

const (
	WishlistTableName = "wishlist"
)

// WishlistModel is a model for entity.Wishlist
type Wishlist struct {
	Credential_id uuid.UUID   `gorm:"type:uuid;not_null" json:"credential_id"`
	Game_id       uuid.UUID   `gorm:"type:uuid;not_null" json:"game_id"`
	Credential    *Credential `gorm:"foreignKey:Credential_id"`
	Game          *Game       `gorm:"foreignKey:Game_id"`
}

type WishlistGame struct {
	Credential_id uuid.UUID `json:"credential_id"`
	Game          string    `json:"game"`
}

func NewWishlist(credential_id, game_id uuid.UUID) *Wishlist {
	return &Wishlist{
		Credential_id: credential_id,
		Game_id:       game_id,
	}
}

// TableName specifies table name for WishlistModel.
func (model *Wishlist) TableName() string {
	return WishlistTableName
}
