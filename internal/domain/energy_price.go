package domain

import (
	"time"
)

type EnergyPrice struct {
	price          float64
	halfHourPeriod time.Time
}

func (e EnergyPrice) GetPrice() float64 {
	return e.price
}

func (e EnergyPrice) GetHalfHourPeriod() time.Time {
	return e.halfHourPeriod
}

func (e EnergyPrice) ToEnergyUsage() EnergyUsage {
	if e.GetPrice() == 0.0 {
		return zero()
	}
	if e.GetPrice() > 0.0 {
		return EnergyUsage{msg: "Being charged for electricity", pricePerKwh: e.price}
	}
	return negative(e.price)
}

func NewEnergyPrice(price float64, halfHourPeriod time.Time) EnergyPrice {
	return EnergyPrice{
		price: price,
		halfHourPeriod: halfHourPeriod,
	}
}
