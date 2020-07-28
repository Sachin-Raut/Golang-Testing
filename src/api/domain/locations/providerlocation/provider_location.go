package providerlocation

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Sachin-Raut/Golang-Testing/src/api/domain/locations"
	"github.com/Sachin-Raut/Golang-Testing/src/api/utils/errors"
	"github.com/mercadolibre/golang-restclient/rest"
)

const (
	urlgetCountry = "https://api.mercadolibre.com/countries/%s"
)

//GetCountry is
func GetCountry(countryID string) (*locations.Country, *errors.APIError) {
	a := fmt.Sprintf(urlgetCountry, countryID)
	fmt.Println(a)
	response := rest.Get(fmt.Sprintf(urlgetCountry, countryID))

	fmt.Println("response.Response =", response.Response)
	fmt.Println("response =", response)
	if response == nil || response.Response == nil {
		fmt.Println("111")
		//i.e we have invalid REST client response fom the API or we have timedout
		return nil, &errors.APIError{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("invalid restclient response when trying to get country %s", countryID),
			Error:   "Country not found",
		}
	}

	if response.StatusCode > 299 {
		fmt.Println("222")
		var apiErr errors.APIError

		if err := json.Unmarshal(response.Bytes(), &apiErr); err != nil {

			return nil, &errors.APIError{
				Status:  http.StatusInternalServerError,
				Message: fmt.Sprintf("invalid error interface when trying to get country %s", countryID),
				Error:   "Country not found",
			}
		}
		return nil, &apiErr
	}

	var result locations.Country

	if err := json.Unmarshal(response.Bytes(), &result); err != nil {
		fmt.Println("333")
		return nil, &errors.APIError{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("error when trying to unmarshal country data for %s", countryID),
		}
	}
	fmt.Println("444")
	return &result, nil
}
