package domain

import "time"

func SaveEnergyPrices(energyPrices []EnergyPrice, persist func(EnergyPrice) error) []error {
	errs := []error{}
	for _, energyPrice := range energyPrices {
		err := persist(energyPrice)
		if err != nil {
			errs = append(errs, err)
		}
		errs = append(errs, err)
	}
	return errs
}

type EnergyPriceRepository interface {
	Save(EnergyPrice) error
	Get(time.Time) (EnergyPrice, error)
}

type NotificationSender func(EnergyUsage) error
