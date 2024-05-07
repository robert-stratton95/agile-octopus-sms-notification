package octopus

import (
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestSuccessfulResponse(t *testing.T) {
	url := "http://localhost/v1/products/AGILE-FLEX-22-11-25/electricity-tariffs/E-1R-AGILE-FLEX-22-11-25-C/standard-unit-rates/"
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		"GET",
		url,
		httpmock.NewStringResponder(200, `{
			"count": 48,
			"next": null,
			"previous": null,
			"results": [
			  {
				"value_exc_vat": 13.64,
				"value_inc_vat": 14.322,
				"valid_from": "2024-05-06T23:30:00Z",
				"valid_to": "2024-05-07T00:00:00Z",
				"payment_method": null
			  }]
			}`))

	ratesResult, err := GetRatesResponse("http://localhost")	


	assert.Nil(t, err, "No error should be thrown")
	assert.Equal(t, ratesResult, RatesResponse{
		Count: 48,
		Results: []EnergyRateResult{{
			ValueExcludingTax: 13.64,
			ValueIncludingTax: 14.322,
			ValueFrom:         "2024-05-06T23:30:00Z",
			ValueTo:           "2024-05-07T00:00:00Z",
		}},
	})
}