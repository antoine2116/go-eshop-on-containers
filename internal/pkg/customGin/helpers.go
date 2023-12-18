package customGin

// Source : https://github.com/mehdihadeli/go-ecommerce-microservices/blob/719f1f5ce1ba079bd5323c05e89739c31ba62ff2/internal/pkg/web/route/helpers.go

import (
	"fmt"
	"go.uber.org/fx"
)

// when we register multiple handlers with output type `echo.HandlerFunc` we get exception `type already provided`, so we should use tags annotation

// AsRoute annotates the given constructor to state that
// it provides a route to the "routes" group.
func AsRoute(handler interface{}, routeGroupName string) interface{} {
	return fx.Annotate(
		handler,
		fx.As(new(Endpoint)),
		fx.ResultTags(fmt.Sprintf(`group:"%s"`, routeGroupName)),
	)
}
