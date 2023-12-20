package endpoints

import (
	"github.com/antoine2116/go-eshop-on-containers/internal/pkg/customGin"
	customErrors "github.com/antoine2116/go-eshop-on-containers/internal/pkg/http/custom_errors"
	"github.com/antoine2116/go-eshop-on-containers/internal/pkg/utils"
	"github.com/antoine2116/go-eshop-on-containers/internal/services/catalogservice/internal/contracts/params"
	dtosV1 "github.com/antoine2116/go-eshop-on-containers/internal/services/catalogservice/internal/dtos/v1"
	"github.com/antoine2116/go-eshop-on-containers/internal/services/catalogservice/internal/features/getting_items_with_name/v1/dtos"
	"github.com/gin-gonic/gin"
	"net/http"
)

type getItemWithNameEndpoint struct {
	params.CatalogRouteParams
}

func NewGetItemWithNameEndpoint(params params.CatalogRouteParams) customGin.Endpoint {
	return &getItemWithNameEndpoint{
		CatalogRouteParams: params,
	}
}

func (ep *getItemWithNameEndpoint) MapEndpoint() {
	ep.CatalogGroup.GET("/items/withname/:name", ep.handler())
}

func (ep *getItemWithNameEndpoint) handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		request := &dtos.GetItemsWithNameRequestDto{}
		if err := c.BindUri(request); err != nil {
			badRequestErr := customErrors.NewBadRequestError(
				"Error in getting data from path parameter.",
				err,
			)
			c.Error(badRequestErr)
			return
		}

		paginationQuery, err := utils.GetPaginationQueryFromCtx(c)
		if err != nil {
			badRequestErr := customErrors.NewBadRequestError(
				"Error in getting data from query string.",
				err,
			)
			c.Error(badRequestErr)
			return
		}

		request.PaginationQuery = paginationQuery

		items, err := ep.ItemRepository.GetItemsWithName(ctx, request.PaginationQuery, request.Name)
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

		c.JSON(http.StatusOK, &dtos.GetItemsWithNameResponseDto{Items: paginationResultDto})
	}
}
