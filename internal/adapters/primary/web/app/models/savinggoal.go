package models

type SavingsGoal struct {
	name    string
	goal    int64
	balance int64
}

func NewSavingsGoal(name string, balance int64, goal int64) SavingsGoal {
	return SavingsGoal{
		name:    name,
		balance: balance,
		goal:    goal,
	}
}

func (sg SavingsGoal) Name() string {
	return sg.name
}

func (sg SavingsGoal) Remaining() string {
	return currencySprintf(sg.goal - sg.balance)
}

func (sg SavingsGoal) Balance() string {
	return currencySprintf(sg.balance)
}

func (sg SavingsGoal) Covered() int64 {
	return int64((float64(sg.balance) / float64(sg.goal) * 100))
}
