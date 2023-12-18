package params

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type CatalogRouteParams struct {
	fx.In

	CatalogGroup *gin.RouterGroup
}
