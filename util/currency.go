package util

// constants for all supported currencies
const (
	USD = "USD"
	EUR = "EUR"
	CAD = "CAD"
)

// isSupportedCurrency return true if currency is supported
func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, CAD:
		return true
	}

	return false
}
