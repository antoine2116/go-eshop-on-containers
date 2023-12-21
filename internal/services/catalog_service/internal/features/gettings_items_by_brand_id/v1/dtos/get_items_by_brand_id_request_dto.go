package dtos

import "github.com/antoine2116/go-eshop-on-containers/internal/pkg/utils"

type GetItemsByBrandIdRequestDto struct {
	*utils.PaginationQuery `binding:"-"`
	BrandId                int `uri:"brandId" binding:"required"`
}
