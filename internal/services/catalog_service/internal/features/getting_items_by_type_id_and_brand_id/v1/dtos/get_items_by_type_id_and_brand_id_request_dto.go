package dtos

import "github.com/antoine2116/go-eshop-on-containers/internal/pkg/utils"

type GetItemsByTypeIdAndBrandItRequestDto struct {
	*utils.PaginationQuery `biding:"-"`
	TypeId                 int `uri:"typeId" biding:"required"`
	BrandId                int `uri:"brandId"`
}
