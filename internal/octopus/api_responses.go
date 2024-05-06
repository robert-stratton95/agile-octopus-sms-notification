package octopus

import "encoding/json"

type EnergyRateResult struct {
	ValueExcludingTax float64 `json:"value_exc_vat"`
	ValueIncludingTax float64 `json:"value_inc_vat"`
	ValueFrom  string `json:"valid_from"`
	ValueTo string `json:"valid_to"`
}

type RatesResponse struct {
	Count int64 `json:"count"`
	Results []EnergyRateResult `json:"results"`
}

func UmarshallRatesRespone(bytes []byte) (RatesResponse, error) {
	var response RatesResponse
	err := json.Unmarshal(bytes, &response)
	if err != nil {
		return RatesResponse{}, err
	}
	return response, nil
}