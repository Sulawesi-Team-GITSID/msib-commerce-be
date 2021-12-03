package entity

import "github.com/google/uuid"

const (
	FileTableName = "file"
)

// ArticleModel is a model for entity.Article
type File struct {
	Id        uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	Entity_id uuid.UUID `gorm:"type:uuid;foreign_key" json:"entity_id"`
	Path      string    `gorm:"type:varchar;not_null" json:"path"`
	Image_url string    `gorm:"type:varchar;null" json:"image_url"`
}

func NewFile(id, entity_id uuid.UUID, path, image_url string) *File {
	return &File{
		Id:        id,
		Entity_id: entity_id,
		Path:      path,
		Image_url: image_url,
	}
}

// TableName specifies table name for ArticleModel.
func (model *File) TableName() string {
	return FileTableName
}
