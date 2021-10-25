package entity

import "github.com/google/uuid"

const (
	CredentialTableName = "credential"
)

// ArticleModel is a model for entity.Article
type Credential struct {
	Id       uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	Username string    `gorm:"type:varchar;not_null" json:"username"`
	Password string    `gorm:"type:varchar;not_null" json:"password"`
	Seller   bool      `gorm:"type:bool;default:false;not_null" json:"seller"`
}

func NewCredential(id uuid.UUID, username, password string, seller bool) *Credential {
	return &Credential{
		Id:       id,
		Username: username,
		Password: password,
		Seller:   seller,
	}
}
func NewSeller(seller bool) *Credential {
	return &Credential{
		Seller: seller,
	}
}

// TableName specifies table name for ArticleModel.
func (model *Credential) TableName() string {
	return CredentialTableName
}

// func (tv *TV) GenerateSlug() string {
// 	return html.EscapeString(strings.ToLower(strings.ReplaceAll(tv.title, " ", "-")))
// }
