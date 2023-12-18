package endpoints

import (
	"github.com/antoine2116/go-eshop-on-containers/internal/pkg/customGin"
	"github.com/antoine2116/go-eshop-on-containers/internal/pkg/mapper"
	"github.com/antoine2116/go-eshop-on-containers/internal/services/catalogservice/internal/contracts/params"
	"github.com/antoine2116/go-eshop-on-containers/internal/services/catalogservice/internal/data/repositories"
	dtosV1 "github.com/antoine2116/go-eshop-on-containers/internal/services/catalogservice/internal/dtos/v1"
	"github.com/antoine2116/go-eshop-on-containers/internal/services/catalogservice/internal/features/getting_item_by_id/v1/dtos"
	"github.com/gin-gonic/gin"
	"net/http"
)

type getItemByIdEndpoint struct {
	params params.CatalogRouteParams
	repo   repositories.ItemRepository
}

func NewGetItemByIdEndpoint(params params.CatalogRouteParams, repo repositories.ItemRepository) customGin.Endpoint {
	return &getItemByIdEndpoint{
		params: params,
		repo:   repo,
	}
}

func (ep *getItemByIdEndpoint) MapEndpoint() {
	ep.params.CatalogGroup.GET("/items/:id", ep.handler())
}

func (ep *getItemByIdEndpoint) handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		request := &dtos.GetItemByIdRequestDto{}
		if err := c.BindUri(request); err != nil {
			c.JSON(http.StatusBadRequest, "Error in the binding request")
		}

		item, err := ep.repo.GetItemById(ctx, request.ItemId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, "Error in requesting the database")
			return
		}

		itemDto, err := mapper.Map[*dtosV1.ItemDto](item)
		if err != nil {
			c.JSON(http.StatusInternalServerError, "Error in mapping to dto")
			return
		}

		c.JSON(http.StatusOK, &dtos.GetItemByIdResponseDto{Item: itemDto})
	}
}
