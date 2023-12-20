package dtos

import "github.com/antoine2116/go-eshop-on-containers/internal/pkg/utils"

type GetItemsWithNameRequestDto struct {
	*utils.PaginationQuery `biding:"-"`
	Name                   string `uri:"name" biding:"required"`
}
