package domain

import "time"

func SaveEnergyPrices(energyPrices []EnergyPrice, persist func(EnergyPrice) error) []error {
	errs := []error{}
	for _, energyPrice := range energyPrices {
		err := persist(energyPrice)
		errs = append(errs, err)
	}
	return errs
}

func EnergyPriceToEnergyUsage(energyPrice EnergyPrice) EnergyUsage {
	if energyPrice.GetPrice() == 0.0 {
		return *zero()
	}
	panic("Not implemented yet")
}

type EnergyPriceRepository interface {
	Save(EnergyPrice) error
	Get(time.Time) (EnergyPrice, error)
}

type NotificationSender func(EnergyUsage) error
