package octopus

import (
	"io"
	"net/http"
)

const ratesEndpoint string = "/v1/products/AGILE-FLEX-22-11-25/electricity-tariffs/E-1R-AGILE-FLEX-22-11-25-C/standard-unit-rates/"
var httpClient http.Client = http.Client{} 

func GetRatesResponse(baseUrl string) (RatesResponse, error) {
	request, requestErr := http.NewRequest("GET", baseUrl + ratesEndpoint, nil)
	if requestErr != nil {
		return RatesResponse{}, requestErr
	}
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
