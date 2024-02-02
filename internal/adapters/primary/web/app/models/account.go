package models

type AccountType string

type Account struct {
	AccType     AccountType
	DisplayName string
}

const (
	NormalAccount   AccountType = "normal"
	SavingsAccount              = "savings"
	LoanAccount                 = "loan"
	CreditAccount               = "credit"
	ExternalAccount             = "external"
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
