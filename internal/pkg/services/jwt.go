package services

import (
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/ilfey/devback/internal/pkg/iservices"
)

type Service struct {
	SecretKey string
	LifeSpan  int
	Issure    string
}

func NewService(key string, lifespan int) iservices.JWT {
	return &Service{
		SecretKey: key,
		LifeSpan:  lifespan,
		Issure:    "devback",
	}
}

func (s *Service) GenerateToken(username string) string {

	claims := jwt.MapClaims{}

	claims["authorized"] = true
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(s.LifeSpan)).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(s.SecretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (s *Service) GetToken(ctx *gin.Context) string {
	token := ctx.Query("token")
	if token != "" {
		return token
	}

	authHeader := ctx.Request.Header.Get("Authorization")
	if authHeader == "" {
		return ""
	}

	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 && headerParts[0] != "Bearer" {
		return ""
	}

	return headerParts[1]
}

func (s *Service) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("invalid token %v", token.Header["alg"])
		}

		return []byte(s.SecretKey), nil
	})
}

func (s *Service) GetTokenId(ctx *gin.Context) (string, error) {
	tokenString := s.GetToken(ctx)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(s.SecretKey), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {

		return claims["username"].(string), nil
	}

	return "", nil
}
