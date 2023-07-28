package server

import (
	"github.com/gin-gonic/gin"

	"github.com/ilfey/devback/internal/app/endpoints"

	v1 "github.com/ilfey/devback/internal/app/endpoints/v1"
	adminV1 "github.com/ilfey/devback/internal/app/endpoints/v1/admin"
	usersV1 "github.com/ilfey/devback/internal/app/endpoints/v1/users"

	linksV2 "github.com/ilfey/devback/internal/app/endpoints/v2/links"
)

type SectionType string

const (
	SectionLinks    SectionType = "/links"
	SectionUsers    SectionType = "/users"
	SectionMessages SectionType = "/messages"
)

const (
	get    = "GET"
	post   = "POST"
	delete = "DELETE"
	patch  = "PATCH"

	V1_GUEST = "GUEST"
	V1_USER  = "USER"
	V1_ADMIN = "ADMIN"
)

type ServerRoute struct {
	/*
		By default - GUEST

		GUEST		- /api/vx/*
		USER 		- /api/vx/user/*
		ADMIN 	- /api/vx/<admin path>/*
	*/
	role     string
	section  SectionType
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

// Api v2 handlers
func (s *Server) getApiRoutesV2() (routes []*ServerRoute) {
	return []*ServerRoute{
		// ##################################################################
		// #                             Links                              #
		// ##################################################################

		/*

			### Get links
			GET https://{{host}}/{{admin_path}}/links
			Authorization: Bearer {{token}}

		*/
		{
			section:  SectionLinks,
			method:   get,
			endpoint: linksV2.GetLinks(s.Store),
		},
		/*

			### Get link by id
			GET https://{{host}}/{{admin_path}}/links/{{link_id}}
			Authorization: Bearer {{token}}

		*/
		{
			section:  SectionLinks,
			method:   get,
			path:     "/:id",
			endpoint: linksV2.GetLink(s.Store),
		},
		/*

			### Delete link by id
			DELETE https://{{host}}/links/{{delete_link_id}}
			Authorization: Bearer {{token}}

		*/
		{
			section:  SectionLinks,
			method:   delete,
			path:     "/:id",
			endpoint: linksV2.DeleteLink(s.Store),
		},
		/*

			### Create link
			POST https://{{host}}/links
			Authorization: Bearer {{token}}

			{
				"url": "{{url}}"
				"description": "{{description}}"
			}

		*/
		{
			section:  SectionLinks,
			method:   post,
			endpoint: linksV2.CreateLink(s.Store),
		},
		/*

			### Restore link by id
			POST https://{{host}}/{{admin_path}}/links/{{restore_link_id}}/restore
			Authorization: Bearer {{token}}

		*/
		{
			section:  SectionLinks,
			method:   post,
			path:     "/:id",
			endpoint: linksV2.RestoreLink(s.Store),
		},

		// ##################################################################
		// #                             Users                              #
		// ##################################################################

	}
}

// Api v1 handlers
func (s *Server) getApiRoutesV1() (routes []*ServerRoute) {
	routes = []*ServerRoute{
		/*

			### Get subinfo about server: start_time, uptime, etc...
			GET https://{{host}}/ping

		*/
		{
			role:     V1_GUEST,
			method:   get,
			path:     "/ping",
			endpoint: v1.Ping(s.Config, s.Services.JWT),
		},
		/*

			### Get all messages
			GET https://{{host}}/messages

		*/
		{
			role:     V1_GUEST,
			method:   get,
			path:     "/messages",
			endpoint: v1.GetMessages(s.Store),
		},
		/*

			### Get message by id
			GET https://{{host}}/messages/{{message_id}}

		*/
		{
			role:     V1_GUEST,
			method:   get,
			path:     "/messages/:id",
			endpoint: v1.GetMessage(s.Store),
		},
		/*

			### Get all contacts
			GET https://{{host}}/contacts

		*/
		{
			role:     V1_GUEST,
			method:   get,
			path:     "/contacts",
			endpoint: v1.GetContacts(s.Store),
		},
		/*

			### Get contact by id
			GET https://{{host}}/contacts/{{contact_id}}

		*/
		{
			role:     V1_GUEST,
			method:   get,
			path:     "/contacts/:id",
			endpoint: v1.GetContact(s.Store),
		},
		/*

			### Login user
			POST https://{{host}}/login
			Content-Type: application/json

			{
				"username": "{{username}}",
				"password": "{{password}}"
			}

		*/
		{
			role:     V1_GUEST,
			method:   post,
			path:     "/login",
			endpoint: usersV1.Login(s.Store, s.Services.JWT),
		},
		/*

			### Register user
			POST https://{{host}}/register
			Content-Type: application/json

			{
				"username": "{{username}}",
				"password": "{{password}}"
			}

		*/
		{
			role:     V1_GUEST,
			method:   post,
			path:     "/register",
			endpoint: usersV1.Register(s.Store),
		},

		// ##################################################################
		// #                             Users                              #
		// ##################################################################

		/*

			### Delete current user
			DELETE https://{{host}}/users/delete
			Authorization: Bearer {{token}}

		*/
		{
			role:     V1_USER,
			method:   delete,
			path:     "/delete",
			endpoint: usersV1.DeleteAccount(s.Store, s.Services.JWT),
		},
		/*

			### Create message
			POST https://{{host}}/users/messages
			Authorization: Bearer {{token}}

			{
				"content": "{{message_content}}",
				"reply_ro": {{reply_message_id}}
			}

		*/
		{
			role:     V1_USER,
			method:   post,
			path:     "/messages",
			endpoint: usersV1.CreateMessage(s.Store, s.Services.JWT),
		},
		/*

			### Delete message by id
			DELETE https://{{host}}/users/messages/{{delete_message_id}}
			Authorization: Bearer {{token}}

		*/
		{
			role:     V1_USER,
			method:   delete,
			path:     "/messages/:id",
			endpoint: usersV1.DeleteMessage(s.Store, s.Services.JWT),
		},
		/*

			### Edit message by id
			PATCH https://{{host}}/users/messages/{{edit_message_id}}
			Authorization: Bearer {{token}}

			{
				"content": "{{new_message_content}}"
			}

		*/
		{
			role:     V1_USER,
			method:   patch,
			path:     "/messages/:id",
			endpoint: usersV1.EditMessage(s.Store, s.Services.JWT),
		},

		// ##################################################################
		// #                             Admin                              #
		// ##################################################################

		/*

			### Delete message by id
			DELETE https://{{host}}/{{admin_path}}/messages/{{delete_message_id}}
			Authorization: Bearer {{token}}

		*/
		{
			role:     V1_ADMIN,
			method:   delete,
			path:     "/users/messages/:id",
			endpoint: adminV1.DeleteMessage(s.Store),
		},
		/*

			### Delete permanently message by id
			DELETE https://{{host}}/{{admin_path}}/messages/{{delete_permanently_message_id}}/permanently
			Authorization: Bearer {{token}}

		*/
		{
			role:     V1_ADMIN,
			method:   delete,
			path:     "/users/messages/:id/permanently",
			endpoint: adminV1.DeleteMessagePermanently(s.Store),
		},
		/*

			### Edit message by id
			PATCH https://{{host}}/{{admin_path}}/messages/{{edit_message_id}}
			Authorization: Bearer {{token}}

			{
				"content": "{{new_message_content}}"
			}

		*/
		{
			role:     V1_ADMIN,
			method:   patch,
			path:     "/users/messages/:id",
			endpoint: adminV1.EditMessage(s.Store),
		},
		/*

			### Restore message by id
			POST https://{{host}}/{{admin_path}}/messages/{{restore_message_id}}
			Authorization: Bearer {{token}}

		*/
		{
			role:     V1_ADMIN,
			method:   post,
			path:     "/users/messages/:id/restore",
			endpoint: adminV1.RestoreMessage(s.Store),
		},
		/*

			### Delete user by id
			DELETE https://{{host}}/{{admin_path}}/users/{{delete_user_id}}
			Authorization: Bearer {{token}}

		*/
		{
			role:     V1_ADMIN,
			method:   delete,
			path:     "/users/:username",
			endpoint: adminV1.DeleteAccount(s.Store),
		},
		/*

			### Delete permanently user by id
			DELETE https://{{host}}/{{admin_path}}/users/{{delete_permanently_user_id}}/permanently
			Authorization: Bearer {{token}}

		*/
		{
			role:     V1_ADMIN,
			method:   delete,
			path:     "/users/:username/permanently",
			endpoint: adminV1.DeleteAccountPermanently(s.Store),
		},
		/*

			### Restore user by id
			POST https://{{host}}/{{admin_path}}/users/{{restore_user_id}}/restore
			Authorization: Bearer {{token}}

		*/
		{
			role:     V1_ADMIN,
			method:   post,
			path:     "/users/:username/restore",
			endpoint: adminV1.RestoreAccount(s.Store),
		},
		/*

			### Delete permanently link by id
			DELETE https://{{host}}/{{admin_path}}/links/{{delete_permanently_link_id}}/permanently
			Authorization: Bearer {{token}}

		*/
		{
			role:     V1_ADMIN,
			method:   delete,
			path:     "/links/:id/permanently",
			endpoint: adminV1.DeleteLinkPermanently(s.Store),
		},
		/*

			### Get contacts
			GET https://{{host}}/{{admin_path}}/links/{{delete_link_id}}
			Authorization: Bearer {{token}}

		*/
		{
			role:     V1_ADMIN,
			method:   post,
			path:     "/contacts",
			endpoint: adminV1.CreateContact(s.Store),
		},
	}

	return routes
}
