package model

import (
	"time"

	"gorm.io/gorm"
)

type Base struct {
	Uuid      string         `json:"uuid" gorm:"index;unique"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
