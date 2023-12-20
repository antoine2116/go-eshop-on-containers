package endpoints

import (
	"fmt"
	"github.com/antoine2116/go-eshop-on-containers/internal/pkg/customGin"
	customErrors "github.com/antoine2116/go-eshop-on-containers/internal/pkg/http/custom_errors"
	"github.com/antoine2116/go-eshop-on-containers/internal/pkg/mapper"
	"github.com/antoine2116/go-eshop-on-containers/internal/services/catalogservice/internal/contracts/params"
	dtosV1 "github.com/antoine2116/go-eshop-on-containers/internal/services/catalogservice/internal/dtos/v1"
	"github.com/antoine2116/go-eshop-on-containers/internal/services/catalogservice/internal/features/getting_item_by_id/v1/dtos"
	"github.com/gin-gonic/gin"
	"net/http"
)

type getItemByIdEndpoint struct {
	params.CatalogRouteParams
}

func NewGetItemByIdEndpoint(params params.CatalogRouteParams) customGin.Endpoint {
	return &getItemByIdEndpoint{
		CatalogRouteParams: params,
	}
}

func (ep *getItemByIdEndpoint) MapEndpoint() {
	ep.CatalogGroup.GET("/items/:id", ep.handler())
}

func (ep *getItemByIdEndpoint) handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		request := &dtos.GetItemByIdRequestDto{}
		if err := c.BindUri(request); err != nil {
			badRequestErr := customErrors.NewBadRequestError(
				"Error in getting data from path parameter.",
				err,
			)
			c.Error(badRequestErr)
			return
		}

		item, err := ep.ItemRepository.GetItemById(ctx, request.ItemId)
		if err != nil {
			internalServerErr := customErrors.NewInternalServerError(err)
			c.Error(internalServerErr)
			return
		}

		if item == nil {
			notFoundErr := customErrors.NewNotFoundError(
				fmt.Sprintf("Can't find the item with id %d into the database.", request.ItemId),
				err,
			)
			c.Error(notFoundErr)
			return
		}

		itemDto, err := mapper.Map[*dtosV1.ItemDto](item)
		if err != nil {
			internalServerErr := customErrors.NewInternalServerError(err)
			c.Error(internalServerErr)
			return
		}

		c.JSON(http.StatusOK, &dtos.GetItemByIdResponseDto{Item: itemDto})
	}
}
