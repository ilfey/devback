package server

import (
	"github.com/gin-gonic/gin"
	"github.com/ilfey/devback/internal/app/config"
	"github.com/ilfey/devback/internal/app/middlewares"
	"github.com/ilfey/devback/internal/pkg/services"
	"github.com/ilfey/devback/internal/pkg/iservices"
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

	s.groups.admin = s.Router.Group(config.AdminPath, middlewares.AdminAuthMiddleware(s.Services.JWT, s.Config))
	s.groups.user = s.Router.Group("/", middlewares.JwtAuthMiddleware(s.Services.JWT))

	s.Build()

	return s
}

func (s *Server) Build() {
	routes := s.getRoutes()

	s.Router.LoadHTMLGlob("./web/template/*")

	for _, route := range routes {
		s.HandleRoute(route)

		routeLogger := s.Logger.WithFields(logrus.Fields{
			"Method": route.method,
			"Path":   route.path,
		})

		routeLogger.Info()
	}
}

func (s *Server) Run() error {
	return s.Router.Run(s.Config.Addr)
}

func (s *Server) HandleRoute(r ServerRoute) gin.IRoutes {
	switch r.role {
	case ADMIN:
		return s.groups.admin.Handle(r.method, r.path, r.endpoint)
	case USER:
		return s.groups.user.Handle(r.method, r.path, r.endpoint)
	default:
		return s.Router.Handle(r.method, r.path, r.endpoint)
	}
}
