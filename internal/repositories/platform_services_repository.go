package repositories

import (
	"gorm.io/gorm"
	"icecreambash/tgup_backend/internal/models"
)

type PlatformServiceRepository struct {
	db    *gorm.DB
	model models.PlatformServices
}

func NewPlatformServiceRepository(db *gorm.DB) *PlatformServiceRepository {
	return &PlatformServiceRepository{db: db}
}

func (context PlatformServiceRepository) GetServicesByPlatformID(id int) ([]models.PlatformServices, error) {
	var platformServices []models.PlatformServices
	err := context.db.Where("platform_id = ?", id).Find(&platformServices).Error
	if err != nil {
		return []models.PlatformServices{}, err
	}
	return platformServices, nil
}
