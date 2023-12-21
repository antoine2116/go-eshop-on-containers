package endpoints

import (
	"github.com/antoine2116/go-eshop-on-containers/internal/pkg/customGin"
	customErrors "github.com/antoine2116/go-eshop-on-containers/internal/pkg/http/custom_errors"
	"github.com/antoine2116/go-eshop-on-containers/internal/pkg/mapper"
	"github.com/antoine2116/go-eshop-on-containers/internal/services/catalogservice/internal/contracts/params"
	dtosV1 "github.com/antoine2116/go-eshop-on-containers/internal/services/catalogservice/internal/dtos/v1"
	"github.com/antoine2116/go-eshop-on-containers/internal/services/catalogservice/internal/features/getting_brands/v1/dtos"
	"github.com/gin-gonic/gin"
	"net/http"
)

type getBrandsEndpoint struct {
	params.CatalogRouteParams
}

func NewGetBrandsEndpoint(params params.CatalogRouteParams) customGin.Endpoint {
	return &getBrandsEndpoint{
		CatalogRouteParams: params,
	}
}

func (ep *getBrandsEndpoint) MapEndpoint() {
	ep.CatalogGroup.GET("/brands", ep.handler())
}

func (ep *getBrandsEndpoint) handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		brands, err := ep.Repository.GetAllBrands(ctx)
		if err != nil {
			internalServerErr := customErrors.NewInternalServerError(err)
			c.Error(internalServerErr)
			return
		}

		brandsDto, err := mapper.Map[[]*dtosV1.BrandDto](brands)
		if err != nil {
			internalServerErr := customErrors.NewInternalServerError(err)
			c.Error(internalServerErr)
			return
		}

		c.JSON(http.StatusOK, &dtos.GetBrandsResponseDto{Brands: brandsDto})
	}
}
