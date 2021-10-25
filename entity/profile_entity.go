package entity

import (
	"time"

	"github.com/google/uuid"
)

const (
	ProfileTableName = "credential"
)

// ArticleModel is a model for entity.Article
type Profile struct {
	Id            uuid.UUID   `gorm:"type:uuid;primary_key" json:"id"`
	Credential_id uuid.UUID   `gorm:"type:uuid;not_null" json:"credential_id"`
	First_name    string      `gorm:"type:varchar;not_null" json:"first_name"`
	Last_name     string      `gorm:"type:varchar;not_null" json:"last_name"`
	Phone         string      `gorm:"type:varchar;not_null" json:"phone"`
	Gender        string      `gorm:"type:varchar;not_null" json:"gender"`
	Birthday      time.Time   `gorm:"type:date;not_null" json:"birthday"`
	Credential    *Credential `gorm:"foreignKey:Credential_id"`
}

func NewProfile(id, credential_id uuid.UUID, first_name, last_name, phone, gender string, birthday time.Time) *Profile {
	return &Profile{
		Id:            id,
		Credential_id: credential_id,
		First_name:    first_name,
		Last_name:     last_name,
		Phone:         phone,
		Gender:        gender,
		Birthday:      birthday,
	}
}

// TableName specifies table name for ArticleModel.
func (model *Profile) TableName() string {
	return ProfileTableName
}

// func (tv *TV) GenerateSlug() string {
// 	return html.EscapeString(strings.ToLower(strings.ReplaceAll(tv.title, " ", "-")))
// }
