package entity

import "github.com/google/uuid"

const (
	SuperAdminTableName = "superAdmin"
)

type SuperAdmin struct {
	ID       uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	Nama     string    `gorm:"type:varchar;not_null" json:"nama"`
	Email    string    `gorm:"type:varchar;not_null;unique" json:"email"`
	Password string    `gorm:"type:varchar;not_null" json:"password"`
}

func NewSuperAdmin(id uuid.UUID, name, email, password string) *SuperAdmin {
	return &SuperAdmin{
		ID:       id,
		Nama:     name,
		Email:    email,
		Password: password,
	}
}

func (model *SuperAdmin) TableName() string {
	return SuperAdminTableName
}
