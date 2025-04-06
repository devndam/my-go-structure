package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID              uint           `json:"id" gorm:"primaryKey"`
	UUID            uuid.UUID      `json:"uuid" gorm:"type:char(36);uniqueIndex"`
	Email           string         `json:"email" gorm:"uniqueIndex;not null"`
	FirstName       string         `json:"first_name"`
	LastName        string         `json:"last_name"`
	EmailVerifiedAt *time.Time     `json:"email_verified_at"` // pointer to allow null
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`
}

// BeforeCreate GORM hook to set UUID before creating a new user
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.UUID = uuid.New()
	return
}
