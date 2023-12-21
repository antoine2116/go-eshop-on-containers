package endpoints

import (
	"github.com/antoine2116/go-eshop-on-containers/internal/pkg/customGin"
	customErrors "github.com/antoine2116/go-eshop-on-containers/internal/pkg/http/custom_errors"
	"github.com/antoine2116/go-eshop-on-containers/internal/pkg/mapper"
	"github.com/antoine2116/go-eshop-on-containers/internal/services/catalogservice/internal/contracts/params"
	dtosV1 "github.com/antoine2116/go-eshop-on-containers/internal/services/catalogservice/internal/dtos/v1"
	"github.com/antoine2116/go-eshop-on-containers/internal/services/catalogservice/internal/features/getting_types/v1/dtos"
	"github.com/gin-gonic/gin"
	"net/http"
)

type getTypesEndpoint struct {
	params.CatalogRouteParams
}

func NewGetTypesEndpoint(params params.CatalogRouteParams) customGin.Endpoint {
	return &getTypesEndpoint{
		CatalogRouteParams: params,
	}
}

func (ep *getTypesEndpoint) MapEndpoint() {
	ep.CatalogGroup.GET("/types", ep.handler())
}

func (ep *getTypesEndpoint) handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		types, err := ep.Repository.GetAllTypes(ctx)
		if err != nil {
			internalServerErr := customErrors.NewInternalServerError(err)
			c.Error(internalServerErr)
			return
		}

		typesDto, err := mapper.Map[[]*dtosV1.TypeDto](types)
		if err != nil {
			internalServerErr := customErrors.NewInternalServerError(err)
			c.Error(internalServerErr)
			return
		}

		c.JSON(http.StatusOK, &dtos.GetTypesResponseDto{Types: typesDto})
	}
}
