package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestShouldMapEnergyPriceToEnergyUsage(t *testing.T) {
	enery_prices := make([]EnergyPrice, 3)
	enery_prices = append(enery_prices, EnergyPrice{price: 0.0, halfHourPeriod: time.Date(2009, 11, 17, 0, 0, 0, 0, time.UTC)})
	enery_prices = append(enery_prices, EnergyPrice{price: 1.0, halfHourPeriod: time.Date(2009, 11, 17, 0, 0, 0, 0, time.UTC)})
	enery_prices = append(enery_prices, EnergyPrice{price: -1.0, halfHourPeriod: time.Date(2009, 11, 17, 0, 0, 0, 0, time.UTC)})

	assert.Equal(t, EnergyUsage{msg: "CHAAAARGE!!!", pricePerKwh: 0.0}, EnergyPriceToEnergyUsage(enery_prices[0]), "Should be energy usage with zero price/kWh")
	assert.Equal(t, EnergyUsage{msg: "CHAAAARGE!!!", pricePerKwh: 1.0}, EnergyPriceToEnergyUsage(enery_prices[1]), "Should be energy usage with CORRECT price/kWh")
	assert.Equal(t, EnergyUsage{msg: "USE ELECTRICITY !!!!", pricePerKwh: -1.0}, enery_prices[2], "Should be energy usage with negative price/kWh")
}
