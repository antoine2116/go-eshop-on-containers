package models

import (
	"time"
)

type Item struct {
	Id                int       `json:"id" gorm:"primaryKey"`
	Name              string    `json:"name"`
	Description       string    `json:"description"`
	Price             float64   `json:"price"`
	PictureFileName   string    `json:"pictureFileName"`
	PictureUri        string    `json:"pictureUri"`
	TypeId            int       `json:"TypeId"`
	Type              Type      `json:"Type" gorm:"foreignKey:TypeId"`
	BrandId           int       `json:"BrandId"`
	Brand             Brand     `json:"Brand" gorm:"foreignKey:BrandId"`
	AvailableStock    int       `json:"availableStock"`
	RestockThreshold  int       `json:"restockThreshold"`
	MaxStockThreshold int       `json:"maxStockThreshold"`
	OnReorder         bool      `json:"onReorder"`
	CreatedAt         time.Time `json:"createdAt"` // https://gorm.io/docs/models.html#gorm-Model
	UpdatedAt         time.Time `json:"updatedAt"` // https://gorm.io/docs/models.html#gorm-Model
}
