package endpoints

import (
	"github.com/antoine2116/go-eshop-on-containers/internal/pkg/customGin"
	"github.com/antoine2116/go-eshop-on-containers/internal/pkg/utils"
	"github.com/antoine2116/go-eshop-on-containers/internal/services/catalogservice/internal/contracts/params"
	"github.com/antoine2116/go-eshop-on-containers/internal/services/catalogservice/internal/data/repositories"
	dtosV1 "github.com/antoine2116/go-eshop-on-containers/internal/services/catalogservice/internal/dtos/v1"
	"github.com/antoine2116/go-eshop-on-containers/internal/services/catalogservice/internal/features/getting_items/v1/dtos"
	"github.com/gin-gonic/gin"
	"net/http"
)

type getItemsEndpoint struct {
	params params.CatalogRouteParams
	repo   repositories.ItemRepository
}

func NewGetItemsEndpoint(params params.CatalogRouteParams, repo repositories.ItemRepository) customGin.Endpoint {
	return &getItemsEndpoint{
		params: params,
		repo:   repo,
	}
}

func (ep *getItemsEndpoint) MapEndpoint() {
	ep.params.CatalogGroup.GET("/items", ep.handler())
}

func (ep *getItemsEndpoint) handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		paginationQuery, err := utils.GetPaginationQueryFromCtx(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, "Error in getting data from query string")
			return
		}

		query := dtos.GetItemsRequestDto{PaginationQuery: paginationQuery}

		items, err := ep.repo.GetAllItems(ctx, query.PaginationQuery)
		if err != nil {
			c.JSON(http.StatusInternalServerError, "Error in requesting the database")
			return
		}

		paginationResultDto, err := utils.PaginationResultToPaginationResultDto[*dtosV1.ItemDto](items)
		if err != nil {
			c.JSON(http.StatusInternalServerError, "Error in mapping to dto")
			return
		}

		c.JSON(http.StatusOK, &dtos.GetItemsResponseDto{Items: paginationResultDto})
	}
}
