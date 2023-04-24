package server

import (
	"github.com/gin-gonic/gin"

	"github.com/ilfey/devback/internal/app/endpoints"
	v1 "github.com/ilfey/devback/internal/app/endpoints/v1"
	adminV1 "github.com/ilfey/devback/internal/app/endpoints/v1/admin"
	userV1 "github.com/ilfey/devback/internal/app/endpoints/v1/user"
)

const (
	get    = "GET"
	post   = "POST"
	delete = "DELETE"
	patch  = "PATCH"

	GUEST = "GUEST"
	USER  = "USER"
	ADMIN = "ADMIN"
)

type ServerRoute struct {
	role     string // GUEST (default) - /api/vx/*, USER - /api/vx/user/*, ADMIN - /api/vx/<admin path>/*
	method   string
	path     string
	endpoint gin.HandlerFunc
}

// Handlers with path /*. Can not have role
func (s *Server) getRoutes() (routes []ServerRoute) {
	routes = []ServerRoute{
		{
			method:   get,
			path:     "/",
			endpoint: endpoints.Index(),
		},
	}

	return routes
}

// Api v1 handlers
func (s *Server) getApiRoutesV1() (routes []ServerRoute) {
	routes = []ServerRoute{
		{
			role:     GUEST,
			method:   get,
			path:     "/messages",
			endpoint: v1.ReadMessages(s.Store),
		},
		{
			role:     GUEST,
			method:   post,
			path:     "/user/login",
			endpoint: userV1.Login(s.Store, s.Services.JWT),
		},
		{
			role:     GUEST,
			method:   post,
			path:     "/user/register",
			endpoint: userV1.Register(s.Store),
		},
		{
			role:     GUEST,
			method:   post,
			path:     "/user/delete",
			endpoint: userV1.DeleteAccount(s.Store, s.Services.JWT),
		},
		{
			role:     USER,
			method:   post,
			path:     "/message",
			endpoint: userV1.CreateMessage(s.Store, s.Services.JWT),
		},
		{
			role:     USER,
			method:   delete,
			path:     "/message",
			endpoint: userV1.DeleteMessage(s.Store, s.Services.JWT),
		},
		{
			role:     USER,
			method:   patch,
			path:     "/message",
			endpoint: userV1.EditMessage(s.Store, s.Services.JWT),
		},
		{
			role:     ADMIN,
			method:   delete,
			path:     "/user/message",
			endpoint: adminV1.DeleteMessage(s.Store),
		},
		{
			role:     ADMIN,
			method:   patch,
			path:     "/user/message",
			endpoint: adminV1.EditMessage(s.Store),
		},
		{
			role:     ADMIN,
			method:   post,
			path:     "/user/message/restore",
			endpoint: adminV1.RestoreMessage(s.Store),
		},
		{
			role:     ADMIN,
			method:   delete,
			path:     "/user",
			endpoint: adminV1.DeleteAccount(s.Store),
		},
		{
			role:     ADMIN,
			method:   post,
			path:     "/user/restore",
			endpoint: adminV1.RestoreAccount(s.Store),
		},
	}

	return routes
}
