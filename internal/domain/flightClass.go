package domain

type FlightClass int

const (
	FirstClass FlightClass = iota
	BusinessClass
	EconomyClass
)

func (data FlightClass) String() (result string) {
	switch data {
	case FirstClass:
		result = "First Class"
	case BusinessClass:
		result = "Business Class"
	case EconomyClass:
		result = "Economic Class"
	default:
		result = "Unknown"
	}
	return result
}
