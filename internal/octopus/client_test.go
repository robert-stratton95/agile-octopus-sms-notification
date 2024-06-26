package octopus

import (
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestSuccessfulResponse(t *testing.T) {
	url := "http://localhost/v1/products/AGILE-FLEX-22-11-25/electricity-tariffs/E-1R-AGILE-FLEX-22-11-25-C/standard-unit-rates/?period_from=2024-01-01T00%3A00%3A00Z&period_to=2024-01-02T00%3A00%3A00Z"
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

	ratesResult, err := GetRatesResponse("http://localhost", time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC))

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
