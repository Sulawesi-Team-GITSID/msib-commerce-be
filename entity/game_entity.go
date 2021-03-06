package entity

import "github.com/google/uuid"

const (
	GameTableName = "game"
)

// GameModel is a model for entity.Game
type Game struct {
	Id        uuid.UUID `gorm:"type:uuid;primary_key;unique" json:"id"`
	Shop_id   uuid.UUID `gorm:"type:uuid;not_null" json:"shop_id"`
	Image_url string    `gorm:"type:varchar;null" json:"image_url"`
	NamaGame  string    `gorm:"type:varchar;not_null" json:"nama_game"`
	Harga     int       `gorm:"type:int;not_null" json:"harga"`
	Genre_id  uuid.UUID `gorm:"type:uuid;not_null" json:"genre_id"`
	Deleted   bool      `gorm:"type:bool;default:false;not_null" json:"deleted"`
	Shop      *Shop     `gorm:"foreignKey:Shop_id"`
	Genre     *Genre    `gorm:"foreignKey:Genre_id"`
}

type ListGenre struct {
	Genre string `gorm:"type:varchar;default:false;not_null" json:"genre"`
}

// ListGame struct is a model for collecting query genre.name
type ListGame struct {
	Id        uuid.UUID `json:"id"`
	Shop_id   uuid.UUID `json:"shop_id"`
	NamaGame  string    `json:"nama_game"`
	Image_url string    `json:"image_url"`
	Harga     int       `json:"harga"`
	Genre     string    `json:"genre"`
}

type GameShop struct {
	Id        uuid.UUID `json:"id"`
	Shop_id   uuid.UUID `json:"shop_id"`
	NamaGame  string    `json:"nama_game"`
	Image_url string    `json:"image_url"`
	Harga     int       `json:"harga"`
	Shop      string    `json:"shop"`
}

type TrendGame struct {
	Id        uuid.UUID `json:"id"`
	NamaGame  string    `json:"nama_game"`
	Harga     int       `json:"harga"`
	Image_url string    `json:"image_url"`
	Shop_id   uuid.UUID `json:"shop_id"`
	Rating    float64   `json:"rating"`
}

func NewGame(id, shop_id uuid.UUID, image_url, nama_game string, harga int, genre_id uuid.UUID, deleted bool) *Game {
	return &Game{
		Id:        id,
		Shop_id:   shop_id,
		Image_url: image_url,
		NamaGame:  nama_game,
		Harga:     harga,
		Genre_id:  genre_id,
		Deleted:   deleted,
	}
}

// TableName specifies table name for GameModel.
func (model *Game) TableName() string {
	return GameTableName
}
