package models

import (
	"math"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/number"
)

type AccountType string
type DebtType int8

type Account struct {
	AccType        AccountType
	DebtType       DebtType
	DisplayName    string
	CurrentBalance int64
	Limit          int64
}

const (
	NoneDebt DebtType = iota
	IAmOwedDebt
	IOweDebt

	NormalAccount   AccountType = "normal"
	SavingsAccount  AccountType = "savings"
	LoanAccount     AccountType = "loan"
	CreditAccount   AccountType = "credit"
	ExternalAccount AccountType = "external"

	cents = 100
)

var (
	printer = message.NewPrinter(language.English)
)

func (a Account) Type() string {
	return string(a.AccType)
}

func (a Account) Name() string {
	return a.DisplayName
}

func (a Account) Balance() string {
	return printer.Sprintf(
		"%v",
		number.Decimal(
			float64(a.CurrentBalance)/cents,
			number.MaxFractionDigits(2),
			number.MinFractionDigits(2),
		),
	)
}

func (a Account) Available() string {
	return printer.Sprintf(
		"%v",
		number.Decimal(
			float64(a.Limit+a.CurrentBalance)/cents,
			number.MaxFractionDigits(2),
			number.MinFractionDigits(2),
		),
	)
}

func (a Account) Covered() int16 {
	switch a.AccType {
	case NormalAccount, SavingsAccount, ExternalAccount:
		return 100
	case CreditAccount:
		return int16((math.Floor(float64(a.Limit+a.CurrentBalance) / float64(a.Limit) * 100.0)))
	default:
		return int16((math.Floor(float64(a.CurrentBalance) / float64(a.Limit) * 100.0)))
	}
}

func (a Account) InTheRed() bool {
	switch a.AccType {
	case CreditAccount:
		return (a.Limit + a.CurrentBalance) < 0
	case LoanAccount:
		return a.DebtType == IAmOwedDebt
	default:
		return a.CurrentBalance < 0
	}
}
