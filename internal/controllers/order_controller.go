package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
	"icecreambash/tgup_backend/internal/models"
	orderRequest "icecreambash/tgup_backend/internal/requests/order"
	"icecreambash/tgup_backend/internal/responses"
	"icecreambash/tgup_backend/internal/services"
	"icecreambash/tgup_backend/pkg/database"
	"net/http"
)

type OrderController struct {
	platformService services.PlatformService
}

func NewOrderController(platformService services.PlatformService) *OrderController {
	return &OrderController{
		platformService: platformService,
	}
}

func (order *OrderController) ShowAll(c *gin.Context) {
	user, ok := c.MustGet("user").(models.User)
	if !ok {
		return
	}

	var orders = make([]models.Order, 0)

	database.DB.Model(&models.Order{}).Preload("Service").Where("user_id = ?", user.ID).Find(&orders)

	values := responses.GetOrdersList(orders)

	c.JSON(http.StatusOK, gin.H{
		"data": values,
	})
}

func (order *OrderController) Create(c *gin.Context) {
	var valueRequest orderRequest.CreateRequest

	user, ok := c.MustGet("user").(models.User)
	if !ok {
		return
	}

	err := c.BindJSON(&valueRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	platform, err := order.platformService.GetPlatformByID(valueRequest.PlatformID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(platform) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Platform does not exist"})
		return
	}

	platformOriginal := platform[0]

	service := models.PlatformServices{}

	err = database.DB.Model(&service).Where("id = ?", valueRequest.ServiceID).Where("platform_id = ?", platformOriginal.ID).Find(&service).Error

	/*
		Если что-то пошло не так в базе данных
	*/
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	/*
		Если сервис не найден
	*/
	if service.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Service does not exist"})
		return
	}

	/*
		Количество превышено
	*/
	if valueRequest.Count > int(service.MaxCount) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Max count exceeded"})
		return
	}

	/*
		Количество меньше допустимого
	*/
	if valueRequest.Count < int(service.MinCount) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Min count failure"})
		return
	}

	newOrder := models.Order{}

	newOrder.Link = valueRequest.Link
	newOrder.Count = uint(valueRequest.Count)
	newOrder.ServiceID = int(service.ID)
	newOrder.UserID = user.ID
	newOrder.Status = models.StatusOrder(models.START)

	err = database.DB.Create(&newOrder).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create order"})
		return
	}

	body, _ := json.Marshal(&newOrder)

	err = database.RabbitMQ.Publish(
		"",
		"order_queue",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	c.JSON(200, gin.H{
		"service": newOrder,
	})
}
