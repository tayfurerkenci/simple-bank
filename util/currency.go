package util

// Constants for all supported currencies
const (
	TRY = "TRY"
	USD = "USD"
	EUR = "EUR"
	CAD = "CAD"
)

// IsSupportedCurrency returns true if the currency is supported
func IsSupportedCurrency(currency string) bool {
	switch currency {
	case TRY, USD, EUR, CAD:
		return true
	}

	return false
}