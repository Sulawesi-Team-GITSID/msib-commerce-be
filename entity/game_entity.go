package entity

import "github.com/google/uuid"

const (
	GameTableName = "game"
)

// GameModel is a model for entity.Game
type Game struct {
	Id       uuid.UUID `gorm:"type:uuid;primary_key;unique" json:"id"`
	Shop_id  uuid.UUID `gorm:"type:uuid;not_null" json:"shop_id"`
	NamaGame string    `gorm:"type:varchar;not_null" json:"nama_game"`
	Harga    int       `gorm:"type:int;not_null" json:"harga"`
	Genre_id uuid.UUID `gorm:"type:uuid;not_null" json:"genre_id"`
	Shop     *Shop     `gorm:"foreignKey:Shop_id"`
	Genre    *Genre    `gorm:"foreignKey:Genre_id"`
}

type ListGenre struct {
	Genre string `gorm:"type:varchar;default:false;not_null" json:"genre"`
}

// ListGame struct is a model for collecting query genre.name
type ListGame struct {
	Id       uuid.UUID `json:"id"`
	Shop_id  uuid.UUID `json:"shop_id"`
	NamaGame string    `json:"nama_game"`
	Harga    int       `json:"harga"`
	Genre    string    `json:"genre"`
}

type TrendGame struct {
	Id       uuid.UUID `json:"id"`
	NamaGame string    `json:"nama_game"`
	Harga    int       `json:"harga"`
	Rating   float64   `json:"rating"`
}

func NewGame(id, shop_id uuid.UUID, nama_game string, harga int, genre_id uuid.UUID) *Game {
	return &Game{
		Id:       id,
		Shop_id:  shop_id,
		NamaGame: nama_game,
		Harga:    harga,
		Genre_id: genre_id,
	}
}

// TableName specifies table name for GameModel.
func (model *Game) TableName() string {
	return GameTableName
}
