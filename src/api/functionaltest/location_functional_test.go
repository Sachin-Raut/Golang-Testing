package functionaltest

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/Sachin-Raut/Golang-Testing/src/api/utils/errors"
	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/stretchr/testify/assert"
)

func TestGetCountryNotFound(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "https://api.mercadolibre.com/countries/AR",
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: http.StatusNotFound,
		RespBody: `{"status":404,
		"error":"not_found",
		"message":"no country with id AR"}`,
	})

	//execution
	response, err := http.Get("http://localhost:8080/countries/AR")

	//validation
	assert.Nil(t, err)
	assert.NotNil(t, response)

	bytes, _ := ioutil.ReadAll(response.Body)

	var apiErr errors.APIError
	err = json.Unmarshal(bytes, &apiErr)

	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusNotFound, apiErr.Status)
	assert.EqualValues(t, "not_found", apiErr.Error)
	assert.EqualValues(t, "no country with id ARs", apiErr.Message)
}
