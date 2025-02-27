package middlewares

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"icecreambash/tgup_backend/internal/config"
	"icecreambash/tgup_backend/internal/models"
	"time"
)

func LoadJWTAuth() *jwt.GinJWTMiddleware {
	authMiddleware, _ := jwt.New(&jwt.GinJWTMiddleware{
		Realm:         "Lead",
		Key:           []byte(config.GlobalConfig.JWTPrivateToken),
		Timeout:       time.Hour,
		MaxRefresh:    time.Hour,
		IdentityKey:   "id",
		Authenticator: authenticator(),
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",

		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})
	return authMiddleware
}

func authenticator() func(c *gin.Context) (interface{}, error) {
	return func(c *gin.Context) (interface{}, error) {
		return models.User{}, nil
	}
}
