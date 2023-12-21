package params

import (
	"github.com/antoine2116/go-eshop-on-containers/internal/services/catalogservice/internal/data/repositories"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type CatalogRouteParams struct {
	fx.In

	CatalogGroup *gin.RouterGroup
	Logger       *zap.Logger
	Repository   repositories.CatalogRepository
}
