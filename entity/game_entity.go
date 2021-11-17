package entity

import "github.com/google/uuid"

const (
	GameTableName = "game"
)

// GameModel is a model for entity.Game
type Game struct {
	Id         uuid.UUID   `gorm:"type:uuid;primary_key" json:"id"`
	Shop_id    uuid.UUID   `gorm:"type:uuid;primary_key" json:"shop_id"`
	NamaGame   string      `gorm:"type:varchar;not_null" json:"nama_game"`
	Harga      int         `gorm:"type:int;not_null" json:"harga"`
	Genre      string      `gorm:"type:varchar;default:false;not_null" json:"genre"`
	Credential *Credential `gorm:"foreignKey:Shop_id"`
}

type Genre struct {
	Genre string `gorm:"type:varchar;default:false;not_null" json:"genre"`
}

type TrendGame struct {
	Id       uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	NamaGame string    `gorm:"type:varchar;not_null" json:"nama_game"`
	Harga    int       `gorm:"type:int;not_null" json:"harga"`
	Rating   float64   `gorm:"type:numeric;not_null" json:"rating"`
}

func NewGame(id, shop_id uuid.UUID, nama_game string, harga int, genre string) *Game {
	return &Game{
		Id:       id,
		Shop_id:  shop_id,
		NamaGame: nama_game,
		Harga:    harga,
		Genre:    genre,
	}
}

// TableName specifies table name for GameModel.
func (model *Game) TableName() string {
	return GameTableName
}
