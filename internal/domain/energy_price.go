package domain

import (
	"time"
)

type EnergyPrice struct {
	price float64
	halfHourPeriod time.Time
}

func New(price float64, halfHourPeriod time.Time) *EnergyPrice {
	return &EnergyPrice{
		price: price,
		halfHourPeriod: halfHourPeriod,
	}
}

func (e *EnergyPrice) GetPrice() float64 {
	return e.price
}

func (e *EnergyPrice) GetHalfHourPeriod() time.Time {
	return e.halfHourPeriod
}