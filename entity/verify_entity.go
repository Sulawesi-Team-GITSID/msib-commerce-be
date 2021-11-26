package entity

import (
	"time"

	"github.com/google/uuid"
)

const (
	VerificationTableName = "verification"
)

// VerificationModel is a model for entity.Verification
type Verification struct {
	Id            uuid.UUID   `gorm:"type:uuid;primary_key" json:"id"`
	Credential_id uuid.UUID   `gorm:"type:uuid;not_null" json:"credential_id"`
	Code          string      `gorm:"type:varchar;not_null;unique" json:"code"`
	Expiresat     time.Time   `gorm:"type:date;not_null" json:"expiresat"`
	Credential    *Credential `gorm:"foreignKey:Credential_id"`
}

type Getcode struct {
	Credential_id uuid.UUID `gorm:"type:uuid;not_null" json:"credential_id"`
	Code          string    `gorm:"type:varchar;not_null;unique" json:"code"`
	Expiresat     time.Time `gorm:"type:date;not_null" json:"expiresat"`
	Email         string    `gorm:"type:varchar;not_null;unique" json:"email"`
}

func NewVerification(id, credential_id uuid.UUID, code string, expiresat time.Time) *Verification {
	return &Verification{
		Id:            id,
		Credential_id: credential_id,
		Code:          code,
		Expiresat:     expiresat,
	}
}

// TableName specifies table name for VerificationModel.
func (model *Verification) TableName() string {
	return VerificationTableName
}
