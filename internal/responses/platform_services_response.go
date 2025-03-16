package responses

import "icecreambash/tgup_backend/internal/models"

type PlatformServiceResponse struct {
	ID            uint    `json:"id"`
	PlatformID    uint    `json:"platform_id"`
	Name          string  `json:"name"`
	Slug          string  `json:"slug"`
	MinCount      uint    `json:"min_count"`
	MaxCount      uint    `json:"max_count"`
	Amount        float64 `json:"amount"`
	AmountAbility uint    `json:"amount_ability"`
}

func GetAllPlatformServicesResponse(services []models.PlatformServices) []PlatformServiceResponse {
	var platformServices []PlatformServiceResponse

	for _, item := range services {
		platformServices = append(platformServices, ParsePlatformServiceToResponse(item))
	}

	return platformServices
}

func ParsePlatformServiceToResponse(item models.PlatformServices) PlatformServiceResponse {
	return PlatformServiceResponse{
		ID:            item.ID,
		PlatformID:    item.PlatformID,
		Name:          item.Name,
		Slug:          item.Slug,
		MinCount:      item.MinCount,
		MaxCount:      item.MaxCount,
		Amount:        item.Amount,
		AmountAbility: item.AmountAbility,
	}
}
