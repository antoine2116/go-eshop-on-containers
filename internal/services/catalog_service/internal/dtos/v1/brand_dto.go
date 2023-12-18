package dtosV1

import "time"

type BrandDto struct {
	Id        int       `json:"id"`
	Brand     string    `json:"brand"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
