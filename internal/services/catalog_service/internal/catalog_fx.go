package catalog

import (
	"github.com/antoine2116/go-eshop-on-containers/internal/pkg/customGin"
	gormPostgres "github.com/antoine2116/go-eshop-on-containers/internal/pkg/gorm"
	zapLogger "github.com/antoine2116/go-eshop-on-containers/internal/pkg/logger/zap_logger"
	"github.com/antoine2116/go-eshop-on-containers/internal/pkg/migration"
	"github.com/antoine2116/go-eshop-on-containers/internal/services/catalogservice/config"
	"github.com/antoine2116/go-eshop-on-containers/internal/services/catalogservice/internal/data/repositories"
	getItemByIdV1 "github.com/antoine2116/go-eshop-on-containers/internal/services/catalogservice/internal/features/getting_item_by_id/v1/endpoints"
	getItemsV1 "github.com/antoine2116/go-eshop-on-containers/internal/services/catalogservice/internal/features/getting_items/v1/endpoints"
	getItemsWithNameV1 "github.com/antoine2116/go-eshop-on-containers/internal/services/catalogservice/internal/features/getting_items_with_name/v1/endpoints"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"catalogfx",
	config.Module,
	zapLogger.Module,
	gormPostgres.Module,
	customGin.Module,
	migration.Module,

	fx.Provide(repositories.NewItemRepository),

	fx.Provide(func(s customGin.GinHttpServer) *gin.RouterGroup {
		var g *gin.RouterGroup
		s.RouteBuilder().RegisterGroup("/api/v1", func(v1 *gin.RouterGroup) {
			group := v1.Group("/catalog")
			g = group
		})

		return g
	}),

	fx.Provide(
		customGin.AsRoute(getItemsV1.NewGetItemsEndpoint, "catalog-route"),
		customGin.AsRoute(getItemByIdV1.NewGetItemByIdEndpoint, "catalog-route"),
		customGin.AsRoute(getItemsWithNameV1.NewGetItemWithNameEndpoint, "catalog-route"),
	),
)
