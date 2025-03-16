package responses

import (
	"github.com/google/uuid"
	"icecreambash/tgup_backend/internal/models"
	"time"
)

type OrderResponse struct {
	ID        uuid.UUID               `json:"id"`
	Status    models.StatusOrder      `json:"status"`
	Link      string                  `json:"link"`
	Count     int                     `json:"count"`
	Reserved  bool                    `json:"reserved"`
	CreatedAt time.Time               `json:"created_at"`
	Service   PlatformServiceResponse `json:"service,omitempty"`
}

func GetOrdersList(orders []models.Order) []OrderResponse {
	var ordersResponse = make([]OrderResponse, 0)

	for _, order := range orders {
		ordersResponse = append(ordersResponse, OrderResponse{
			ID:        order.ID,
			Status:    order.Status,
			Link:      order.Link,
			Count:     int(order.Count),
			Reserved:  order.Reserved,
			CreatedAt: order.CreatedAt,
			Service:   ParsePlatformServiceToResponse(order.Service),
		})
	}

	return ordersResponse
}

func ParseOrderToResponse(order models.Order) OrderResponse {
	return OrderResponse{}
}
