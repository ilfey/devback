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

	s.Services.JWT = services.NewService(config.Key, config.LifeSpan)

	s.groups.api = s.Router.Group(config.ApiPath)
	s.groups.admin = s.groups.api.Group("/", middlewares.AdminAuthMiddleware(s.Services.JWT, s.Config))
	s.groups.user = s.groups.api.Group("/", middlewares.JwtAuthMiddleware(s.Services.JWT))

	s.Build()

	return s
}

func (s *Server) Build() {
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
}

func (s *Server) Run() error {
	return s.Router.Run(s.Config.Addr)
}

func (s *Server) HandleApiRoutes(fn func() []ServerRoute, v string) {
	routes := fn()
	for _, route := range routes {

		switch route.role {
		case ADMIN:
			route.path = "/" + v + s.Config.AdminPath + route.path
			s.groups.admin.Handle(route.method, route.path, route.endpoint)
		case USER:
			route.path = "/" + v + "/user" + route.path
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
}

// TODO: not used
func (s *Server) HandleApiRoute(r ServerRoute) gin.IRoutes {
	switch r.role {
	case ADMIN:
		return s.groups.admin.Handle(r.method, s.Config.AdminPath+r.path, r.endpoint)
	case USER:
		return s.groups.user.Handle(r.method, r.path, r.endpoint)
	default:
		return s.groups.api.Handle(r.method, r.path, r.endpoint)
	}
}
