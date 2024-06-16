package octopus

import (
	"agile-octopus-sms-notification/internal/domain"
	"agile-octopus-sms-notification/internal/utils"
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
		return utils.MapSlice(getRateResponse.Results, energyResultToEnergyPrice), nil
	}
}

func energyResultToEnergyPrice(result EnergyRateResult) domain.EnergyPrice {
	halfHourPeriod, _ := time.Parse(time.RFC3339, result.ValueFrom)
	return domain.NewEnergyPrice(
		result.ValueIncludingTax,
		halfHourPeriod,
	)
}
