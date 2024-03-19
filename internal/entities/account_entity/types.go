package account_entity

type Kind string

const (
	NormalAccount  Kind = "capital.normal"
	SavingsAccount      = "capital.savings"
	DebtAccount         = "debt.normal"
	CreditAccount       = "debt.credit"
	IncomeAccount       = "external.income"
	ExpenseAccount      = "external.expense"
)
