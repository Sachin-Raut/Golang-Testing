package services

import (
	"github.com/Sachin-Raut/Golang-Testing/src/api/domain/locations"
	"github.com/Sachin-Raut/Golang-Testing/src/api/domain/locations/providerlocation"
	"github.com/Sachin-Raut/Golang-Testing/src/api/utils/errors"
)

type locationService struct{}

type locationServiceInterface interface {
	GetCountry(countryID string) (*locations.Country, *errors.APIError)
}

var (
	//LocationService is
	LocationService locationServiceInterface
)

func init() {
	LocationService = &locationService{}
}

func (s *locationService) GetCountry(countryID string) (*locations.Country, *errors.APIError) {
	return providerlocation.GetCountry(countryID)
}
