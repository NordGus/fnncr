package models

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/number"
)

const (
	cents = 100
)

var (
	printer = message.NewPrinter(language.English)
)

func currencySprintf(amount int64, currency string) string {
	return printer.Sprintf(
		"%v %s",
		number.Decimal(
			float64(amount)/cents,
			number.MaxFractionDigits(2),
			number.MinFractionDigits(2),
		),
		currency,
	)
}
