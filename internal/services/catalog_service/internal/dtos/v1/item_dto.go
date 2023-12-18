package dtosV1

import "time"

type ItemDto struct {
	Id                int       `json:"id"`
	Name              string    `json:"name"`
	Description       string    `json:"description"`
	Price             float64   `json:"price"`
	PictureFileName   string    `json:"pictureFileName"`
	PictureUri        string    `json:"pictureUri"`
	CatalogType       TypeDto   `json:"catalogType"`
	CatalogBrand      BrandDto  `json:"catalogBrand"`
	AvailableStock    int       `json:"availableStock"`
	RestockThreshold  int       `json:"restockThreshold"`
	MaxStockThreshold int       `json:"maxStockThreshold"`
	OnReorder         bool      `json:"onReorder"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
}
