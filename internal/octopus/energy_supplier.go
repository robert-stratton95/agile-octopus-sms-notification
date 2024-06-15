package octopus

import (
	"agile-octopus-sms-notification/internal/domain"
)

func OctopusEnergyPriceSupplier(baseUrl string) domain.EnergyPriceSupplier {
	return func(domain.Clock) []domain.EnergyPrice {
		panic("not implemented")
	}
}
