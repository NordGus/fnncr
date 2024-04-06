package account_entity

type Kind string

const (
	NormalAccount  Kind = "capital.normal"
	SavingAccount       = "capital.savings"
	LoanAccount         = "debt.loan"
	CreditAccount       = "debt.credit"
	IncomeAccount       = "external.income"
	ExpenseAccount      = "external.expense"
	HistoryAccount      = "system.history"
	InvalidAccount      = "system.invalid"
)

func ParseKind(kind string) Kind {
	switch kind {
	case "capital.normal":
		return NormalAccount
	case "capital.savings":
		return SavingAccount
	case "debt.loan":
		return LoanAccount
	case "debt.credit":
		return CreditAccount
	case "external.income":
		return IncomeAccount
	case "external.expense":
		return ExpenseAccount
	case "system.history":
		return HistoryAccount
	default:
		return InvalidAccount
	}
}
