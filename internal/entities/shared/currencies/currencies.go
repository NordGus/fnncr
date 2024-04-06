package currencies

type Currency string

const (
	Euro         Currency = "eur"
	USDollar              = "usd"
	BritishPound          = "gbp"
	Invalid               = "invalid"
)

func ParseCurrency(currency string) Currency {
	switch currency {
	case "euro":
		return Euro
	case "usd":
		return USDollar
	case "gbp":
		return BritishPound
	default:
		return Invalid
	}
}
