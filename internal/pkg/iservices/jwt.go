package iservices

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type JWT interface {
	GenerateToken(string) string
	GetToken(*gin.Context) string
	GetTokenId(*gin.Context) (string, error)
	ValidateToken(string) (*jwt.Token, error)
}
