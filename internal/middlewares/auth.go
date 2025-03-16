package middlewares

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"icecreambash/tgup_backend/internal/configs"
	"icecreambash/tgup_backend/internal/models"
	"icecreambash/tgup_backend/pkg/database"
	"net/http"
	"time"
)

func LoadJWTAuth() *jwt.GinJWTMiddleware {
	authMiddleware, _ := jwt.New(&jwt.GinJWTMiddleware{
		Realm:        "Lead",
		Key:          []byte(configs.GlobalConfig.JWTPrivateToken),
		Timeout:      time.Hour,
		MaxRefresh:   time.Hour,
		IdentityKey:  "id",
		Authorizator: authorizator,
		TokenLookup:  "header: Authorization, query: token, cookie: jwt",
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized", "message": message})
		},
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})
	return authMiddleware
}

func authorizator(data interface{}, c *gin.Context) bool {
	claims := jwt.ExtractClaims(c)

	var user models.User

	database.DB.Model(&user).Where("id = ?", claims["id"]).First(&user)

	if claims["signature"] != user.Password {
		return false
	}

	c.Set("user", user)

	return true
}
