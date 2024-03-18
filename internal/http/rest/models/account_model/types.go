package account_model

type Type string

const (
	NormalAccount  Type = "capital.normal"
	SavingsAccount      = "capital.savings"
	DebtAccount         = "debt.normal"
	CreditAccount       = "debt.credit"
	IncomeAccount       = "external.income"
	ExpenseAccount      = "external.expense"
)
