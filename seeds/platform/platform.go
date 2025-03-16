package platform

import (
	"icecreambash/tgup_backend/internal/models"
	"icecreambash/tgup_backend/pkg/database"
	"log"
)

var servicesPlatform = map[string][]models.PlatformServices{
	"tg": []models.PlatformServices{
		models.PlatformServices{
			Name:          "Подписчики в группу(3Days+ drop)",
			Slug:          "subscribers_in_groups_3d_drop",
			MinCount:      10,
			MaxCount:      10000,
			Drip:          false,
			Amount:        0.5,
			AmountAbility: 1000,
		},
		models.PlatformServices{
			Name:          "Подписчики в группу(7Days+ drop)",
			Slug:          "subscribers_in_groups_7d_drop",
			MinCount:      10,
			MaxCount:      10000,
			Drip:          false,
			Amount:        0.75,
			AmountAbility: 1000,
		},
		models.PlatformServices{
			Name:          "Подписчики в группу(30 Days drop)",
			Slug:          "subscribers_in_groups_30d_drop",
			MinCount:      10,
			MaxCount:      10000,
			Drip:          false,
			Amount:        1,
			AmountAbility: 1000,
		},
		models.PlatformServices{
			Name:          "Подписчики в группу(No drop)",
			Slug:          "subscribers_in_groups_no_drop",
			MinCount:      10,
			MaxCount:      10000,
			Drip:          false,
			Amount:        1.4,
			AmountAbility: 1000,
		},
	},
}

func PlatformSeeds() {
	platforms := []models.Platform{
		{Name: "Telegram", Slug: "tg"},
	}

	for key, platform := range platforms {
		database.DB.Where("slug = ?", platform.Slug).First(&platform)
		if platform.ID == 0 {
			err := database.DB.Save(&platform).Error
			if err != nil {
				log.Fatal(err)
			}
		}
		platforms[key] = platform
	}

	for _, platform := range platforms {
		services := servicesPlatform[platform.Slug]
		if services == nil {
			continue
		}
		for _, service := range services {
			err := database.DB.Model(&service).Where("slug = ?", service.Slug).Where("platform_id = ?", platform.ID).Find(&service).Error
			if err != nil {
				log.Fatal(err)
			}
			if service.ID == 0 {
				service.PlatformID = platform.ID
				err := database.DB.Save(&service).Error
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	}
}
