package models

import (
	"math"

	"golang.org/x/text/number"
)

type SavingsGoal struct {
	GoalName string
	Goal     int64
	Saved    int64
}

func (sg SavingsGoal) Name() string {
	return sg.GoalName
}

func (sg SavingsGoal) Remaining() string {
	return printer.Sprintf(
		"%v",
		number.Decimal(
			float32(sg.Goal-sg.Saved)/cents,
			number.MaxFractionDigits(2),
			number.MinFractionDigits(2),
		),
	)
}

func (sg SavingsGoal) Balance() string {
	return printer.Sprintf(
		"%v",
		number.Decimal(
			float32(sg.Saved)/cents,
			number.MaxFractionDigits(2),
			number.MinFractionDigits(2),
		),
	)
}

func (sg SavingsGoal) Covered() int16 {
	return int16((math.Floor(float64(sg.Saved) / float64(sg.Goal) * 100.0)))
}
