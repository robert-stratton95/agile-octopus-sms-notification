package domain

import (
	"time"
)

type EnergyPrice struct {
	price          float64
	halfHourPeriod time.Time
}

func (e *EnergyPrice) GetPrice() float64 {
	return e.price
}

func (e *EnergyPrice) GetHalfHourPeriod() time.Time {
	return e.halfHourPeriod
}
