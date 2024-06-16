package octopus

import (
	"agile-octopus-sms-notification/internal/domain"
)

func OctopusEnergyPriceSupplier(baseUrl string) domain.EnergyPriceSupplier {
	return func(c domain.Clock) ([]domain.EnergyPrice, error) {
		panic("not implemented")
	}
}
