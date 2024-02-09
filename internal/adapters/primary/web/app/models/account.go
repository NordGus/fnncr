package models

import (
	"math"

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
)

func (a Account) Type() string {
	return string(a.AccType)
}

func (a Account) Name() string {
	return a.DisplayName
}

func (a Account) Balance() string {
	switch a.AccType {
	case LoanAccount:
		return printer.Sprintf(
			"%v",
			number.Decimal(
				a.loanAccountBalance(),
				number.MaxFractionDigits(2),
				number.MinFractionDigits(2),
			),
		)
	default:
		return printer.Sprintf(
			"%v",
			number.Decimal(
				float64(a.CurrentBalance)/cents,
				number.MaxFractionDigits(2),
				number.MinFractionDigits(2),
			),
		)
	}
}

func (a Account) Available() string {
	switch a.DebtType {
	case IOweDebt:
		return printer.Sprintf(
			"%v",
			number.Decimal(
				float64((a.Limit+a.CurrentBalance))/cents,
				number.MaxFractionDigits(2),
				number.MinFractionDigits(2),
			),
		)
	case IAmOwedDebt:
		return printer.Sprintf(
			"%v",
			number.Decimal(
				float64(a.Limit-a.CurrentBalance)/cents,
				number.MaxFractionDigits(2),
				number.MinFractionDigits(2),
			),
		)
	default:
		return a.Balance()

	}
}

func (a Account) Covered() int16 {
	switch a.AccType {
	case NormalAccount, SavingsAccount, ExternalAccount:
		return 100
	case CreditAccount:
		return int16(a.creditAccountCovered())
	default:
		return int16(a.loanAccountCovered())
	}
}

func (a Account) InTheRed() bool {
	switch a.AccType {
	case CreditAccount:
		return (a.Limit + a.CurrentBalance) < 0
	case LoanAccount:
		return a.DebtType == IOweDebt
	default:
		return a.CurrentBalance < 0
	}
}

func (a Account) creditAccountCovered() float64 {
	covered := math.Floor(float64(a.Limit+a.CurrentBalance) / float64(a.Limit) * 100.0)

	if a.DebtType == IOweDebt {
		return covered
	}

	return -covered
}

func (a Account) loanAccountCovered() float64 {
	covered := math.Floor(float64(a.CurrentBalance) / float64(a.Limit) * 100.0)

	if a.DebtType == IOweDebt {
		return -covered
	}

	return covered
}

func (a Account) loanAccountBalance() float64 {
	switch a.DebtType {
	case IOweDebt:
		return -float64(a.Limit+a.CurrentBalance) / cents
	default:
		return float64(a.Limit-a.CurrentBalance) / cents
	}
}
