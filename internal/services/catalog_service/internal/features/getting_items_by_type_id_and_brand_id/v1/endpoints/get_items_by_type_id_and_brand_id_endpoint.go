package endpoints

import (
	"github.com/antoine2116/go-eshop-on-containers/internal/pkg/customGin"
	customErrors "github.com/antoine2116/go-eshop-on-containers/internal/pkg/http/custom_errors"
	"github.com/antoine2116/go-eshop-on-containers/internal/pkg/utils"
	"github.com/antoine2116/go-eshop-on-containers/internal/services/catalogservice/internal/contracts/params"
	dtosV1 "github.com/antoine2116/go-eshop-on-containers/internal/services/catalogservice/internal/dtos/v1"
	"github.com/antoine2116/go-eshop-on-containers/internal/services/catalogservice/internal/features/getting_items_by_type_id_and_brand_id/v1/dtos"
	"github.com/gin-gonic/gin"
	"net/http"
)

type getItemsByTypeIdAndBrandIdEndpoint struct {
	params.CatalogRouteParams
}

func NewGetItemsByTypeIdAndBrandIdEndpoint(params params.CatalogRouteParams) customGin.Endpoint {
	return &getItemsByTypeIdAndBrandIdEndpoint{
		CatalogRouteParams: params,
	}
}

func (ep *getItemsByTypeIdAndBrandIdEndpoint) MapEndpoint() {
	ep.CatalogGroup.GET("/items/type/:typeId/brand/:brandId", ep.handler())
}

func (ep *getItemsByTypeIdAndBrandIdEndpoint) handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		request := &dtos.GetItemsByTypeIdAndBrandIdRequestDto{}
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

		items, err := ep.ItemRepository.GetItemsByTypeIdAndBrandId(ctx, request.PaginationQuery, request.TypeId, request.BrandId)
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

		c.JSON(http.StatusOK, &dtos.GetItemsByTypeIdAndBrandIdResponseDto{Items: paginationResultDto})
	}
}
