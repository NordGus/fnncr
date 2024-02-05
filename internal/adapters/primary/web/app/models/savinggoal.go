package models

import (
	"fmt"
	"math"
)

type SavingsGoal struct {
	GoalName string
	Goal     int64
	Achieved int64
}

func (sg SavingsGoal) Name() string {
	return sg.GoalName
}

func (sg SavingsGoal) Remaining() string {
	return fmt.Sprintf("%d", sg.Goal-sg.Achieved)
}

func (sg SavingsGoal) Balance() string {
	return fmt.Sprintf("%d", sg.Achieved)
}

func (sg SavingsGoal) Covered() int16 {
	return int16((math.Floor(float64(sg.Achieved) / float64(sg.Goal) * 100.0)))
}
