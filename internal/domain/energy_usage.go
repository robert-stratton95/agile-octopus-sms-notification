package domain

type EnergyUsage struct {
	msg         string
	pricePerKwh float64
}

func (e *EnergyUsage) GetMsg() string {
	return e.msg
}

func (e *EnergyUsage) GetPricePerKwh() float64 {
	return e.pricePerKwh
}

func zero() EnergyUsage {
	return EnergyUsage{
		msg:         "CHAAAARGE!!!",
		pricePerKwh: 0.0,
	}
}

func negative(pricePerKwh float64) EnergyUsage {
	return EnergyUsage{
		msg:         "USE ELECTRICITY !!!!",
		pricePerKwh: pricePerKwh,
	}
}
