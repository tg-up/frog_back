package utils

import (
	jwtlib "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"icecreambash/tgup_backend/internal/configs"
	"icecreambash/tgup_backend/internal/models"
	"time"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateToken(user *models.User) string {

	token := jwtlib.New(jwtlib.GetSigningMethod("HS256"))

	// Set some claims
	token.Claims = jwtlib.MapClaims{
		"id":         user.ID,
		"name":       user.Name,
		"email":      user.Email,
		"role":       user.Role,
		"created_at": user.CreatedAt.Unix(),
		"signature":  user.Password,
		"exp":        time.Now().Add(time.Hour * 1).Unix(),
	}
	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte(configs.GlobalConfig.JWTPrivateToken))

	if err != nil {
		panic(err)
	}

	return tokenString
}
