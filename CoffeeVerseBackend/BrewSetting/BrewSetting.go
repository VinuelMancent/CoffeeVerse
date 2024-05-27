package BrewSetting

type BrewSetting struct {
	BrewType         BrewTypes
	Coffee           string
	Grinder          string
	GrindLevel       int
	WaterTemperature int
}

type BrewTypes int

const (
	Drip BrewTypes = iota
	Espresso
	FrenchPress
	AeroPress
	PourOver
)
