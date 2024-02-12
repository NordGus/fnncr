package models

type SavingsGoal struct {
	GoalName string
	Goal     int64
	Saved    int64
}

func (sg SavingsGoal) Name() string {
	return sg.GoalName
}

func (sg SavingsGoal) Remaining() string {
	return currencySprintf(sg.Goal - sg.Saved)
}

func (sg SavingsGoal) Balance() string {
	return currencySprintf(sg.Saved)
}

func (sg SavingsGoal) Covered() int64 {
	return int64((float64(sg.Saved) / float64(sg.Goal) * 100))
}
