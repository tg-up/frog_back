package responses

import "icecreambash/tgup_backend/internal/models"

type PlatformServiceResponse struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	Slug       string `json:"slug"`
	PlatformID uint   `json:"platform_id"`
}

func GetAllPlatformServicesResponse(services []models.PlatformServices) []PlatformServiceResponse {
	var platformServices []PlatformServiceResponse

	for _, item := range services {
		platformServices = append(platformServices, PlatformServiceResponse{
			ID:         item.ID,
			Name:       item.Name,
			Slug:       item.Slug,
			PlatformID: item.PlatformID,
		})
	}

	return platformServices
}
