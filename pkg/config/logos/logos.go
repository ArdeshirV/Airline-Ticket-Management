package logos

import (
	"strings"
)

// Public Interface ----------------------------------------------------------------------

type AirlineID int
type AirlineName string
type AirlineLogo string

const (
	AmericanAirline AirlineID = iota
	DeltaAirLine
	UnitedAirline
	Lufthansa
	Emirate
	BritishAirway
	AirFrance
	CathayPacificAirway
	QantasAirway
	SingaporeAirline
)

const (
	AmericanAirlineName     AirlineName = "american_airlines"
	DeltaAirLineName        AirlineName = "delta_air_lines"
	UnitedAirlineName       AirlineName = "united_airlines"
	LufthansaName           AirlineName = "lufthansa"
	EmirateName             AirlineName = "emirates"
	BritishAirwayName       AirlineName = "british_airways"
	AirFranceName           AirlineName = "air_france"
	CathayPacificAirwayName AirlineName = "cathay_pacific_airways"
	QantasAirwayName        AirlineName = "qantas_airways"
	SingaporeAirlineName    AirlineName = "singapore_airlines"
)

func GetAirlineName(id AirlineID) AirlineName {
	load()
	return airlineNames[id]
}

func GetAirlineNameList() []AirlineName {
	load()
	return airlineNames[:]
}

// Private Implementation ----------------------------------------------------------------

var (
	airlineNames []AirlineName
)

func load() {
	if airlineNames == nil {
		airlineNames = make([]AirlineName, 10)
		airlineNames[AmericanAirline] = AmericanAirlineName
		airlineNames[DeltaAirLine] = DeltaAirLineName
		airlineNames[UnitedAirline] = UnitedAirlineName
		airlineNames[Lufthansa] = LufthansaName
		airlineNames[Emirate] = EmirateName
		airlineNames[BritishAirway] = BritishAirwayName
		airlineNames[AirFrance] = AirFranceName
		airlineNames[CathayPacificAirway] = CathayPacificAirwayName
		airlineNames[QantasAirway] = QantasAirwayName
		airlineNames[SingaporeAirline] = SingaporeAirlineName
	}
}

func normalize(text AirlineName) AirlineName {
	return AirlineName(strings.ToLower(strings.TrimSpace(string(text))))
}
