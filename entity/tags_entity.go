package entity

import "github.com/google/uuid"

const (
	TagsTableName = "tags"
)

// TagsModel is a model for entity.Tags
type Tags struct {
	Id   uuid.UUID `gorm:"type:uuid;primary_key;unique" json:"id"`
	Name string    `gorm:"type:varchar;not_null;unique" json:"name"`
}

func NewTags(id uuid.UUID, name string) *Tags {
	return &Tags{
		Id:   id,
		Name: name,
	}
}

// TableName specifies table name for TagsModel.
func (model *Tags) TableName() string {
	return TagsTableName
}
