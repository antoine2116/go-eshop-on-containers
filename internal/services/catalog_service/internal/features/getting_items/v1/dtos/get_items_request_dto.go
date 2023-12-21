package dtos

import "github.com/antoine2116/go-eshop-on-containers/internal/pkg/utils"

type GetItemsRequestDto struct {
	*utils.PaginationQuery `binding:"-"`
	Ids                    string `form:"ids"`
}
