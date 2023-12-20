package endpoints

import (
	"github.com/antoine2116/go-eshop-on-containers/internal/pkg/customGin"
	customErrors "github.com/antoine2116/go-eshop-on-containers/internal/pkg/http/custom_errors"
	"github.com/antoine2116/go-eshop-on-containers/internal/pkg/utils"
	"github.com/antoine2116/go-eshop-on-containers/internal/services/catalogservice/internal/contracts/params"
	dtosV1 "github.com/antoine2116/go-eshop-on-containers/internal/services/catalogservice/internal/dtos/v1"
	"github.com/antoine2116/go-eshop-on-containers/internal/services/catalogservice/internal/features/getting_items/v1/dtos"
	"github.com/gin-gonic/gin"
	"net/http"
)

type getItemsEndpoint struct {
	params.CatalogRouteParams
}

func NewGetItemsEndpoint(params params.CatalogRouteParams) customGin.Endpoint {
	return &getItemsEndpoint{
		CatalogRouteParams: params,
	}
}

func (ep *getItemsEndpoint) MapEndpoint() {
	ep.CatalogGroup.GET("/items", ep.handler())
}

func (ep *getItemsEndpoint) handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		paginationQuery, err := utils.GetPaginationQueryFromCtx(c)
		if err != nil {
			badRequestErr := customErrors.NewBadRequestError(
				"Error in getting data from query string.",
				err,
			)
			c.Error(badRequestErr)
			return
		}

		query := dtos.GetItemsRequestDto{PaginationQuery: paginationQuery}

		items, err := ep.ItemRepository.GetAllItems(ctx, query.PaginationQuery)
		if err != nil {
			internalServerErr := customErrors.NewInternalServerError(err)
			c.Error(internalServerErr)
			return
		}

		paginationResultDto, err := utils.PaginationResultToPaginationResultDto[*dtosV1.ItemDto](items)
		if err != nil {
			internalServerErr := customErrors.NewInternalServerError(err)
			c.Error(internalServerErr)
			return
		}

		c.JSON(http.StatusOK, &dtos.GetItemsResponseDto{Items: paginationResultDto})
	}
}
