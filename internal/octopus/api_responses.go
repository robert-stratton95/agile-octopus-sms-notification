package octopus

import "encoding/json"

type OctopusEnergyRateResult struct {
	ValueExcludingTax float64 `json:"value_exc_vat"`
	ValueIncludingTax float64 `json:"value_inc_vat"`
	ValueFrom  string `json:"valid_from"`
	ValueTo string `json:"valid_to"`
}

type OctopusRatesResponse struct {
	Count int64 `json:"count"`
	Results []OctopusEnergyRateResult `json:"results"`
}

func UmarshallRatesRespone(bytes []byte) (OctopusRatesResponse, error) {
	var response OctopusRatesResponse
	err := json.Unmarshal(bytes, &response)
	if err != nil {
		return OctopusRatesResponse{}, err
	}
	return response, nil
}