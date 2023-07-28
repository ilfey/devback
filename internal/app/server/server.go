package server

import (
	"github.com/gin-gonic/gin"
	"github.com/ilfey/devback/internal/app/config"
	"github.com/ilfey/devback/internal/app/middlewares"
	"github.com/ilfey/devback/internal/pkg/iservices"
	"github.com/ilfey/devback/internal/pkg/services"
	"github.com/ilfey/devback/internal/pkg/store"
	"github.com/sirupsen/logrus"
)

type Server struct {
	Config *config.Config
	Router *gin.Engine
	Logger *logrus.Logger
	Store  *store.Store

	Services struct {
		JWT iservices.JWT
	}

	groups struct {
		api   *gin.RouterGroup
		admin *gin.RouterGroup
		user  *gin.RouterGroup
	}
}

func NewServer(config *config.Config, logger *logrus.Logger, store *store.Store) *Server {
	s := new(Server)

	s.Config = config
	s.Router = gin.New()
	s.Logger = logger
	s.Store = store

	s.Build()

	return s
}

func (s *Server) Build() {

	// Create services
	s.Services.JWT = services.NewService(s.Config.Key, s.Config.LifeSpan)

	// Add cors middleware
	s.Router.Use(middlewares.CorsMiddleware())

	// Add logging middleware
	s.Router.Use(middlewares.LoggingMiddleWare(s.Logger))

	// Create groups
	// Api group - group with authorization middleware
	s.groups.api = s.Router.Group(s.Config.ApiPath, middlewares.AuthorizationMiddleware(s.Config.AdminUsername, s.Services.JWT, s.Logger))
	s.groups.admin = s.groups.api.Group("/", middlewares.AdminAuthMiddleware())
	s.groups.user = s.groups.api.Group("/", middlewares.JwtAuthMiddleware())

	// Load templates for public endpoints
	s.Router.LoadHTMLGlob("./web/template/*")

	// Load public endpoints
	routes := s.getRoutes()

	for _, route := range routes {
		s.Router.Handle(route.method, route.path, route.endpoint)

		routeLogger := s.Logger.WithFields(logrus.Fields{
			"Method": route.method,
			"Path":   route.path,
		})

		routeLogger.Info()
	}

	// Load api endpoints
	s.HandleApiRoutes(s.getApiRoutesV1, "v1")

	// Load api endpoints
	s.HandleApiRoutes(s.getApiRoutesV2, "v2")
}

func (s *Server) Run() error {
	s.Logger.Info("Server is live!")
	return s.Router.Run(s.Config.Addr)
}

func (s *Server) HandleApiRoutes(fn func() []*ServerRoute, v string) {
	routes := fn()
	
	s.Logger.Infof("Api %s", v)

	switch v {
	case "v1":
		for _, route := range routes {

			switch route.role {
			case V1_ADMIN:
				route.path = "/" + v + s.Config.AdminPath + route.path
				s.groups.admin.Handle(route.method, route.path, route.endpoint)
			case V1_USER:
				route.path = "/" + v + "/users" + route.path
				s.groups.user.Handle(route.method, route.path, route.endpoint)
			default:
				route.path = "/" + v + route.path
				s.groups.api.Handle(route.method, route.path, route.endpoint)
			}

			routeLogger := s.Logger.WithFields(logrus.Fields{
				"Method": route.method,
				"Path":   s.Config.ApiPath + route.path,
			})

			routeLogger.Info()
		}
	case "v2":
		for _, route := range routes {

			route.path = "/" + v + string(route.section) + route.path
			s.groups.api.Handle(route.method, route.path, route.endpoint)

			routeLogger := s.Logger.WithFields(logrus.Fields{
				"Method": route.method,
				"Path":   s.Config.ApiPath + route.path,
			})

			routeLogger.Info()
		}
	}

}
