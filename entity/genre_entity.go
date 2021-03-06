package entity

import (
	"github.com/google/uuid"
)

const (
	GenreTableName = "genre"
)

// GenreModel is a model for entity.Genre
type Genre struct {
	Id      uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	Name    string    `gorm:"type:varchar;not_null;unique" json:"name"`
	Deleted bool      `gorm:"type:bool;default:false;not_null" json:"deleted"`
}

func NewGenre(id uuid.UUID, name string, deleted bool) *Genre {
	return &Genre{
		Id:      id,
		Name:    name,
		Deleted: deleted,
	}
}

// TableName specifies table name for GenreModel.
func (model *Genre) TableName() string {
	return GenreTableName
}
