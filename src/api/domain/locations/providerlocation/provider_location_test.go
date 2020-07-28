package providerlocation

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
	assert.EqualValues(t, "", err.Message)
}
