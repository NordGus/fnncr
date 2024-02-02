package models

type AccountType string

type Account struct {
	AccType     AccountType
	DisplayName string
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
	return "420.69"
}

func (a Account) Covered() int16 {
	return 37
}
