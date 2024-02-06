package models

import (
	"fmt"
	"math"
)

type AccountType string

type Account struct {
	AccType        AccountType
	DisplayName    string
	CurrentBalance int64
	Limit          int64
}

const (
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
	return fmt.Sprintf("%d", a.CurrentBalance)
}

func (a Account) Available() string {
	return fmt.Sprintf("%d", a.Limit-a.CurrentBalance)
}

func (a Account) Covered() int16 {
	switch a.AccType {
	case NormalAccount, SavingsAccount, ExternalAccount:
		return 100
	case CreditAccount:
		return int16((math.Floor(float64(a.Limit-a.CurrentBalance) / float64(a.Limit) * 100.0)))
	default:
		return int16((math.Floor(float64(a.CurrentBalance) / float64(a.Limit) * 100.0)))
	}
}
