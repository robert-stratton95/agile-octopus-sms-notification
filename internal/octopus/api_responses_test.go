package octopus

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldUnmarshallResponse(t *testing.T) {
	response := `{
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
		}`

	ratesResult, err := UmarshallRatesRespone([]byte(response))

	assert.Nil(t, err, "No error should be thrown")
	assert.Equal(t, ratesResult, OctopusRatesResponse{
		Count: 48,
		Results: []OctopusEnergyRateResult{{
			ValueExcludingTax: 13.64,
			ValueIncludingTax: 14.322,
			ValueFrom:         "2024-05-06T23:30:00Z",
			ValueTo:           "2024-05-07T00:00:00Z",
		}},
	})
}
