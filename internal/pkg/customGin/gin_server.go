package customGin

import (
	"context"
	"github.com/antoine2116/go-eshop-on-containers/internal/pkg/customGin/config"
	"github.com/antoine2116/go-eshop-on-containers/internal/pkg/customGin/middlewares"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type GinHttpServer interface {
	Run() error
	GracefulShutdown(context context.Context) error
	Cfg() *config.GinHttpServerOptions
	RouteBuilder() *RouteBuilder
}

type ginHttpServer struct {
	router       *gin.Engine
	server       *http.Server
	config       *config.GinHttpServerOptions
	logger       *zap.Logger
	routeBuilder *RouteBuilder
}

func NewGinHttpServer(logger *zap.Logger, config *config.GinHttpServerOptions) GinHttpServer {
	router := gin.New()

	router.Use(middlewares.GinLogger(logger))
	router.Use(middlewares.ErrorHandler(logger))

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	server := &http.Server{
		Addr:    config.Port,
		Handler: router,
	}

	return &ginHttpServer{
		router:       router,
		server:       server,
		config:       config,
		logger:       logger,
		routeBuilder: NewRouteBuilder(router),
	}
}

func (s *ginHttpServer) Run() error {
	return s.server.ListenAndServe()
}

func (s *ginHttpServer) GracefulShutdown(ctx context.Context) error {
	if err := s.server.Shutdown(ctx); err != nil {
		return err
	}

	return nil
}

func (s *ginHttpServer) Cfg() *config.GinHttpServerOptions {
	return s.config
}

func (s *ginHttpServer) RouteBuilder() *RouteBuilder {
	return s.routeBuilder
}
