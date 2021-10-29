package entity

import "github.com/google/uuid"

const (
	UsersTableName = "users"
)

type Users struct {
	ID       uuid.UUID `gorm:"primary_key;auto_increment" json:"id"`
	Nama     string    `gorm:"not null" json:"nama"`
	Email    string    `gorm:"not null" json:"email"`
	Password string    `gorm:"not null" json:"password"`
}

func NewUsers(id uuid.UUID, name, email, password string) *Users {
	return &Users{
		ID:       id,
		Nama:     name,
		Email:    email,
		Password: password,
	}
}

func (model *Users) TableName() string {
	return UsersTableName
}
