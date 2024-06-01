package domain

import "time"

func FindAndNotifyEnergyPrices(energyPrices []EnergyPrice, clock Clock, notifier NotificationSender) []error {
	errs := []error{}
	now := clock.Now()

	currentPrices := filter(energyPrices, func(e EnergyPrice) bool {
		return e.halfHourPeriod.Before(now) && e.halfHourPeriod.Add(30 * time.Minute).After(now)
	})

	if len(currentPrices) == 1 {
		notifier.Notify(currentPrices[0].ToEnergyUsage())
	} 

	return errs
}

type EnergyPriceRepository interface {
	Save(EnergyPrice) error
	Get(time.Time) (EnergyPrice, error)
}

type NotificationSender interface {
	Notify(EnergyUsage) error
}

type Clock interface {
	Now() time.Time
}

type LocalClock struct {

}

func (c LocalClock) Now() time.Time {
	return time.Now()
}

func filter[T any] (list []T, predicate func (T) bool) []T {
	result := make([]T, 0)
	for _, t := range list {
		if predicate(t) {
			result = append(result, t)
		}
	}
	return result
}
