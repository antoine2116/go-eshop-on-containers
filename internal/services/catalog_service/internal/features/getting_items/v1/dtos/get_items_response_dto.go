package dtos

import (
	"github.com/antoine2116/go-eshop-on-containers/internal/pkg/utils"
	dtosV1 "github.com/antoine2116/go-eshop-on-containers/internal/services/catalogservice/internal/dtos/v1"
)

type GetItemsResponseDto struct {
	Items *utils.PaginationResult[*dtosV1.ItemDto] `json:"items"`
}
