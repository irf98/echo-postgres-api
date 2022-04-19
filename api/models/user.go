package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        int32          `json:"id" gorm:"primaryKey;type:not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Email     string         `json:"email" gorm:"type:not null"`
	Password  []byte         `json:"password" gorm:"type:not null"`
}
