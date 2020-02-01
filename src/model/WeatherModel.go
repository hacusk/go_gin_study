package model

type Weather struct {
	Location  Location    `json:"location"`
	Forecasts []Forecasts `json:"forecasts"`
}

type Location struct {
	City       string `json:"city"`
	Area       string `json:"area"`
	Prefecture string `json:"prefecture"`
}

type Forecasts struct {
	Telop       string      `json:"telop"`
	Temperature Temperature `json:"temperature"`
}

type Temperature struct {
	MinTemperature MinTemperature `json:"min"`
	MaxTemperature MaxTemperature `json:"max"`
}

type MinTemperature struct {
	Celsius string `json:"celsius"`
}

type MaxTemperature struct {
	Celsius string `json:"celsius"`
}
