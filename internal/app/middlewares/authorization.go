package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/ilfey/devback/internal/pkg/iservices"
	"github.com/sirupsen/logrus"
)

const (
	AUTH_CONTEXT = "auth-context"
)

type AuthorizationContext struct {
	admin string
	*gin.Context
	jwt      iservices.JWT
	logger   *logrus.Logger
	token    string
	username string
}

func AuthorizationMiddleware(admin string, jwt iservices.JWT, logger *logrus.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		aCtx := new(AuthorizationContext)
		aCtx.admin = admin
		aCtx.Context = ctx
		aCtx.jwt = jwt
		aCtx.logger = logger

		ctx.Set(AUTH_CONTEXT, aCtx)

		ctx.Next()
	}
}

func (ctx *AuthorizationContext) IsAuthorized() bool {

	// Get token from "Authorization" header
	token := ctx.jwt.GetToken(ctx.Context)

	// Validate token
	if _, err := ctx.jwt.ValidateToken(token); err != nil {
		return false
	}

	// Parse username from token
	username, err := ctx.jwt.GetTokenId(ctx.Context)
	if err != nil {
		return false
	}

	// Save token and username
	ctx.token = token
	ctx.username = username

	return true
}

func (ctx *AuthorizationContext) IsAdmin() bool {

	// Check stored token and username or get it
	if ctx.token == "" && ctx.username == "" && !ctx.IsAuthorized() {
		return false
	}

	// Check username is equals admin's username
	return ctx.username == ctx.admin
}
