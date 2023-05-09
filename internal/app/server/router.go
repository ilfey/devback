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
			path:     "/ping",
			endpoint: v1.Ping(s.Config, s.Services.JWT),
		},
		{
			role:     GUEST,
			method:   get,
			path:     "/messages",
			endpoint: v1.ReadMessages(s.Store),
		},
		{
			role:     GUEST,
			method:   get,
			path:     "/contacts",
			endpoint: v1.GetContacts(s.Store),
		},
		{
			role:     GUEST,
			method:   get,
			path:     "/contacts/:id",
			endpoint: v1.GetContact(s.Store),
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

		// USER

		{
			role:     USER,
			method:   post,
			path:     "/delete",
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
			path:     "/message/:id",
			endpoint: userV1.DeleteMessage(s.Store, s.Services.JWT),
		},
		{
			role:     USER,
			method:   patch,
			path:     "/message/:id",
			endpoint: userV1.EditMessage(s.Store, s.Services.JWT),
		},

		// ADMIN

		{
			role:     ADMIN,
			method:   delete,
			path:     "/users/messages/:id",
			endpoint: adminV1.DeleteMessage(s.Store),
		},
		{
			role:     ADMIN,
			method:   delete,
			path:     "/users/messages/:id/permanently",
			endpoint: adminV1.DeleteMessagePermanently(s.Store),
		},
		{
			role:     ADMIN,
			method:   patch,
			path:     "/users/messages/:id",
			endpoint: adminV1.EditMessage(s.Store),
		},
		{
			role:     ADMIN,
			method:   post,
			path:     "/users/messages/:id/restore",
			endpoint: adminV1.RestoreMessage(s.Store),
		},
		{
			role:     ADMIN,
			method:   delete,
			path:     "/users/:username",
			endpoint: adminV1.DeleteAccount(s.Store),
		},
		{
			role:     ADMIN,
			method:   delete,
			path:     "/users/:username/permanently",
			endpoint: adminV1.DeleteAccountPermanently(s.Store),
		},
		{
			role:     ADMIN,
			method:   post,
			path:     "/users/:username/restore",
			endpoint: adminV1.RestoreAccount(s.Store),
		},
		{
			role:     ADMIN,
			method:   post,
			path:     "/links",
			endpoint: adminV1.CreateLink(s.Store),
		},
		{
			role:     ADMIN,
			method:   get,
			path:     "/links",
			endpoint: adminV1.GetLinks(s.Store),
		},
		{
			role:     ADMIN,
			method:   get,
			path:     "/links/:id",
			endpoint: adminV1.GetLink(s.Store),
		},
		{
			role:     ADMIN,
			method:   delete,
			path:     "/links/:id",
			endpoint: adminV1.DeleteLink(s.Store),
		},
		{
			role:     ADMIN,
			method:   delete,
			path:     "/links/:id/permanently",
			endpoint: adminV1.DeleteLinkPermanently(s.Store),
		},
		{
			role:     ADMIN,
			method:   post,
			path:     "/links/:id/restore",
			endpoint: adminV1.RestoreLink(s.Store),
		},
		{
			role:     ADMIN,
			method:   post,
			path:     "/contacts",
			endpoint: adminV1.CreateContact(s.Store),
		},
	}

	return routes
}
