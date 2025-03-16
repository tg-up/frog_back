package services

import (
	"icecreambash/tgup_backend/internal/models"
	"icecreambash/tgup_backend/internal/repositories"
	"icecreambash/tgup_backend/internal/responses"
)

type PlatformService struct {
	platformRepository        *repositories.PlatformRepository
	platformServiceRepository *repositories.PlatformServiceRepository
}

func NewPlatformService(platformRepository *repositories.PlatformRepository, platformServiceRepository *repositories.PlatformServiceRepository) *PlatformService {
	return &PlatformService{
		platformRepository:        platformRepository,
		platformServiceRepository: platformServiceRepository,
	}
}

func (context PlatformService) GetAllPlatforms() ([]responses.PlatformResponse, error) {
	values, err := context.platformRepository.GetAllPlatforms()
	return responses.GetAllPlatformResponse(values), err
}

func (context PlatformService) GetPlatformByID(id int) ([]models.Platform, error) {
	values, err := context.platformRepository.GetAllPlatforms()
	return values, err
}

func (context PlatformService) GetPlatformServicesByID(id int) ([]responses.PlatformServiceResponse, error) {
	values, err := context.platformServiceRepository.GetServicesByPlatformID(id)
	if err != nil {
		return []responses.PlatformServiceResponse{}, err
	}
	return responses.GetAllPlatformServicesResponse(values), nil
}
