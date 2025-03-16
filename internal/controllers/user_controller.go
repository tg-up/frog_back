package controllers

import (
	"github.com/gin-gonic/gin"
	"icecreambash/tgup_backend/internal/models"
	"icecreambash/tgup_backend/internal/requests/user"
	"icecreambash/tgup_backend/internal/services"
	"icecreambash/tgup_backend/pkg/database"
	"icecreambash/tgup_backend/pkg/utils"
	"net/http"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) *UserController {
	return &UserController{userService: services.UserService{}}
}

// Login godoc
// @Summary      Авторизация в системе
// @Description  Метод для получения JWT токена и использования его в дальнейшем для подписи действий в API
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request body user.LoginRequest true "Параметры для авторизации"
// @Router       /api_gateway/auth/login [post]
func (userController *UserController) Login(c *gin.Context) {
	request := user.LoginRequest{}

	err := c.ShouldBindBodyWithJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Validation failed",
			"details": err.Error(),
		})
		return
	}

	var userFind models.User

	database.DB.Model(&userFind).Where("email = ?", request.Email).First(&userFind)

	if userFind == (models.User{}) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Validation failed",
			"details": "User or password not found",
		})
		return
	}

	if utils.CheckPasswordHash(request.Password, userFind.Password) == false {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Validation failed",
			"details": "User or password not found",
		})
	}

	c.JSON(201, gin.H{
		"token": utils.GenerateToken(&userFind),
	})
}

// Register godoc
// @Summary      Регистрация в системе
// @Description  Метод для создания пользователя
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request body user.RegisterRequest true "Параметры для регистрации"
// @Router       /api_gateway/auth/register [post]
func (userController *UserController) Register(c *gin.Context) {
	request := user.RegisterRequest{}

	err := c.ShouldBindBodyWithJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Validation failed",
			"details": err.Error(),
		})
		return
	}

	var userFind models.User

	database.DB.Model(userFind).Where("email = ?", request.Email).First(&userFind)

	if userFind != (models.User{}) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Validation failed",
			"details": "Email already taken",
		})
		return
	}

	password, _ := utils.HashPassword(request.Password)

	var newUser = models.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: password,
		Role:     models.USER,
	}

	database.DB.Create(&newUser)

	c.JSON(http.StatusOK, newUser)

}
