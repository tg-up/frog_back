package repositories

import (
	"errors"
	"gorm.io/gorm"
	"icecreambash/tgup_backend/internal/models"
)

type PlatformRepository struct {
	db *gorm.DB
}

func NewPlatformRepository(db *gorm.DB) *PlatformRepository {
	return &PlatformRepository{db: db}
}

func (context PlatformRepository) GetAllPlatforms() ([]models.Platform, error) {

	var platforms []models.Platform

	err := context.db.Model(&models.Platform{}).Preload("Services").Find(&platforms).Error

	if err != nil {
		return []models.Platform{}, errors.New("Failed to get platforms")
	}

	return platforms, nil
}

func (context PlatformRepository) GetPlatformByID(id int) ([]models.Platform, error) {
	var platform []models.Platform
	err := context.db.Model(&models.Platform{}).Where("id = ?", id).Find(&platform).Error
	if err != nil {
		return []models.Platform{}, errors.New("Failed to get platform")
	}
	return platform, nil
}
