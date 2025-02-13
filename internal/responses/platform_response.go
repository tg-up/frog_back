package responses

import "icecreambash/tgup_backend/internal/models"

type PlatformResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

func GetAllPlatformResponse(platforms []models.Platform) []PlatformResponse {

	var values []PlatformResponse

	for _, platform := range platforms {
		values = append(values, PlatformResponse{
			ID:   platform.ID,
			Name: platform.Name,
			Slug: platform.Slug,
		})
	}

	return values
}
