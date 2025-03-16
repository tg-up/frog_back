package order

type CreateRequest struct {
	PlatformID int    `json:"platform_id" binding:"required"`
	ServiceID  int    `json:"service_id" binding:"required"`
	Count      int    `json:"count" binding:"required"`
	Link       string `json:"link" binding:"required"`
}
