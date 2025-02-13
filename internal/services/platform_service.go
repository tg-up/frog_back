package services

import (
	"icecreambash/tgup_backend/internal/repositories"
	"icecreambash/tgup_backend/internal/responses"
)

type PlatformService struct {
	platformRepository        *repositories.PlatformRepository
	platformServiceRepository *repositories.PlatformServiceRepository
	serviceFieldsRepository   *repositories.ServiceFieldsRepository
}

func NewPlatformService(platformRepository *repositories.PlatformRepository, platformServiceRepository *repositories.PlatformServiceRepository, serviceFieldsRepository *repositories.ServiceFieldsRepository) *PlatformService {
	return &PlatformService{
		platformRepository:        platformRepository,
		platformServiceRepository: platformServiceRepository,
		serviceFieldsRepository:   serviceFieldsRepository,
	}
}

func (context PlatformService) GetAllPlatforms() ([]responses.PlatformResponse, error) {
	values, err := context.platformRepository.GetAllPlatforms()
	return responses.GetAllPlatformResponse(values), err
}

func (context PlatformService) GetPlatformServicesByID(id int) ([]responses.PlatformServiceResponse, error) {
	values, err := context.platformServiceRepository.GetServicesByPlatformID(id)
	if err != nil {
		return []responses.PlatformServiceResponse{}, err
	}
	return responses.GetAllPlatformServicesResponse(values), nil
}

func (context PlatformService) GetFieldsByServiceID(id int) ([]responses.ServiceFieldsResponse, error) {
	values, err := context.serviceFieldsRepository.GetFieldsByServiceID(id)
	return responses.GetFieldsResponse(values), err
}
