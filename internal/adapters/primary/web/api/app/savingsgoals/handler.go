package savingsgoals

import "github.com/labstack/echo/v4"

const (
	// Routes

	// NOTE: Saving Goals are a completely different data structure from accounts but it depends on accounts,
	// specifically savings accounts. As the balance changes in the different savings account so does the completion in
	// each savings goals. This goals contains an priority value order so the lower the priority value the earlier the
	// the goal is fullfil.
	SavingsGoalsRoute = "/savings_goals"
)

type Handler interface {
	ListSavingsGoalsHandlerFunc(c echo.Context) error
}

type handler struct {
}

func New() Handler {
	return &handler{}
}
