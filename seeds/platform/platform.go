package platform

import (
	"encoding/json"
	"icecreambash/tgup_backend/internal/models"
	"icecreambash/tgup_backend/pkg/database"
	"log"
)

var servicesPlatform = map[string][]models.PlatformServices{
	"vk": []models.PlatformServices{
		models.PlatformServices{
			Name: "Подписчики в группу",
			Slug: "subscribers_in_groups",
		},
		models.PlatformServices{
			Name: "Друзья",
			Slug: "subscribers_in_friends",
		},
	},
}

var calls = map[string][]models.ServiceField{
	"vk_subscribers_in_groups": []models.ServiceField{
		models.ServiceField{
			FieldName:    "Тип накрутки",
			FieldSlug:    "type_worker",
			FieldType:    "radio",
			DefaultValue: "5",
			TempOptions: []models.Option{
				models.Option{
					Key:   "Боты",
					Value: 5,
				},
				models.Option{
					Key:   "Живые",
					Value: 4,
				},
			},
			CostType: "coeff",
			Size:     "grow",
		},
		models.ServiceField{
			FieldName:    "Качество",
			FieldSlug:    "quality",
			FieldType:    "radio",
			DefaultValue: "5",
			TempOptions: []models.Option{
				models.Option{
					Key:   "Низкое - 3",
					Value: 5,
				},
				models.Option{
					Key:   "Среднее - 5",
					Value: 4,
				},
				models.Option{
					Key:   "Высокое - 10",
					Value: 10,
				},
			},
			CostType: "base",
			Size:     "grow",
		},
	},
}

func PlatformSeeds() {
	platforms := []models.Platform{
		{Name: "VK", Slug: "vk"},
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

			fields := calls[platform.Slug+"_"+service.Slug]

			if fields == nil {
				continue
			}

			for _, field := range fields {
				var tempField models.ServiceField
				database.DB.Model(&tempField).Where("field_slug = ?", field.FieldSlug).Where("service_id = ?", service.ID).Find(&tempField)

				if tempField.ID == 0 {
					data, err := json.Marshal(field.TempOptions)
					if err != nil {
						log.Fatal(err)
					}
					field.Options.Scan(data)
					field.ServiceID = service.ID
					database.DB.Save(&field)
				}

			}
		}
	}
}
