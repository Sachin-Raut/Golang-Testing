package providerlocation

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
run all the tests after disconnecting & reconnecting the internet
*/

func TestGetCountryRestClientErrorNoInternet(t *testing.T) {
	// rest.StartMockupServer()
	// rest.FlushMockups()
	// rest.AddMockups(&rest.Mock{
	// 	URL:          "https://api.mercadolibre.com/countries/AR",
	// 	HTTPMethod:   http.MethodGet,
	// 	RespHTTPCode: 0,
	// })

	country, err := GetCountry("AR")

	assert.Nil(t, country)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid rest client response when trying to get country AR", err.Message)
}

func TestGetCountryNotFound(t *testing.T) {
	country, err := GetCountry("AR")

	assert.Nil(t, country)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, "Country not found", err.Message)
}
