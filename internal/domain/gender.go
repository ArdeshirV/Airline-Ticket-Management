package domain

type Gender int

const (
	Male Gender = iota
	Female
)

func (data Gender) String() (result string) {
	switch data {
	case Male:
		result = "Male"
	case Female:
		result = "Female"
	default:
		result = "Unknown"
	}
	return result
}
