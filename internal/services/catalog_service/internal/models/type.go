package models

import (
	"time"
)

type Type struct {
	Id        int       `json:"id" gorm:"primaryKey"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"createdAt"` // https://gorm.io/docs/models.html#gorm-Model
	UpdatedAt time.Time `json:"updatedAt"` // https://gorm.io/docs/models.html#gorm-Model
}
