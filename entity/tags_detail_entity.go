package entity

import "github.com/google/uuid"

const (
	Tags_detailTableName = "tags_detail"
)

// Tags_detailModel is a model for entity.Tags_detail
type Tags_detail struct {
	Game_id uuid.UUID `gorm:"type:uuid;not_null" json:"game_id"`
	Tags_id uuid.UUID `gorm:"type:uuid;not_null" json:"tags_id"`
	Game    *Game     `gorm:"foreignKey:Game_id"`
	Tags    *Tags     `gorm:"foreignKey:Tags_id"`
}

type GameTags struct {
	Game_id uuid.UUID `gorm:"type:uuid;not_null" json:"game_id"`
	Tags    string    `gorm:"type:uuid;not_null" json:"tags"`
}

func NewTags_detail(game_id, tags_id uuid.UUID) *Tags_detail {
	return &Tags_detail{
		Game_id: game_id,
		Tags_id: tags_id,
	}
}

// TableName specifies table name for Tags_detailModel.
func (model *Tags_detail) TableName() string {
	return Tags_detailTableName
}
