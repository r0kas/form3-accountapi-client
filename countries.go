package account

// Country type represents supported countries list by the API.
// It ensures correct country codes are used.
// Has Bank ID Codes linked to the country.
/////
// Have chosen to have this custom implementation
// as were there no package which would fulfill ISO 3166.
// Additionally this provided the opportunity to group countries with Bank ID Codes.
/////
type Country string

// List of countries supported by Form3
const (
	UnitedKingdom Country = "GB"
	Australia             = "AU"
	Belgium               = "BE"
	Canada                = "CA"
	France                = "FR"
	Germany               = "DE"
	Greece                = "GR"
	HongKong              = "HK"
	Italy                 = "IT"
	Luxembourg            = "LU"
	Netherlands           = "NL"
	Poland                = "PL"
	Portugal              = "PT"
	Spain                 = "ES"
	Switzerland           = "CH"
	UnitedStates          = "US"
)

// Code returns string representation of ISO 3166 country code.
func (country Country) Code() string {
	return string(country)
}

// BankIDCode returns associated Bank ID Code for given country
func (country Country) BankIDCode() string {
	var code string
	switch country {
	case UnitedKingdom:
		code = "GBDSC"
	case Australia:
		code = "AUBSB"
	case Belgium:
		code = "BE"
	case Canada:
		code = "CACPA"
	case France:
		code = "FR"
	case Germany:
		code = "DEBLZ"
	case Greece:
		code = "GRBIC"
	case HongKong:
		code = "HKNCC"
	case Italy:
		code = "ITNCC"
	case Luxembourg:
		code = "LULUX"
	case Netherlands:
		code = ""
	case Poland:
		code = "PLKNR"
	case Portugal:
		code = "PTNCC"
	case Spain:
		code = "ESNCC"
	case Switzerland:
		code = "CHBCC"
	case UnitedStates:
		code = "USABA"
	}
	return code
}
