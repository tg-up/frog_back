package repositories

import (
	"gorm.io/gorm"
	"icecreambash/tgup_backend/internal/models"
)

type ServiceFieldsRepository struct {
	db    *gorm.DB
	model models.PlatformServices
}

func NewServiceFieldsRepository(db *gorm.DB) *ServiceFieldsRepository {
	return &ServiceFieldsRepository{db: db}
}

func (context ServiceFieldsRepository) GetFieldsByServiceID(serviceID int) ([]models.ServiceField, error) {
	var tempModels []models.ServiceField

	err := context.db.Model(&models.ServiceField{}).Where("service_id = ?", serviceID).Find(&tempModels).Error

	if err != nil {
		return []models.ServiceField{}, err
	}

	return tempModels, nil
}
