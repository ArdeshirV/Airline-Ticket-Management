package logos

import (
	"strings"
)

// Public Interface ----------------------------------------------------------------------

type AirlineID int
type AirlineTag int
type AirlineName string
type AirlineLogo string

const (
	AirlineLogoTag AirlineTag = 0
	AirlineNameTag AirlineTag = 1
)

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

func (airlineID AirlineID) GetName() string {
	return string(airlineNames[airlineID][int(AirlineNameTag)])
}

func (airlineID AirlineID) GetLogo() (result string) {
	return string(airlineNames[airlineID][int(AirlineLogoTag)])
}

func GetAirlineNames() (result [10][2]string) {
	return airlineNames
}

func GetAirlineLogoByName(name string) (string, error) {
	return "", nil
}

// Private Implementation ----------------------------------------------------------------

var (
	airlineNames [10][2]string
)

func load() {
	if len(airlineNames) == 0 {
		airlineNames := make([][2]string, 10)
		airlineNames[AmericanAirline] = [2]string{string(AmericanAirlineName), "American Airline"}
		airlineNames[DeltaAirLine] = [2]string{string(DeltaAirLineName), "Delta Airline"}
		airlineNames[UnitedAirline] = [2]string{string(UnitedAirlineName), "United Airline"}
		airlineNames[Lufthansa] = [2]string{string(LufthansaName), "Lufthansa"}
		airlineNames[Emirate] = [2]string{string(EmirateName), "Emirate"}
		airlineNames[BritishAirway] = [2]string{string(BritishAirwayName), "British Airways"}
		airlineNames[AirFrance] = [2]string{string(AirFranceName), "Air France"}
		airlineNames[CathayPacificAirway] = [2]string{string(CathayPacificAirwayName), "Cathay Pacific Airways"}
		airlineNames[QantasAirway] = [2]string{string(QantasAirwayName), "Qantas Airways"}
		airlineNames[SingaporeAirline] = [2]string{string(SingaporeAirlineName), "Singapore Airline"}
	}
}

func normalize(text AirlineName) AirlineName {
	return AirlineName(strings.ToLower(strings.TrimSpace(string(text))))
}
