package entity

import "github.com/google/uuid"

const (
	ReviewTableName = "review"
)

// ReviewModel is a model for entity.Review
type Review struct {
	Id      uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	Game_id uuid.UUID `gorm:"type:uuid;not_null" json:"game_id"`
	Rating  float64   `gorm:"type:numeric;not_null" json:"rating"`
	Comment string    `gorm:"type:varchar;not_null" json:"comment"`
	Game    *Game     `gorm:"foreignKey:Game_id"`
}

func NewReview(id, game_id uuid.UUID, rating float64, comment string) *Review {
	return &Review{
		Id:      id,
		Game_id: game_id,
		Rating:  rating,
		Comment: comment,
	}
}

// TableName specifies table name for ReviewModel.
func (model *Review) TableName() string {
	return ReviewTableName
}
