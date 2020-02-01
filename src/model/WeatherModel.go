package model

// Weather return Location and Forecasts Info
type Weather struct {
	Location  Location    `json:"location"`
	Forecasts []Forecasts `json:"forecasts"`
}

// Location return Location Info
type Location struct {
	City       string `json:"city"`
	Area       string `json:"area"`
	Prefecture string `json:"prefecture"`
}

// Forecasts return Today and Tommorow Forecast Info
type Forecasts struct {
	DateLabel   string      `json:"dateLabel"`
	Date        string      `json:"date"`
	Telop       string      `json:"telop"`
	Temperature Temperature `json:"temperature"`
}

// Temperature return Temperature Info
type Temperature struct {
	MinTemperature MinTemperature `json:"min"`
	MaxTemperature MaxTemperature `json:"max"`
}

// MinTemperature return MinTemperature Celsius
type MinTemperature struct {
	Celsius string `json:"celsius"`
}

// MaxTemperature return MaxTemperature Celsius
type MaxTemperature struct {
	Celsius string `json:"celsius"`
}
