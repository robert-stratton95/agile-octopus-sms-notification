package domain

import (
	"fmt"
	"time"
)

func NotifyEnergyPrices(energyPrices []EnergyPrice, clock Clock, notifier NotificationSender) []error {
	errs := []error{}
	now := clock.Now()

	currentPrices := filter(energyPrices, func(e EnergyPrice) bool {
		return e.halfHourPeriod.Before(now) && e.halfHourPeriod.Add(30*time.Minute).After(now)
	})

	if len(currentPrices) == 1 {
		notifier.Notify(currentPrices[0].ToEnergyUsage())
	} else {
		errorMsg := fmt.Sprintf("Could not find current energy price at %s: %v", now, currentPrices)
		price_error := CurrentEnergyPriceError{msg: errorMsg}
		errs = append(errs, price_error)
	}

	return errs
}

type CurrentEnergyPriceError struct {
	msg string
}

func (c CurrentEnergyPriceError) Error() string {
	return c.msg
}

type EnergyPriceSupplier func(Clock) ([]EnergyPrice, error)

type NotificationSender interface {
	Notify(EnergyUsage) error
}

func filter[T any](list []T, predicate func(T) bool) []T {
	result := make([]T, 0)
	for _, t := range list {
		if predicate(t) {
			result = append(result, t)
		}
	}
	return result
}
