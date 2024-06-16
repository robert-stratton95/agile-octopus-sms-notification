package octopus

import (
	"agile-octopus-sms-notification/internal/domain"
	"time"
)

func OctopusEnergyPriceSupplier(baseUrl string) domain.EnergyPriceSupplier {
	return func(clock domain.Clock) ([]domain.EnergyPrice, error) {
		now := clock.Now()
		from := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
		to := from.Add(24 * time.Hour)
		getRateResponse, err := GetRatesResponse(baseUrl, from, to)
		if err != nil {
			return make([]domain.EnergyPrice, 0), err
		}
		return mapFunc(getRateResponse.Results, energyResultToEnergyPrice), nil
	}
}

func energyResultToEnergyPrice(result EnergyRateResult) domain.EnergyPrice {
	halfHourPeriod, _ := time.Parse(time.RFC3339, result.ValueFrom)
	return domain.NewEnergyPrice(
		result.ValueIncludingTax,
		halfHourPeriod,
	)
}

func mapFunc[T any, U any](slice []T, mappingFunc func(T) U) []U {
	resultSlice := make([]U, 0)
	for _, t := range slice {
		resultSlice = append(resultSlice, mappingFunc(t))
	}
	return resultSlice
}
