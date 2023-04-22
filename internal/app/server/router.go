package server

import (
	"github.com/gin-gonic/gin"
	"github.com/ilfey/devback/internal/app/endpoints"
)

const (
	get  = "GET"
	post = "POST"

	GUEST = "GUEST"
	USER  = "USER"
	ADMIN = "ADMIN"
)

type ServerRoute struct {
	role     string
	method   string
	path     string
	endpoint gin.HandlerFunc
}

func (s *Server) getRoutes() (routes []ServerRoute) {
	routes = []ServerRoute{
		{
			role:     GUEST,
			method:   get,
			path:     "/",
			endpoint: endpoints.Index(),
		},
		{
			role:     GUEST,
			method:   post,
			path:     "/user/login",
			endpoint: endpoints.Login(s.Store, s.Services.JWT),
		},
		{
			role:     GUEST,
			method:   post,
			path:     "/user/register",
			endpoint: endpoints.Register(s.Store),
		},
		{
			role:     GUEST,
			method:   get,
			path:     "/messages",
			endpoint: endpoints.ReadMessages(s.Store),
		},
		{
			role:     USER,
			method:   post,
			path:     "/message",
			endpoint: endpoints.CreateMessage(s.Store, s.Services.JWT),
		},
	}

	return routes
}
