package models

type (
	GoalCurrency string

	SavingsGoal struct {
		currency GoalCurrency
		name     string
		goal     int64
		balance  int64
	}
)

const (
	USD_ GoalCurrency = "USD"
	EUR_ GoalCurrency = "EUR"
	GBP_ GoalCurrency = "GBP"
	AUD_ GoalCurrency = "AUD"
	CAD_ GoalCurrency = "CAD"
)

func NewSavingsGoal(name string, balance int64, goal int64, currency GoalCurrency) SavingsGoal {
	return SavingsGoal{
		currency: currency,
		name:     name,
		balance:  balance,
		goal:     goal,
	}
}

func (sg SavingsGoal) Name() string {
	return sg.name
}

func (sg SavingsGoal) Remaining() string {
	return currencySprintf(sg.goal-sg.balance, string(sg.currency))
}

func (sg SavingsGoal) Balance() string {
	return currencySprintf(sg.balance, string(sg.currency))
}

func (sg SavingsGoal) Covered() int64 {
	return int64((float64(sg.balance) / float64(sg.goal) * 100))
}
