package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestShouldMapEnergyPriceToEnergyUsage(t *testing.T) {
	energy_prices := make([]EnergyPrice, 0, 3)
	energy_prices = append(energy_prices, EnergyPrice{price: 0.0, halfHourPeriod: time.Date(2009, 11, 17, 0, 0, 0, 0, time.UTC)})
	energy_prices = append(energy_prices, EnergyPrice{price: 1.0, halfHourPeriod: time.Date(2009, 11, 17, 0, 0, 0, 0, time.UTC)})
	energy_prices = append(energy_prices, EnergyPrice{price: -1.0, halfHourPeriod: time.Date(2009, 11, 17, 0, 0, 0, 0, time.UTC)})

	assert.Equal(t, EnergyUsage{msg: "CHAAAARGE!!!", pricePerKwh: 0.0}, energy_prices[0].ToEnergyUsage(), "Should be energy usage with zero price/kWh")
	assert.Equal(t, EnergyUsage{msg: "Being charged for electricity", pricePerKwh: 1.0}, energy_prices[1].ToEnergyUsage(), "Should be energy usage with correct price/kWh")
	assert.Equal(t, EnergyUsage{msg: "USE ELECTRICITY !!!!", pricePerKwh: -1.0}, energy_prices[2].ToEnergyUsage(), "Should be energy usage with negative price/kWh")
}
