package models

import (
	"math"
)

type (
	AccountType string

	AccountCurrency string

	Account struct {
		model    AccountType
		currency AccountCurrency
		name     string
		balance  int64
		limit    int64
	}
)

func NewAccount(at AccountType, name string, balance int64, limit int64, currency AccountCurrency) Account {
	return Account{
		model:    at,
		currency: currency,
		name:     name,
		balance:  balance,
		limit:    limit,
	}
}

const (
	NormalAccount   AccountType = "normal"
	SavingsAccount  AccountType = "savings"
	LoanAccount     AccountType = "loan"
	CreditAccount   AccountType = "credit"
	ExternalAccount AccountType = "external"

	USD AccountCurrency = "USD"
	EUR AccountCurrency = "EUR"
	GBP AccountCurrency = "GBP"
	AUD AccountCurrency = "AUD"
	CAD AccountCurrency = "CAD"
)

func (a Account) Type() string {
	return string(a.model)
}

func (a Account) Name() string {
	return a.name
}

func (a Account) Balance() string {
	return currencySprintf(a.balance, string(a.currency))
}

func (a Account) Available() string {
	switch a.model {
	case CreditAccount:
		return currencySprintf(a.limit+a.balance, string(a.currency))
	default:
		return currencySprintf(a.balance, string(a.currency))
	}
}

func (a Account) Covered() int64 {
	var covered int64 = 100

	if a.model == LoanAccount || a.model == CreditAccount {
		covered = int64((math.Abs(float64(a.balance)) - math.Abs(float64(a.limit))) / math.Abs(float64(a.limit)) * 100)
	}

	if covered < 0 {
		return -covered
	}

	return covered
}

func (a Account) InTheRed() bool {
	switch a.model {
	case CreditAccount:
		return (a.limit + a.balance) < 0
	default:
		return a.balance < 0
	}
}
