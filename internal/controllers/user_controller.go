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

	c.JSON(200, gin.H{
		"password": request.Password,
		"valid":    utils.GenerateToken(&userFind),
	})
}

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
