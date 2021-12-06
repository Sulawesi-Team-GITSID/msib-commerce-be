package entity

import "github.com/google/uuid"

const (
	CredentialTableName = "credential"
)

// CredentialModel is a model for entity.Credential
type Credential struct {
	Id        uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	Image_url string    `gorm:"type:varchar;null" json:"image_url"`
	Username  string    `gorm:"type:varchar;not_null;unique" json:"username"`
	Email     string    `gorm:"type:varchar;not_null;unique" json:"email"`
	Password  string    `gorm:"type:varchar;not_null" json:"password"`
	Seller    bool      `gorm:"type:bool;default:false;not_null" json:"seller"`
	Verified  bool      `gorm:"type:bool;default:false;not_null" json:"verified"`
}

type CredentialEmail struct {
	Id    uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	Email string    `gorm:"type:varchar;not_null;unique" json:"email"`
}

func NewCredential(id uuid.UUID, image_url, username, email, password string, seller, verified bool) *Credential {
	return &Credential{
		Id:        id,
		Image_url: image_url,
		Username:  username,
		Email:     email,
		Password:  password,
		Seller:    seller,
		Verified:  verified,
	}
}

func UpdatePasswordEntity(id uuid.UUID, password string) *Credential {
	return &Credential{
		Id:       id,
		Password: password,
	}
}

// TableName specifies table name for CredentialModel.
func (model *Credential) TableName() string {
	return CredentialTableName
}
