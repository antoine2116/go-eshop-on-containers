package customGin

import "github.com/gin-gonic/gin"

type RouteBuilder struct {
	gin *gin.Engine
}

func NewRouteBuilder(gin *gin.Engine) *RouteBuilder {
	return &RouteBuilder{gin: gin}
}

func (r *RouteBuilder) RegisterRoute(builder func(g *gin.Engine)) *RouteBuilder {
	builder(r.gin)
	return r
}

func (r *RouteBuilder) RegisterGroup(groupName string, builder func(g *gin.RouterGroup)) *RouteBuilder {
	builder(r.gin.Group(groupName))
	return r
}
