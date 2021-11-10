package entity

import "github.com/google/uuid"

const (
	GameTableName = "game"
)

// GameModel is a model for entity.Game
type Game struct {
	Id       uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	NamaGame string    `gorm:"type:varchar;not_null" json:"nama_game"`
	Harga    int       `gorm:"type:int;not_null" json:"harga"`
	Genre    string    `gorm:"type:varchar;default:false;not_null" json:"genre"`
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

func NewGame(id uuid.UUID, nama_game string, harga int, genre string) *Game {
	return &Game{
		Id:       id,
		NamaGame: nama_game,
		Harga:    harga,
		Genre:    genre,
	}
}

// TableName specifies table name for GameModel.
func (model *Game) TableName() string {
	return GameTableName
}
