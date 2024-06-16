package octopus

import (
	"io"
	"net/http"
	"time"
)

const ratesEndpoint string = "/v1/products/AGILE-FLEX-22-11-25/electricity-tariffs/E-1R-AGILE-FLEX-22-11-25-C/standard-unit-rates/"

var httpClient http.Client = http.Client{}

func GetRatesResponse(baseUrl string, from time.Time, to time.Time) (RatesResponse, error) {
	request, requestErr := http.NewRequest("GET", baseUrl+ratesEndpoint, nil)
	if requestErr != nil {
		return RatesResponse{}, requestErr
	}

	query := request.URL.Query()
	query.Add("period_from", from.Format(time.RFC3339))
	query.Add("period_to", to.Format(time.RFC3339))

	request.URL.RawQuery = query.Encode()

	response, responseErr := httpClient.Do(request)
	if responseErr != nil {
		return RatesResponse{}, responseErr
	}
	responseBytes, bytesError := io.ReadAll(response.Body)
	if bytesError != nil {
		return RatesResponse{}, bytesError
	}
	ratesResponse, unmarshallErr := UmarshallRatesRespone(responseBytes)
	if unmarshallErr != nil {
		return RatesResponse{}, unmarshallErr
	}

	return ratesResponse, nil
}
