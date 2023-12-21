package dtos

import (
	dtosV1 "github.com/antoine2116/go-eshop-on-containers/internal/services/catalogservice/internal/dtos/v1"
)

type GetBrandsResponseDto struct {
	Brands []*dtosV1.BrandDto `json:"brands"`
}
