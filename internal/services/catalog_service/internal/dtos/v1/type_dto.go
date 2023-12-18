package dtosV1

import "time"

type TypeDto struct {
	Id        int       `json:"id"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
